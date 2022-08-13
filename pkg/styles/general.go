package styles

import "github.com/charmbracelet/lipgloss"

var (
	StyleSubtle lipgloss.Style
)

func init() {
	StyleSubtle = lipgloss.NewStyle().Faint(true)
}
