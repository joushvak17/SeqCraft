package sequence

// GCContent calculates the GC content of a sequence as a percentage.
func GCContent(seq string) float64 {
	// TODO: Consider using goroutines to parallelize the calculation of GC content.
	// FIXME: Reduce Goroutine overhead by processing multiple sequences in a single goroutine.
	// FIXME: Batch multiple sequences together to reduce the number of goroutines.

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
