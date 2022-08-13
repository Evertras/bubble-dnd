package world

import tea "github.com/charmbracelet/bubbletea"

func (w World) Update(msg tea.Msg) World {
	switch w.state {
	case StateGeneral:
		break

	case StateBattle:
		break
	}

	return w
}

func (w *World) updateGeneral(msg tea.Msg) {
}
