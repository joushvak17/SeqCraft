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

	var records []Record
	var currentRecord *Record

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue // Skip empty lines
		}

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

	if currentRecord != nil {
		records = append(records, *currentRecord)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return records, nil
}
