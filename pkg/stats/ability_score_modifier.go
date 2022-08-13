package stats

type AbilityModifier int

func (m AbilityModifier) String() string {
	return modifierString(m)
}
