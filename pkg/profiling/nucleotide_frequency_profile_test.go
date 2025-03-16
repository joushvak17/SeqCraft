package profiling

import (
	"os"
	"runtime/pprof"
	"testing"

	"github.com/joushvak17/SeqCraft/pkg/sequence"
)

func TestNucleotideFrequencyProfile(t *testing.T) {
	file, err := os.Create("nucleotide_frequency.prof")
	if err != nil {
		t.Fatal("Could not create CPU profile: ", err)
	}
	defer file.Close()

	if err := pprof.StartCPUProfile(file); err != nil {
		t.Fatal("Could not start CPU profile: ", err)
	}
	defer pprof.StopCPUProfile()

	for range 1000000 {
		// TODO: Add more real-world test cases to improve profiling data.
		nucleotideFreq := sequence.NucleotideFrequency("ATGCATGCATGCATGC")
		if nucleotideFreq == nil {
			t.Error("NucleotideFrequency returned nil unexpectedly")
		}
	}
}
