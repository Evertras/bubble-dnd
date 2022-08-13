package stats

import "fmt"

type Kind int

const (
	Strength Kind = iota
	Dexterity
	Constitution
	Wisdom
	Intelligence
	Charisma
)

func (k Kind) String() string {
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
