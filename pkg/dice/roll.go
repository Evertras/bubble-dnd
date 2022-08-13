package dice

import "math/rand"

func (p *Pool) Roll() PoolResult {
	return p.RollWithRerollStrategy(nil)
}

func (p *Pool) RollWithRerollStrategy(rerollStrategy func(int) bool) PoolResult {
	if p.count <= 0 {
		return PoolResult{
			total:       0,
			rollResults: []RollResult{},
		}
	}

	total := 0
	rollResults := make([]RollResult, 0, p.count)

	sides := int(p.kind)

	rollDie := func() int { return rand.Intn(sides) + 1 }

	for i := 0; i < p.count; i++ {
		rerolled := false
		value := rollDie()

		if rerollStrategy != nil && rerollStrategy(value) {
			value = rollDie()
			rerolled = true
		}

		total += value

		rollResults = append(rollResults, RollResult{
			value:    value,
			rerolled: rerolled,
		})
	}

	return PoolResult{
		total:       total,
		rollResults: rollResults,
	}
}
