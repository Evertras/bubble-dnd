package styles

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/evertras/bubble-dnd/pkg/stats"
)

var (
	StyleAbilityScores map[stats.Kind]lipgloss.Style

	ColorAbilityScores map[stats.Kind]lipgloss.AdaptiveColor
)

func init() {
	StyleAbilityScores = make(map[stats.Kind]lipgloss.Style)
	ColorAbilityScores = make(map[stats.Kind]lipgloss.AdaptiveColor)

	ColorAbilityScores[stats.Strength] = lipgloss.AdaptiveColor{
		Light: "#330000",
		Dark:  "#dd5555",
	}

	ColorAbilityScores[stats.Dexterity] = lipgloss.AdaptiveColor{
		Light: "#003300",
		Dark:  "#55cc88",
	}

	ColorAbilityScores[stats.Constitution] = lipgloss.AdaptiveColor{
		Light: "#000033",
		Dark:  "#5577dd",
	}

	ColorAbilityScores[stats.Wisdom] = lipgloss.AdaptiveColor{
		Light: "#330033",
		Dark:  "#aa88aa",
	}

	ColorAbilityScores[stats.Intelligence] = lipgloss.AdaptiveColor{
		Light: "#003333",
		Dark:  "#66aaaa",
	}

	ColorAbilityScores[stats.Charisma] = lipgloss.AdaptiveColor{
		Light: "#333300",
		Dark:  "#cccc88",
	}

	abilityScoreStyle := func(color lipgloss.AdaptiveColor) lipgloss.Style {
		return lipgloss.NewStyle().Foreground(color)
	}

	StyleAbilityScores[stats.Strength] = abilityScoreStyle(ColorAbilityScores[stats.Strength])
	StyleAbilityScores[stats.Dexterity] = abilityScoreStyle(ColorAbilityScores[stats.Dexterity])
	StyleAbilityScores[stats.Constitution] = abilityScoreStyle(ColorAbilityScores[stats.Constitution])
	StyleAbilityScores[stats.Wisdom] = abilityScoreStyle(ColorAbilityScores[stats.Wisdom])
	StyleAbilityScores[stats.Intelligence] = abilityScoreStyle(ColorAbilityScores[stats.Intelligence])
	StyleAbilityScores[stats.Charisma] = abilityScoreStyle(ColorAbilityScores[stats.Charisma])
}
