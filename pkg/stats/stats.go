package stats

import "fmt"

type Stat struct {
	kind Kind
	base int
}

func newStat(kind Kind, base int) Stat {
	if kind < Strength || kind > Charisma {
		// Should never happen, some code somewhere is being very bad so stop here
		panic(fmt.Sprintf("stat kind %d is out of bounds", kind))
	}

	return Stat{
		kind: kind,
		base: base,
	}
}

func (s Stat) Kind() Kind {
	return s.kind
}

func (s Stat) Base() int {
	return s.base
}

func (s Stat) Modifier() int {
	var modifier int

	if s.base >= 10 {
		modifier = (s.base - 10) / 2
	} else {
		modifier = (s.base - 11) / 2
	}

	return modifier
}
