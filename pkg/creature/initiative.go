package creature

import "github.com/evertras/bubble-dnd/pkg/dice"

type InitiativeResult struct {
	modifier int
	roll     int
	total    int
}

func (r *InitiativeResult) Total() int {
	return r.total
}

func (r *InitiativeResult) Modifier() int {
	return r.modifier
}

func (r *InitiativeResult) Roll() int {
	return r.roll
}

func (c *Creature) RollInitiative() InitiativeResult {
	roll := dice.RollD20()
	modifier := c.stats.Dexterity().Modifier()

	return InitiativeResult{
		roll:     roll,
		modifier: modifier,
		total:    roll + modifier,
	}
}
