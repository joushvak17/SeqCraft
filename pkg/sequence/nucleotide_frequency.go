package sequence

// Define these as package-level constants to avoid recreating them on each function call
var validBase = [256]byte{
	'A': 1, 'T': 1, 'G': 1, 'C': 1, 'U': 1, 'N': 1,
	'a': 1, 't': 1, 'g': 1, 'c': 1, 'u': 1, 'n': 1,
}

var upperCase = [256]byte{
	'a': 'A', 't': 'T', 'g': 'G', 'c': 'C', 'n': 'N', 'u': 'U',
}

// Common nucleotide bases we care about
var nucleotideBases = [6]byte{'A', 'C', 'G', 'T', 'U', 'N'}

// NucleotideFrequency calculates the frequency of each relevant nucleotide in a sequence.
func NucleotideFrequency(seq string) map[rune]float64 {
	if len(seq) == 0 {
		return map[rune]float64{}
	}

	// Use a smaller array for just the bases we care about (to reduce allocation cost)
	// Index 0=A, 1=C, 2=G, 3=T, 4=U, 5=N
	counts := [6]int{}
	validBases := 0

	// Direct string access with byte conversion is faster
	for i := 0; i < len(seq); i++ {
		base := seq[i]
		if validBase[base] == 1 {
			if base >= 'a' && base <= 'z' {
				base = upperCase[base]
			}

			// Map the base to an index in our small array
			switch base {
			case 'A':
				counts[0]++
			case 'C':
				counts[1]++
			case 'G':
				counts[2]++
			case 'T':
				counts[3]++
			case 'U':
				counts[4]++
			case 'N':
				counts[5]++
			}
			validBases++
		}
	}

	// Create final result map
	freq := make(map[rune]float64, 6)

	if validBases > 0 {
		total := float64(validBases)
		// We know exactly which bases to process now
		for i, count := range counts {
			if count > 0 {
				freq[rune(nucleotideBases[i])] = float64(count) / total * 100
			}
		}
	}

	return freq
}
