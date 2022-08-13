package dice

type PoolResult struct {
	total       int
	rollResults []RollResult
}

type RollResult struct {
	value    int
	rerolled bool
}

func (rr *RollResult) Value() int {
	return rr.value
}

func (pr *PoolResult) Total() int {
	return pr.total
}

func (pr *PoolResult) RollResults() []RollResult {
	return pr.rollResults
}
