package uisheet

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/evertras/bubble-dnd/pkg/stats"
)

func (m Model) View() string {
	statLine := func(s stats.Stat) string {
		var modifierStr string

		modifier := s.Modifier()

		if modifier >= 0 {
			modifierStr = fmt.Sprintf("+%d", modifier)
		} else {
			modifierStr = fmt.Sprintf("%d", modifier)
		}

		return fmt.Sprintf("%s: %d (%s)", s.Kind(), s.Base(), modifierStr)
	}

	statsCol := m.data.Stats()

	return lipgloss.JoinVertical(
		lipgloss.Left,
		statLine(statsCol.Strength()),
		statLine(statsCol.Dexterity()),
		statLine(statsCol.Constitution()),
		statLine(statsCol.Wisdom()),
		statLine(statsCol.Intelligence()),
		statLine(statsCol.Charisma()),
	)
}
