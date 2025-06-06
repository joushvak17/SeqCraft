package parse

import (
	"fmt"
	"log/slog"
	"os"
	"strings"
	"time"

	"github.com/AlecAivazis/survey/v2"
	"github.com/briandowns/spinner"
	"github.com/joushvak17/SeqCraft/pkg/parse"
	"github.com/joushvak17/SeqCraft/pkg/sequence"
	"github.com/joushvak17/SeqCraft/pkg/utils"
	"github.com/spf13/cobra"
)

// NewParseCmd creates and returns the `parse` command.
func NewParseCmd() *cobra.Command {
	var (
		sequenceLength bool
		gcContent      bool
		reverseComp    bool
		nucleotideFreq bool
		interactive    bool
		outputFile     string
	)

	parseCmd := &cobra.Command{
		Use:   "parse <file>",
		Short: "Parse and analyze a FASTA file",
		Long:  "Parse and analyze a FASTA file, including sequence length, GC content, reverse complement, and nucleotide frequency.",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
				Level: slog.LevelInfo,
			}))
			slog.SetDefault(logger)

			// Start the timer
			startTime := time.Now()

			filename := args[0]

			// Input validation, check if the file exists and is a valid FASTA file
			if _, err := os.Stat(filename); os.IsNotExist(err) {
				slog.Error("File does not exist", slog.String("filename", filename))
				return
			}
			if !strings.HasSuffix(filename, ".fasta") && !strings.HasSuffix(filename, ".fa") {
				slog.Error("File is not a valid FASTA file", slog.String("filename", filename))
				return
			}

			// Interactive mode
			if interactive {
				options := []string{"Sequence Length", "GC Content", "Reverse Complement", "Nucleotide Frequency"}
				var selectedOptions []string
				prompt := &survey.MultiSelect{
					Message: "Select analysis options:",
					Options: options,
				}
				err := survey.AskOne(prompt, &selectedOptions)
				if err != nil {
					slog.Error("Prompt failed", slog.String("error", err.Error()))
					return
				}
				for _, option := range selectedOptions {
					switch option {
					case "Sequence Length":
						sequenceLength = true
					case "GC Content":
						gcContent = true
					case "Reverse Complement":
						reverseComp = true
					case "Nucleotide Frequency":
						nucleotideFreq = true
					}
				}
				var outputFilePrompt string
				outputPrompt := &survey.Input{
					Message: "Enter output file (leave blank for no output file):",
				}
				err = survey.AskOne(outputPrompt, &outputFilePrompt)
				if err != nil {
					slog.Error("Prompt failed", slog.String("error", err.Error()))
					return
				}
				if outputFilePrompt != "" {
					outputFile = outputFilePrompt
				}
			}

			// Start the spinner
			s := spinner.New(spinner.CharSets[11], 100*time.Millisecond)
			s.Suffix = " Processing ... "
			s.Start()

			// Prepare the output string
			output := "\nSeqCraft Parse Output(s):\n"

			// Parse the FASTA file
			records, err := parse.Parse(filename)
			if err != nil {
				slog.Error("Error parsing FASTA file", slog.String("filename", filename), slog.String("error", err.Error()))
				s.Stop()
				return
			}

			// Initialize variables for aggregate statistics
			totalLength := 0
			totalGCContent := 0.0
			totalNucleotideFreq := make(map[rune]float64)
			var lengths []int

			for _, record := range records {
				// Print record information
				output += fmt.Sprintf("\nID: %s\n", record.ID)
				output += fmt.Sprintf("Description: %s\n", record.Description)
				output += fmt.Sprintf("Sequence: %s\n", record.Sequence)

				// Sequence length
				if sequenceLength {
					length := len(record.Sequence)
					totalLength += length
					lengths = append(lengths, length)
					output += fmt.Sprintf("Sequence Length: %d\n", length)
				}

				// GC content
				if gcContent {
					gc := sequence.GCContent(record.Sequence)
					totalGCContent += gc
					output += fmt.Sprintf("GC Content: %.2f%%\n", gc)
				}

				// Reverse complement
				if reverseComp {
					reverse := sequence.ReverseComplement(record.Sequence)
					output += fmt.Sprintf("Reverse Complement: %s\n", reverse)
				}

				// Nucleotide frequency
				if nucleotideFreq {
					freq := sequence.NucleotideFrequency(record.Sequence)
					output += "Nucleotide Frequency:\n"
					for nucleotide, count := range freq {
						totalNucleotideFreq[nucleotide] += count
						output += fmt.Sprintf("%s: %.4f\n", string(nucleotide), count)
					}
				}
			}

			// Print aggregate statistics
			output += "\nSeqCraft Aggregate Statistics:\n"

			// Sequence length
			if sequenceLength {
				averageLength := float64(totalLength) / float64(len(records))
				output += fmt.Sprintf("Total Sequence Length: %d\n", totalLength)
				output += fmt.Sprintf("Average Sequence Length: %.2f\n", averageLength)
				minLength, maxLength, medianLength := utils.CalculateLengthStats(lengths)
				output += fmt.Sprintf("Min Sequence Length: %d\n", minLength)
				output += fmt.Sprintf("Max Sequence Length: %d\n", maxLength)
				output += fmt.Sprintf("Median Sequence Length: %.2f\n", medianLength)
			}

			// GC content
			if gcContent {
				averageGCContent := totalGCContent / float64(len(records))
				output += fmt.Sprintf("Average GC Content: %.2f%%\n", averageGCContent)
			}

			// Nucleotide frequency
			if nucleotideFreq {
				output += "Total Nucleotide Frequency:\n"
				for nucleotide, count := range totalNucleotideFreq {
					output += fmt.Sprintf("%s: %.4f\n", string(nucleotide), count)
				}
			}

			// Get the current date/time and create the output string
			currentTime := time.Now().Format("January 2, 2006 3:04 PM")
			output += fmt.Sprintf("File: %s\n", filename)
			output += fmt.Sprintf("Date and Time: %s\n", currentTime)
			output += fmt.Sprintf("Number of Records Parsed: %d\n", len(records))

			// Print the time taken for processing
			elapsedTime := time.Since(startTime)
			output += fmt.Sprintf("Time taken for processing: %s\n", elapsedTime)

			// Print or save output
			if outputFile != "" {
				err := os.WriteFile(outputFile, []byte(output), 0644)
				if err != nil {
					slog.Error("Error writing to file", slog.String("outputFile", outputFile), slog.String("error", err.Error()))
					s.Stop()
				} else {
					fmt.Printf("Output written to %s\n", outputFile)
					s.Stop()
				}
			} else {
				fmt.Println(output)
				s.Stop()
			}
		},
	}

	// Add flags to the command
	parseCmd.Flags().BoolVarP(&sequenceLength, "length", "l", false, "Calculate sequence length")
	parseCmd.Flags().BoolVarP(&gcContent, "gc", "g", false, "Calculate GC content")
	parseCmd.Flags().BoolVarP(&reverseComp, "reverse", "r", false, "Calculate reverse complement")
	parseCmd.Flags().BoolVarP(&nucleotideFreq, "freq", "f", false, "Calculate nucleotide frequency")
	parseCmd.Flags().BoolVarP(&interactive, "interactive", "i", false, "Enable interactive mode")
	parseCmd.Flags().StringVarP(&outputFile, "output", "o", "", "Output file to save results")

	return parseCmd
}
