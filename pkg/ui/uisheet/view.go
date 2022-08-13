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

func viewAbilityScore(s stats.AbilityScore) string {
	statStyle := styleBoxed.Copy().Inherit(styles.StyleAbilityScores[s.Kind()])

	return statStyle.Render(
		lipgloss.JoinVertical(
			0.49,
			s.Kind().String(),
			fmt.Sprintf("%d", s.Base()),
			s.Modifier().String(),
		),
	)
}

func (m Model) viewAbilityScoreBlock() string {
	statsCol := m.data.AbilityScores()

	return lipgloss.JoinVertical(
		lipgloss.Left,
		viewAbilityScore(statsCol.Strength()),
		viewAbilityScore(statsCol.Dexterity()),
		viewAbilityScore(statsCol.Constitution()),
		viewAbilityScore(statsCol.Wisdom()),
		viewAbilityScore(statsCol.Intelligence()),
		viewAbilityScore(statsCol.Charisma()),
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
	statBlock := m.viewAbilityScoreBlock()

	return lipgloss.JoinVertical(
		lipgloss.Left,
		nameBlock,
		statBlock,
	)
}
