package character

import "github.com/evertras/bubble-dnd/pkg/stats"

type Character struct {
	name      string
	statBlock stats.Block
}

func New(name string, statBlock stats.Block) Character {
	return Character{
		name:      name,
		statBlock: statBlock,
	}
}

func (c *Character) Name() string {
	return c.name
}

func (c *Character) StatBlock() stats.Block {
	return c.statBlock
}

func (c *Character) AbilityScoreBlock() stats.AbilityScoreBlock {
	return c.statBlock.AbilityScores()
}
