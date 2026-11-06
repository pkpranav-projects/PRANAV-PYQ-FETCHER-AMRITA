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
			if len(subj.Papers) == 0 {
				t.Errorf("expected papers for subject %s, got 0", subj.Code)
			}
		}
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
