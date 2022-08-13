package stats

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCollectionTracksCorrectStats(t *testing.T) {
	expectedStrength := 9
	expectedDexterity := 8
	expectedConstitution := 13
	expectedWisdom := 16
	expectedIntelligence := 14
	expectedCharisma := 11

	c := NewCollection(
		expectedStrength,
		expectedDexterity,
		expectedConstitution,
		expectedWisdom,
		expectedIntelligence,
		expectedCharisma,
	)

	assert.Equal(t, expectedStrength, c.Strength().Base())
	assert.Equal(t, expectedDexterity, c.Dexterity().Base())
	assert.Equal(t, expectedConstitution, c.Constitution().Base())
	assert.Equal(t, expectedWisdom, c.Wisdom().Base())
	assert.Equal(t, expectedIntelligence, c.Intelligence().Base())
	assert.Equal(t, expectedCharisma, c.Charisma().Base())
}
