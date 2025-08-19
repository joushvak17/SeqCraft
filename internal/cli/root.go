// Package cli provides the root command and subcommands for the SeqCraft CLI tool.
//
// For more information, visit: https://github.com/joushvak17/SeqCraft
// Copyright 2025 Joushva Kamble
// Licensed under the MIT License (see LICENSE file)
package cli

import (
	"github.com/joushvak17/SeqCraft/internal/parse"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "seqcraft",
	Short: "SeqCraft - Bioinformatics CLI Tool written in Go",
	// FIXME: Provide a more detailed description of the SeqCraft CLI tool.
	Long:  "SeqCraft - Bioinformatics CLI tool written in Go that provides efficient features for FASTA file analysis, including sequence parsing, alignment, and structure prediction.",
}

// Execute runs the root command of the SeqCraft CLI tool.
func Execute() error {
	return rootCmd.Execute()
}

// Initialize the root command and add subcommands.
func init() {
	rootCmd.AddCommand(parse.NewParseCmd())
}
