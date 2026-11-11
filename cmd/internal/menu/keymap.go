package menu

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/huh"
)

// NewMenuKeyMap returns a keymap configured so that pressing Esc or Ctrl+C
// aborts the current form, allowing the application to navigate back to the previous section.
func NewMenuKeyMap() *huh.KeyMap {
	km := huh.NewDefaultKeyMap()
	km.Quit = key.NewBinding(
		key.WithKeys("esc", "ctrl+c"),
		key.WithHelp("esc", "back"),
	)
	return km
}

// NewMenuForm creates a new huh.Form pre-configured with the Esc-to-back keymap.
func NewMenuForm(groups ...*huh.Group) *huh.Form {
	return huh.NewForm(groups...).WithKeyMap(NewMenuKeyMap())
}
