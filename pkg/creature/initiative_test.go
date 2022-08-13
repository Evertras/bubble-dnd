package creature

import (
	"testing"

	"github.com/evertras/bubble-dnd/pkg/stats"
	"github.com/stretchr/testify/assert"
)

func TestInitiativeTracksRoll(t *testing.T) {
	const baseScore = 10
	const dexterityScore = 14
	const expectedModifier = 2

	creature := New('c', stats.NewCollection(
		baseScore,
		dexterityScore,
		baseScore,
		baseScore,
		baseScore,
		baseScore,
	))

	roll := creature.RollInitiative()

	assert.Equal(t, roll.Modifier(), expectedModifier)
}
