package parse

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

// Record represents a single FASTA record.
type Record struct {
	ID          string // Sequence ID (text after '>')
	Description string // Optional description (rest of the header line)
	Sequence    string // Sequence data
}

// Parse reads a FASTA file and returns a slice of Record.
func Parse(filename string) ([]Record, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Create a slice to store the records and a pointer to the current record
	var records []Record
	var currentRecord *Record

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Remove leading and trailing whitespaces
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue // Skip empty lines
		}

		// Check if the line is a header
		if line[0] == '>' {
			// Append the processed record to the slice if it exists
			if currentRecord != nil {
				records = append(records, *currentRecord)
			}
			// Create a new record
			currentRecord = &Record{}
			header := strings.TrimSpace(line[1:]) // Remove '>'
			// Split ID and description
			fields := strings.SplitN(header, " ", 2)
			currentRecord.ID = fields[0]
			if len(fields) > 1 {
				currentRecord.Description = fields[1]
			}
		} else if currentRecord != nil {
			// Append sequence data, since it is not a header
			currentRecord.Sequence += strings.TrimSpace(line)
		} else {
			return nil, errors.New("Invalid FASTA format: Sequence data without header")
		}
	}

	// Add the last record, since we exit the loop without appending it
	if currentRecord != nil {
		records = append(records, *currentRecord)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return records, nil
}
