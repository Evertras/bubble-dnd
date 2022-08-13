package creature

import (
	"sync/atomic"

	"github.com/evertras/bubble-dnd/pkg/stats"
)

type ID int64

var currentID int64 = 1

type Creature struct {
	id          ID
	stats       stats.Collection
	displayRune rune
}

func New(displayRune rune, s stats.Collection) Creature {
	id := atomic.AddInt64(&currentID, 1)
	return Creature{
		displayRune: displayRune,
		stats:       s,
		id:          ID(id),
	}
}

func (c Creature) ID() ID {
	return c.id
}

func (c Creature) Stats() stats.Collection {
	return c.stats
}
