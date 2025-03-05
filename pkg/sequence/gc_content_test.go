package sequence

import "testing"

func TestGCContent(t *testing.T) {
	tests := []struct {
		seq      string
		expected float64
	}{
		// TODO: Come up with more test cases
		{"ATGC", 50.0},
		{"GGCC", 100.0},
		{"ATAT", 0.0},
		{"", 0.0}, // Edge case: empty sequence
	}

	for _, test := range tests {
		result := GCContent(test.seq)
		if result != test.expected {
			t.Errorf("GCContent(%s) = %.2f, want %.2f", test.seq, result, test.expected)
		}
	}
}
