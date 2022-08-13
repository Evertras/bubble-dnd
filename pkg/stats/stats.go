package stats

import "fmt"

type Modifier int

func (m Modifier) String() string {
	if m >= 0 {
		return fmt.Sprintf("+%d", m)
	}

	return fmt.Sprintf("%d", m)
}

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

func (s Stat) Modifier() Modifier {
	var modifier int

	if s.base >= 10 {
		modifier = (s.base - 10) / 2
	} else {
		modifier = (s.base - 11) / 2
	}

	return Modifier(modifier)
}
