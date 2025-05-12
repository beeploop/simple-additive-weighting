package saw

type Criteria struct {
	Title string
	Value float64
}

// e.g., Service with criterion price, rating, distance, completed_bookings
type Alternative struct {
	Title     string
	Criterion []Criteria
}

// return Criteria and bool determining if found or not
func (a *Alternative) CriteriaWithTitle(title string) (Criteria, bool) {
	if title == "" {
		return Criteria{}, false
	}

	for _, criteria := range a.Criterion {
		if criteria.Title == title {
			return criteria, true
		}
	}

	return Criteria{}, false
}
