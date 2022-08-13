package creature

import (
	"testing"

	"github.com/evertras/bubble-dnd/pkg/stats"
	"github.com/stretchr/testify/assert"
)

func TestNewCreaturesGetDifferentIDs(t *testing.T) {
	makeCreature := func() Creature {
		basicStats := stats.NewCollection(10, 10, 10, 10, 10, 10)
		return New('c', basicStats)
	}

	a := makeCreature()
	b := makeCreature()

	assert.NotEqual(t, a.ID(), b.ID())
}
