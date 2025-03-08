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

			// Add additional information to the output message
			fmt.Printf("\nSeqCraft Parse Output(s):\n")
			fmt.Printf("Number of records parsed: %d\n", len(records))

			totalLength := 0
			totalGCContent := 0.0
			totalNucleotideFreq := make(map[rune]float64)

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
					totalLength += length
					fmt.Printf("Sequence Length: %d\n", length)
				}

				if gcContent {
					gc := sequence.GCContent(record.Sequence)
					totalGCContent += gc
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
					for nucleotide, count := range freq {
						totalNucleotideFreq[nucleotide] += count
					}
					fmt.Println("Nucleotide Frequency:")
					for nucleotide, count := range freq {
						fmt.Printf("%s: %.4f\n", string(nucleotide), count)
					}
				}
			}

			// Print aggregate statistics
			if sequenceLength {
				averageLength := float64(totalLength) / float64(len(records))
				fmt.Printf("\nTotal Sequence Length: %d\n", totalLength)
				fmt.Printf("Average Sequence Length: %.2f\n", averageLength)
			}

			if gcContent {
				averageGCContent := totalGCContent / float64(len(records))
				fmt.Printf("Average GC Content: %.2f%%\n", averageGCContent)
			}

			if nucleotideFreq {
				fmt.Println("Total Nucleotide Frequency:")
				for nucleotide, count := range totalNucleotideFreq {
					fmt.Printf("%s: %.4f\n", string(nucleotide), count)
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
