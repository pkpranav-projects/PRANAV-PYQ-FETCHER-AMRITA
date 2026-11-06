package menu

import (
	"fmt"
	"strings"
)

// PaperItem represents an individual question paper file
type PaperItem struct {
	Name string
	Path string
	Year string
	Type string // EndSem, P1, P2
}

// SubjectItem represents an academic subject with code and question papers
type SubjectItem struct {
	Code   string
	Name   string
	Papers []PaperItem
}

// BranchItem represents an engineering branch or academic department
type BranchItem struct {
	Code        string
	Name        string
	Description string
	Subjects    []SubjectItem
}

func makePapers(branchCode, subjectCode string) []PaperItem {
	cleanSubject := strings.ReplaceAll(subjectCode, " ", "_")
	return []PaperItem{
		{
			Name: "2023 - End Semester Examination (Solved)",
			Path: fmt.Sprintf("/xmlui/bitstream/handle/123456789/%s_%s_2023_EndSem_Solved.pdf", branchCode, cleanSubject),
			Year: "2023",
			Type: "EndSem",
		},
		{
			Name: "2023 - Periodical Examination 2 (P2)",
			Path: fmt.Sprintf("/xmlui/bitstream/handle/123456789/%s_%s_2023_P2.pdf", branchCode, cleanSubject),
			Year: "2023",
			Type: "P2",
		},
		{
			Name: "2023 - Periodical Examination 1 (P1)",
			Path: fmt.Sprintf("/xmlui/bitstream/handle/123456789/%s_%s_2023_P1.pdf", branchCode, cleanSubject),
			Year: "2023",
			Type: "P1",
		},
		{
			Name: "2022 - End Semester Examination",
			Path: fmt.Sprintf("/xmlui/bitstream/handle/123456789/%s_%s_2022_EndSem.pdf", branchCode, cleanSubject),
			Year: "2022",
			Type: "EndSem",
		},
		{
			Name: "2022 - Periodical Examination 2 (P2)",
			Path: fmt.Sprintf("/xmlui/bitstream/handle/123456789/%s_%s_2022_P2.pdf", branchCode, cleanSubject),
			Year: "2022",
			Type: "P2",
		},
		{
			Name: "2022 - Periodical Examination 1 (P1)",
			Path: fmt.Sprintf("/xmlui/bitstream/handle/123456789/%s_%s_2022_P1.pdf", branchCode, cleanSubject),
			Year: "2022",
			Type: "P1",
		},
		{
			Name: "2021 - End Semester Examination",
			Path: fmt.Sprintf("/xmlui/bitstream/handle/123456789/%s_%s_2021_EndSem.pdf", branchCode, cleanSubject),
			Year: "2021",
			Type: "EndSem",
		},
		{
			Name: "2020 - End Semester Examination",
			Path: fmt.Sprintf("/xmlui/bitstream/handle/123456789/%s_%s_2020_EndSem.pdf", branchCode, cleanSubject),
			Year: "2020",
			Type: "EndSem",
		},
	}
}

// GetBranchCatalog returns the full curriculum mapping of branches, subjects, and papers
func GetBranchCatalog() []BranchItem {
	return []BranchItem{
		{
			Code:        "CYS",
			Name:        "Cyber Security",
			Description: "B.Tech Cyber Security",
			Subjects: []SubjectItem{
				{Code: "19CYS201", Name: "Foundations of Cyber Security", Papers: makePapers("CYS", "19CYS201")},
				{Code: "19CYS202", Name: "Cryptography and Network Security", Papers: makePapers("CYS", "19CYS202")},
				{Code: "19CYS203", Name: "Operating Systems and System Security", Papers: makePapers("CYS", "19CYS203")},
				{Code: "19CYS204", Name: "Data Structures and Algorithms", Papers: makePapers("CYS", "19CYS204")},
				{Code: "19CYS205", Name: "Computer Networks and Protocols", Papers: makePapers("CYS", "19CYS205")},
				{Code: "19CYS301", Name: "Secure Coding and Software Security", Papers: makePapers("CYS", "19CYS301")},
				{Code: "19CYS302", Name: "Digital Forensics and Incident Response", Papers: makePapers("CYS", "19CYS302")},
				{Code: "19CYS303", Name: "Penetration Testing and Ethical Hacking", Papers: makePapers("CYS", "19CYS303")},
				{Code: "19CYS304", Name: "Malware Analysis and Reverse Engineering", Papers: makePapers("CYS", "19CYS304")},
				{Code: "19CYS305", Name: "Database Systems and Database Security", Papers: makePapers("CYS", "19CYS305")},
				{Code: "19CYS306", Name: "Cloud and Virtualization Security", Papers: makePapers("CYS", "19CYS306")},
				{Code: "19CYS307", Name: "Web Application Security", Papers: makePapers("CYS", "19CYS307")},
			},
		},
		{
			Code:        "CSE",
			Name:        "Computer Science & Engineering",
			Description: "B.Tech Computer Science & Engineering",
			Subjects: []SubjectItem{
				{Code: "19CSE201", Name: "Data Structures and Algorithms", Papers: makePapers("CSE", "19CSE201")},
				{Code: "19CSE202", Name: "Operating Systems", Papers: makePapers("CSE", "19CSE202")},
				{Code: "19CSE203", Name: "Database Management Systems", Papers: makePapers("CSE", "19CSE203")},
				{Code: "19CSE204", Name: "Computer Organization and Architecture", Papers: makePapers("CSE", "19CSE204")},
				{Code: "19CSE205", Name: "Object Oriented Programming in Java/C++", Papers: makePapers("CSE", "19CSE205")},
				{Code: "19CSE301", Name: "Computer Networks", Papers: makePapers("CSE", "19CSE301")},
				{Code: "19CSE302", Name: "Theory of Computation & Automata", Papers: makePapers("CSE", "19CSE302")},
				{Code: "19CSE303", Name: "Design and Analysis of Algorithms", Papers: makePapers("CSE", "19CSE303")},
				{Code: "19CSE304", Name: "Compiler Design", Papers: makePapers("CSE", "19CSE304")},
				{Code: "19CSE305", Name: "Software Engineering & Agile Practices", Papers: makePapers("CSE", "19CSE305")},
				{Code: "19CSE306", Name: "Artificial Intelligence Foundations", Papers: makePapers("CSE", "19CSE306")},
				{Code: "19CSE307", Name: "Distributed Systems and Cloud", Papers: makePapers("CSE", "19CSE307")},
			},
		},
		{
			Code:        "ECE",
			Name:        "Electronics & Communication",
			Description: "B.Tech Electronics & Communication Engineering",
			Subjects: []SubjectItem{
				{Code: "19ECE201", Name: "Digital Signal Processing", Papers: makePapers("ECE", "19ECE201")},
				{Code: "19ECE202", Name: "Signals and Systems", Papers: makePapers("ECE", "19ECE202")},
				{Code: "19ECE203", Name: "Electronic Devices and Circuits", Papers: makePapers("ECE", "19ECE203")},
				{Code: "19ECE204", Name: "Digital System Design and Verilog HDL", Papers: makePapers("ECE", "19ECE204")},
				{Code: "19ECE205", Name: "Microprocessors and Microcontrollers", Papers: makePapers("ECE", "19ECE205")},
				{Code: "19ECE301", Name: "Analog Communication Systems", Papers: makePapers("ECE", "19ECE301")},
				{Code: "19ECE302", Name: "Digital Communication Systems", Papers: makePapers("ECE", "19ECE302")},
				{Code: "19ECE303", Name: "VLSI Design and Technology", Papers: makePapers("ECE", "19ECE303")},
				{Code: "19ECE304", Name: "Electromagnetic Waves and Transmission Lines", Papers: makePapers("ECE", "19ECE304")},
				{Code: "19ECE305", Name: "Control Systems Engineering", Papers: makePapers("ECE", "19ECE305")},
				{Code: "19ECE306", Name: "Embedded Systems and IoT Applications", Papers: makePapers("ECE", "19ECE306")},
			},
		},
		{
			Code:        "AIE",
			Name:        "Artificial Intelligence",
			Description: "B.Tech Artificial Intelligence & Data Science",
			Subjects: []SubjectItem{
				{Code: "19AIE201", Name: "Foundations of Artificial Intelligence", Papers: makePapers("AIE", "19AIE201")},
				{Code: "19AIE202", Name: "Mathematics for Machine Learning", Papers: makePapers("AIE", "19AIE202")},
				{Code: "19AIE203", Name: "Machine Learning Algorithms & Applications", Papers: makePapers("AIE", "19AIE203")},
				{Code: "19AIE204", Name: "Data Structures & Algorithms for AI", Papers: makePapers("AIE", "19AIE204")},
				{Code: "19AIE301", Name: "Deep Learning and Neural Networks", Papers: makePapers("AIE", "19AIE301")},
				{Code: "19AIE302", Name: "Natural Language Processing (NLP)", Papers: makePapers("AIE", "19AIE302")},
				{Code: "19AIE303", Name: "Computer Vision and Image Processing", Papers: makePapers("AIE", "19AIE303")},
				{Code: "19AIE304", Name: "Reinforcement Learning & Decision Making", Papers: makePapers("AIE", "19AIE304")},
				{Code: "19AIE305", Name: "Big Data Analytics and Processing Pipelines", Papers: makePapers("AIE", "19AIE305")},
				{Code: "19AIE306", Name: "AI Ethics, Governance and Trustworthy AI", Papers: makePapers("AIE", "19AIE306")},
			},
		},
		{
			Code:        "EEE",
			Name:        "Electrical & Electronics",
			Description: "B.Tech Electrical & Electronics Engineering",
			Subjects: []SubjectItem{
				{Code: "19EEE201", Name: "Electric Circuit Analysis", Papers: makePapers("EEE", "19EEE201")},
				{Code: "19EEE202", Name: "Electrical Machines - I & II", Papers: makePapers("EEE", "19EEE202")},
				{Code: "19EEE203", Name: "Power Systems Generation & Transmission", Papers: makePapers("EEE", "19EEE203")},
				{Code: "19EEE204", Name: "Linear Control Systems", Papers: makePapers("EEE", "19EEE204")},
				{Code: "19EEE301", Name: "Power Electronics and Drives", Papers: makePapers("EEE", "19EEE301")},
				{Code: "19EEE302", Name: "Microcontroller and Embedded Systems", Papers: makePapers("EEE", "19EEE302")},
				{Code: "19EEE303", Name: "Renewable Energy Systems and Smart Grid", Papers: makePapers("EEE", "19EEE303")},
				{Code: "19EEE304", Name: "High Voltage Engineering", Papers: makePapers("EEE", "19EEE304")},
			},
		},
		{
			Code:        "MECH",
			Name:        "Mechanical Engineering",
			Description: "B.Tech Mechanical Engineering",
			Subjects: []SubjectItem{
				{Code: "19MEC201", Name: "Engineering Thermodynamics", Papers: makePapers("MECH", "19MEC201")},
				{Code: "19MEC202", Name: "Fluid Mechanics and Hydraulic Machinery", Papers: makePapers("MECH", "19MEC202")},
				{Code: "19MEC203", Name: "Mechanics of Solids and Materials Strength", Papers: makePapers("MECH", "19MEC203")},
				{Code: "19MEC204", Name: "Kinematics and Dynamics of Machinery", Papers: makePapers("MECH", "19MEC204")},
				{Code: "19MEC301", Name: "Manufacturing Technology and CNC", Papers: makePapers("MECH", "19MEC301")},
				{Code: "19MEC302", Name: "Heat and Mass Transfer", Papers: makePapers("MECH", "19MEC302")},
				{Code: "19MEC303", Name: "Design of Machine Elements", Papers: makePapers("MECH", "19MEC303")},
				{Code: "19MEC304", Name: "Automobile Engineering and IC Engines", Papers: makePapers("MECH", "19MEC304")},
			},
		},
		{
			Code:        "MCA",
			Name:        "Master of Computer Applications",
			Description: "Master of Computer Applications (MCA)",
			Subjects: []SubjectItem{
				{Code: "20MCA201", Name: "Advanced Data Structures & Algorithms", Papers: makePapers("MCA", "20MCA201")},
				{Code: "20MCA202", Name: "Full Stack Web Application Development", Papers: makePapers("MCA", "20MCA202")},
				{Code: "20MCA203", Name: "Advanced Database Management Systems", Papers: makePapers("MCA", "20MCA203")},
				{Code: "20MCA204", Name: "Cloud Computing and Virtualization", Papers: makePapers("MCA", "20MCA204")},
				{Code: "20MCA205", Name: "Software Quality Assurance and Testing", Papers: makePapers("MCA", "20MCA205")},
			},
		},
		{
			Code:        "MTECH",
			Name:        "M.Tech Computer Science",
			Description: "M.Tech Computer Science & Engineering",
			Subjects: []SubjectItem{
				{Code: "21CS601", Name: "Advanced Algorithms and Complexity", Papers: makePapers("MTECH", "21CS601")},
				{Code: "21CS602", Name: "Advanced Computer Architecture & Parallelism", Papers: makePapers("MTECH", "21CS602")},
				{Code: "21CS603", Name: "Cloud Computing and Distributed Systems", Papers: makePapers("MTECH", "21CS603")},
				{Code: "21CS604", Name: "Machine Learning and Data Mining", Papers: makePapers("MTECH", "21CS604")},
				{Code: "21CS605", Name: "Cryptographic Protocols and Cyber Defense", Papers: makePapers("MTECH", "21CS605")},
			},
		},
	}
}
