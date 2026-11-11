package webclient

import (
	"amrita_pyq/cmd/internal/configs"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestFetchHTML(t *testing.T) {
	client := DefaultWebClient{}

	tests := []struct {
		name    string
		url     string
		wantErr bool
	}{
		{
			name:    "SuccessValidURLFetch",
			url:     "https://httpbin.org/get",
			wantErr: false,
		},
		{
			name:    "FailInvalidURLFetch",
			url:     "http://invalid.invalid",
			wantErr: true,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			_, gotErr := client.FetchHTML(tc.url)
			if tc.wantErr && gotErr == nil {
				t.Errorf("expected error for URL %s, got nil", tc.url)
			} else if !tc.wantErr && gotErr != nil {
				t.Errorf("unexpected error for URL %s: %v", tc.url, gotErr)
			}
		})
	}
}

func TestOpenBrowser(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping TestOpenBrowser in short mode")
	}

	client := DefaultWebClient{}

	tests := []struct {
		name    string
		url     string
		wantErr bool
	}{
		{
			name:    "SuccessOpenValidURLInBrowser",
			url:     "https://httpbin.org/get",
			wantErr: false,
		},
		// TODO: Add test cases to cover error scenarios.
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			if gotErr := client.OpenBrowser(tc.url); (gotErr != nil) != tc.wantErr {
				t.Errorf("for URL %s, expected error: %v, got: %v", tc.url, tc.wantErr, gotErr)
			}
		})
	}
}

func TestDownloadFileDemoMode(t *testing.T) {
	configs.IsDemoMode = true
	defer func() { configs.IsDemoMode = false }()

	client := DefaultWebClient{}
	tmpFile := filepath.Join(t.TempDir(), "test_demo.pdf")

	err := client.DownloadFile("http://example.com/demo.pdf", tmpFile)
	if err != nil {
		t.Fatalf("DownloadFile failed in demo mode: %v", err)
	}

	data, err := os.ReadFile(tmpFile)
	if err != nil {
		t.Fatalf("failed to read downloaded file: %v", err)
	}

	if !strings.HasPrefix(string(data), "%PDF-1.4") {
		t.Errorf("demo file is not a valid PDF: expected prefix '%%PDF-1.4', got %q", string(data[:min(10, len(data))]))
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
