package sequence

// GCContent calculates the G and C nucleotides (case-insensitive) of a sequence as a percentage.
func GCContent(seq string) float64 {
	if len(seq) == 0 {
		return 0.0
	}

	gcCount := 0

	// Use a standard for loop for index-based access
	for i := range seq {
		switch seq[i] {
		case 'g', 'c', 'G', 'C':
			gcCount++
		}
	}

	return float64(gcCount) / float64(len(seq)) * 100.0
}
