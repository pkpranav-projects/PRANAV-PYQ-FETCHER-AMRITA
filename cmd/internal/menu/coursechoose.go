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

	// Add Easy Mode (Fast Past 5-Year Pack Downloader) at the top
	options = append(options, huh.NewOption("⚡      │ Easy Mode: Fast Past 5-Year Pack Downloader (Midterms + EndSem by 2-Year Range)", "easy_mode"))

	// Add quick branch options
	for _, b := range branches {
		label := fmt.Sprintf("%-6s │ %s", b.Code, b.Name)
		options = append(options, huh.NewOption(label, b.Code))
	}

	// Add classic explorer and quit options
	options = append(options, huh.NewOption("📂      │ Browse Full DSpace Directory (Semester by Semester)", "dspace_full"))
	options = append(options, huh.NewOption("❌      │ Quit", "quit"))

	selectField := huh.NewSelect[string]().
		Title("Select Course / Branch or Mode (Press '/' to search)").
		Options(options...).
		Value(&selectedBranch)

	if len(options) > 8 {
		selectField = selectField.WithHeight(16).(*huh.Select[string])
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

	if selectedBranch == "easy_mode" {
		em := EasyModeSelect{ReqClient: cs.ReqClient}
		em.StartEasyMode()
		return
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

		// Add "Download ALL Papers" and 2-Year Ranges
		if len(subject.Papers) > 0 {
			options = append(options, huh.NewOption(
				fmt.Sprintf("📥 [DOWNLOAD ALL] Past 5 Years Papers (2021 – 2026: Midterms + EndSems) [%d papers]", len(subject.Papers)),
				"download_all",
			))
			papers2526 := FilterPapersByRange(subject.Papers, 2025, 2026)
			papers2324 := FilterPapersByRange(subject.Papers, 2023, 2024)
			papers2122 := FilterPapersByRange(subject.Papers, 2021, 2022)
			papers1920 := FilterPapersByRange(subject.Papers, 2019, 2020)

			options = append(options, huh.NewOption(
				fmt.Sprintf("📦 [2-YEAR RANGE] 2025 – 2026 Papers (Midterm & EndSem) [%d papers]", len(papers2526)),
				"range_25_26",
			))
			options = append(options, huh.NewOption(
				fmt.Sprintf("📦 [2-YEAR RANGE] 2023 – 2024 Papers (Midterm & EndSem) [%d papers]", len(papers2324)),
				"range_23_24",
			))
			options = append(options, huh.NewOption(
				fmt.Sprintf("📦 [2-YEAR RANGE] 2021 – 2022 Papers (Midterm & EndSem) [%d papers]", len(papers2122)),
				"range_21_22",
			))
			options = append(options, huh.NewOption(
				fmt.Sprintf("📦 [2-YEAR RANGE] 2019 – 2020 Papers (Midterm & EndSem) [%d papers]", len(papers1920)),
				"range_19_20",
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
			targetDir := filepath.Join("downloads", branch.Code, SanitizeFileName(subject.Code), "Past_5_Years_All")
			DownloadAllFiles(cs.ReqClient.WebClient, fileList, targetDir)

		case "range_25_26":
			papers := FilterPapersByRange(subject.Papers, 2025, 2026)
			var fileList []File
			for _, p := range papers {
				fileList = append(fileList, File{name: p.Name, path: p.Path})
			}
			targetDir := filepath.Join("downloads", branch.Code, SanitizeFileName(subject.Code), "2025-2026")
			DownloadAllFiles(cs.ReqClient.WebClient, fileList, targetDir)

		case "range_23_24":
			papers := FilterPapersByRange(subject.Papers, 2023, 2024)
			var fileList []File
			for _, p := range papers {
				fileList = append(fileList, File{name: p.Name, path: p.Path})
			}
			targetDir := filepath.Join("downloads", branch.Code, SanitizeFileName(subject.Code), "2023-2024")
			DownloadAllFiles(cs.ReqClient.WebClient, fileList, targetDir)

		case "range_21_22":
			papers := FilterPapersByRange(subject.Papers, 2021, 2022)
			var fileList []File
			for _, p := range papers {
				fileList = append(fileList, File{name: p.Name, path: p.Path})
			}
			targetDir := filepath.Join("downloads", branch.Code, SanitizeFileName(subject.Code), "2021-2022")
			DownloadAllFiles(cs.ReqClient.WebClient, fileList, targetDir)

		case "range_19_20":
			papers := FilterPapersByRange(subject.Papers, 2019, 2020)
			var fileList []File
			for _, p := range papers {
				fileList = append(fileList, File{name: p.Name, path: p.Path})
			}
			targetDir := filepath.Join("downloads", branch.Code, SanitizeFileName(subject.Code), "2019-2020")
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
