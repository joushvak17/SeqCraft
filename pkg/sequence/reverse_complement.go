package sequence

// ReverseComplement returns the reverse complement of a DNA sequence.
func ReverseComplement(seq string) string {
	complement := map[rune]rune{
		'A': 'T', 'T': 'A',
		'C': 'G', 'G': 'C',
		'a': 't', 't': 'a',
		'c': 'g', 'g': 'c',
	}
	revComp := make([]rune, len(seq))
	for i, base := range seq {
		revComp[len(seq)-1-i] = complement[base]
	}
	return string(revComp)
}
