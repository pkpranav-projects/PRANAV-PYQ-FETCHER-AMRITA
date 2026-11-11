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

// makePapers generates past 5+ years of question papers covering both Midterm and End Semester exams
func makePapers(branchCode, subjectCode string) []PaperItem {
	cleanSubject := strings.ReplaceAll(subjectCode, " ", "_")
	return []PaperItem{
		// 2025 - 2026 Range
		{
			Name: fmt.Sprintf("[%s] 2026 - Mid-Term Examination (P1)", subjectCode),
			Path: fmt.Sprintf("/xmlui/bitstream/handle/123456789/%s_%s_2026_MidTerm_P1.pdf", branchCode, cleanSubject),
			Year: "2026",
			Type: "Midterm",
			Exam: "P1",
		},
		{
			Name: fmt.Sprintf("[%s] 2025 - End Semester Examination (Regular)", subjectCode),
			Path: fmt.Sprintf("/xmlui/bitstream/handle/123456789/%s_%s_2025_EndSem_Regular.pdf", branchCode, cleanSubject),
			Year: "2025",
			Type: "EndSem",
			Exam: "Regular",
		},
		{
			Name: fmt.Sprintf("[%s] 2025 - Periodical Examination 2 (Mid Term 2)", subjectCode),
			Path: fmt.Sprintf("/xmlui/bitstream/handle/123456789/%s_%s_2025_P2.pdf", branchCode, cleanSubject),
			Year: "2025",
			Type: "Midterm",
			Exam: "P2",
		},
		{
			Name: fmt.Sprintf("[%s] 2025 - Periodical Examination 1 (Mid Term 1)", subjectCode),
			Path: fmt.Sprintf("/xmlui/bitstream/handle/123456789/%s_%s_2025_P1.pdf", branchCode, cleanSubject),
			Year: "2025",
			Type: "Midterm",
			Exam: "P1",
		},

		// 2023 - 2024 Range
		{
			Name: fmt.Sprintf("[%s] 2024 - End Semester Examination (Regular)", subjectCode),
			Path: fmt.Sprintf("/xmlui/bitstream/handle/123456789/%s_%s_2024_EndSem.pdf", branchCode, cleanSubject),
			Year: "2024",
			Type: "EndSem",
			Exam: "Regular",
		},
		{
			Name: fmt.Sprintf("[%s] 2024 - Periodical Examination 2 (Mid Term 2)", subjectCode),
			Path: fmt.Sprintf("/xmlui/bitstream/handle/123456789/%s_%s_2024_P2.pdf", branchCode, cleanSubject),
			Year: "2024",
			Type: "Midterm",
			Exam: "P2",
		},
		{
			Name: fmt.Sprintf("[%s] 2024 - Periodical Examination 1 (Mid Term 1)", subjectCode),
			Path: fmt.Sprintf("/xmlui/bitstream/handle/123456789/%s_%s_2024_P1.pdf", branchCode, cleanSubject),
			Year: "2024",
			Type: "Midterm",
			Exam: "P1",
		},
		{
			Name: fmt.Sprintf("[%s] 2023 - End Semester Examination (Solved)", subjectCode),
			Path: fmt.Sprintf("/xmlui/bitstream/handle/123456789/%s_%s_2023_EndSem_Solved.pdf", branchCode, cleanSubject),
			Year: "2023",
			Type: "EndSem",
			Exam: "Solved",
		},
		{
			Name: fmt.Sprintf("[%s] 2023 - Periodical Examination 2 (Mid Term 2)", subjectCode),
			Path: fmt.Sprintf("/xmlui/bitstream/handle/123456789/%s_%s_2023_P2.pdf", branchCode, cleanSubject),
			Year: "2023",
			Type: "Midterm",
			Exam: "P2",
		},
		{
			Name: fmt.Sprintf("[%s] 2023 - Periodical Examination 1 (Mid Term 1)", subjectCode),
			Path: fmt.Sprintf("/xmlui/bitstream/handle/123456789/%s_%s_2023_P1.pdf", branchCode, cleanSubject),
			Year: "2023",
			Type: "Midterm",
			Exam: "P1",
		},

		// 2021 - 2022 Range
		{
			Name: fmt.Sprintf("[%s] 2022 - End Semester Examination (Regular)", subjectCode),
			Path: fmt.Sprintf("/xmlui/bitstream/handle/123456789/%s_%s_2022_EndSem.pdf", branchCode, cleanSubject),
			Year: "2022",
			Type: "EndSem",
			Exam: "Regular",
		},
		{
			Name: fmt.Sprintf("[%s] 2022 - Periodical Examination 2 (Mid Term 2)", subjectCode),
			Path: fmt.Sprintf("/xmlui/bitstream/handle/123456789/%s_%s_2022_P2.pdf", branchCode, cleanSubject),
			Year: "2022",
			Type: "Midterm",
			Exam: "P2",
		},
		{
			Name: fmt.Sprintf("[%s] 2022 - Periodical Examination 1 (Mid Term 1)", subjectCode),
			Path: fmt.Sprintf("/xmlui/bitstream/handle/123456789/%s_%s_2022_P1.pdf", branchCode, cleanSubject),
			Year: "2022",
			Type: "Midterm",
			Exam: "P1",
		},
		{
			Name: fmt.Sprintf("[%s] 2021 - End Semester Examination (Regular)", subjectCode),
			Path: fmt.Sprintf("/xmlui/bitstream/handle/123456789/%s_%s_2021_EndSem.pdf", branchCode, cleanSubject),
			Year: "2021",
			Type: "EndSem",
			Exam: "Regular",
		},
		{
			Name: fmt.Sprintf("[%s] 2021 - Periodical Examination 1 (Mid Term 1)", subjectCode),
			Path: fmt.Sprintf("/xmlui/bitstream/handle/123456789/%s_%s_2021_P1.pdf", branchCode, cleanSubject),
			Year: "2021",
			Type: "Midterm",
			Exam: "P1",
		},

		// 2019 - 2020 Range
		{
			Name: fmt.Sprintf("[%s] 2020 - End Semester Examination (Regular)", subjectCode),
			Path: fmt.Sprintf("/xmlui/bitstream/handle/123456789/%s_%s_2020_EndSem.pdf", branchCode, cleanSubject),
			Year: "2020",
			Type: "EndSem",
			Exam: "Regular",
		},
		{
			Name: fmt.Sprintf("[%s] 2019 - End Semester Examination (Regular)", subjectCode),
			Path: fmt.Sprintf("/xmlui/bitstream/handle/123456789/%s_%s_2019_EndSem.pdf", branchCode, cleanSubject),
			Year: "2019",
			Type: "EndSem",
			Exam: "Regular",
		},
		{
			Name: fmt.Sprintf("[%s] 2019 - Periodical Examination 2 (Mid Term 2)", subjectCode),
			Path: fmt.Sprintf("/xmlui/bitstream/handle/123456789/%s_%s_2019_P2.pdf", branchCode, cleanSubject),
			Year: "2019",
			Type: "Midterm",
			Exam: "P2",
		},
		{
			Name: fmt.Sprintf("[%s] 2019 - Periodical Examination 1 (Mid Term 1)", subjectCode),
			Path: fmt.Sprintf("/xmlui/bitstream/handle/123456789/%s_%s_2019_P1.pdf", branchCode, cleanSubject),
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
			Description: "B.Tech Cyber Security",
			Subjects: []SubjectItem{
				{Code: "19CYS101", Name: "Introduction to Cybersecurity & Computing", Semester: 1, Papers: makePapers("CYS", "19CYS101")},
				{Code: "19CYS102", Name: "Data Structures & Algorithms Fundamentals", Semester: 2, Papers: makePapers("CYS", "19CYS102")},
				{Code: "19CYS201", Name: "Foundations of Cyber Security", Semester: 3, Papers: makePapers("CYS", "19CYS201")},
				{Code: "19CYS204", Name: "Data Structures and Algorithms", Semester: 3, Papers: makePapers("CYS", "19CYS204")},
				{Code: "19CYS202", Name: "Cryptography and Network Security", Semester: 4, Papers: makePapers("CYS", "19CYS202")},
				{Code: "19CYS203", Name: "Operating Systems and System Security", Semester: 4, Papers: makePapers("CYS", "19CYS203")},
				{Code: "19CYS205", Name: "Computer Networks and Protocols", Semester: 4, Papers: makePapers("CYS", "19CYS205")},
				{Code: "19CYS301", Name: "Secure Coding and Software Security", Semester: 5, Papers: makePapers("CYS", "19CYS301")},
				{Code: "19CYS305", Name: "Database Systems and Database Security", Semester: 5, Papers: makePapers("CYS", "19CYS305")},
				{Code: "19CYS303", Name: "Penetration Testing and Ethical Hacking", Semester: 6, Papers: makePapers("CYS", "19CYS303")},
				{Code: "19CYS306", Name: "Cloud and Virtualization Security", Semester: 6, Papers: makePapers("CYS", "19CYS306")},
				{Code: "19CYS302", Name: "Digital Forensics and Incident Response", Semester: 7, Papers: makePapers("CYS", "19CYS302")},
				{Code: "19CYS304", Name: "Malware Analysis and Reverse Engineering", Semester: 7, Papers: makePapers("CYS", "19CYS304")},
				{Code: "19CYS307", Name: "Web Application Security", Semester: 7, Papers: makePapers("CYS", "19CYS307")},
				{Code: "19CYS401", Name: "Cyber Threat Intelligence & Incident Response", Semester: 8, Papers: makePapers("CYS", "19CYS401")},
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
