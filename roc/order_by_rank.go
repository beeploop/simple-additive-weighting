package roc

import "slices"

func orderByRank(criterion []Criteria) []Criteria {
	cloned := slices.Clone(criterion)

	slices.SortFunc(cloned, func(a, b Criteria) int {
		return a.Rank - b.Rank
	})

	return cloned
}
