package utils

import "testing"

func TestCalculateLengthStats(t *testing.T) {
	tests := []struct {
		name     string
		lengths  []int
		expected [3]float64
	}{
		{"Empty slice", []int{}, [3]float64{0, 0, 0}},
		{"Single element", []int{5}, [3]float64{5, 5, 5}},
		{"Two elements", []int{5, 10}, [3]float64{5, 10, 7.5}},
		{"Odd number of elements", []int{1, 2, 3}, [3]float64{1, 3, 2}},
		{"Even number of elements", []int{1, 2, 3, 4}, [3]float64{1, 4, 2.5}},
	}

	for _, test := range tests {
		min, max, median := CalculateLengthStats(test.lengths)
		if float64(min) != test.expected[0] || float64(max) != test.expected[1] || float64(median) != test.expected[2] {
			t.Errorf("CalculateLengthStats(%v) = %v, %v, %v; want %v", test.lengths, min, max, median, test.expected)
		}
	}
}
