package sequence

// GCContent calculates the GC content of a sequence
func GCContent(seq string) float64 {
	if len(seq) == 0 {
		return 0.0
	}

	gcCount := 0
	for _, base := range seq {
		if base == 'g' || base == 'c' || base == 'G' || base == 'C' {
			gcCount++
		}
	}

	return float64(gcCount) / float64(len(seq)) * 100
}
