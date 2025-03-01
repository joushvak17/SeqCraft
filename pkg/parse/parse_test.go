package parse

import (
	"os"
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	// Create a temporary FASTA file
	fastaContent := `>seq1 description1
	ATCG
	>seq2 description2
	GGTA`
	tmpfile, err := os.CreateTemp("", "test.fasta")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.Write([]byte(fastaContent)); err != nil {
		t.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatal(err)
	}

	// Call the Parse function
	records, err := Parse(tmpfile.Name())
	if err != nil {
		t.Fatalf("Parse() error = %v", err)
	}

	// Expected records
	expectedRecords := []Record{
		{ID: "seq1", Description: "description1", Sequence: "ATCG"},
		{ID: "seq2", Description: "description2", Sequence: "GGTA"},
	}

	// Verify the results
	if !reflect.DeepEqual(records, expectedRecords) {
		t.Errorf("Parse() = %v, want %v", records, expectedRecords)
	}
}