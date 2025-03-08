package sequence

// ReverseComplement returns the reverse complement of a DNA sequence.
func ReverseComplement(seq string) string {
	complement := map[rune]rune{
		'A': 'T',
		'T': 'A',
		'C': 'G',
		'G': 'C',
		'a': 't',
		't': 'a',
		'c': 'g',
		'g': 'c',
	}

	// Create a slice to hold the reverse complement
	reverseComp := make([]rune, len(seq))

	// Iterate over the sequence in reverse order
	for i, nucleotide := range seq {
		if comp, ok := complement[nucleotide]; ok {
			reverseComp[len(seq)-1-i] = comp
		} else {
			// If the nucleotide is not recognized, keep it as is
			reverseComp[len(seq)-1-i] = nucleotide
		}
	}

	return string(reverseComp)
}
