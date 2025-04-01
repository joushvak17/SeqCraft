package main

import (
	"log/slog"

	"github.com/joushvak17/SeqCraft/internal/cli"
)

func main() {
	if err := cli.Execute(); err != nil {
		slog.Error("Error executing CLI command", "error", err)
	}
}
