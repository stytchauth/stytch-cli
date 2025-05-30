package main

import (
	"log/slog"
	"os"

	"github.com/stytchauth/stytch-cli/cmd"
)

func main() {
	err := cmd.NewRootCommand().Execute()
	if err != nil {
		slog.Error("Failed to execute command", "error", err)
		os.Exit(1)
	}
}
