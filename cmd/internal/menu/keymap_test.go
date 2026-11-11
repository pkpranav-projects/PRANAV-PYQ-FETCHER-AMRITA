package menu

import (
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

func TestEscKeyTakesBackToPreviousSection(t *testing.T) {
	var val string
	selectField := huh.NewSelect[string]().
		Title("Select Test").
		Options(
			huh.NewOption("Semester 1", "sem1"),
			huh.NewOption("Semester 2", "sem2"),
		).
		Value(&val)

	form := NewMenuForm(huh.NewGroup(selectField))

	// Simulate Esc key press
	m, _ := form.Update(tea.KeyMsg{Type: tea.KeyEsc})
	updatedForm := m.(*huh.Form)

	if updatedForm.State != huh.StateAborted {
		t.Fatalf("expected form to be in StateAborted after pressing Esc, got %v", updatedForm.State)
	}
}
