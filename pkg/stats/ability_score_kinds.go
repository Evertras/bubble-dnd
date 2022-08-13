package stats

import "fmt"

type AbilityKind int

const (
	Strength AbilityKind = iota
	Dexterity
	Constitution
	Wisdom
	Intelligence
	Charisma
)

func (k AbilityKind) String() string {
	switch k {
	case Strength:
		return "STR"

	case Dexterity:
		return "DEX"

	case Constitution:
		return "CON"

	case Wisdom:
		return "WIS"

	case Intelligence:
		return "INT"

	case Charisma:
		return "CHA"
	}

	panic(fmt.Sprintf("stat out of bounds: %d", k))
}
