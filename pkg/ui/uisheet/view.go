package uisheet

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/evertras/bubble-dnd/pkg/stats"
	"github.com/evertras/bubble-dnd/pkg/styles"
)

var (
	styleBoxed      = lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).Padding(0, 1)
	styleUnderlined = lipgloss.NewStyle().Underline(true)
)

func viewStat(s stats.Stat) string {
	statStyle := styleBoxed.Copy().Inherit(styles.StyleStats[s.Kind()])

	return statStyle.Render(
		lipgloss.JoinVertical(
			0.49,
			s.Kind().String(),
			fmt.Sprintf("%d", s.Base()),
			s.Modifier().String(),
		),
	)
}

func (m Model) viewStatBlock() string {
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

func (m Model) viewNameBlock() string {
	block := lipgloss.JoinHorizontal(
		lipgloss.Top,
		styles.StyleSubtle.Render("Name: "),
		styleUnderlined.Render(m.data.Name()),
	)

	return styleBoxed.Render(block)
}

func (m Model) View() string {
	nameBlock := m.viewNameBlock()
	statBlock := m.viewStatBlock()

	return lipgloss.JoinVertical(
		lipgloss.Left,
		nameBlock,
		statBlock,
	)
}
