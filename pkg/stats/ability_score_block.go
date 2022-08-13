package stats

type AbilityScoreBlock struct {
	strength     AbilityScore
	dexterity    AbilityScore
	constitution AbilityScore
	wisdom       AbilityScore
	intelligence AbilityScore
	charisma     AbilityScore
}

func BaseAbilityScoreBlock() AbilityScoreBlock {
	return NewAbilityScoreBlock(10, 10, 10, 10, 10, 10)
}

func NewAbilityScoreBlock(strength, dexterity, constitution, wisdom, intelligence, charisma int) AbilityScoreBlock {
	return AbilityScoreBlock{
		strength:     newAbilityScore(Strength, strength),
		dexterity:    newAbilityScore(Dexterity, dexterity),
		constitution: newAbilityScore(Constitution, constitution),
		wisdom:       newAbilityScore(Wisdom, wisdom),
		intelligence: newAbilityScore(Intelligence, intelligence),
		charisma:     newAbilityScore(Charisma, charisma),
	}
}

func (c AbilityScoreBlock) WithStrength(base int) AbilityScoreBlock {
	c.strength = newAbilityScore(Strength, base)

	return c
}

func (c AbilityScoreBlock) WithDexterity(base int) AbilityScoreBlock {
	c.dexterity = newAbilityScore(Dexterity, base)

	return c
}

func (c AbilityScoreBlock) WithConstitution(base int) AbilityScoreBlock {
	c.constitution = newAbilityScore(Constitution, base)

	return c
}

func (c AbilityScoreBlock) WithWisdom(base int) AbilityScoreBlock {
	c.wisdom = newAbilityScore(Wisdom, base)

	return c
}

func (c AbilityScoreBlock) WithIntelligence(base int) AbilityScoreBlock {
	c.intelligence = newAbilityScore(Intelligence, base)

	return c
}

func (c AbilityScoreBlock) WithCharisma(base int) AbilityScoreBlock {
	c.charisma = newAbilityScore(Charisma, base)

	return c
}

func (c AbilityScoreBlock) Strength() AbilityScore {
	return c.strength
}

func (c AbilityScoreBlock) Dexterity() AbilityScore {
	return c.dexterity
}

func (c AbilityScoreBlock) Constitution() AbilityScore {
	return c.constitution
}

func (c AbilityScoreBlock) Wisdom() AbilityScore {
	return c.wisdom
}

func (c AbilityScoreBlock) Intelligence() AbilityScore {
	return c.intelligence
}

func (c AbilityScoreBlock) Charisma() AbilityScore {
	return c.charisma
}
