package sequence

// ReverseComplement returns the reverse complement of a DNA sequence.
func ReverseComplement(seq string) string {
	reverseComp := make([]byte, len(seq))

	for i := range len(seq) {
		var comp byte
		switch seq[len(seq)-1-i] {
		case 'A':
			comp = 'T'
		case 'T':
			comp = 'A'
		case 'C':
			comp = 'G'
		case 'G':
			comp = 'C'
		case 'a':
			comp = 't'
		case 't':
			comp = 'a'
		case 'c':
			comp = 'g'
		case 'g':
			comp = 'c'
		default:
			comp = seq[len(seq)-1-i]
		}

		reverseComp[i] = comp
	}

	return string(reverseComp)
}
