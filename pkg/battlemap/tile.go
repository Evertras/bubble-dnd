package battlemap

type Tile struct {
	terrain Terrain
}

func NewTile(terrain Terrain) Tile {
	return Tile{
		terrain: terrain,
	}
}

func (t Tile) Terrain() Terrain {
	return t.terrain
}
