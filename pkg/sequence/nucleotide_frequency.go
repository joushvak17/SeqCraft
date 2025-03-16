package sequence

// NucleotideFrequency calculates the frequency of each nucleotide in a sequence.
func NucleotideFrequency(seq string) map[rune]float64 {
	if len(seq) == 0 {
		return map[rune]float64{}
	}

	// Use byte array for faster access
	// Index corresponds to ASCII value of the base
	var validBase [256]byte
	validBase['A'] = 1
	validBase['T'] = 1
	validBase['G'] = 1
	validBase['C'] = 1
	validBase['U'] = 1
	validBase['N'] = 1
	validBase['a'] = 1
	validBase['t'] = 1
	validBase['g'] = 1
	validBase['c'] = 1
	validBase['u'] = 1
	validBase['n'] = 1

	// Storing uppercase bases for case-insensitive comparison
	var upperCase [256]byte
	upperCase['a'] = 'A'
	upperCase['t'] = 'T'
	upperCase['g'] = 'G'
	upperCase['c'] = 'C'
	upperCase['n'] = 'N'
	upperCase['n'] = 'N'
	upperCase['u'] = 'U'

	// Pre-allocate map with all possible bases
	freq := make(map[rune]float64, 6)
	validBases := 0

	// Process as byte array for faster access
	for i := range len(seq) {
		base := seq[i]
		if validBase[base] == 1 {
			if base >= 'a' && base <= 'z' {
				base = upperCase[base]
			}
			freq[rune(base)]++
			validBases++
		}
	}

	// Calculate frequencies
	total := float64(validBases)
	if total > 0 {
		for base, count := range freq {
			freq[base] = count / total * 100
		}
	}

	return freq
}
