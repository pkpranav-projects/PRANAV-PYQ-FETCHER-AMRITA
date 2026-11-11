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
)

type EasyModeSelect struct {
	ReqClient requestclient.RequestClient
}

// StartEasyMode runs the fast pack downloader: Semester -> Branch -> Subject -> 2-Year Range / 5-Year Bulk Downloader
func (em *EasyModeSelect) StartEasyMode() {
	em.ChooseSemester()
}

// ChooseSemester lets the student pick their current semester (1 to 8) or choose all
func (em *EasyModeSelect) ChooseSemester() {
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
		huh.NewOption("All Semesters (Browse all subjects across 4 years)", 0),
		huh.NewOption("🔙 Back to Main Menu", -1),
		huh.NewOption("❌ Quit", -2),
	}

	configs.SelectionHistory = []string{"⚡ Easy Mode"}
	selectionDisplay := "Mode:\n⚡ Easy Mode (Fast Past 5-Year Pack Downloader)"

	selectField := huh.NewSelect[int]().
		Title("Select Semester (Mid-term & End-term papers for your semester)").
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
			configs.SelectionHistory = []string{}
			cs := CourseSelect{ReqClient: em.ReqClient}
			cs.ChooseCourse()
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
		cs := CourseSelect{ReqClient: em.ReqClient}
		cs.ChooseCourse()
		return
	}

	em.ChooseBranch(selectedSem)
}

// ChooseBranch lets the student select their course / branch (CYS, CSE, ECE, etc.)
func (em *EasyModeSelect) ChooseBranch(sem int) {
	branches := GetBranchCatalog()

	var selectedBranchCode string
	var options []huh.Option[string]

	for _, b := range branches {
		label := fmt.Sprintf("%-6s │ %s", b.Code, b.Name)
		options = append(options, huh.NewOption(label, b.Code))
	}
	options = append(options, huh.NewOption("🔙 Back to Semester Selection", "back"))
	options = append(options, huh.NewOption("❌ Quit", "quit"))

	semLabel := "All Semesters"
	if sem > 0 {
		semLabel = fmt.Sprintf("Semester %d", sem)
	}
	configs.SelectionHistory = []string{"⚡ Easy Mode", semLabel}
	selectionDisplay := "Selection(s):\n" + strings.Join(configs.SelectionHistory, " → ")

	selectField := huh.NewSelect[string]().
		Title(fmt.Sprintf("Select Branch for %s (Press '/' to search)", semLabel)).
		Options(options...).
		Value(&selectedBranchCode)

	if len(options) > 8 {
		selectField = selectField.WithHeight(14).(*huh.Select[string])
	}

	form := NewMenuForm(
		huh.NewGroup(
			huh.NewNote().TitleFunc(func() string { return selectionDisplay }, &configs.SelectionHistory),
			selectField,
		),
	)

	if err := form.Run(); err != nil {
		if errors.Is(err, huh.ErrUserAborted) {
			em.ChooseSemester()
			return
		}
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	switch selectedBranchCode {
	case "quit":
		util.QuitWithSpinner()
	case "back":
		em.ChooseSemester()
		return
	default:
		for _, b := range branches {
			if b.Code == selectedBranchCode {
				em.ChooseSubject(sem, b)
				return
			}
		}
	}
}

// ChooseSubject filters subjects by the chosen branch and semester
func (em *EasyModeSelect) ChooseSubject(sem int, branch BranchItem) {
	for {
		matchingSubjects := FilterSubjectsBySemester(branch.Subjects, sem)

		var selectedCode string
		var options []huh.Option[string]

		for _, subj := range matchingSubjects {
			label := fmt.Sprintf("%-10s (Sem %d) │ %s", subj.Code, subj.Semester, subj.Name)
			options = append(options, huh.NewOption(label, subj.Code))
		}
		options = append(options, huh.NewOption("🔙 Back to Branch Selection", "back"))
		options = append(options, huh.NewOption("❌ Quit", "quit"))

		semLabel := "All Semesters"
		if sem > 0 {
			semLabel = fmt.Sprintf("Semester %d", sem)
		}
		configs.SelectionHistory = []string{"⚡ Easy Mode", branch.Code, semLabel}
		selectionDisplay := "Selection(s):\n" + strings.Join(configs.SelectionHistory, " → ")

		selectField := huh.NewSelect[string]().
			Title(fmt.Sprintf("Select Subject for %s - %s (Press '/' to search)", branch.Code, semLabel)).
			Options(options...).
			Value(&selectedCode)

		if len(options) > 8 {
			selectField = selectField.WithHeight(16).(*huh.Select[string])
		}

		form := NewMenuForm(
			huh.NewGroup(
				huh.NewNote().TitleFunc(func() string { return selectionDisplay }, &configs.SelectionHistory),
				selectField,
			),
		)

		if err := form.Run(); err != nil {
			if errors.Is(err, huh.ErrUserAborted) {
				em.ChooseBranch(sem)
				return
			}
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}

		switch selectedCode {
		case "quit":
			util.QuitWithSpinner()
		case "back":
			em.ChooseBranch(sem)
			return
		default:
			for _, subj := range branch.Subjects {
				if subj.Code == selectedCode {
					em.ShowEasyModeDownloader(branch, subj, sem)
					return
				}
			}
		}
	}
}

// ShowEasyModeDownloader displays the 2-year range download buttons and 1-click Past 5 Years pack
func (em *EasyModeSelect) ShowEasyModeDownloader(branch BranchItem, subject SubjectItem, sem int) {
	for {
		// Filter papers by 2-year ranges
		papers5Yr := FilterPapersByRange(subject.Papers, 2021, 2026)
		papers2526 := FilterPapersByRange(subject.Papers, 2025, 2026)
		papers2324 := FilterPapersByRange(subject.Papers, 2023, 2024)
		papers2122 := FilterPapersByRange(subject.Papers, 2021, 2022)
		papers1920 := FilterPapersByRange(subject.Papers, 2019, 2020)

		var selectedAction string
		var options []huh.Option[string]

		// 1-Click Past 5 Years Bulk Download (Midterms + EndSems)
		options = append(options, huh.NewOption(
			fmt.Sprintf("📥 [DOWNLOAD ALL] Past 5 Years Papers (2021 – 2026: Midterms + EndSems) [%d papers]", len(papers5Yr)),
			"download_5yr",
		))

		// 2-Year Range Download Buttons
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

		// Individual paper options
		for _, p := range subject.Papers {
			label := fmt.Sprintf("📄 [%s %-7s] %s", p.Year, p.Type, p.Name)
			options = append(options, huh.NewOption(label, p.Path))
		}

		options = append(options, huh.NewOption("🔙 Back to Subjects", "back_subjects"))
		options = append(options, huh.NewOption("🏠 Main Menu", "main_menu"))
		options = append(options, huh.NewOption("❌ Quit", "quit"))

		configs.SelectionHistory = []string{"⚡ Easy Mode", branch.Code, subject.Code + " " + subject.Name}
		selectionDisplay := fmt.Sprintf(
			"⚡ FAST EXAM PACK DOWNLOADER\nSubject: %s - %s (Semester %d)\nTotal Papers Available: %d (Midterms + EndSemesters)",
			subject.Code, subject.Name, subject.Semester, len(subject.Papers),
		)

		selectField := huh.NewSelect[string]().
			Title("Choose Download Action or 2-Year Range:").
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
				em.ChooseSubject(sem, branch)
				return
			}
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}

		baseFolder := filepath.Join("downloads", branch.Code, SanitizeFileName(subject.Code))

		switch selectedAction {
		case "download_5yr":
			var fileList []File
			for _, p := range papers5Yr {
				fileList = append(fileList, File{name: p.Name, path: p.Path})
			}
			targetDir := filepath.Join(baseFolder, "Past_5_Years_All")
			DownloadAllFiles(em.ReqClient.WebClient, fileList, targetDir)

		case "range_25_26":
			var fileList []File
			for _, p := range papers2526 {
				fileList = append(fileList, File{name: p.Name, path: p.Path})
			}
			targetDir := filepath.Join(baseFolder, "2025-2026")
			DownloadAllFiles(em.ReqClient.WebClient, fileList, targetDir)

		case "range_23_24":
			var fileList []File
			for _, p := range papers2324 {
				fileList = append(fileList, File{name: p.Name, path: p.Path})
			}
			targetDir := filepath.Join(baseFolder, "2023-2024")
			DownloadAllFiles(em.ReqClient.WebClient, fileList, targetDir)

		case "range_21_22":
			var fileList []File
			for _, p := range papers2122 {
				fileList = append(fileList, File{name: p.Name, path: p.Path})
			}
			targetDir := filepath.Join(baseFolder, "2021-2022")
			DownloadAllFiles(em.ReqClient.WebClient, fileList, targetDir)

		case "range_19_20":
			var fileList []File
			for _, p := range papers1920 {
				fileList = append(fileList, File{name: p.Name, path: p.Path})
			}
			targetDir := filepath.Join(baseFolder, "2019-2020")
			DownloadAllFiles(em.ReqClient.WebClient, fileList, targetDir)

		case "back_subjects":
			em.ChooseSubject(sem, branch)
			return

		case "main_menu":
			configs.SelectionHistory = []string{}
			cs := CourseSelect{ReqClient: em.ReqClient}
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
						if err := em.ReqClient.WebClient.OpenBrowser(paperURL); err != nil {
							fmt.Println(configs.ErrorStyle.Render(fmt.Sprintf("Error: %v\n", err)))
						}
					} else if action == "download" {
						cleanName := SanitizeFileName(paper.Name)
						if !strings.HasSuffix(strings.ToLower(cleanName), ".pdf") {
							cleanName += ".pdf"
						}
						targetFolder := filepath.Join(baseFolder, paper.Year)
						downloadPath := filepath.Join(targetFolder, cleanName)
						fmt.Printf("Downloading to %s...\n", downloadPath)

						if err := em.ReqClient.WebClient.DownloadFile(paperURL, downloadPath); err != nil {
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
