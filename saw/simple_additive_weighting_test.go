package saw

import (
	"slices"
	"testing"

	"github.com/beeploop/simple-additive-weighting/utils"
	"github.com/stretchr/testify/assert"
)

func TestSimpleAdditiveWeighting(t *testing.T) {
	alternatives := []Alternative{
		{
			Title: "service a",
			Criterion: []Criteria{
				{Title: "price", Value: 100},
				{Title: "rating", Value: 4.5},
				{Title: "distance", Value: 2.0},
				{Title: "completed_booking", Value: 100},
			},
		},
		{
			Title: "service b",
			Criterion: []Criteria{
				{Title: "price", Value: 80},
				{Title: "rating", Value: 4.7},
				{Title: "distance", Value: 1.0},
				{Title: "completed_booking", Value: 120},
			},
		},
		{
			Title: "service a",
			Criterion: []Criteria{
				{Title: "price", Value: 90},
				{Title: "rating", Value: 4.2},
				{Title: "distance", Value: 5.0},
				{Title: "completed_booking", Value: 50},
			},
		},
	}

	t.Run("test get criterias with title", func(t *testing.T) {
		tests := []struct {
			input    string
			expected []Criteria
		}{
			{
				input: "price",
				expected: []Criteria{
					{Title: "price", Value: 100},
					{Title: "price", Value: 80},
					{Title: "price", Value: 90},
				},
			},
			{
				input: "rating",
				expected: []Criteria{
					{Title: "rating", Value: 4.5},
					{Title: "rating", Value: 4.7},
					{Title: "rating", Value: 4.2},
				},
			},
			{
				input:    "foo",
				expected: []Criteria{},
			},
		}

		for _, test := range tests {
			clonedAlternatives := slices.Clone(alternatives)
			s := NewSAW(clonedAlternatives)

			result := s.CriteriasWithTitle(test.input)

			assert.EqualValues(t, test.expected, result)
		}

	})

	t.Run("test compute weighted sum", func(t *testing.T) {
		tests := []struct {
			input    []WeightAndNormalizedPair
			expected float64
		}{
			{
				input: []WeightAndNormalizedPair{
					{Weight: 0.3, Normalized: 0.80},
					{Weight: 0.3, Normalized: 0.96},
					{Weight: 0.2, Normalized: 0.50},
					{Weight: 0.2, Normalized: 0.60},
				},
				expected: 0.748,
			},
			{
				input: []WeightAndNormalizedPair{
					{Weight: 0.3, Normalized: 1.00},
					{Weight: 0.3, Normalized: 1.00},
					{Weight: 0.2, Normalized: 0.20},
					{Weight: 0.2, Normalized: 1.00},
				},
				expected: 0.84,
			},
			{
				input: []WeightAndNormalizedPair{
					{Weight: 0.3, Normalized: 0.89},
					{Weight: 0.3, Normalized: 0.89},
					{Weight: 0.2, Normalized: 1.00},
					{Weight: 0.2, Normalized: 0.25},
				},
				expected: 0.784,
			},
		}

		for _, test := range tests {
			clonedAlternatives := slices.Clone(alternatives)
			s := NewSAW(clonedAlternatives)

			result := s.ComputeWeightedSum(test.input)
			isApproximatelyEqual := utils.ApproximateEqual(test.expected, result)

			assert.True(t, isApproximatelyEqual, "expected: %v, got: %v", test.expected, result)
		}
	})
}

