package menu

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"amrita_pyq/cmd/internal/configs"
	"amrita_pyq/cmd/internal/requestclient"
	"amrita_pyq/cmd/util"

	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/huh/spinner"
)

type (
	SubCommunity struct {
		name string
		path string
	}

	SubCommunityTable struct {
		ReqClient requestclient.RequestClient
	}
)

func (sct *SubCommunityTable) ChooseSubCommunity(url string) {
	for {
		action := func() {
			time.Sleep(2 * time.Second)
		}
		if err := spinner.New().Title("Fetching Sub Communities...").Action(action).Run(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		subCommunities, err := sct.ReqClient.SubComReq(url)

		if err != nil {
			fmt.Print(configs.ErrorStyle.Render(fmt.Sprintf("Error: %v\n", err)))
			return
		}

		var selectedOption string
		var options []huh.Option[string]
		var subComList []SubCommunity

		// Convert subCommunities to huh options.
		for _, subCommunity := range subCommunities {
			subComItem := SubCommunity{subCommunity.Name, subCommunity.Path}
			subComList = append(subComList, subComItem)
			options = append(options, huh.NewOption(subComItem.name, subComItem.path))
		}
		// Add back option.
		options = append(options, huh.NewOption("Back to Main Menu", "back"))
		options = append(options, huh.NewOption("Quit", "quit"))
		selectionDisplay := "Selection(s):\n" + strings.Join(configs.SelectionHistory, " → ")

		// Create the select field.
		selectField := huh.NewSelect[string]().
			Title("Select Sub Community to view").
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
		case "back":
			cs := CourseSelect{
				ReqClient: sct.ReqClient,
			}
			cs.ChooseCourse()
		case "quit":
			util.QuitWithSpinner() // Quit the program.
		default:
			// Find selected subCommunity and process it.
			for _, subComItem := range subComList {
				if subComItem.path == selectedOption {
					url := configs.BASE_URL + subComItem.path

					// Ask whether to download all papers immediately or view them
					var action string
					actionForm := huh.NewForm(
						huh.NewGroup(
							huh.NewSelect[string]().
								Title(fmt.Sprintf("Subject: %s", subComItem.name)).
								Options(
									huh.NewOption("📥 Download ALL Papers for this Subject Immediately", "download_all"),
									huh.NewOption("📋 View / Select Individual Papers", "view"),
									huh.NewOption("🔙 Back", "back"),
								).
								Value(&action),
						),
					)

					if err := actionForm.Run(); err != nil {
						fmt.Println(err)
						continue
					}

					if action == "download_all" {
						files, err := sct.ReqClient.YearReq(url)
						if err != nil {
							fmt.Println(configs.ErrorStyle.Render(fmt.Sprintf("Error fetching papers: %v\n", err)))
							continue
						}
						var fileList []File
						for _, f := range files {
							fileList = append(fileList, File{name: f.Name, path: f.Path})
						}
						targetDir := filepath.Join("downloads", SanitizeFileName(subComItem.name))
						DownloadAllFiles(sct.ReqClient.WebClient, fileList, targetDir)
					} else if action == "view" {
						yearTable := YearTable{
							ReqClient: sct.ReqClient,
						}
						yearTable.ChooseQP(url)
					}
					break
				}
			}
		}
	}
}
