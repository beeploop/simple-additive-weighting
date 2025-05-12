package saw

import (
	"slices"

	"github.com/beeploop/simple-additive-weighting/utils"
)

type WeightAndNormalizedPair struct {
	Weight     float64
	Normalized float64
}

type SimpleAdditiveWeighting struct {
	Alternatives []Alternative
}

func NewSAW(alternatives []Alternative) *SimpleAdditiveWeighting {
	return &SimpleAdditiveWeighting{
		Alternatives: alternatives,
	}
}

func (s *SimpleAdditiveWeighting) CriteriasWithTitle(title string) []Criteria {
	criterias := make([]Criteria, 0)

	for _, alternative := range s.Alternatives {
		for _, criteria := range alternative.Criterion {
			if criteria.Title == title {
				criterias = append(criterias, criteria)
			}
		}
	}

	return criterias
}

func (s *SimpleAdditiveWeighting) ComputeWeightedSum(input []WeightAndNormalizedPair) float64 {
	initial := utils.Map(input, func(pair WeightAndNormalizedPair) float64 {
		return pair.Weight * pair.Normalized
	})

	sum := slices.Collect(utils.Reduce(initial, func(acc, b float64) float64 {
		return acc + b
	}, 0))

	return sum[0]
}
