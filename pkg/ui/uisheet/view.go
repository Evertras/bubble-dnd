package uisheet

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/evertras/bubble-dnd/pkg/stats"
	"github.com/evertras/bubble-dnd/pkg/styles"
)

func viewStat(s stats.Stat) string {
	boldStyle := lipgloss.NewStyle().Bold(true)

	return styles.StyleStats[s.Kind()].Render(
		lipgloss.JoinVertical(
			lipgloss.Center,
			s.Kind().String(),
			fmt.Sprintf("%d", s.Base()),
			boldStyle.Render(s.Modifier().String()),
		),
	)
}

func (m Model) View() string {
	statsCol := m.data.Stats()

	return lipgloss.JoinVertical(
		lipgloss.Left,
		viewStat(statsCol.Strength()),
		viewStat(statsCol.Dexterity()),
		viewStat(statsCol.Constitution()),
		viewStat(statsCol.Wisdom()),
		viewStat(statsCol.Intelligence()),
		viewStat(statsCol.Charisma()),
	)
}
