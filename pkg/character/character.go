package character

import "github.com/evertras/bubble-dnd/pkg/stats"

type Character struct {
	name          string
	abilityScores stats.AbilityScores
}

func New(name string, statsCol stats.AbilityScores) Character {
	return Character{
		name:          name,
		abilityScores: statsCol,
	}
}

func (c *Character) Name() string {
	return c.name
}

func (c *Character) AbilityScores() stats.AbilityScores {
	return c.abilityScores
}
