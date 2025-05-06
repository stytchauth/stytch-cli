package cmd

import (
	"github.com/spf13/cobra"

	project "github.com/stytchauth/stytch-cli/cmd/projects"
)

func NewRootCommand() *cobra.Command {
	// rootCmd represents the base command when called without any subcommands
	var command = &cobra.Command{
		Use:   "stytch-cli",
		Short: "A brief description of your application",
	}

	command.AddCommand(project.NewRootCommand())
	command.AddCommand(NewVersionCommand())
	command.AddCommand(NewAuthenticateCommand())

	return command
}
