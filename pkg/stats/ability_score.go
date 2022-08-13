package stats

import "fmt"

type AbilityScore struct {
	kind AbilityKind
	base int
}

func newAbilityScore(kind AbilityKind, base int) AbilityScore {
	if kind < Strength || kind > Charisma {
		// Should never happen, some code somewhere is being very bad so stop here
		panic(fmt.Sprintf("stat kind %d is out of bounds", kind))
	}

	return AbilityScore{
		kind: kind,
		base: base,
	}
}

func (s AbilityScore) Kind() AbilityKind {
	return s.kind
}

func (s AbilityScore) Base() int {
	return s.base
}

func (s AbilityScore) Modifier() Modifier {
	var modifier int

	if s.base >= 10 {
		modifier = (s.base - 10) / 2
	} else {
		modifier = (s.base - 11) / 2
	}

	return Modifier(modifier)
}
