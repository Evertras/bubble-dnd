package styles

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/evertras/bubble-dnd/pkg/stats"
)

var (
	StyleStats map[stats.Kind]lipgloss.Style

	ColorStats map[stats.Kind]lipgloss.AdaptiveColor
)

func init() {
	StyleStats = make(map[stats.Kind]lipgloss.Style)
	ColorStats = make(map[stats.Kind]lipgloss.AdaptiveColor)

	ColorStats[stats.Strength] = lipgloss.AdaptiveColor{
		Light: "#330000",
		Dark:  "#dd5555",
	}

	ColorStats[stats.Dexterity] = lipgloss.AdaptiveColor{
		Light: "#003300",
		Dark:  "#55cc88",
	}

	ColorStats[stats.Constitution] = lipgloss.AdaptiveColor{
		Light: "#000033",
		Dark:  "#5577dd",
	}

	ColorStats[stats.Wisdom] = lipgloss.AdaptiveColor{
		Light: "#330033",
		Dark:  "#aa88aa",
	}

	ColorStats[stats.Intelligence] = lipgloss.AdaptiveColor{
		Light: "#003333",
		Dark:  "#66aaaa",
	}

	ColorStats[stats.Charisma] = lipgloss.AdaptiveColor{
		Light: "#333300",
		Dark:  "#cccc88",
	}

	statStyle := func(color lipgloss.AdaptiveColor) lipgloss.Style {
		return lipgloss.NewStyle().Foreground(color)
	}

	StyleStats[stats.Strength] = statStyle(ColorStats[stats.Strength])
	StyleStats[stats.Dexterity] = statStyle(ColorStats[stats.Dexterity])
	StyleStats[stats.Constitution] = statStyle(ColorStats[stats.Constitution])
	StyleStats[stats.Wisdom] = statStyle(ColorStats[stats.Wisdom])
	StyleStats[stats.Intelligence] = statStyle(ColorStats[stats.Intelligence])
	StyleStats[stats.Charisma] = statStyle(ColorStats[stats.Charisma])
}
