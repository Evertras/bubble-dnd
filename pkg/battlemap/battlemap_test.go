package battlemap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewOpenBattlemap(t *testing.T) {
	const (
		width  = 10
		height = 15
	)

	testMap := NewOpen(width, height)

	assert.Equal(t, width, testMap.Width())
	assert.Equal(t, height, testMap.Height())

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			tile := testMap.TileAt(x, y)

			assert.Equal(t, TerrainOpen, tile.Terrain())
		}
	}
}
