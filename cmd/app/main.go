package main

import (
	"fmt"

	"github.com/joushvak17/Bioinformatics-CLI-Tool/internal/cli"
)

func main() {
	if err := cli.Execute(); err != nil {
		fmt.Println("Error:", err)
	}
}