package stats

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAbilityScoreStrings(t *testing.T) {
	assert.Equal(t, "STR", Strength.String())
	assert.Equal(t, "DEX", Dexterity.String())
	assert.Equal(t, "CON", Constitution.String())
	assert.Equal(t, "WIS", Wisdom.String())
	assert.Equal(t, "INT", Intelligence.String())
	assert.Equal(t, "CHA", Charisma.String())
}
