package menu

import (
	"testing"
)

func TestGetBranchCatalog(t *testing.T) {
	branches := GetBranchCatalog()
	if len(branches) == 0 {
		t.Fatalf("expected non-empty branch catalog")
	}

	branchMap := make(map[string]BranchItem)
	for _, b := range branches {
		branchMap[b.Code] = b
	}

	expectedBranches := []string{"CYS", "CSE", "ECE", "AIE", "EEE", "MECH", "MCA", "MTECH"}
	for _, expected := range expectedBranches {
		branch, found := branchMap[expected]
		if !found {
			t.Errorf("expected branch %s in catalog, but was missing", expected)
			continue
		}
		if len(branch.Subjects) == 0 {
			t.Errorf("expected subjects for branch %s, got 0", expected)
		}
		for _, subj := range branch.Subjects {
			if subj.Code == "" || subj.Name == "" {
				t.Errorf("subject in branch %s has empty code or name: %+v", expected, subj)
			}
			if subj.Semester <= 0 {
				t.Errorf("subject %s in branch %s has invalid semester: %d", subj.Code, expected, subj.Semester)
			}
			if len(subj.Papers) == 0 {
				t.Errorf("expected papers for subject %s, got 0", subj.Code)
			}
		}
	}
}

func TestFilterPapersByRange(t *testing.T) {
	papers := makePapers("CYS", "19CYS202")

	// Test 2025-2026 Range
	p2526 := FilterPapersByRange(papers, 2025, 2026)
	if len(p2526) == 0 {
		t.Errorf("expected papers in 2025-2026 range, got 0")
	}
	for _, p := range p2526 {
		if p.Year != "2025" && p.Year != "2026" {
			t.Errorf("unexpected year %s in 2025-2026 range", p.Year)
		}
	}

	// Test 2023-2024 Range
	p2324 := FilterPapersByRange(papers, 2023, 2024)
	if len(p2324) == 0 {
		t.Errorf("expected papers in 2023-2024 range, got 0")
	}
	for _, p := range p2324 {
		if p.Year != "2023" && p.Year != "2024" {
			t.Errorf("unexpected year %s in 2023-2024 range", p.Year)
		}
	}

	// Test 2021-2022 Range
	p2122 := FilterPapersByRange(papers, 2021, 2022)
	if len(p2122) == 0 {
		t.Errorf("expected papers in 2021-2022 range, got 0")
	}

	// Test Past 5 Years Range (2021 - 2026)
	p5Yr := FilterPapersByRange(papers, 2021, 2026)
	if len(p5Yr) < len(p2526)+len(p2324)+len(p2122) {
		t.Errorf("expected 5-year pack to contain sum of 2-year ranges, got %d", len(p5Yr))
	}
}

func TestFilterSubjectsBySemester(t *testing.T) {
	branches := GetBranchCatalog()
	var cys BranchItem
	for _, b := range branches {
		if b.Code == "CYS" {
			cys = b
			break
		}
	}

	// Filter semester 4
	sem4Subjects := FilterSubjectsBySemester(cys.Subjects, 4)
	if len(sem4Subjects) == 0 {
		t.Fatalf("expected subjects for CYS semester 4, got 0")
	}
	for _, s := range sem4Subjects {
		if s.Semester != 4 {
			t.Errorf("expected semester 4, got %d for subject %s", s.Semester, s.Code)
		}
	}

	// Filter all semesters (sem = 0)
	allSubjects := FilterSubjectsBySemester(cys.Subjects, 0)
	if len(allSubjects) != len(cys.Subjects) {
		t.Errorf("expected all %d subjects when sem=0, got %d", len(cys.Subjects), len(allSubjects))
	}

	// Verify Semester 3 CYS Subjects match official Amrita curriculum
	sem3Subjects := FilterSubjectsBySemester(cys.Subjects, 3)
	if len(sem3Subjects) < 6 {
		t.Errorf("expected at least 6 subjects for CYS semester 3, got %d", len(sem3Subjects))
	}

	expectedSem3Codes := map[string]string{
		"20CYS201": "Optimization Techniques",
		"20CYS202": "User Interface Design",
		"20CYS203": "Operating Systems",
		"20CYS204": "Database Management Systems",
		"20CYS205": "Modern Cryptography",
		"21MAT205": "Differential Equations and Numerical Methods",
	}

	sem3Map := make(map[string]string)
	for _, s := range sem3Subjects {
		sem3Map[s.Code] = s.Name
	}

	for code, expectedName := range expectedSem3Codes {
		name, found := sem3Map[code]
		if !found {
			t.Errorf("missing expected CYS Sem 3 subject %s (%s)", code, expectedName)
		} else if name != expectedName {
			t.Errorf("subject %s has name %q, expected %q", code, name, expectedName)
		}
	}

	for _, s := range sem3Subjects {
		t.Logf("CYS Sem 3: %s | %s (%d papers)", s.Code, s.Name, len(s.Papers))
	}
}

func TestSanitizeFileName(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"2023 End Semester Examination (Solved)", "2023_End_Semester_Examination_(Solved)"},
		{"Data Structures / Algorithms", "Data_Structures_-_Algorithms"},
		{"Course: 19CYS202 *Special? \"Quote\" <Tag> | Test", "Course-_19CYS202_Special_Quote_Tag_-_Test"},
	}

	for _, tc := range tests {
		actual := SanitizeFileName(tc.input)
		if actual != tc.expected {
			t.Errorf("SanitizeFileName(%q) = %q; expected %q", tc.input, actual, tc.expected)
		}
	}
}
