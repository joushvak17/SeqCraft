package parse

import (
	"bufio"   // Buffered I/O operations
	"errors"  // Error handling
	"os"      // OS functionality
	"strings" // String manipulation
)

// Record represents a single FASTA record
type Record struct {
	ID          string // Sequence ID (text after '>')
	Description string // Optional description (rest of the header line)
	Sequence    string // Sequence data
}

// Parse reads a FASTA file and returns a slice of Record
func Parse(filename string) ([]Record, error) {
	// Open the file
	file, err := os.Open(filename)
	// Handle errors
	if err != nil {
		return nil, err
	}
	// Close the file when the function returns
	defer file.Close()

	// Create a slice to store the records and a pointer to the current record
	var records []Record
	var currentRecord *Record

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Remove leading and trailing whitespaces
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue // Skip empty lines
		}

		// Check if the line is a header
		if line[0] == '>' {
			// New record
			if currentRecord != nil {
				records = append(records, *currentRecord)
			}
			currentRecord = &Record{}
			header := strings.TrimSpace(line[1:]) // Remove '>'
			// Split ID and description
			fields := strings.SplitN(header, " ", 2)
			currentRecord.ID = fields[0]
			if len(fields) > 1 {
				currentRecord.Description = fields[1]
			}
		} else if currentRecord != nil {
			// Append sequence data
			currentRecord.Sequence += line
		} else {
			return nil, errors.New("Invalid FASTA format: sequence data without header")
		}
	}

	// Add the last record
	if currentRecord != nil {
		records = append(records, *currentRecord)
	}

	// Check for errors during the scan
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return records, nil
}
