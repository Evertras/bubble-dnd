package uisheet

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/evertras/bubble-dnd/pkg/creature"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case creature.Creature:
		m.data = msg
	}

	return m, nil
}
