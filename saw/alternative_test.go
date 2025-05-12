package saw

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAlternative(t *testing.T) {
	t.Run("get a criteria with title", func(t *testing.T) {
		tests := []struct {
			input            string
			alternative      Alternative
			expectedCriteria Criteria
			expectedFound    bool
		}{
			{
				input: "price",
				alternative: Alternative{
					Title: "service a",
					Criterion: []Criteria{
						{Title: "price", Value: 10},
						{Title: "rating", Value: 4.5},
						{Title: "distance", Value: 2.0},
						{Title: "bookings_completed", Value: 20},
					},
				},
				expectedCriteria: Criteria{Title: "price", Value: 10},
				expectedFound:    true,
			},
			{
				input: "rating",
				alternative: Alternative{
					Title: "service b",
					Criterion: []Criteria{
						{Title: "price", Value: 10},
						{Title: "rating", Value: 4.5},
						{Title: "distance", Value: 2.0},
						{Title: "bookings_completed", Value: 20},
					},
				},
				expectedCriteria: Criteria{Title: "rating", Value: 4.5},
				expectedFound:    true,
			},
			{
				input: "foo",
				alternative: Alternative{
					Title: "service b",
					Criterion: []Criteria{
						{Title: "price", Value: 10},
						{Title: "rating", Value: 4.5},
						{Title: "distance", Value: 2.0},
						{Title: "bookings_completed", Value: 20},
					},
				},
				expectedCriteria: Criteria{},
				expectedFound:    false,
			},
		}

		for _, test := range tests {
			alternative := Alternative{
				Title:     test.alternative.Title,
				Criterion: test.alternative.Criterion,
			}

			result, found := alternative.CriteriaWithTitle(test.input)

			assert.Equal(t, test.expectedFound, found)
			assert.EqualValues(t, test.expectedCriteria, result)
		}
	})
}
