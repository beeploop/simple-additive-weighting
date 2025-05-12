package roc

import (
	"testing"

	"github.com/beeploop/simple-additive-weighting/utils"
	"github.com/stretchr/testify/assert"
)

func TestCalculateROC(t *testing.T) {
	tests := []struct {
		input     string
		criterion []Criteria
		expected  float64
	}{
		{
			input: "price",
			criterion: []Criteria{
				{Title: "price", Rank: 1},
				{Title: "rating", Rank: 2},
				{Title: "distance", Rank: 3},
				{Title: "completed_bookings", Rank: 4},
			},
			expected: 0.521,
		},
		{
			input: "rating",
			criterion: []Criteria{
				{Title: "price", Rank: 1},
				{Title: "rating", Rank: 2},
				{Title: "distance", Rank: 3},
				{Title: "completed_bookings", Rank: 4},
			},
			expected: 0.27,
		},
		{
			input: "distance",
			criterion: []Criteria{
				{Title: "price", Rank: 1},
				{Title: "rating", Rank: 2},
				{Title: "distance", Rank: 3},
				{Title: "completed_bookings", Rank: 4},
			},
			expected: 0.145,
		},
		{
			input: "completed_bookings",
			criterion: []Criteria{
				{Title: "price", Rank: 1},
				{Title: "rating", Rank: 2},
				{Title: "distance", Rank: 3},
				{Title: "completed_bookings", Rank: 4},
			},
			expected: 0.0625,
		},
	}

	for _, test := range tests {
		r := NewRankOrderCentroid(test.criterion)

		weight := r.CalculateWeightOf(test.input)

		approximatelyEqual := utils.ApproximateEqual(test.expected, weight)
		assert.True(t, approximatelyEqual, "expected: %v , got: %v", test.expected, weight)
	}
}
