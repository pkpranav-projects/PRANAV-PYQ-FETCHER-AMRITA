package menu

import (
	"errors"
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

// ChooseCourse presents top-level selection with Cyber Security as the primary featured course
func (cs *CourseSelect) ChooseCourse() {
	var selectedOption string
	var options []huh.Option[string]

	// 1. Easy Mode
	options = append(options, huh.NewOption("⚡      │ Easy Mode: Fast Past 5-Year Pack Downloader (Midterms + EndSem by 2-Year Range)", "easy_mode"))

	// 2. Primary featured program: Cyber Security (CYS)
	options = append(options, huh.NewOption("🛡️  CYS │ Cyber Security (B.Tech CYS)", "CYS"))

	// 3. Other Courses & Branches
	options = append(options, huh.NewOption("📚      │ Other Courses & Branches (CSE, ECE, AIE, EEE, MECH, MCA, M.Tech)", "other_courses"))

	// 4. Classic DSpace Explorer
	options = append(options, huh.NewOption("📂      │ Browse Full DSpace Directory (Semester by Semester)", "dspace_full"))

	// 5. Quit
	options = append(options, huh.NewOption("❌      │ Quit", "quit"))

	selectField := huh.NewSelect[string]().
		Title("Select Course or Mode (Press '/' to search)").
		Options(options...).
		Value(&selectedOption)

	form := NewMenuForm(
		huh.NewGroup(selectField),
	)

	if err := form.Run(); err != nil {
		if errors.Is(err, huh.ErrUserAborted) {
			util.QuitWithSpinner()
		}
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	switch selectedOption {
	case "quit":
		util.QuitWithSpinner()
	case "easy_mode":
		em := EasyModeSelect{ReqClient: cs.ReqClient}
		em.StartEasyMode()
		return
	case "dspace_full":
		configs.SelectionHistory = []string{"DSpace Full Directory"}
		cs.browseDSpaceDirectory()
		return
	case "other_courses":
		cs.ChooseOtherCourse()
		return
	case "CYS":
		branches := GetBranchCatalog()
		for _, b := range branches {
			if b.Code == "CYS" {
				configs.SelectionHistory = []string{"Cyber Security (CYS)"}
				cs.ChooseSemester(b)
				return
			}
		}
	}
}

// ChooseOtherCourse lets the user access other engineering and master's courses
func (cs *CourseSelect) ChooseOtherCourse() {
	branches := GetBranchCatalog()

	var selectedBranch string
	var options []huh.Option[string]

	for _, b := range branches {
		if b.Code == "CYS" {
			continue // CYS is featured on the main menu
		}
		label := fmt.Sprintf("%-6s │ %s", b.Code, b.Name)
		options = append(options, huh.NewOption(label, b.Code))
	}
	options = append(options, huh.NewOption("🔙 Back to Main Menu", "back"))
	options = append(options, huh.NewOption("❌ Quit", "quit"))

	selectField := huh.NewSelect[string]().
		Title("Select Other Course / Branch (Press '/' to search)").
		Options(options...).
		Value(&selectedBranch)

	if len(options) > 8 {
		selectField = selectField.WithHeight(14).(*huh.Select[string])
	}

	form := NewMenuForm(
		huh.NewGroup(selectField),
	)

	if err := form.Run(); err != nil {
		if errors.Is(err, huh.ErrUserAborted) {
			cs.ChooseCourse()
			return
		}
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	if selectedBranch == "quit" {
		util.QuitWithSpinner()
	}
	if selectedBranch == "back" {
		cs.ChooseCourse()
		return
	}

	for _, b := range branches {
		if b.Code == selectedBranch {
			configs.SelectionHistory = []string{b.Code + " (" + b.Name + ")"}
			cs.ChooseSemester(b)
			return
		}
	}
}

// ChooseSemester asks for the semester first (Semester 1 to 8, or All)
func (cs *CourseSelect) ChooseSemester(branch BranchItem) {
	for {
		var selectedSem int
		options := []huh.Option[int]{
			huh.NewOption("Semester 1 (1st Year)", 1),
			huh.NewOption("Semester 2 (1st Year)", 2),
			huh.NewOption("Semester 3 (2nd Year)", 3),
			huh.NewOption("Semester 4 (2nd Year)", 4),
			huh.NewOption("Semester 5 (3rd Year)", 5),
			huh.NewOption("Semester 6 (3rd Year)", 6),
			huh.NewOption("Semester 7 (4th Year)", 7),
			huh.NewOption("Semester 8 (4th Year)", 8),
			huh.NewOption("All Semesters (Show all subjects)", 0),
			huh.NewOption("🔙 Back", -1),
			huh.NewOption("❌ Quit", -2),
		}

		configs.SelectionHistory = []string{branch.Code + " (" + branch.Name + ")"}
		selectionDisplay := "Program:\n" + branch.Code + " - " + branch.Name

		selectField := huh.NewSelect[int]().
			Title(fmt.Sprintf("Select Semester for %s:", branch.Name)).
			Options(options...).
			Value(&selectedSem)

		form := NewMenuForm(
			huh.NewGroup(
				huh.NewNote().TitleFunc(func() string { return selectionDisplay }, &configs.SelectionHistory),
				selectField,
			),
		)

		if err := form.Run(); err != nil {
			if errors.Is(err, huh.ErrUserAborted) {
				return
			}
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}

		if selectedSem == -2 {
			util.QuitWithSpinner()
		}
		if selectedSem == -1 {
			configs.SelectionHistory = []string{}
			cs.ChooseCourse()
			return
		}

		cs.ChooseSemesterPapers(branch, selectedSem)
	}
}

// ChooseSemesterPapers displays subjects and question papers for the selected semester
func (cs *CourseSelect) ChooseSemesterPapers(branch BranchItem, sem int) {
	for {
		subjects := FilterSubjectsBySemester(branch.Subjects, sem)

		// Collect all papers across these semester subjects
		var allSemPapers []PaperItem
		for _, s := range subjects {
			allSemPapers = append(allSemPapers, s.Papers...)
		}

		var selectedAction string
		var options []huh.Option[string]

		semLabel := "All Semesters"
		if sem > 0 {
			semLabel = fmt.Sprintf("Semester %d", sem)
			// Option to download all papers for this entire semester at once
			if len(allSemPapers) > 0 {
				options = append(options, huh.NewOption(
					fmt.Sprintf("📥 [DOWNLOAD ALL] Download ALL Question Papers for Semester %d (%d papers)", sem, len(allSemPapers)),
					"download_sem_all",
				))
			}
		}

		// Option to pick a subject and download its papers immediately
		options = append(options, huh.NewOption("📥 [1-CLICK BULK] Select a Subject to Download ALL Its Papers", "bulk_download_subject"))

		// List of subjects for this semester
		for _, subj := range subjects {
			label := fmt.Sprintf("%-10s │ %s (%d papers)", subj.Code, subj.Name, len(subj.Papers))
			options = append(options, huh.NewOption(label, subj.Code))
		}

		options = append(options, huh.NewOption("🔙 Back to Semesters", "back"))
		options = append(options, huh.NewOption("🏠 Main Menu", "main_menu"))
		options = append(options, huh.NewOption("❌ Quit", "quit"))

		configs.SelectionHistory = []string{branch.Code, semLabel}
		selectionDisplay := fmt.Sprintf(
			"Program: %s (%s) → %s\nTotal Available Subjects: %d | Total Papers: %d",
			branch.Code, branch.Name, semLabel, len(subjects), len(allSemPapers),
		)

		selectField := huh.NewSelect[string]().
			Title(fmt.Sprintf("%s - %s: Select Subject or Download All", branch.Code, semLabel)).
			Options(options...).
			Value(&selectedAction)

		if len(options) > 10 {
			selectField = selectField.WithHeight(18).(*huh.Select[string])
		}

		form := NewMenuForm(
			huh.NewGroup(
				huh.NewNote().TitleFunc(func() string { return selectionDisplay }, &configs.SelectionHistory),
				selectField,
			),
		)

		if err := form.Run(); err != nil {
			if errors.Is(err, huh.ErrUserAborted) {
				return
			}
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}

		switch selectedAction {
		case "download_sem_all":
			var fileList []File
			for _, p := range allSemPapers {
				fileList = append(fileList, File{name: p.Name, path: p.Path})
			}
			targetDir := filepath.Join("downloads", branch.Code, fmt.Sprintf("Semester_%d_All", sem))
			DownloadAllFiles(cs.ReqClient.WebClient, fileList, targetDir)

		case "bulk_download_subject":
			cs.quickBulkDownloadSubject(branch)

		case "back":
			return

		case "main_menu":
			configs.SelectionHistory = []string{}
			cs.ChooseCourse()
			return

		case "quit":
			util.QuitWithSpinner()

		default:
			for _, subj := range branch.Subjects {
				if subj.Code == selectedAction {
					// Ask whether to download all papers immediately or browse papers
					var action string
					actionForm := NewMenuForm(
						huh.NewGroup(
							huh.NewSelect[string]().
								Title(fmt.Sprintf("Subject: %s - %s", subj.Code, subj.Name)).
								Options(
									huh.NewOption(fmt.Sprintf("📥 Download ALL Papers for this Subject (%d papers)", len(subj.Papers)), "download_all"),
									huh.NewOption("📋 View / Select Papers & 2-Year Ranges", "browse"),
									huh.NewOption("🔙 Back", "back"),
								).
								Value(&action),
						),
					)

					if err := actionForm.Run(); err != nil {
						if errors.Is(err, huh.ErrUserAborted) {
							continue
						}
						fmt.Println(err)
						continue
					}

					if action == "download_all" {
						var fileList []File
						for _, p := range subj.Papers {
							fileList = append(fileList, File{name: p.Name, path: p.Path})
						}
						targetDir := filepath.Join("downloads", branch.Code, SanitizeFileName(subj.Code), "All_Papers")
						DownloadAllFiles(cs.ReqClient.WebClient, fileList, targetDir)
					} else if action == "browse" {
						cs.ChoosePaper(branch, subj)
					}
					break
				}
			}
		}
	}
}

// quickBulkDownloadSubject allows picking any subject to download all of its question papers in one action
func (cs *CourseSelect) quickBulkDownloadSubject(branch BranchItem) {
	var chosenCode string
	var options []huh.Option[string]

	for _, subj := range branch.Subjects {
		label := fmt.Sprintf("%-10s (Sem %d) │ %s (%d papers)", subj.Code, subj.Semester, subj.Name, len(subj.Papers))
		options = append(options, huh.NewOption(label, subj.Code))
	}
	options = append(options, huh.NewOption("🔙 Back", "back"))

	selectField := huh.NewSelect[string]().
		Title("Select Subject to Download ALL Papers (Press '/' to search):").
		Options(options...).
		Value(&chosenCode)

	if len(options) > 10 {
		selectField = selectField.WithHeight(18).(*huh.Select[string])
	}

	form := NewMenuForm(
		huh.NewGroup(selectField),
	)

	if err := form.Run(); err != nil {
		return
	}

	if chosenCode == "back" {
		return
	}

	for _, subj := range branch.Subjects {
		if subj.Code == chosenCode {
			var fileList []File
			for _, p := range subj.Papers {
				fileList = append(fileList, File{name: p.Name, path: p.Path})
			}
			targetDir := filepath.Join("downloads", branch.Code, SanitizeFileName(subj.Code), "All_Papers")
			DownloadAllFiles(cs.ReqClient.WebClient, fileList, targetDir)
			break
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

		form := NewMenuForm(
			huh.NewGroup(
				huh.NewNote().TitleFunc(func() string { return selectionDisplay }, &configs.SelectionHistory),
				selectField,
			),
		)

		if err := form.Run(); err != nil {
			if errors.Is(err, huh.ErrUserAborted) {
				return
			}
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
			cs.ChooseSemesterPapers(branch, subject.Semester)
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
					actionForm := NewMenuForm(
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
						if errors.Is(err, huh.ErrUserAborted) {
							continue
						}
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
						targetFolder := filepath.Join("downloads", branch.Code, SanitizeFileName(subject.Code), paper.Year)
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

	form := NewMenuForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("DSpace Course Communities").
				Options(options...).
				Value(&selectedOption),
		),
	)

	err = form.Run()
	if err != nil {
		if errors.Is(err, huh.ErrUserAborted) {
			configs.SelectionHistory = []string{}
			cs.ChooseCourse()
			return
		}
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
