package uisheet

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/evertras/bubble-dnd/pkg/creature"
)

type Model struct {
	data creature.Creature
}

func New(data creature.Creature) Model {
	return Model{
		data: data,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}
