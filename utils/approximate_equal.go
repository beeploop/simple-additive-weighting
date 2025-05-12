package utils

import "math"

func ApproximateEqual(a, b, epsilon float64) bool {
	return math.Abs(a-b) < epsilon
}
