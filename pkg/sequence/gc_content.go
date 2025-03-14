package sequence

// GCContent calculates the GC content of a sequence as a percentage.
func GCContent(seq string) float64 {
	// TODO: Consider using the bytes package when dealing with long sequences.
	// TODO: Consider using map based cache for previously calculated GC content.
	// TODO: Consider using goroutines to parallelize the calculation of GC content.

	if len(seq) == 0 {
		return 0.0
	}

	gcCount, seqBytes := 0, []byte(seq)
	for i := range len(seqBytes) {
		switch seq[i] {
		case 'g', 'c', 'G', 'C':
			gcCount++
		}
	}

	return float64(gcCount) / float64(len(seq)) * 100
}
