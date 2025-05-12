package utils

import (
	"math"

	"github.com/beeploop/simple-additive-weighting/internal/constants"
)

func ApproximateEqual(a, b float64) bool {
	return math.Abs(a-b) < constants.Epsilon
}
