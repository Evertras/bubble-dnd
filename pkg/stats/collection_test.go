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

	assertStats := func() {
		assert.Equal(t, expectedStrength, c.Strength().Base())
		assert.Equal(t, expectedDexterity, c.Dexterity().Base())
		assert.Equal(t, expectedConstitution, c.Constitution().Base())
		assert.Equal(t, expectedWisdom, c.Wisdom().Base())
		assert.Equal(t, expectedIntelligence, c.Intelligence().Base())
		assert.Equal(t, expectedCharisma, c.Charisma().Base())

		assert.Equal(t, Strength, c.Strength().Kind())
		assert.Equal(t, Dexterity, c.Dexterity().Kind())
		assert.Equal(t, Constitution, c.Constitution().Kind())
		assert.Equal(t, Wisdom, c.Wisdom().Kind())
		assert.Equal(t, Intelligence, c.Intelligence().Kind())
		assert.Equal(t, Charisma, c.Charisma().Kind())
	}

	assertStats()

	expectedStrength = expectedStrength + 1
	c = c.WithStrength(expectedStrength)
	assertStats()

	expectedDexterity = expectedDexterity + 1
	c = c.WithDexterity(expectedDexterity)
	assertStats()

	expectedConstitution = expectedConstitution + 1
	c = c.WithConstitution(expectedConstitution)
	assertStats()

	expectedWisdom = expectedWisdom + 1
	c = c.WithWisdom(expectedWisdom)
	assertStats()

	expectedIntelligence = expectedIntelligence + 1
	c = c.WithIntelligence(expectedIntelligence)
	assertStats()

	expectedCharisma = expectedCharisma + 1
	c = c.WithCharisma(expectedCharisma)
	assertStats()
}
