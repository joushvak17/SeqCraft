package sequence

// GCContent calculates the GC content of a sequence as a percentage.
func GCContent(seq string) float64 {
	if len(seq) == 0 {
		return 0.0
	}

	gcCount := 0
	seqBytes := []byte(seq)

	for i := range len(seqBytes) {
		switch seq[i] {
		case 'g', 'c', 'G', 'C':
			gcCount++
		}
	}

	return float64(gcCount) / float64(len(seq)) * 100
}
