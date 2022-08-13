package creature

import (
	"testing"

	"github.com/evertras/bubble-dnd/pkg/stats"
	"github.com/stretchr/testify/assert"
)

func TestInitiativeTracksRoll(t *testing.T) {
	const (
		dexterityScore   = 14
		expectedModifier = 2
	)

	creature := New('c', stats.BaseAbilityScores().WithDexterity(dexterityScore))

	roll := creature.RollInitiative()

	assert.Equal(t, roll.Modifier(), expectedModifier)
}
