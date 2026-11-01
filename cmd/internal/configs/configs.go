package configs

import (
	"github.com/charmbracelet/lipgloss"
)

const (
	APP_NAME        = "PRANAV PYQ FETCHER AMRITA"
	BASE_URL        = "http://dspace.amritanet.edu:8080"
	COURSE_URL      = BASE_URL + "/xmlui/handle/123456789/"
	COURSE_LIST_URL = COURSE_URL + "16"
)

// Styling for the logo and messages.
var (
	LogoStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#01FAC6")).
			Bold(true)

	SuccessStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#10B981")).
			Bold(true).
			Padding(0, 1).
			Margin(0, 0, 1, 0)

	WarningStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#F59E0B")).
			Bold(true).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#F59E0B")).
			Padding(1, 2).
			Margin(1, 0)

	ErrorStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("9")).
			Bold(true).
			Underline(true).
			Padding(0, 1).
			Margin(1, 0, 1, 0).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("1"))

	FetchStatusStyle = lipgloss.NewStyle().
				PaddingLeft(2).
				Foreground(lipgloss.Color("6")).
				Bold(true).
				Margin(1, 0)

	SelectionHistory []string
	IsDemoMode       bool
)
