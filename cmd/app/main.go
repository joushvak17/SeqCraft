// Package main is the entry point for the SeqCraft application.
// It initializes the CLI and executes the root command.
// The application is designed for bioinformatics tasks, particularly for FASTA file analysis.
//
// For more information, visit: https://github.com/joushvak17/SeqCraft
// Copyright 2025 Joushva Kamble
// Licensed under the MIT License (see LICENSE file)
package main

import (
	"log/slog"
	"os"

	"github.com/joushvak17/SeqCraft/internal/cli"
)

func main() {
	// TODO: Figure out additional setup for the CLI, such as configuration files or environment variables.

	// FIXME: Figure out how to handle errors that can occur during command execution.
	// FIXME: Add more detailed logging for debugging purposes.
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))
	slog.SetDefault(logger)

	if err := cli.Execute(); err != nil {
		slog.Error("Error executing CLI command", "error", err)
		os.Exit(1)
	}
}
