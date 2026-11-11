package webclient

import (
	"amrita_pyq/cmd/internal/configs"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

// WebClient defines methods for fetching HTML, opening a browser, and downloading files.
type WebClient interface {
	FetchHTML(url string) (string, error)
	OpenBrowser(url string) error
	DownloadFile(url string, filename string) error
}

// DefaultWebClient implements WebClient using real network calls or demo data.
type DefaultWebClient struct{}

// FetchHTML fetches and parses HTML from the given URL.
func (d DefaultWebClient) FetchHTML(url string) (string, error) {
	if configs.IsDemoMode {
		return getDemoHTML(url), nil
	}

	client := &http.Client{
		Timeout: 12 * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		fmt.Println(configs.ErrorStyle.Render("Error fetching the URL. Make sure you're connected to Amrita WiFi or VPN."))
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected HTTP status: %s", resp.Status)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(bodyBytes), nil
}

// OpenBrowser opens a URL in the default web browser.
func (d DefaultWebClient) OpenBrowser(url string) error {
	if configs.IsDemoMode {
		// In demo mode, create a local preview HTML page
		previewPath := filepath.Join("downloads", "demo_paper_preview.html")
		_ = os.MkdirAll("downloads", 0755)
		htmlContent := fmt.Sprintf(`<!DOCTYPE html>
<html>
<head>
    <title>PRANAV PYQ FETCHER AMRITA - Demo Preview</title>
    <style>
        body { font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif; background: #0f172a; color: #f8fafc; padding: 40px; }
        .card { background: #1e293b; border-radius: 12px; padding: 32px; max-width: 700px; margin: 0 auto; box-shadow: 0 10px 25px rgba(0,0,0,0.5); border: 1px solid #334155; }
        h1 { color: #38bdf8; margin-top: 0; }
        .badge { background: #0284c7; color: white; padding: 4px 12px; border-radius: 9999px; font-size: 13px; font-weight: bold; display: inline-block; }
        .info { background: #334155; border-radius: 8px; padding: 16px; margin: 20px 0; font-size: 14px; line-height: 1.6; }
        .btn { display: inline-block; background: #38bdf8; color: #0f172a; font-weight: bold; padding: 10px 20px; border-radius: 6px; text-decoration: none; margin-top: 10px; }
    </style>
</head>
<body>
    <div class="card">
        <span class="badge">OFFLINE DEMO MODE</span>
        <h1>PRANAV PYQ FETCHER AMRITA</h1>
        <p>This is a simulated question paper preview viewed while not connected to Amrita campus Wi-Fi.</p>
        <div class="info">
            <strong>Target Resource:</strong> %s<br>
            <strong>Note:</strong> When you connect to <strong>Amrita Wi-Fi</strong> or <strong>Amrita VPN</strong>, this will automatically open the real official PDF directly from the DSpace university repository!
        </div>
        <p>✨ Everything in your app is configured and ready to go!</p>
    </div>
</body>
</html>`, url)
		_ = os.WriteFile(previewPath, []byte(htmlContent), 0644)
		absPath, _ := filepath.Abs(previewPath)
		url = "file:///" + filepath.ToSlash(absPath)
	}

	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	}

	if err != nil {
		styledMessage := configs.ErrorStyle.Render("failed to open browser")
		return fmt.Errorf("%s: %w", styledMessage, err)
	}

	return nil
}

// createDemoPDF generates a syntactically valid PDF 1.4 document containing the offline demo notice.
func createDemoPDF(filename, targetURL string) error {
	streamContent := fmt.Sprintf("BT\n/F1 18 Tf\n50 720 Td\n(PRANAV PYQ FETCHER AMRITA) Tj\n/F1 12 Tf\n0 -30 Td\n(Offline Demo Mode - Simulated Question Paper) Tj\n0 -25 Td\n(Connect to Amrita campus Wi-Fi or VPN to download live PDFs.) Tj\nET\n")

	pdfContent := fmt.Sprintf("%%PDF-1.4\n"+
		"1 0 obj\n<< /Type /Catalog /Pages 2 0 R >>\nendobj\n"+
		"2 0 obj\n<< /Type /Pages /Kids [3 0 R] /Count 1 >>\nendobj\n"+
		"3 0 obj\n<< /Type /Page /Parent 2 0 R /MediaBox [0 0 612 792] /Resources << /Font << /F1 4 0 R >> >> /Contents 5 0 R >>\nendobj\n"+
		"4 0 obj\n<< /Type /Font /Subtype /Type1 /BaseFont /Helvetica >>\nendobj\n"+
		"5 0 obj\n<< /Length %d >>\nstream\n%sendstream\nendobj\n"+
		"xref\n0 6\n0000000000 65535 f \n0000000009 00000 n \n0000000058 00000 n \n0000000115 00000 n \n0000000236 00000 n \n0000000305 00000 n \n"+
		"trailer\n<< /Size 6 /Root 1 0 R >>\nstartxref\n480\n%%%%EOF", len(streamContent), streamContent)

	return os.WriteFile(filename, []byte(pdfContent), 0644)
}

// DownloadFile downloads a file from the URL and saves it to the local disk inside the downloads/ folder.
func (d DefaultWebClient) DownloadFile(url string, filename string) error {
	// Ensure directory exists
	dir := filepath.Dir(filename)
	if dir != "" && dir != "." {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
	}

	if configs.IsDemoMode {
		return createDemoPDF(filename, url)
	}

	// Fetch data from server
	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check server response
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("server returned status: %s", resp.Status)
	}

	// Read initial chunk to verify it is a valid PDF
	headerBuf := make([]byte, 512)
	n, err := io.ReadFull(resp.Body, headerBuf)
	if err != nil && err != io.EOF && err != io.ErrUnexpectedEOF {
		return err
	}
	headerBuf = headerBuf[:n]

	// Verify PDF magic bytes (%PDF)
	if !strings.HasPrefix(string(headerBuf), "%PDF") {
		return fmt.Errorf("file received from server is not a valid PDF")
	}

	// Create output file only after successful response and PDF verification
	out, err := os.Create(filename)
	if err != nil {
		return err
	}

	// Write verified header chunk
	if _, err := out.Write(headerBuf); err != nil {
		out.Close()
		_ = os.Remove(filename)
		return err
	}

	// Stream remainder of response body
	if _, err := io.Copy(out, resp.Body); err != nil {
		out.Close()
		_ = os.Remove(filename)
		return err
	}

	out.Close()
	return nil
}

// getDemoHTML provides mock responses reflecting the DSpace XMLUI schema.
func getDemoHTML(url string) string {
	switch {
	case strings.Contains(url, "16") || url == configs.COURSE_LIST_URL:
		return `<html><body>
			<div id="aspect_artifactbrowser_CommunityViewer_div_community-view">
				<div class="artifact-title"><span>B.Tech Cyber Security (CYS)</span><a href="/xmlui/handle/123456789/btech_cys"></a></div>
				<div class="artifact-title"><span>B.Tech Computer Science & Engineering (CSE)</span><a href="/xmlui/handle/123456789/btech_cse"></a></div>
				<div class="artifact-title"><span>B.Tech Electronics & Communication (ECE)</span><a href="/xmlui/handle/123456789/btech_ece"></a></div>
				<div class="artifact-title"><span>B.Tech Artificial Intelligence (AIE)</span><a href="/xmlui/handle/123456789/btech_ai"></a></div>
				<div class="artifact-title"><span>B.Tech Electrical & Electronics (EEE)</span><a href="/xmlui/handle/123456789/btech_eee"></a></div>
				<div class="artifact-title"><span>B.Tech Mechanical Engineering (MECH)</span><a href="/xmlui/handle/123456789/btech_mech"></a></div>
				<div class="artifact-title"><span>Master of Computer Applications (MCA)</span><a href="/xmlui/handle/123456789/mca"></a></div>
				<div class="artifact-title"><span>M.Tech Computer Science</span><a href="/xmlui/handle/123456789/mtech_cse"></a></div>
			</div>
		</body></html>`

	case strings.Contains(url, "btech") || strings.Contains(url, "cys") || strings.Contains(url, "cse") || strings.Contains(url, "ece") || strings.Contains(url, "mca") || strings.Contains(url, "mtech"):
		return `<html><body>
			<div id="aspect_artifactbrowser_CommunityViewer_div_community-view">
				<ul>
					<li><a href="/xmlui/handle/123456789/sem1"><span>Semester 1</span></a></li>
					<li><a href="/xmlui/handle/123456789/sem2"><span>Semester 2</span></a></li>
					<li><a href="/xmlui/handle/123456789/sem3"><span>Semester 3</span></a></li>
					<li><a href="/xmlui/handle/123456789/sem4"><span>Semester 4</span></a></li>
					<li><a href="/xmlui/handle/123456789/sem5"><span>Semester 5</span></a></li>
					<li><a href="/xmlui/handle/123456789/sem6"><span>Semester 6</span></a></li>
					<li><a href="/xmlui/handle/123456789/sem7"><span>Semester 7</span></a></li>
					<li><a href="/xmlui/handle/123456789/sem8"><span>Semester 8</span></a></li>
				</ul>
			</div>
		</body></html>`

	case strings.Contains(url, "sem"):
		return `<html><body>
			<div id="aspect_artifactbrowser_CommunityViewer_div_community-view">
				<ul>
					<li><a href="/xmlui/handle/123456789/assessment_p1"><span>Periodical 1</span></a></li>
					<li><a href="/xmlui/handle/123456789/assessment_p2"><span>Periodical 2</span></a></li>
					<li><a href="/xmlui/handle/123456789/assessment_endsem"><span>End Semester Examination</span></a></li>
				</ul>
			</div>
		</body></html>`

	case strings.Contains(url, "assessment"):
		return `<html><body>
			<div xmlns="http://di.tamu.edu/DRI/1.0/">
				<ul>
					<li><a href="/xmlui/handle/123456789/subj_cys201">20CYS201 - Optimization Techniques</a></li>
					<li><a href="/xmlui/handle/123456789/subj_cys202">20CYS202 - User Interface Design</a></li>
					<li><a href="/xmlui/handle/123456789/subj_cys203">20CYS203 - Operating Systems</a></li>
					<li><a href="/xmlui/handle/123456789/subj_cys204">20CYS204 - Database Management Systems</a></li>
					<li><a href="/xmlui/handle/123456789/subj_cys205">20CYS205 - Modern Cryptography</a></li>
					<li><a href="/xmlui/handle/123456789/subj_mat205">21MAT205 - Differential Equations and Numerical Methods</a></li>
					<li><a href="/xmlui/handle/123456789/subj_dsa">19CSE201 - Data Structures & Algorithms</a></li>
					<li><a href="/xmlui/handle/123456789/subj_os">19CSE202 - Operating Systems</a></li>
					<li><a href="/xmlui/handle/123456789/subj_dbms">19CSE203 - Database Management Systems</a></li>
					<li><a href="/xmlui/handle/123456789/subj_cn">19CSE301 - Computer Networks</a></li>
				</ul>
			</div>
		</body></html>`

	default:
		return `<html><body>
			<div class="file-list">
				<div class="file-wrapper">
					<div><a href="/xmlui/bitstream/handle/123456789/demo_2023_endsem.pdf"></a></div>
					<div><div><span title="ignored"></span><span title="2023_End_Semester_Examination_Solved"></span></div></div>
				</div>
				<div class="file-wrapper">
					<div><a href="/xmlui/bitstream/handle/123456789/demo_2023_p2.pdf"></a></div>
					<div><div><span title="ignored"></span><span title="2023_Periodical_2"></span></div></div>
				</div>
				<div class="file-wrapper">
					<div><a href="/xmlui/bitstream/handle/123456789/demo_2023_p1.pdf"></a></div>
					<div><div><span title="ignored"></span><span title="2023_Periodical_1"></span></div></div>
				</div>
				<div class="file-wrapper">
					<div><a href="/xmlui/bitstream/handle/123456789/demo_2022_endsem.pdf"></a></div>
					<div><div><span title="ignored"></span><span title="2022_End_Semester_Examination"></span></div></div>
				</div>
				<div class="file-wrapper">
					<div><a href="/xmlui/bitstream/handle/123456789/demo_2021_endsem.pdf"></a></div>
					<div><div><span title="ignored"></span><span title="2021_End_Semester_Examination"></span></div></div>
				</div>
			</div>
		</body></html>`
	}
}
