package battlemap

type Battlemap struct {
	width  int
	height int
	data   [][]Tile
}

func NewOpen(width, height int) Battlemap {
	data := make([][]Tile, width)

	for x := 0; x < width; x++ {
		data[x] = make([]Tile, height)

		for y := 0; y < height; y++ {
			data[x][y] = NewTile(TerrainOpen)
		}
	}

	return Battlemap{
		width:  width,
		height: height,
		data:   data,
	}
}

func (b *Battlemap) Width() int {
	return b.width
}

func (b *Battlemap) Height() int {
	return b.height
}

func (b *Battlemap) TileAt(x, y int) Tile {
	return b.data[x][y]
}
