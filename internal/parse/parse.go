package parse

import (
	"fmt"
	"os"
	"sort"
	"time"

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
		outputFile     string
		verbose        bool
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

			// Get the current date and time
			currentTime := time.Now().Format("January 2, 2006 3:04 PM")

			// Prepare output
			output := "\nSeqCraft Parse Output(s):\n"
			output += fmt.Sprintf("Date and Time: %s\n", currentTime)
			output += fmt.Sprintf("Number of records parsed: %d\n", len(records))

			totalLength := 0
			totalGCContent := 0.0
			totalNucleotideFreq := make(map[rune]float64)
			var lengths []int

			for _, record := range records {
				output += fmt.Sprintf("\nID: %s\n", record.ID)
				output += fmt.Sprintf("Description: %s\n", record.Description)

				// Print the sequence with a maximum width based on the terminal width and the prefix length
				sequenceValue := record.Sequence
				prefixLength := len("Sequence: ")
				width -= prefixLength
				if len(sequenceValue) > width {
					sequenceValue = sequenceValue[:width-3] + "..."
				}
				output += fmt.Sprintf("Sequence: %s\n\n", sequenceValue)

				if sequenceLength {
					length := len(record.Sequence)
					totalLength += length
					lengths = append(lengths, length)
					output += fmt.Sprintf("Sequence Length: %d\n", length)
				}

				if gcContent {
					gc := sequence.GCContent(record.Sequence)
					totalGCContent += gc
					output += fmt.Sprintf("GC Content: %.2f%%\n", gc)
				}

				if reverseComp {
					// Calculate reverse complement
					reverse := sequence.ReverseComplement(record.Sequence)
					output += fmt.Sprintf("Reverse Complement: %s\n", reverse)
				}

				if nucleotideFreq {
					// Calculate nucleotide frequency
					freq := sequence.NucleotideFrequency(record.Sequence)
					for nucleotide, count := range freq {
						totalNucleotideFreq[nucleotide] += count
					}
					output += "Nucleotide Frequency:\n"
					for nucleotide, count := range freq {
						output += fmt.Sprintf("%s: %.4f\n", string(nucleotide), count)
					}
				}
			}

			// Print aggregate statistics
			output += "\nSeqCraft Aggregate Statistic(s):\n"

			if sequenceLength {
				averageLength := float64(totalLength) / float64(len(records))
				output += fmt.Sprintf("\nTotal Sequence Length: %d\n", totalLength)
				output += fmt.Sprintf("Average Sequence Length: %.2f\n", averageLength)
				// Add min, max, and median length calculations
				minLength, maxLength, medianLength := calculateLengthStats(lengths)
				output += fmt.Sprintf("Min Sequence Length: %d\n", minLength)
				output += fmt.Sprintf("Max Sequence Length: %d\n", maxLength)
				output += fmt.Sprintf("Median Sequence Length: %.2f\n", medianLength)
			}

			if gcContent {
				averageGCContent := totalGCContent / float64(len(records))
				output += fmt.Sprintf("Average GC Content: %.2f%%\n", averageGCContent)
			}

			if nucleotideFreq {
				output += "Total Nucleotide Frequency:\n"
				for nucleotide, count := range totalNucleotideFreq {
					output += fmt.Sprintf("%s: %.4f\n", string(nucleotide), count)
				}
			}

			// output += "\n" // Add a newline at the end

			// Print or save output
			if outputFile != "" {
				err := os.WriteFile(outputFile, []byte(output), 0644)
				if err != nil {
					fmt.Printf("Error writing to file: %v\n", err)
				} else {
					fmt.Printf("Output written to %s\n", outputFile)
				}
			} else {
				fmt.Print(output)
			}
		},
	}

	// Add flags to the command
	parseCmd.Flags().BoolVarP(&sequenceLength, "length", "l", false, "Calculate sequence length")
	parseCmd.Flags().BoolVarP(&gcContent, "gc", "g", false, "Calculate GC content")
	parseCmd.Flags().BoolVarP(&reverseComp, "reverse", "r", false, "Calculate reverse complement")
	parseCmd.Flags().BoolVarP(&nucleotideFreq, "freq", "f", false, "Calculate nucleotide frequency")
	parseCmd.Flags().StringVarP(&outputFile, "output", "o", "", "Output file to save results")
	parseCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose mode")

	return parseCmd
}

// calculateLengthStats calculates the minimum, maximum, and median lengths from a slice of lengths.
func calculateLengthStats(lengths []int) (min, max int, median float64) {
	if len(lengths) == 0 {
		return 0, 0, 0
	}
	sort.Ints(lengths)
	min = lengths[0]
	max = lengths[len(lengths)-1]
	if len(lengths)%2 == 0 {
		median = float64(lengths[len(lengths)/2-1]+lengths[len(lengths)/2]) / 2
	} else {
		median = float64(lengths[len(lengths)/2])
	}
	return min, max, median
}
