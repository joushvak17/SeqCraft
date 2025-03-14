package sequence

import "testing"

func TestGCContent(t *testing.T) {
	tests := []struct {
		name     string
		seq      string
		expected float64
	}{
		{"Balanced", "ATGC", 50.0},
		{"Full GC", "GGCC", 100.0},
		{"No GC", "ATAT", 0.0},
		{"Empty sequence", "", 0.0},
		{"Mixed case", "ATgc", 50.0},
		{"Non-standard characters", "ATN-", 0.0},
		{"Long sequence", "ATGCATGCATGCATGC", 50.0},
	}

	for _, test := range tests {
		// TODO: Consider if calling the name of the test is necessary.
		t.Run(test.name, func(t *testing.T) {
			result := GCContent(test.seq)
			if result != test.expected {
				t.Errorf("GCContent(%q) = %.2f, want %.2f", test.seq, result, test.expected)
			}
		})
	}
}
