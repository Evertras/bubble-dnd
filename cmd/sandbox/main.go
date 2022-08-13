package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/evertras/bubble-dnd/pkg/character"
	"github.com/evertras/bubble-dnd/pkg/stats"
	"github.com/evertras/bubble-dnd/pkg/ui/uisheet"
)

type Model struct {
	sheet uisheet.Model
}

func NewModel() Model {
	fighter := character.New(
		"Fighter McFightface",
		stats.NewAbilityScores(16, 13, 14, 11, 8, 9),
	)

	return Model{
		sheet: uisheet.New(fighter),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m Model) View() string {
	sheetView := m.sheet.View()

	return lipgloss.JoinVertical(
		lipgloss.Left,
		"Press 'q' or 'ctrl+c' to quit",
		sheetView,
		"\n",
	)
}

func main() {
	m := NewModel()

	p := tea.NewProgram(m)

	if err := p.Start(); err != nil {
		log.Fatal(err)
	}
}
