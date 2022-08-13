package uisheet

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/evertras/bubble-dnd/pkg/character"
)

type Model struct {
	data character.Character
}

func New(data character.Character) Model {
	return Model{
		data: data,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}
