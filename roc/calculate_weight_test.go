package roc

import (
	"testing"

	"github.com/beeploop/simple-additive-weighting/utils"
	"github.com/stretchr/testify/assert"
)

func TestCalculateROC(t *testing.T) {
	tests := []struct {
		criteria  Criteria
		criterion []Criteria
		expected  float64
	}{
		{
			criteria: Criteria{Title: "price", Rank: 1},
			criterion: []Criteria{
				{Title: "price", Rank: 1},
				{Title: "rating", Rank: 2},
				{Title: "distance", Rank: 3},
				{Title: "completed bookings", Rank: 4},
			},
			expected: 0.521,
		},
		{
			criteria: Criteria{Title: "rating", Rank: 2},
			criterion: []Criteria{
				{Title: "price", Rank: 1},
				{Title: "rating", Rank: 2},
				{Title: "distance", Rank: 3},
				{Title: "completed bookings", Rank: 4},
			},
			expected: 0.27,
		},
		{
			criteria: Criteria{Title: "distance", Rank: 3},
			criterion: []Criteria{
				{Title: "price", Rank: 1},
				{Title: "rating", Rank: 2},
				{Title: "distance", Rank: 3},
				{Title: "completed bookings", Rank: 4},
			},
			expected: 0.145,
		},
		{
			criteria: Criteria{Title: "completed bookings", Rank: 4},
			criterion: []Criteria{
				{Title: "price", Rank: 1},
				{Title: "rating", Rank: 2},
				{Title: "distance", Rank: 3},
				{Title: "completed bookings", Rank: 4},
			},
			expected: 0.0625,
		},
	}

	for _, test := range tests {
		score := CalculateWeight(test.criteria, test.criterion)

		approximatelyEqual := utils.ApproximateEqual(test.expected, score)
		assert.True(t, approximatelyEqual, "expected: %v , got: %v", test.expected, score)
	}
}
