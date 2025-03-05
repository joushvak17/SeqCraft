package sequence

// NucleotideFrequency calculates the frequency of each nucleotide in a sequence.
func NucleotideFrequency(seq string) map[rune]float64 {
	freq := make(map[rune]float64)
	total := float64(len(seq))

	for _, base := range seq {
		freq[base]++
	}

	// Convert counts to frequencies
	for base := range freq {
		freq[base] = freq[base] / total * 100
	}

	return freq
}
