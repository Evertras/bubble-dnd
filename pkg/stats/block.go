package stats

type ProficiencyBonus int

func (p ProficiencyBonus) String() string {
	return modifierString(p)
}

type Block struct {
	abilityScores    AbilityScoreBlock
	proficiencyBonus ProficiencyBonus
}

func NewBlock(abilityScores AbilityScoreBlock, proficiencyBonus ProficiencyBonus) Block {
	return Block{
		abilityScores:    abilityScores,
		proficiencyBonus: proficiencyBonus,
	}
}

func (b Block) AbilityScores() AbilityScoreBlock {
	return b.abilityScores
}

func (b Block) ProficiencyBonus() ProficiencyBonus {
	return b.proficiencyBonus
}
