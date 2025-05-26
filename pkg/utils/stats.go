package utils

import "sort"

// CalculateLengthStats calculates the minimum, maximum, and median values from a slice of integers.
// Returns min, max, and median in that order.
func CalculateLengthStats(lengths []int) (min int, max int, median float64) {
	if len(lengths) == 0 {
		return 0, 0, 0
	}
	sort.Ints(lengths)
	min, max = lengths[0], lengths[len(lengths)-1]
	if len(lengths)%2 == 0 {
		median = float64(lengths[len(lengths)/2-1]+lengths[len(lengths)/2]) / 2
	} else {
		median = float64(lengths[len(lengths)/2])
	}
	return min, max, median
}
