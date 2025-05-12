package saw

import "math"

type normalizer struct{}

func NewNormalizer() *normalizer {
	return &normalizer{}
}

func (n *normalizer) NormalizeCost(value float64, otherValues []Criteria) float64 {
	minValue := n.getMin(otherValues)
	return minValue / value
}

func (n *normalizer) NormalizeBenefit(value float64, otherValues []Criteria) float64 {
	maxValue := n.getMax(otherValues)
	return value / maxValue
}

func (n *normalizer) getMax(criterion []Criteria) float64 {
	maxV := math.SmallestNonzeroFloat64

	for _, criteria := range criterion {
		if criteria.Value >= maxV {
			maxV = criteria.Value
		}
	}

	return maxV
}

func (n *normalizer) getMin(criterion []Criteria) float64 {
	minV := math.MaxFloat32

	for _, criteria := range criterion {
		if criteria.Value < minV {
			minV = criteria.Value
		}
	}

	return minV
}
