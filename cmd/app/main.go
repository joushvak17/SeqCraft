package main

import (
	"log/slog"
	"os"

	"github.com/joushvak17/SeqCraft/internal/cli"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))
	slog.SetDefault(logger)

	if err := cli.Execute(); err != nil {
		slog.Error("Error executing CLI command", "error", err)
		os.Exit(1)
	}
}
