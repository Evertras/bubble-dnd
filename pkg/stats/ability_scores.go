package stats

type AbilityScores struct {
	strength     AbilityScore
	dexterity    AbilityScore
	constitution AbilityScore
	wisdom       AbilityScore
	intelligence AbilityScore
	charisma     AbilityScore
}

func BaseAbilityScores() AbilityScores {
	return NewAbilityScores(10, 10, 10, 10, 10, 10)
}

func NewAbilityScores(strength, dexterity, constitution, wisdom, intelligence, charisma int) AbilityScores {
	return AbilityScores{
		strength:     newAbilityScore(Strength, strength),
		dexterity:    newAbilityScore(Dexterity, dexterity),
		constitution: newAbilityScore(Constitution, constitution),
		wisdom:       newAbilityScore(Wisdom, wisdom),
		intelligence: newAbilityScore(Intelligence, intelligence),
		charisma:     newAbilityScore(Charisma, charisma),
	}
}

func (c AbilityScores) WithStrength(base int) AbilityScores {
	c.strength = newAbilityScore(Strength, base)

	return c
}

func (c AbilityScores) WithDexterity(base int) AbilityScores {
	c.dexterity = newAbilityScore(Dexterity, base)

	return c
}

func (c AbilityScores) WithConstitution(base int) AbilityScores {
	c.constitution = newAbilityScore(Constitution, base)

	return c
}

func (c AbilityScores) WithWisdom(base int) AbilityScores {
	c.wisdom = newAbilityScore(Wisdom, base)

	return c
}

func (c AbilityScores) WithIntelligence(base int) AbilityScores {
	c.intelligence = newAbilityScore(Intelligence, base)

	return c
}

func (c AbilityScores) WithCharisma(base int) AbilityScores {
	c.charisma = newAbilityScore(Charisma, base)

	return c
}

func (c AbilityScores) Strength() AbilityScore {
	return c.strength
}

func (c AbilityScores) Dexterity() AbilityScore {
	return c.dexterity
}

func (c AbilityScores) Constitution() AbilityScore {
	return c.constitution
}

func (c AbilityScores) Wisdom() AbilityScore {
	return c.wisdom
}

func (c AbilityScores) Intelligence() AbilityScore {
	return c.intelligence
}

func (c AbilityScores) Charisma() AbilityScore {
	return c.charisma
}
