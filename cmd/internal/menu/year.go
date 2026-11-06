package menu

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"amrita_pyq/cmd/internal/configs"
	"amrita_pyq/cmd/internal/requestclient"
	"amrita_pyq/cmd/internal/webclient"
	"amrita_pyq/cmd/util"

	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/huh/spinner"
)

type (
	File struct {
		name string
		path string
	}

	YearTable struct {
		ReqClient requestclient.RequestClient
	}
)

// SanitizeFileName cleans invalid characters from file names for local saving.
func SanitizeFileName(name string) string {
	replacer := strings.NewReplacer(
		" ", "_",
		"/", "-",
		"\\", "-",
		":", "-",
		"*", "",
		"?", "",
		"\"", "",
		"<", "",
		">", "",
		"|", "-",
	)
	return replacer.Replace(name)
}

// DownloadAllFiles downloads a list of question paper files in bulk with a progress log.
func DownloadAllFiles(client webclient.WebClient, fileList []File, targetDir string) {
	if targetDir == "" {
		targetDir = "downloads"
	}
	_ = os.MkdirAll(targetDir, 0755)

	fmt.Println()
	fmt.Println(configs.FetchStatusStyle.Render(fmt.Sprintf("📥 Bulk Downloading %d question paper(s) into ./%s...", len(fileList), targetDir)))
	fmt.Println()

	successCount := 0
	for i, fileItem := range fileList {
		cleanName := SanitizeFileName(fileItem.name)
		if !strings.HasSuffix(strings.ToLower(cleanName), ".pdf") {
			cleanName += ".pdf"
		}
		targetPath := filepath.Join(targetDir, cleanName)
		url := configs.BASE_URL + fileItem.path
		if strings.HasPrefix(fileItem.path, "http://") || strings.HasPrefix(fileItem.path, "https://") {
			url = fileItem.path
		}

		fmt.Printf("  [%d/%d] Downloading %s... ", i+1, len(fileList), fileItem.name)
		if err := client.DownloadFile(url, targetPath); err != nil {
			fmt.Printf("❌ %v\n", err)
		} else {
			fmt.Printf("✓ Saved\n")
			successCount++
		}
		time.Sleep(100 * time.Millisecond)
	}

	fmt.Println()
	if successCount == len(fileList) {
		fmt.Println(configs.SuccessStyle.Render(fmt.Sprintf("✓ Bulk Download Complete! All %d papers saved in ./%s", successCount, targetDir)))
	} else {
		fmt.Println(configs.WarningStyle.Render(fmt.Sprintf("Downloaded %d of %d papers into ./%s", successCount, len(fileList), targetDir)))
	}
	fmt.Println()
	time.Sleep(1 * time.Second)
}

func (yt *YearTable) ChooseQP(url string) {
	for {
		action := func() {
			time.Sleep(2 * time.Second)
		}
		if err := spinner.New().Title("Fetching question papers...").Action(action).Run(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		files, err := yt.ReqClient.YearReq(url)
		if err != nil {
			fmt.Print(configs.ErrorStyle.Render(fmt.Sprintf("Error: %v\n", err)))
			return
		}

		var selectedOption string
		var options []huh.Option[string]
		var fileList []File

		// Add "Download ALL Papers" option if files exist
		if len(files) > 1 {
			options = append(options, huh.NewOption(fmt.Sprintf("📥 Download ALL Papers (%d available)", len(files)), "download_all"))
		}

		// Convert files to huh options.
		for _, file := range files {
			fileItem := File{file.Name, file.Path}
			fileList = append(fileList, fileItem)
			options = append(options, huh.NewOption(fileItem.name, fileItem.path))
		}
		// Add back option.
		options = append(options, huh.NewOption("🔙 Back to Main Menu", "back"))
		options = append(options, huh.NewOption("❌ Quit", "quit"))
		selectionDisplay := "Selection(s):\n" + strings.Join(configs.SelectionHistory, " -> ")

		// Create the select field.
		selectField := huh.NewSelect[string]().
			Title("Select Question Paper to view or download").
			Options(options...).
			Value(&selectedOption)

		// Limit number of options displayed when
		// there are more than 10 options to prevent overflow.
		if len(options) > 10 {
			selectField = selectField.WithHeight(20).(*huh.Select[string])
		}

		// Create the form.
		form := huh.NewForm(
			huh.NewGroup(
				huh.NewNote().
					TitleFunc(func() string { return selectionDisplay }, &configs.SelectionHistory),
				selectField,
			),
		)

		// Run the form.
		err = form.Run()
		if err != nil {
			fmt.Printf("Error: %v", err)
			os.Exit(1)
		}

		// Handle selection.
		switch selectedOption {
		case "download_all":
			DownloadAllFiles(yt.ReqClient.WebClient, fileList, "downloads")
		case "back":
			cs := CourseSelect{
				ReqClient: yt.ReqClient,
			}
			cs.ChooseCourse()
			return
		case "quit":
			util.QuitWithSpinner() // Quit the program.
		default:
			// Find selected file and process it.
			for _, fileItem := range fileList {
				if fileItem.path == selectedOption {
					paperURL := configs.BASE_URL + fileItem.path
					if strings.HasPrefix(fileItem.path, "http://") || strings.HasPrefix(fileItem.path, "https://") {
						paperURL = fileItem.path
					}

					// --- Ask to Open or Download ---
					var action string
					actionForm := huh.NewForm(
						huh.NewGroup(
							huh.NewSelect[string]().
								Title(fmt.Sprintf("Action for: %s", fileItem.name)).
								Options(
									huh.NewOption("🌐 Open in Browser", "open"),
									huh.NewOption("💾 Download PDF", "download"),
									huh.NewOption("🔙 Back", "back"),
								).
								Value(&action),
						),
					)

					if err := actionForm.Run(); err != nil {
						fmt.Println(err)
						continue
					}

					if action == "open" {
						// Open the file in the browser.
						if err := yt.ReqClient.WebClient.OpenBrowser(paperURL); err != nil {
							fmt.Println(configs.ErrorStyle.Render(fmt.Sprintf("Error: %v\n", err)))
						}
					} else if action == "download" {
						// Download the file into downloads folder
						cleanName := SanitizeFileName(fileItem.name)
						if !strings.HasSuffix(strings.ToLower(cleanName), ".pdf") {
							cleanName += ".pdf"
						}
						downloadPath := filepath.Join("downloads", cleanName)
						fmt.Printf("Downloading to %s...\n", downloadPath)

						if err := yt.ReqClient.WebClient.DownloadFile(paperURL, downloadPath); err != nil {
							fmt.Println(configs.ErrorStyle.Render(fmt.Sprintf("Download Error: %v\n", err)))
						} else {
							fmt.Println(configs.SuccessStyle.Render(fmt.Sprintf("✓ Download Complete: %s", downloadPath)))
							time.Sleep(1 * time.Second)
						}
					}
					break
				}
			}
		}
	}
}
