package stats

type Collection struct {
	strength     Stat
	dexterity    Stat
	constitution Stat
	wisdom       Stat
	intelligence Stat
	charisma     Stat
}

func NewCollection(strength, dexterity, constitution, wisdom, intelligence, charisma int) Collection {
	return Collection {
		strength: newStat(Strength, strength),
		dexterity: newStat(Dexterity, dexterity),
		constitution: newStat(Constitution, constitution),
		wisdom: newStat(Wisdom, wisdom),
		intelligence: newStat(Intelligence, intelligence),
		charisma: newStat(Charisma, charisma),
	}
}

func (c *Collection) Strength() Stat {
	return c.strength
}

func (c *Collection) Dexterity() Stat {
	return c.dexterity
}

func (c *Collection) Constitution() Stat {
	return c.constitution
}

func (c *Collection) Wisdom() Stat {
	return c.wisdom
}

func (c *Collection) Intelligence() Stat {
	return c.intelligence
}

func (c *Collection) Charisma() Stat {
	return c.charisma
}
