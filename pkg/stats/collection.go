package stats

type Collection struct {
	strength     Stat
	dexterity    Stat
	constitution Stat
	wisdom       Stat
	intelligence Stat
	charisma     Stat
}

func BaseStats() Collection {
	return NewCollection(10, 10, 10, 10, 10, 10)
}

func NewCollection(strength, dexterity, constitution, wisdom, intelligence, charisma int) Collection {
	return Collection{
		strength:     newStat(Strength, strength),
		dexterity:    newStat(Dexterity, dexterity),
		constitution: newStat(Constitution, constitution),
		wisdom:       newStat(Wisdom, wisdom),
		intelligence: newStat(Intelligence, intelligence),
		charisma:     newStat(Charisma, charisma),
	}
}

func (c Collection) WithStrength(base int) Collection {
	c.strength = newStat(Strength, base)

	return c
}

func (c Collection) WithDexterity(base int) Collection {
	c.dexterity = newStat(Dexterity, base)

	return c
}

func (c Collection) WithConstitution(base int) Collection {
	c.constitution = newStat(Constitution, base)

	return c
}

func (c Collection) WithWisdom(base int) Collection {
	c.wisdom = newStat(Wisdom, base)

	return c
}

func (c Collection) WithIntelligence(base int) Collection {
	c.intelligence = newStat(Intelligence, base)

	return c
}

func (c Collection) WithCharisma(base int) Collection {
	c.charisma = newStat(Charisma, base)

	return c
}

func (c Collection) Strength() Stat {
	return c.strength
}

func (c Collection) Dexterity() Stat {
	return c.dexterity
}

func (c Collection) Constitution() Stat {
	return c.constitution
}

func (c Collection) Wisdom() Stat {
	return c.wisdom
}

func (c Collection) Intelligence() Stat {
	return c.intelligence
}

func (c Collection) Charisma() Stat {
	return c.charisma
}
