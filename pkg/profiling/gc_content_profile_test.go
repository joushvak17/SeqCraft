package profiling

import (
	"os"
	"runtime/pprof"
	"testing"

	"github.com/joushvak17/SeqCraft/pkg/sequence"
)

func TestGCContentProfile(t *testing.T) {
	file, err := os.Create("gc_content.prof")
	if err != nil {
		t.Fatal("Could not create CPU profile: ", err)
	}
	defer file.Close()

	if err := pprof.StartCPUProfile(file); err != nil {
		t.Fatal("Could not start CPU profile: ", err)
	}
	defer pprof.StopCPUProfile()

	for i := 0; i < 1000000; i++ {
		sequence.GCContent("ATGCATGCATGCATGC")
	}
}
