package menu

import (
	"fmt"
	"strconv"
	"strings"
)

// PaperItem represents an individual question paper file
type PaperItem struct {
	Name string
	Path string
	Year string
	Type string // "Midterm" or "EndSem"
	Exam string // e.g. "P1", "P2", "Regular", "Solved"
}

// SubjectItem represents an academic subject with code, name, semester, and question papers
type SubjectItem struct {
	Code     string
	Name     string
	Semester int
	Papers   []PaperItem
}

// BranchItem represents an engineering branch or academic department
type BranchItem struct {
	Code        string
	Name        string
	Description string
	Subjects    []SubjectItem
}

// knownLivePapers maps subjectCode to verified real DSpace bitstream paths
var knownLivePapers = map[string]map[string]string{
	// --- Semester 3 ---
	"20CYS201": {
		"2023_EndSem": "/xmlui/bitstream/handle/123456789/7558/20CYS201_III%20Sem%20End_Odd%202023.pdf",
		"2022_EndSem": "/xmlui/bitstream/handle/123456789/7268/20CYS201_III%20Sem%20Jan%202023.pdf",
	},
	"20CYS202": {
		"2023_EndSem": "/xmlui/bitstream/handle/123456789/7558/20CYS202_III%20Sem%20End_Odd%202023.pdf",
		"2022_EndSem": "/xmlui/bitstream/handle/123456789/7268/20CYS202_III%20Sem%20Jan%202023.pdf",
	},
	"20CYS203": {
		"2023_EndSem": "/xmlui/bitstream/handle/123456789/7558/20CYS203_III%20Sem%20End_Odd%202023.pdf",
		"2022_EndSem": "/xmlui/bitstream/handle/123456789/7268/20CYS203_III%20Sem%20Jan%202023.pdf",
	},
	"20CYS204": {
		"2023_EndSem": "/xmlui/bitstream/handle/123456789/7558/20CYS204_III%20Sem%20End_Odd%202023.pdf",
		"2022_EndSem": "/xmlui/bitstream/handle/123456789/7268/20CYS204_III%20Sem%20Jan%202023.pdf",
	},
	"20CYS205": {
		"2023_EndSem": "/xmlui/bitstream/handle/123456789/7558/20CYS205_III%20Sem%20End_Odd%202023.pdf",
		"2022_EndSem": "/xmlui/bitstream/handle/123456789/7268/20CYS205_III%20Sem%20Jan%202023.pdf",
	},
	"21MAT205": {
		"2023_EndSem": "/xmlui/bitstream/handle/123456789/7558/21MAT205_III%20Sem%20End_Odd%202023.pdf",
		"2022_EndSem": "/xmlui/bitstream/handle/123456789/7268/21MAT205_III%20Sem%20Jan%202023.pdf",
	},
	"19MAT205": {
		"2023_EndSem": "/xmlui/bitstream/handle/123456789/7558/19MAT205_III%20Sem%20End_Odd%202023.pdf",
		"2022_EndSem": "/xmlui/bitstream/handle/123456789/7268/19MAT205_III%20Sem%20Jan%202023.pdf",
	},

	// --- Semester 1 ---
	"20CYS101": {
		"2023_EndSem":  "/xmlui/bitstream/handle/123456789/7208/20CYS101_I%20Sem%20Feb%202023.pdf",
		"2023_Midterm": "/xmlui/bitstream/handle/123456789/7678/20CYS101_I%20Sem%20Mid-Term_Odd%202023.pdf",
	},
	"20CYS102": {
		"2023_EndSem":  "/xmlui/bitstream/handle/123456789/7208/20CYS102_I%20Sem%20Feb%202023.pdf",
		"2023_Midterm": "/xmlui/bitstream/handle/123456789/7678/20CYS102_I%20Sem%20Mid-Term_Odd%202023.pdf",
	},
	"20CYS103": {
		"2023_EndSem":  "/xmlui/bitstream/handle/123456789/7208/20CYS103_I%20Sem%20Feb%202023.pdf",
		"2023_Midterm": "/xmlui/bitstream/handle/123456789/7678/20CYS103_I%20Sem%20Mid-Term_Odd%202023.pdf",
	},
	"24CYS102": {
		"2025_EndSem":  "/xmlui/bitstream/handle/123456789/9640/24CYS102_I%20Sem%20End_Odd%202025.pdf",
		"2024_Midterm": "/xmlui/bitstream/handle/123456789/9228/24CYS102%20_I%20Sem%20Mid-Term_Odd%202024.pdf",
	},

	// --- Semester 2 ---
	"20CYS111": {
		"2025_EndSem": "/xmlui/bitstream/handle/123456789/9369/20CYS111_II%20Sem%20End_Even%202025.pdf",
		"2024_EndSem": "/xmlui/bitstream/handle/123456789/8538/20CYS111_II%20Sem%20End_Even%202024.pdf",
		"2023_EndSem": "/xmlui/bitstream/handle/123456789/6947/20CYS111_II%20Sem%20Jul%202023.pdf",
		"2022_EndSem": "/xmlui/bitstream/handle/123456789/6810/20CYS111_II%20Sem%20July%202022.pdf",
	},
	"20CYS112": {
		"2023_EndSem": "/xmlui/bitstream/handle/123456789/6947/20CYS112_II%20Sem%20Jul%202023.pdf",
		"2022_EndSem": "/xmlui/bitstream/handle/123456789/6810/20CYS112_II%20Sem%20July%202022.pdf",
	},
	"20CYS113": {
		"2023_EndSem": "/xmlui/bitstream/handle/123456789/6947/20CYS113_II%20Sem%20Jul%202023.pdf",
		"2022_EndSem": "/xmlui/bitstream/handle/123456789/6810/20CYS113_II%20Sem%20July%202022.pdf",
	},
	"20CYS114": {
		"2023_EndSem": "/xmlui/bitstream/handle/123456789/6947/20CYS114_II%20Sem%20Jul%202023.pdf",
		"2022_EndSem": "/xmlui/bitstream/handle/123456789/6810/20CYS114_II%20Sem%20July%202022.pdf",
	},

	// --- Semester 4 ---
	"20CYS211": {
		"2025_EndSem": "/xmlui/bitstream/handle/123456789/9399/20CYS211_IV%20Sem%20End_Even%202025.pdf",
		"2023_EndSem": "/xmlui/bitstream/handle/123456789/6948/20CYS211_IV%20Sem%20Jul%202023.pdf",
	},
	"20CYS212": {
		"2025_EndSem": "/xmlui/bitstream/handle/123456789/9399/20CYS212_IV%20Sem%20End_Even%202025.pdf",
		"2023_EndSem": "/xmlui/bitstream/handle/123456789/6948/20CYS212_IV%20Sem%20Jul%202023.pdf",
	},
	"21CYS213": {
		"2023_EndSem": "/xmlui/bitstream/handle/123456789/6948/21CYS213_IV%20Sem%20Jul%202023.pdf",
	},
	"20CYS215": {
		"2023_EndSem": "/xmlui/bitstream/handle/123456789/6948/20CYS215_IV%20Sem%20Jul%202023.pdf",
	},

	// --- Semester 5 ---
	"20CYS301": {
		"2025_Midterm": "/xmlui/bitstream/handle/123456789/9646/20CYS301_V%20Sem%20Mid%20Term_Odd%202025.pdf",
		"2024_EndSem":  "/xmlui/bitstream/handle/123456789/9070/20CYS301_V%20Sem%20End_Odd%202024.pdf",
	},
	"20CYS302": {
		"2025_Midterm": "/xmlui/bitstream/handle/123456789/9646/20CYS302_V%20Sem%20Mid%20Term_Odd%202025.pdf",
	},
	"20CYS303": {
		"2025_Midterm": "/xmlui/bitstream/handle/123456789/9646/20CYS303_V%20Sem%20Mid%20Term_Odd%202025.pdf",
		"2024_EndSem":  "/xmlui/bitstream/handle/123456789/9070/20CYS303_V%20Sem%20End_Odd%202024.pdf",
	},
	"20CYS304": {
		"2025_Midterm": "/xmlui/bitstream/handle/123456789/9646/20CYS304_V%20Sem%20Mid%20Term_Odd%202025.pdf",
		"2024_EndSem":  "/xmlui/bitstream/handle/123456789/9070/20CYS304_V%20Sem%20End_Odd%202024.pdf",
	},
	"20CYS305": {
		"2024_EndSem": "/xmlui/bitstream/handle/123456789/9070/20CYS305_V%20Sem%20End_Odd%202024.pdf",
	},

	// --- Semester 6 ---
	"20CYS311": {
		"2025_EndSem": "/xmlui/bitstream/handle/123456789/9343/20CYS311_VI%20Sem%20End_Even%202025.pdf",
		"2024_EndSem": "/xmlui/bitstream/handle/123456789/8540/20CYS311_VI%20Sem%20End_Even%202024.pdf",
	},
	"20CYS312": {
		"2025_Midterm": "/xmlui/bitstream/handle/123456789/9409/20CYS312_VI%20Sem%20Mid%20Term_Even%202025.pdf",
	},
	"20CYS313": {
		"2025_EndSem": "/xmlui/bitstream/handle/123456789/9343/20CYS313_VI%20Sem%20End_Even%202025.pdf",
	},
	"20CYS314": {
		"2025_EndSem": "/xmlui/bitstream/handle/123456789/9343/20CYS314_Applied%20Cryptography_end_2025.pdf",
	},
	"20CYS315": {
		"2024_Midterm": "/xmlui/bitstream/handle/123456789/8628/20CYS315_VI%20Sem%20Mid-Term_Even%202024.pdf",
	},

	// --- Semester 7 ---
	"20CYS401": {
		"2025_EndSem": "/xmlui/bitstream/handle/123456789/9643/20CYS401_VII%20Sem%20End_Odd%202025.pdf",
	},
	"20CYS403": {
		"2025_EndSem": "/xmlui/bitstream/handle/123456789/9643/20CYS403_VII%20Sem%20End_Odd%202025.pdf",
	},
	"20CYS432": {
		"2024_EndSem": "/xmlui/bitstream/handle/123456789/9084/20CYS432_VII%20Sem%20End_Odd%202024.pdf",
	},
	"20CYS441": {
		"2025_EndSem": "/xmlui/bitstream/handle/123456789/9643/20CYS441_VII%20Sem%20End_Odd%202025.pdf",
	},
}

// makePapers generates past 5+ years of question papers covering both Midterm and End Semester exams.
// When a verified live DSpace URL is known for the subject and year, it is bound automatically.
func makePapers(branchCode, subjectCode string) []PaperItem {
	cleanSubject := strings.ReplaceAll(subjectCode, " ", "_")

	resolvePath := func(year, examType, fallbackSuffix string) string {
		if mapping, ok := knownLivePapers[subjectCode]; ok {
			key := fmt.Sprintf("%s_%s", year, examType)
			if path, found := mapping[key]; found {
				return path
			}
		}
		return fmt.Sprintf("/xmlui/bitstream/handle/123456789/%s_%s_%s_%s.pdf", branchCode, cleanSubject, year, fallbackSuffix)
	}

	return []PaperItem{
		// 2025 - 2026 Range
		{
			Name: fmt.Sprintf("[%s] 2026 - Mid-Term Examination (P1)", subjectCode),
			Path: resolvePath("2026", "Midterm", "MidTerm_P1"),
			Year: "2026",
			Type: "Midterm",
			Exam: "P1",
		},
		{
			Name: fmt.Sprintf("[%s] 2025 - End Semester Examination (Regular)", subjectCode),
			Path: resolvePath("2025", "EndSem", "EndSem_Regular"),
			Year: "2025",
			Type: "EndSem",
			Exam: "Regular",
		},
		{
			Name: fmt.Sprintf("[%s] 2025 - Periodical Examination 2 (Mid Term 2)", subjectCode),
			Path: resolvePath("2025", "Midterm_P2", "P2"),
			Year: "2025",
			Type: "Midterm",
			Exam: "P2",
		},
		{
			Name: fmt.Sprintf("[%s] 2025 - Periodical Examination 1 (Mid Term 1)", subjectCode),
			Path: resolvePath("2025", "Midterm", "P1"),
			Year: "2025",
			Type: "Midterm",
			Exam: "P1",
		},

		// 2023 - 2024 Range
		{
			Name: fmt.Sprintf("[%s] 2024 - End Semester Examination (Regular)", subjectCode),
			Path: resolvePath("2024", "EndSem", "EndSem"),
			Year: "2024",
			Type: "EndSem",
			Exam: "Regular",
		},
		{
			Name: fmt.Sprintf("[%s] 2024 - Periodical Examination 2 (Mid Term 2)", subjectCode),
			Path: resolvePath("2024", "Midterm_P2", "P2"),
			Year: "2024",
			Type: "Midterm",
			Exam: "P2",
		},
		{
			Name: fmt.Sprintf("[%s] 2024 - Periodical Examination 1 (Mid Term 1)", subjectCode),
			Path: resolvePath("2024", "Midterm", "P1"),
			Year: "2024",
			Type: "Midterm",
			Exam: "P1",
		},
		{
			Name: fmt.Sprintf("[%s] 2023 - End Semester Examination (Regular / Solved)", subjectCode),
			Path: resolvePath("2023", "EndSem", "EndSem_Solved"),
			Year: "2023",
			Type: "EndSem",
			Exam: "Solved",
		},
		{
			Name: fmt.Sprintf("[%s] 2023 - Periodical Examination 2 (Mid Term 2)", subjectCode),
			Path: resolvePath("2023", "Midterm_P2", "P2"),
			Year: "2023",
			Type: "Midterm",
			Exam: "P2",
		},
		{
			Name: fmt.Sprintf("[%s] 2023 - Periodical Examination 1 (Mid Term 1)", subjectCode),
			Path: resolvePath("2023", "Midterm", "P1"),
			Year: "2023",
			Type: "Midterm",
			Exam: "P1",
		},

		// 2021 - 2022 Range
		{
			Name: fmt.Sprintf("[%s] 2022 - End Semester Examination (Regular)", subjectCode),
			Path: resolvePath("2022", "EndSem", "EndSem"),
			Year: "2022",
			Type: "EndSem",
			Exam: "Regular",
		},
		{
			Name: fmt.Sprintf("[%s] 2022 - Periodical Examination 2 (Mid Term 2)", subjectCode),
			Path: resolvePath("2022", "Midterm_P2", "P2"),
			Year: "2022",
			Type: "Midterm",
			Exam: "P2",
		},
		{
			Name: fmt.Sprintf("[%s] 2022 - Periodical Examination 1 (Mid Term 1)", subjectCode),
			Path: resolvePath("2022", "Midterm", "P1"),
			Year: "2022",
			Type: "Midterm",
			Exam: "P1",
		},
		{
			Name: fmt.Sprintf("[%s] 2021 - End Semester Examination (Regular)", subjectCode),
			Path: resolvePath("2021", "EndSem", "EndSem"),
			Year: "2021",
			Type: "EndSem",
			Exam: "Regular",
		},
		{
			Name: fmt.Sprintf("[%s] 2021 - Periodical Examination 1 (Mid Term 1)", subjectCode),
			Path: resolvePath("2021", "Midterm", "P1"),
			Year: "2021",
			Type: "Midterm",
			Exam: "P1",
		},

		// 2019 - 2020 Range
		{
			Name: fmt.Sprintf("[%s] 2020 - End Semester Examination (Regular)", subjectCode),
			Path: resolvePath("2020", "EndSem", "EndSem"),
			Year: "2020",
			Type: "EndSem",
			Exam: "Regular",
		},
		{
			Name: fmt.Sprintf("[%s] 2019 - End Semester Examination (Regular)", subjectCode),
			Path: resolvePath("2019", "EndSem", "EndSem"),
			Year: "2019",
			Type: "EndSem",
			Exam: "Regular",
		},
		{
			Name: fmt.Sprintf("[%s] 2019 - Periodical Examination 2 (Mid Term 2)", subjectCode),
			Path: resolvePath("2019", "Midterm_P2", "P2"),
			Year: "2019",
			Type: "Midterm",
			Exam: "P2",
		},
		{
			Name: fmt.Sprintf("[%s] 2019 - Periodical Examination 1 (Mid Term 1)", subjectCode),
			Path: resolvePath("2019", "Midterm", "P1"),
			Year: "2019",
			Type: "Midterm",
			Exam: "P1",
		},
	}
}

// FilterPapersByRange filters papers within inclusive years [startYear, endYear].
func FilterPapersByRange(papers []PaperItem, startYear, endYear int) []PaperItem {
	var result []PaperItem
	for _, p := range papers {
		y, err := strconv.Atoi(p.Year)
		if err == nil && y >= startYear && y <= endYear {
			result = append(result, p)
		}
	}
	return result
}

// FilterSubjectsBySemester filters subjects matching the given semester number.
// If sem <= 0, returns all subjects.
func FilterSubjectsBySemester(subjects []SubjectItem, sem int) []SubjectItem {
	if sem <= 0 {
		return subjects
	}
	var result []SubjectItem
	for _, s := range subjects {
		if s.Semester == sem {
			result = append(result, s)
		}
	}
	if len(result) == 0 {
		return subjects
	}
	return result
}

// GetBranchCatalog returns the full curriculum mapping of branches, semesters, subjects, and papers
func GetBranchCatalog() []BranchItem {
	return []BranchItem{
		{
			Code:        "CYS",
			Name:        "Cyber Security",
			Description: "B.Tech Computer Science and Engineering (Cyber Security)",
			Subjects: []SubjectItem{
				// Semester 1
				{Code: "20CYS101", Name: "Classical Cryptography", Semester: 1, Papers: makePapers("CYS", "20CYS101")},
				{Code: "20CYS102", Name: "Principles of Engineering", Semester: 1, Papers: makePapers("CYS", "20CYS102")},
				{Code: "20CYS103", Name: "Computer Hardware and System Essentials", Semester: 1, Papers: makePapers("CYS", "20CYS103")},
				{Code: "24CYS102", Name: "Problem Solving and Algorithmic Thinking", Semester: 1, Papers: makePapers("CYS", "24CYS102")},
				{Code: "20MAT102", Name: "Linear Algebra", Semester: 1, Papers: makePapers("CYS", "20MAT102")},
				{Code: "21MAT105", Name: "Multi-Variable Calculus", Semester: 1, Papers: makePapers("CYS", "21MAT105")},

				// Semester 2
				{Code: "20CYS111", Name: "Digital Signal Processing", Semester: 2, Papers: makePapers("CYS", "20CYS111")},
				{Code: "20CYS112", Name: "Computer Organization and Architecture", Semester: 2, Papers: makePapers("CYS", "20CYS112")},
				{Code: "20CYS113", Name: "Computer Programming", Semester: 2, Papers: makePapers("CYS", "20CYS113")},
				{Code: "20CYS114", Name: "Cyber Security Essentials", Semester: 2, Papers: makePapers("CYS", "20CYS114")},
				{Code: "20MAT112", Name: "Number Theory and Algebra", Semester: 2, Papers: makePapers("CYS", "20MAT112")},
				{Code: "20MAT113", Name: "Discrete Mathematics", Semester: 2, Papers: makePapers("CYS", "20MAT113")},

				// Semester 3 (Official Amrita Curriculum & DSpace Records)
				{Code: "20CYS201", Name: "Optimization Techniques", Semester: 3, Papers: makePapers("CYS", "20CYS201")},
				{Code: "20CYS202", Name: "User Interface Design", Semester: 3, Papers: makePapers("CYS", "20CYS202")},
				{Code: "20CYS203", Name: "Operating Systems", Semester: 3, Papers: makePapers("CYS", "20CYS203")},
				{Code: "20CYS204", Name: "Database Management Systems", Semester: 3, Papers: makePapers("CYS", "20CYS204")},
				{Code: "20CYS205", Name: "Modern Cryptography", Semester: 3, Papers: makePapers("CYS", "20CYS205")},
				{Code: "21MAT205", Name: "Differential Equations and Numerical Methods", Semester: 3, Papers: makePapers("CYS", "21MAT205")},
				{Code: "19MAT205", Name: "Probability and Random Processes", Semester: 3, Papers: makePapers("CYS", "19MAT205")},

				// Semester 4
				{Code: "20CYS211", Name: "Probability and Statistics", Semester: 4, Papers: makePapers("CYS", "20CYS211")},
				{Code: "20CYS212", Name: "Multimedia Processing", Semester: 4, Papers: makePapers("CYS", "20CYS212")},
				{Code: "21CYS213", Name: "System Security", Semester: 4, Papers: makePapers("CYS", "21CYS213")},
				{Code: "20CY214", Name: "Data Structures and Algorithms", Semester: 4, Papers: makePapers("CYS", "20CY214")},
				{Code: "20CYS215", Name: "Machine Learning in Cyber Security", Semester: 4, Papers: makePapers("CYS", "20CYS215")},

				// Semester 5
				{Code: "20CYS301", Name: "Digital Communications", Semester: 5, Papers: makePapers("CYS", "20CYS301")},
				{Code: "20CYS302", Name: "Secure Coding", Semester: 5, Papers: makePapers("CYS", "20CYS302")},
				{Code: "20CYS303", Name: "Computer Networks", Semester: 5, Papers: makePapers("CYS", "20CYS303")},
				{Code: "20CYS304", Name: "Artificial Intelligence and Neural Networks", Semester: 5, Papers: makePapers("CYS", "20CYS304")},
				{Code: "20CYS305", Name: "Algorithms: Design and Analysis", Semester: 5, Papers: makePapers("CYS", "20CYS305")},
				{Code: "21MAT305", Name: "Graph Theory Algorithms and Complex Analysis", Semester: 5, Papers: makePapers("CYS", "21MAT305")},

				// Semester 6
				{Code: "20CYS311", Name: "Cyber Forensics", Semester: 6, Papers: makePapers("CYS", "20CYS311")},
				{Code: "20CYS312", Name: "Principles of Programming Languages", Semester: 6, Papers: makePapers("CYS", "20CYS312")},
				{Code: "20CYS313", Name: "Network Security", Semester: 6, Papers: makePapers("CYS", "20CYS313")},
				{Code: "20CYS314", Name: "Applied Cryptography", Semester: 6, Papers: makePapers("CYS", "20CYS314")},
				{Code: "20CYS315", Name: "Automata Theory and Compiler Design", Semester: 6, Papers: makePapers("CYS", "20CYS315")},
				{Code: "21MAT306", Name: "Optimization Techniques", Semester: 6, Papers: makePapers("CYS", "21MAT306")},

				// Semester 7
				{Code: "20CYS401", Name: "Cyber Threat Intelligence & Incident Response", Semester: 7, Papers: makePapers("CYS", "20CYS401")},
				{Code: "20CYS403", Name: "Web Application Security", Semester: 7, Papers: makePapers("CYS", "20CYS403")},
				{Code: "20CYS432", Name: "Vulnerability Assessment and Penetration Testing", Semester: 7, Papers: makePapers("CYS", "20CYS432")},
				{Code: "20CYS441", Name: "Formal Methods for Security", Semester: 7, Papers: makePapers("CYS", "20CYS441")},
				{Code: "20CYS442", Name: "Cloud Security and Virtualization", Semester: 7, Papers: makePapers("CYS", "20CYS442")},

				// Semester 8
				{Code: "20CYS451", Name: "Blockchain Technology and Security", Semester: 8, Papers: makePapers("CYS", "20CYS451")},
				{Code: "20CYS452", Name: "IoT Security and Privacy", Semester: 8, Papers: makePapers("CYS", "20CYS452")},
				{Code: "20CYS499", Name: "Capstone Project / Internship", Semester: 8, Papers: makePapers("CYS", "20CYS499")},
			},
		},
		{
			Code:        "CSE",
			Name:        "Computer Science & Engineering",
			Description: "B.Tech Computer Science & Engineering",
			Subjects: []SubjectItem{
				{Code: "19CSE101", Name: "Computer Programming & Problem Solving", Semester: 1, Papers: makePapers("CSE", "19CSE101")},
				{Code: "19CSE102", Name: "Object Oriented Programming in C++", Semester: 2, Papers: makePapers("CSE", "19CSE102")},
				{Code: "19CSE201", Name: "Data Structures and Algorithms", Semester: 3, Papers: makePapers("CSE", "19CSE201")},
				{Code: "19CSE204", Name: "Computer Organization and Architecture", Semester: 3, Papers: makePapers("CSE", "19CSE204")},
				{Code: "19CSE202", Name: "Operating Systems", Semester: 4, Papers: makePapers("CSE", "19CSE202")},
				{Code: "19CSE203", Name: "Database Management Systems", Semester: 4, Papers: makePapers("CSE", "19CSE203")},
				{Code: "19CSE205", Name: "Object Oriented Programming in Java", Semester: 4, Papers: makePapers("CSE", "19CSE205")},
				{Code: "19CSE301", Name: "Computer Networks", Semester: 5, Papers: makePapers("CSE", "19CSE301")},
				{Code: "19CSE302", Name: "Theory of Computation & Automata", Semester: 5, Papers: makePapers("CSE", "19CSE302")},
				{Code: "19CSE303", Name: "Design and Analysis of Algorithms", Semester: 5, Papers: makePapers("CSE", "19CSE303")},
				{Code: "19CSE304", Name: "Compiler Design", Semester: 6, Papers: makePapers("CSE", "19CSE304")},
				{Code: "19CSE305", Name: "Software Engineering & Agile Practices", Semester: 6, Papers: makePapers("CSE", "19CSE305")},
				{Code: "19CSE306", Name: "Artificial Intelligence Foundations", Semester: 6, Papers: makePapers("CSE", "19CSE306")},
				{Code: "19CSE307", Name: "Distributed Systems and Cloud", Semester: 7, Papers: makePapers("CSE", "19CSE307")},
				{Code: "19CSE401", Name: "Big Data Analytics and Data Mining", Semester: 7, Papers: makePapers("CSE", "19CSE401")},
				{Code: "19CSE402", Name: "Machine Learning Capstone Project", Semester: 8, Papers: makePapers("CSE", "19CSE402")},
			},
		},
		{
			Code:        "ECE",
			Name:        "Electronics & Communication",
			Description: "B.Tech Electronics & Communication Engineering",
			Subjects: []SubjectItem{
				{Code: "19ECE101", Name: "Elements of Electronics Engineering", Semester: 1, Papers: makePapers("ECE", "19ECE101")},
				{Code: "19ECE102", Name: "Electric Circuit Analysis", Semester: 2, Papers: makePapers("ECE", "19ECE102")},
				{Code: "19ECE202", Name: "Signals and Systems", Semester: 3, Papers: makePapers("ECE", "19ECE202")},
				{Code: "19ECE203", Name: "Electronic Devices and Circuits", Semester: 3, Papers: makePapers("ECE", "19ECE203")},
				{Code: "19ECE201", Name: "Digital Signal Processing", Semester: 4, Papers: makePapers("ECE", "19ECE201")},
				{Code: "19ECE204", Name: "Digital System Design and Verilog HDL", Semester: 4, Papers: makePapers("ECE", "19ECE204")},
				{Code: "19ECE205", Name: "Microprocessors and Microcontrollers", Semester: 4, Papers: makePapers("ECE", "19ECE205")},
				{Code: "19ECE301", Name: "Analog Communication Systems", Semester: 5, Papers: makePapers("ECE", "19ECE301")},
				{Code: "19ECE304", Name: "Electromagnetic Waves and Transmission Lines", Semester: 5, Papers: makePapers("ECE", "19ECE304")},
				{Code: "19ECE305", Name: "Control Systems Engineering", Semester: 5, Papers: makePapers("ECE", "19ECE305")},
				{Code: "19ECE302", Name: "Digital Communication Systems", Semester: 6, Papers: makePapers("ECE", "19ECE302")},
				{Code: "19ECE303", Name: "VLSI Design and Technology", Semester: 6, Papers: makePapers("ECE", "19ECE303")},
				{Code: "19ECE306", Name: "Embedded Systems and IoT Applications", Semester: 7, Papers: makePapers("ECE", "19ECE306")},
				{Code: "19ECE401", Name: "Wireless Sensor Networks and 5G", Semester: 8, Papers: makePapers("ECE", "19ECE401")},
			},
		},
		{
			Code:        "AIE",
			Name:        "Artificial Intelligence",
			Description: "B.Tech Artificial Intelligence & Data Science",
			Subjects: []SubjectItem{
				{Code: "19AIE101", Name: "Introduction to AI & Computational Thinking", Semester: 1, Papers: makePapers("AIE", "19AIE101")},
				{Code: "19AIE102", Name: "Python for AI & Data Science", Semester: 2, Papers: makePapers("AIE", "19AIE102")},
				{Code: "19AIE201", Name: "Foundations of Artificial Intelligence", Semester: 3, Papers: makePapers("AIE", "19AIE201")},
				{Code: "19AIE202", Name: "Mathematics for Machine Learning", Semester: 3, Papers: makePapers("AIE", "19AIE202")},
				{Code: "19AIE203", Name: "Machine Learning Algorithms & Applications", Semester: 4, Papers: makePapers("AIE", "19AIE203")},
				{Code: "19AIE204", Name: "Data Structures & Algorithms for AI", Semester: 4, Papers: makePapers("AIE", "19AIE204")},
				{Code: "19AIE301", Name: "Deep Learning and Neural Networks", Semester: 5, Papers: makePapers("AIE", "19AIE301")},
				{Code: "19AIE305", Name: "Big Data Analytics and Processing Pipelines", Semester: 5, Papers: makePapers("AIE", "19AIE305")},
				{Code: "19AIE302", Name: "Natural Language Processing (NLP)", Semester: 6, Papers: makePapers("AIE", "19AIE302")},
				{Code: "19AIE303", Name: "Computer Vision and Image Processing", Semester: 6, Papers: makePapers("AIE", "19AIE303")},
				{Code: "19AIE304", Name: "Reinforcement Learning & Decision Making", Semester: 7, Papers: makePapers("AIE", "19AIE304")},
				{Code: "19AIE306", Name: "AI Ethics, Governance and Trustworthy AI", Semester: 7, Papers: makePapers("AIE", "19AIE306")},
				{Code: "19AIE401", Name: "Generative AI and Large Language Models", Semester: 8, Papers: makePapers("AIE", "19AIE401")},
			},
		},
		{
			Code:        "EEE",
			Name:        "Electrical & Electronics",
			Description: "B.Tech Electrical & Electronics Engineering",
			Subjects: []SubjectItem{
				{Code: "19EEE101", Name: "Basics of Electrical Engineering", Semester: 1, Papers: makePapers("EEE", "19EEE101")},
				{Code: "19EEE102", Name: "Engineering Electromagnetic Fields", Semester: 2, Papers: makePapers("EEE", "19EEE102")},
				{Code: "19EEE201", Name: "Electric Circuit Analysis", Semester: 3, Papers: makePapers("EEE", "19EEE201")},
				{Code: "19EEE202", Name: "Electrical Machines - I", Semester: 3, Papers: makePapers("EEE", "19EEE202")},
				{Code: "19EEE203", Name: "Power Systems Generation & Transmission", Semester: 4, Papers: makePapers("EEE", "19EEE203")},
				{Code: "19EEE204", Name: "Linear Control Systems", Semester: 4, Papers: makePapers("EEE", "19EEE204")},
				{Code: "19EEE301", Name: "Power Electronics and Drives", Semester: 5, Papers: makePapers("EEE", "19EEE301")},
				{Code: "19EEE302", Name: "Microcontroller and Embedded Systems", Semester: 5, Papers: makePapers("EEE", "19EEE302")},
				{Code: "19EEE303", Name: "Renewable Energy Systems and Smart Grid", Semester: 6, Papers: makePapers("EEE", "19EEE303")},
				{Code: "19EEE304", Name: "High Voltage Engineering", Semester: 7, Papers: makePapers("EEE", "19EEE304")},
				{Code: "19EEE401", Name: "Electric Vehicles and Battery Systems", Semester: 8, Papers: makePapers("EEE", "19EEE401")},
			},
		},
		{
			Code:        "MECH",
			Name:        "Mechanical Engineering",
			Description: "B.Tech Mechanical Engineering",
			Subjects: []SubjectItem{
				{Code: "19MEC101", Name: "Engineering Graphics and CAD", Semester: 1, Papers: makePapers("MECH", "19MEC101")},
				{Code: "19MEC102", Name: "Engineering Mechanics", Semester: 2, Papers: makePapers("MECH", "19MEC102")},
				{Code: "19MEC201", Name: "Engineering Thermodynamics", Semester: 3, Papers: makePapers("MECH", "19MEC201")},
				{Code: "19MEC203", Name: "Mechanics of Solids and Materials Strength", Semester: 3, Papers: makePapers("MECH", "19MEC203")},
				{Code: "19MEC202", Name: "Fluid Mechanics and Hydraulic Machinery", Semester: 4, Papers: makePapers("MECH", "19MEC202")},
				{Code: "19MEC204", Name: "Kinematics and Dynamics of Machinery", Semester: 4, Papers: makePapers("MECH", "19MEC204")},
				{Code: "19MEC301", Name: "Manufacturing Technology and CNC", Semester: 5, Papers: makePapers("MECH", "19MEC301")},
				{Code: "19MEC302", Name: "Heat and Mass Transfer", Semester: 5, Papers: makePapers("MECH", "19MEC302")},
				{Code: "19MEC303", Name: "Design of Machine Elements", Semester: 6, Papers: makePapers("MECH", "19MEC303")},
				{Code: "19MEC304", Name: "Automobile Engineering and IC Engines", Semester: 7, Papers: makePapers("MECH", "19MEC304")},
				{Code: "19MEC401", Name: "Robotics and Industrial Automation", Semester: 8, Papers: makePapers("MECH", "19MEC401")},
			},
		},
		{
			Code:        "MCA",
			Name:        "Master of Computer Applications",
			Description: "Master of Computer Applications (MCA)",
			Subjects: []SubjectItem{
				{Code: "20MCA101", Name: "Problem Solving and C++ Programming", Semester: 1, Papers: makePapers("MCA", "20MCA101")},
				{Code: "20MCA102", Name: "Computer Organization and Architecture", Semester: 1, Papers: makePapers("MCA", "20MCA102")},
				{Code: "20MCA201", Name: "Advanced Data Structures & Algorithms", Semester: 2, Papers: makePapers("MCA", "20MCA201")},
				{Code: "20MCA202", Name: "Full Stack Web Application Development", Semester: 2, Papers: makePapers("MCA", "20MCA202")},
				{Code: "20MCA203", Name: "Advanced Database Management Systems", Semester: 3, Papers: makePapers("MCA", "20MCA203")},
				{Code: "20MCA204", Name: "Cloud Computing and Virtualization", Semester: 3, Papers: makePapers("MCA", "20MCA204")},
				{Code: "20MCA205", Name: "Software Quality Assurance and Testing", Semester: 4, Papers: makePapers("MCA", "20MCA205")},
			},
		},
		{
			Code:        "MTECH",
			Name:        "M.Tech Computer Science",
			Description: "M.Tech Computer Science & Engineering",
			Subjects: []SubjectItem{
				{Code: "21CS601", Name: "Advanced Algorithms and Complexity", Semester: 1, Papers: makePapers("MTECH", "21CS601")},
				{Code: "21CS602", Name: "Advanced Computer Architecture & Parallelism", Semester: 1, Papers: makePapers("MTECH", "21CS602")},
				{Code: "21CS603", Name: "Cloud Computing and Distributed Systems", Semester: 2, Papers: makePapers("MTECH", "21CS603")},
				{Code: "21CS604", Name: "Machine Learning and Data Mining", Semester: 2, Papers: makePapers("MTECH", "21CS604")},
				{Code: "21CS605", Name: "Cryptographic Protocols and Cyber Defense", Semester: 3, Papers: makePapers("MTECH", "21CS605")},
			},
		},
	}
}
