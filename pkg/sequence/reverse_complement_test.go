package sequence

import "testing"

func TestReverseComplement(t *testing.T) {
	tests := []struct {
		name     string
		seq      string
		expected string
	}{
		{"Basic example", "ATGC", "GCAT"},
		{"Palindromic sequence", "GGCC", "GGCC"},
		{"Palindromic sequence", "ATAT", "ATAT"},
		{"Empty sequence", "", ""}, // Edge case: empty sequence
		{"Lower case", "atgc", "gcat"},
		{"Mixed case", "ATgc", "gcAT"},
		{"Non-standard characters", "ATN-", "-NAT"},
		{"Long sequence", "ATGCATGCATGCATGC", "GCATGCATGCATGCAT"},
	}

	for _, test := range tests {
		result := ReverseComplement(test.seq)
		if result != test.expected {
			t.Errorf("ReverseComplement(%q) = %s, want %s", test.seq, result, test.expected)
		}
	}
}
