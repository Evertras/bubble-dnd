package character

import "github.com/evertras/bubble-dnd/pkg/stats"

type Character struct {
	name          string
	abilityScores stats.AbilityScoreBlock
}

func New(name string, statsCol stats.AbilityScoreBlock) Character {
	return Character{
		name:          name,
		abilityScores: statsCol,
	}
}

func (c *Character) Name() string {
	return c.name
}

func (c *Character) AbilityScoreBlock() stats.AbilityScoreBlock {
	return c.abilityScores
}
