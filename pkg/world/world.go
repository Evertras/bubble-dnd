package world

import (
	"github.com/evertras/bubble-dnd/pkg/battlemap"
	"github.com/evertras/bubble-dnd/pkg/creature"
)

type World struct {
	creatures map[creature.ID]creature.Creature

	state State

	battlemap battlemap.Battlemap
}

func New() World {
	return World{
		creatures: make(map[creature.ID]creature.Creature),
	}
}

func (w World) WithCreatures(creatures []creature.Creature) World {
	m := make(map[creature.ID]creature.Creature)

	for _, c := range creatures {
		m[c.ID()] = c
	}

	w.creatures = m

	return w
}
