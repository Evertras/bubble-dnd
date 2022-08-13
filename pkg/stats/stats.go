package stats

import "fmt"

type Stat struct {
	kind     Kind
	base     int
	modifier int
}

func New(kind Kind, base int) Stat {
	if kind < Strength || kind > Charisma {
		// Should never happen, some code somewhere is being very bad so stop here
		panic(fmt.Sprintf("stat kind %d is out of bounds", kind))
	}

	var modifier int

	if base >= 10 {
		modifier = (base - 10) / 2
	} else {
		modifier = (base - 11) / 2
	}

	return Stat{
		kind:     kind,
		base:     base,
		modifier: modifier,
	}
}

func (s *Stat) Kind() Kind {
	return s.kind
}

func (s *Stat) Base() int {
	return s.base
}

func (s *Stat) Modifier() int {
	return s.modifier
}
