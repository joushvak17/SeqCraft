package sequence

import "testing"

func TestReverseComplement(t *testing.T) {
	tests := []struct {
		seq      string
		expected string
	}{
		{"ATGC", "GCAT"},
		{"GGCC", "GGCC"},
		{"ATAT", "ATAT"},
		{"", ""}, // Edge case: empty sequence
	}

	for _, test := range tests {
		result := ReverseComplement(test.seq)
		if result != test.expected {
			t.Errorf("ReverseComplement(%s) = %s, want %s", test.seq, result, test.expected)
		}
	}
}
