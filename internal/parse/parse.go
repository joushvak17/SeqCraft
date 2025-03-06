package parse

import (
	"fmt"
	"os"

	"github.com/joushvak17/SeqCraft/pkg/parse"
	"github.com/joushvak17/SeqCraft/pkg/sequence"
	"github.com/spf13/cobra"
	"golang.org/x/term"
)

// NewParseCmd creates and returns the `parse` command.
func NewParseCmd() *cobra.Command {
	var (
		// Define flags for the command
		sequenceLength bool
		gcContent      bool
		reverseComp    bool
		nucleotideFreq bool
	)

	parseCmd := &cobra.Command{
		Use:   "parse <file>",
		Short: "Parse and analyze a FASTA file",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			filename := args[0]
			records, err := parse.Parse(filename)
			if err != nil {
				fmt.Printf("Error parsing FASTA file: %v\n", err)
				return
			}

			// Get the terminal width
			width, _, err := term.GetSize(int(os.Stdout.Fd()))
			if err != nil {
				width = 80 // Standard terminal width
			}

			// TODO: Add additional information to the output message, such as the number of records parsed, etc
			fmt.Printf("\nSeqCraft Parse Output(s):\n")

			for _, record := range records {
				fmt.Printf("\nID: %s\n", record.ID)
				fmt.Printf("Description: %s\n", record.Description)

				// Print the sequence with a maximum width based on the terminal width and the prefix length
				sequenceValue := record.Sequence
				prefixLength := len("Sequence: ")
				width -= prefixLength
				if len(sequenceValue) > width {
					sequenceValue = sequenceValue[:width-3] + "..."
				}
				fmt.Printf("Sequence: %s\n\n", sequenceValue)

				if sequenceLength {
					length := len(record.Sequence)
					fmt.Printf("Sequence Length: %d\n", length)
				}

				if gcContent {
					gc := sequence.GCContent(record.Sequence)
					fmt.Printf("GC Content: %.2f%%\n", gc)
				}

				if reverseComp {
					// Calculate reverse complement
					reverse := sequence.ReverseComplement(record.Sequence)
					fmt.Printf("Reverse Complement: %s\n", reverse)
				}

				if nucleotideFreq {
					// Calculate nucleotide frequency
					freq := sequence.NucleotideFrequency(record.Sequence)
					fmt.Println("Nucleotide Frequency:")
					for nucleotide, count := range freq {
						fmt.Printf("%s: %f\n", string(nucleotide), count)
					}
				}
			}

			fmt.Println() // Add a newline at the end
		},
	}

	// Add flags to the command
	parseCmd.Flags().BoolVarP(&sequenceLength, "length", "l", false, "Calculate sequence length")
	parseCmd.Flags().BoolVarP(&gcContent, "gc", "g", false, "Calculate GC content")
	parseCmd.Flags().BoolVarP(&reverseComp, "reverse", "r", false, "Calculate reverse complement")
	parseCmd.Flags().BoolVarP(&nucleotideFreq, "freq", "f", false, "Calculate nucleotide frequency")

	return parseCmd
}
