package parse

import (
	"os"
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	// Create a temporary FASTA file
	fastaContent := `>sp|P01308|INS_HUMAN Insulin OS=Homo sapiens OX=9606 GN=INS PE=1 SV=1
MALWMRLLPLLALLALWGPDPAAAFVNQHLCGSHLVEALYLVCGERGFFYTPKTRREAED
LQVGQVELGGGPGAGSLQPLALEGSLQKRGIVEQCCTSICSLYQLENYCN
>sp|A0PK11|CLRN2_HUMAN Clarin-2 OS=Homo sapiens OX=9606 GN=CLRN2 PE=1 SV=1
MPGWFKKAWYGLASLLSFSSFILIIVALVVPHWLSGKILCQTGVDLVNATDRELVKFIGD
IYYGLFRGCKVRQCGLGGRQSQFTIFPHLVKELNAGLHVMILLLLFLALALALVSMGFAI
LNMIQVPYRAVSGPGGICLWNVLAGGVVALAIASFVAAVKFHDLTERIANFQEKLFQFVV
VEEQYEESFWICVASASAHAANLVVVAISQIPLPEIKTKIEEATVTAEDILY`

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
		{
			ID:          "sp|P01308|INS_HUMAN",
			Description: "Insulin OS=Homo sapiens OX=9606 GN=INS PE=1 SV=1",
			Sequence:    "MALWMRLLPLLALLALWGPDPAAAFVNQHLCGSHLVEALYLVCGERGFFYTPKTRREAEDLQVGQVELGGGPGAGSLQPLALEGSLQKRGIVEQCCTSICSLYQLENYCN",
		},
		{
			ID:          "sp|A0PK11|CLRN2_HUMAN",
			Description: "Clarin-2 OS=Homo sapiens OX=9606 GN=CLRN2 PE=1 SV=1",
			Sequence:    "MPGWFKKAWYGLASLLSFSSFILIIVALVVPHWLSGKILCQTGVDLVNATDRELVKFIGDIYYGLFRGCKVRQCGLGGRQSQFTIFPHLVKELNAGLHVMILLLLFLALALALVSMGFAILNMIQVPYRAVSGPGGICLWNVLAGGVVALAIASFVAAVKFHDLTERIANFQEKLFQFVVVEEQYEESFWICVASASAHAANLVVVAISQIPLPEIKTKIEEATVTAEDILY",
		},
	}

	// Verify the parsed records
	if !reflect.DeepEqual(records, expectedRecords) {
		t.Errorf("Parse() = %v, want %v", records, expectedRecords)
	}

	// Test sequence length
	for i, record := range records {
		expectedLength := len(expectedRecords[i].Sequence)
		actualLength := len(record.Sequence)
		if actualLength != expectedLength {
			t.Errorf("Sequence length mismatch for record %d: got %d, want %d", i, actualLength, expectedLength)
		}
	}
}
