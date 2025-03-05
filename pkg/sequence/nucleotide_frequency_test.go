package sequence

import (
	"reflect"
	"testing"
)

func TestNucleotideFrequency(t *testing.T) {
	tests := []struct {
		seq      string
		expected map[rune]float64
	}{
		{"ATGC", map[rune]float64{'A': 25.0, 'T': 25.0, 'G': 25.0, 'C': 25.0}},
		{"GGCC", map[rune]float64{'G': 50.0, 'C': 50.0}},
		{"ATAT", map[rune]float64{'A': 50.0, 'T': 50.0}},
		{"", map[rune]float64{}}, // Edge case: empty sequence
	}

	for _, test := range tests {
		result := NucleotideFrequency(test.seq)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("NucleotideFrequency(%s) = %v, want %v", test.seq, result, test.expected)
		}
	}
}
