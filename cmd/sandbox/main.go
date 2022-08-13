package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/evertras/bubble-dnd/pkg/creature"
	"github.com/evertras/bubble-dnd/pkg/stats"
	"github.com/evertras/bubble-dnd/pkg/ui/uisheet"
)

type Model struct {
	sheet uisheet.Model
}

func NewModel() Model {
	fighter := creature.New('F', stats.NewCollection(16, 13, 14, 11, 8, 9))
	//goblin := creature.New('G', stats.NewCollection(9, 13, 10, 10, 8, 7))

	//sandboxWorld := world.New().WithCreatures([]creature.Creature{fighter, goblin})

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
