package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/stytchauth/stytch-cli/cmd/internal"
)

func NewVersionCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the version number",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("stytch-cli v" + internal.Version)
		},
	}
}
