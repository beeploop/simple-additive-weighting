package roc

type rankOrderCentroid struct {
	Criterion []Criteria
}

func NewRankOrderCentroid(criterion []Criteria) *rankOrderCentroid {
	return &rankOrderCentroid{
		Criterion: criterion,
	}
}

func (roc *rankOrderCentroid) CalculateWeightOf(target string) float64 {
	initial := 1 / float64(len(roc.Criterion))

	criteria, found := roc.getCriteriaWithTitle(target)
	if !found {
		return 0
	}

	ordered := orderByRank(roc.Criterion)
	start := criteria.Rank - 1

	var sum float64
	for i := start; i < len(ordered); i++ {
		sum += 1 / float64(ordered[i].Rank)
	}

	return initial * sum
}

func (roc *rankOrderCentroid) getCriteriaWithTitle(title string) (Criteria, bool) {
	if title == "" {
		return Criteria{}, false
	}

	for _, criteria := range roc.Criterion {
		if criteria.Title == title {
			return criteria, true
		}
	}

	return Criteria{}, false
}
