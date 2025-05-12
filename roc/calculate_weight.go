package roc

func CalculateWeight(target Criteria, criterion []Criteria) float64 {
	initial := 1 / float64(len(criterion))

	ordered := orderByRank(criterion)
	startingPos := target.Rank - 1

	var sum float64
	for i := startingPos; i < len(ordered); i++ {
		sum += 1 / float64(ordered[i].Rank)
	}

	return initial * sum
}
