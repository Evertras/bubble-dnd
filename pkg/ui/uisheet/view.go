package uisheet

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/evertras/bubble-dnd/pkg/stats"
	"github.com/evertras/bubble-dnd/pkg/styles"
)

func viewStat(s stats.Stat) string {
	statStyle := styles.StyleStats[s.Kind()].Copy().Border(lipgloss.RoundedBorder()).Padding(0, 1)

	return statStyle.Render(
		lipgloss.JoinVertical(
			lipgloss.Center,
			s.Kind().String(),
			fmt.Sprintf("%d", s.Base()),
			s.Modifier().String(),
		),
	)
}

func (m Model) View() string {
	statsCol := m.data.Stats()

	statBlock := lipgloss.JoinVertical(
		lipgloss.Left,
		viewStat(statsCol.Strength()),
		viewStat(statsCol.Dexterity()),
		viewStat(statsCol.Constitution()),
		viewStat(statsCol.Wisdom()),
		viewStat(statsCol.Intelligence()),
		viewStat(statsCol.Charisma()),
	)

	return statBlock
}
