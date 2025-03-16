package sequence

import (
	"reflect"
	"testing"
)

func TestNucleotideFrequency(t *testing.T) {
	tests := []struct {
		name     string
		seq      string
		expected map[rune]float64
	}{
		{
			name:     "Equal Distribution",
			seq:      "ATGC",
			expected: map[rune]float64{'A': 25.0, 'T': 25.0, 'G': 25.0, 'C': 25.0},
		},
		{
			name:     "GC-Rich Sequence",
			seq:      "GGCC",
			expected: map[rune]float64{'G': 50.0, 'C': 50.0},
		},
		{
			name:     "AT-Rich Sequence",
			seq:      "ATAT",
			expected: map[rune]float64{'A': 50.0, 'T': 50.0},
		},
		{
			name:     "Empty Sequence",
			seq:      "",
			expected: map[rune]float64{},
		},
		{
			name:     "Mixed Case Sequence",
			seq:      "AtGc",
			expected: map[rune]float64{'A': 25.0, 'T': 25.0, 'G': 25.0, 'C': 25.0},
		},
		{
			name:     "Sequence with Non-Nucleotide Characters",
			seq:      "ATGC-N",
			expected: map[rune]float64{'A': 20.0, 'T': 20.0, 'G': 20.0, 'C': 20.0, 'N': 20.0},
		},
		{
			name:     "Long Sequence",
			seq:      "AAAAAAAAAAAAAAAAAAAAGGGGGGGGGGCCCCCCCCCCTTTTTTTTTT",
			expected: map[rune]float64{'A': 40.0, 'G': 20.0, 'C': 20.0, 'T': 20.0},
		},
	}

	for _, test := range tests {
		result := NucleotideFrequency(test.seq)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("NucleotideFrequency(%s) = %v, want %v", test.seq, result, test.expected)
		}
	}
}
