package character

import "github.com/evertras/bubble-dnd/pkg/stats"

type Character struct {
	name  string
	stats stats.Collection
}

func New(name string, statsCol stats.Collection) Character {
	return Character{
		name:  name,
		stats: statsCol,
	}
}

func (c *Character) Name() string {
	return c.name
}

func (c *Character) Stats() stats.Collection {
	return c.stats
}
