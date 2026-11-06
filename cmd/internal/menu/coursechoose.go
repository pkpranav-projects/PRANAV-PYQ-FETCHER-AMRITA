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
	Subject struct {
		name string
		path string
	}

	CourseSelect struct {
		ReqClient requestclient.RequestClient
	}
)

// ChooseCourse presents top-level branch selection (CYS, CSE, ECE, AIE, etc.)
// and the option to browse the classic full DSpace directory.
func (cs *CourseSelect) ChooseCourse() {
	branches := GetBranchCatalog()

	var selectedBranch string
	var options []huh.Option[string]

	// Add quick branch options
	for _, b := range branches {
		label := fmt.Sprintf("%-6s │ %s", b.Code, b.Name)
		options = append(options, huh.NewOption(label, b.Code))
	}

	// Add classic explorer and quit options
	options = append(options, huh.NewOption("📂      │ Browse Full DSpace Directory (Semester by Semester)", "dspace_full"))
	options = append(options, huh.NewOption("❌      │ Quit", "quit"))

	selectField := huh.NewSelect[string]().
		Title("Select Course / Branch (Press '/' to search)").
		Options(options...).
		Value(&selectedBranch)

	if len(options) > 8 {
		selectField = selectField.WithHeight(14).(*huh.Select[string])
	}

	form := huh.NewForm(
		huh.NewGroup(
			selectField,
		),
	)

	if err := form.Run(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	if selectedBranch == "quit" {
		util.QuitWithSpinner()
	}

	if selectedBranch == "dspace_full" {
		configs.SelectionHistory = []string{"DSpace Full Directory"}
		cs.browseDSpaceDirectory()
		return
	}

	// Find the chosen branch
	for _, b := range branches {
		if b.Code == selectedBranch {
			configs.SelectionHistory = []string{b.Code + " (" + b.Name + ")"}
			cs.ChooseSubject(b)
			return
		}
	}
}

// ChooseSubject displays all subjects under the selected course/branch.
func (cs *CourseSelect) ChooseSubject(branch BranchItem) {
	for {
		var selectedSubjectCode string
		var options []huh.Option[string]

		for _, subj := range branch.Subjects {
			label := fmt.Sprintf("%-10s │ %s", subj.Code, subj.Name)
			options = append(options, huh.NewOption(label, subj.Code))
		}

		options = append(options, huh.NewOption("🔙 Back to Courses", "back"))
		options = append(options, huh.NewOption("❌ Quit", "quit"))

		selectionDisplay := "Selection(s):\n" + strings.Join(configs.SelectionHistory, " → ")

		selectField := huh.NewSelect[string]().
			Title(fmt.Sprintf("Select Subject for %s (Press '/' to filter)", branch.Code)).
			Options(options...).
			Value(&selectedSubjectCode)

		if len(options) > 10 {
			selectField = selectField.WithHeight(18).(*huh.Select[string])
		}

		form := huh.NewForm(
			huh.NewGroup(
				huh.NewNote().TitleFunc(func() string { return selectionDisplay }, &configs.SelectionHistory),
				selectField,
			),
		)

		if err := form.Run(); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}

		switch selectedSubjectCode {
		case "back":
			configs.SelectionHistory = []string{}
			cs.ChooseCourse()
			return
		case "quit":
			util.QuitWithSpinner()
		default:
			for _, subj := range branch.Subjects {
				if subj.Code == selectedSubjectCode {
					cs.ChoosePaper(branch, subj)
					return
				}
			}
		}
	}
}

// ChoosePaper displays question papers for the selected subject and provides bulk download.
func (cs *CourseSelect) ChoosePaper(branch BranchItem, subject SubjectItem) {
	for {
		var selectedAction string
		var options []huh.Option[string]

		// Add "Download ALL Papers" option at top
		if len(subject.Papers) > 0 {
			options = append(options, huh.NewOption(
				fmt.Sprintf("📥 Download ALL Papers (%d available)", len(subject.Papers)),
				"download_all",
			))
		}

		// Individual papers
		for _, paper := range subject.Papers {
			options = append(options, huh.NewOption(paper.Name, paper.Path))
		}

		options = append(options, huh.NewOption("🔙 Back to Subjects", "back_subjects"))
		options = append(options, huh.NewOption("🏠 Main Menu", "main_menu"))
		options = append(options, huh.NewOption("❌ Quit", "quit"))

		configs.SelectionHistory = []string{branch.Code, subject.Code + " " + subject.Name}
		selectionDisplay := "Selection(s):\n" + strings.Join(configs.SelectionHistory, " → ")

		selectField := huh.NewSelect[string]().
			Title(fmt.Sprintf("Papers for %s: %s", subject.Code, subject.Name)).
			Options(options...).
			Value(&selectedAction)

		if len(options) > 10 {
			selectField = selectField.WithHeight(18).(*huh.Select[string])
		}

		form := huh.NewForm(
			huh.NewGroup(
				huh.NewNote().TitleFunc(func() string { return selectionDisplay }, &configs.SelectionHistory),
				selectField,
			),
		)

		if err := form.Run(); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}

		switch selectedAction {
		case "download_all":
			var fileList []File
			for _, p := range subject.Papers {
				fileList = append(fileList, File{name: p.Name, path: p.Path})
			}
			targetDir := filepath.Join("downloads", branch.Code, SanitizeFileName(subject.Code))
			DownloadAllFiles(cs.ReqClient.WebClient, fileList, targetDir)

		case "back_subjects":
			configs.SelectionHistory = []string{branch.Code + " (" + branch.Name + ")"}
			cs.ChooseSubject(branch)
			return

		case "main_menu":
			configs.SelectionHistory = []string{}
			cs.ChooseCourse()
			return

		case "quit":
			util.QuitWithSpinner()

		default:
			// Individual paper action
			for _, paper := range subject.Papers {
				if paper.Path == selectedAction {
					paperURL := configs.BASE_URL + paper.Path
					if strings.HasPrefix(paper.Path, "http://") || strings.HasPrefix(paper.Path, "https://") {
						paperURL = paper.Path
					}

					var action string
					actionForm := huh.NewForm(
						huh.NewGroup(
							huh.NewSelect[string]().
								Title(fmt.Sprintf("Action for: %s", paper.Name)).
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
						if err := cs.ReqClient.WebClient.OpenBrowser(paperURL); err != nil {
							fmt.Println(configs.ErrorStyle.Render(fmt.Sprintf("Error: %v\n", err)))
						}
					} else if action == "download" {
						cleanName := SanitizeFileName(paper.Name)
						if !strings.HasSuffix(strings.ToLower(cleanName), ".pdf") {
							cleanName += ".pdf"
						}
						targetFolder := filepath.Join("downloads", branch.Code, SanitizeFileName(subject.Code))
						downloadPath := filepath.Join(targetFolder, cleanName)
						fmt.Printf("Downloading to %s...\n", downloadPath)

						if err := cs.ReqClient.WebClient.DownloadFile(paperURL, downloadPath); err != nil {
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

// browseDSpaceDirectory queries the raw DSpace handles
func (cs *CourseSelect) browseDSpaceDirectory() {
	action := func() {
		time.Sleep(1 * time.Second)
	}
	if err := spinner.New().Title("Fetching DSpace Communities...").Action(action).Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	resources, err := cs.ReqClient.GetCoursesReq(configs.COURSE_LIST_URL)
	if err != nil {
		fmt.Println(configs.ErrorStyle.Render(fmt.Sprintf("Error: %v\n", err)))
		os.Exit(1)
	}

	var selectedOption string
	var subjects []Subject
	var options []huh.Option[string]

	for _, res := range resources {
		subject := Subject{res.Name, res.Path}
		subjects = append(subjects, subject)
		options = append(options, huh.NewOption(subject.name, subject.name))
	}
	options = append(options, huh.NewOption("🔙 Back to Main Menu", "back"))
	options = append(options, huh.NewOption("❌ Quit", "quit"))

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("DSpace Course Communities").
				Options(options...).
				Value(&selectedOption),
		),
	)

	err = form.Run()
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}

	if selectedOption == "back" {
		configs.SelectionHistory = []string{}
		cs.ChooseCourse()
		return
	}

	if selectedOption == "quit" {
		util.QuitWithSpinner()
	}

	configs.SelectionHistory = []string{selectedOption}

	for _, subject := range subjects {
		if subject.name == selectedOption {
			url := configs.BASE_URL + subject.path
			semTable := SemTable{
				ReqClient: cs.ReqClient,
			}
			semTable.ChooseQuestionSetFromSemester(url)
			break
		}
	}
}
