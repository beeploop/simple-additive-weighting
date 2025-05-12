package saw

import (
	"testing"

	"github.com/beeploop/simple-additive-weighting/utils"
	"github.com/stretchr/testify/assert"
)

func TestNormalizer(t *testing.T) {
	t.Run("test normalize cost", func(t *testing.T) {
		tests := []struct {
			input     float64
			criterion []Criteria
			expected  float64
		}{
			{
				input: 100,
				criterion: []Criteria{
					{Title: "price", Value: 100},
					{Title: "price", Value: 80},
					{Title: "price", Value: 90},
				},
				expected: 0.80,
			},
			{
				input: 2.0,
				criterion: []Criteria{
					{Title: "distance", Value: 2.0},
					{Title: "distance", Value: 5.0},
					{Title: "distance", Value: 1.0},
				},
				expected: 0.5,
			},
		}

		for _, test := range tests {
			n := NewNormalizer()

			normalized := n.NormalizeCost(test.input, test.criterion)
			isApproximatelyEqual := utils.ApproximateEqual(test.expected, normalized)

			assert.True(t, isApproximatelyEqual, "expected: %v , got: %v", normalized)
		}
	})

	t.Run("test normalize benefit", func(t *testing.T) {
		tests := []struct {
			input     float64
			criterion []Criteria
			expected  float64
		}{
			{
				input: 4.5,
				criterion: []Criteria{
					{Title: "rating", Value: 4.5},
					{Title: "rating", Value: 4.7},
					{Title: "rating", Value: 4.2},
				},
				expected: 0.957,
			},
			{
				input: 120,
				criterion: []Criteria{
					{Title: "completed_bookings", Value: 120},
					{Title: "completed_bookings", Value: 200},
					{Title: "completed_bookings", Value: 50},
				},
				expected: 0.6,
			},
		}

		for _, test := range tests {
			n := NewNormalizer()

			normalized := n.NormalizeBenefit(test.input, test.criterion)
			isApproximatelyEqual := utils.ApproximateEqual(test.expected, normalized)

			assert.True(t, isApproximatelyEqual, "expected: %v , got: %v", normalized)
		}
	})
}
