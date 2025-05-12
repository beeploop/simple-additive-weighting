package roc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrderByRank(t *testing.T) {
	tests := []struct {
		input    []Criteria
		expected []Criteria
	}{
		{
			input: []Criteria{
				{Title: "price", Rank: 1},
				{Title: "rating", Rank: 2},
				{Title: "distance", Rank: 3},
				{Title: "completed bookings", Rank: 4},
			},
			expected: []Criteria{
				{Title: "price", Rank: 1},
				{Title: "rating", Rank: 2},
				{Title: "distance", Rank: 3},
				{Title: "completed bookings", Rank: 4},
			},
		},
		{
			input: []Criteria{
				{Title: "price", Rank: 1},
				{Title: "completed bookings", Rank: 4},
				{Title: "rating", Rank: 2},
				{Title: "distance", Rank: 3},
			},
			expected: []Criteria{
				{Title: "price", Rank: 1},
				{Title: "rating", Rank: 2},
				{Title: "distance", Rank: 3},
				{Title: "completed bookings", Rank: 4},
			},
		},
	}

	for _, test := range tests {
		sorted := orderByRank(test.input)

		assert.Equal(t, test.expected, sorted)
	}
}
