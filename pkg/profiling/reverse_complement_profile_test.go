package profiling

import (
	"os"
	"runtime/pprof"
	"testing"

	"github.com/joushvak17/SeqCraft/pkg/sequence"
)

func TestReverseComplementProfile(t *testing.T) {
	file, err := os.Create("reverse_complement.prof")
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
		reverseComp := sequence.ReverseComplement("ATGCATGCATGCATGC")
		if reverseComp == "" {
			t.Error("ReverseComplement returned empty string unexpectedly")
		}
	}
}
