package cmd

import (
	"log"
	project "stytch-cli/cmd/projects"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

func NewRootCommand() *cobra.Command {
	// rootCmd represents the base command when called without any subcommands
	var command = &cobra.Command{
		Use:   "stytch-cli",
		Short: "A brief description of your application",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			err := godotenv.Load()
			if err != nil {
				log.Fatal("Error loading .env file")
			}
		},
	}

	command.AddCommand(project.NewRootCommand())
	command.AddCommand(NewVersionCommand())
	command.AddCommand(NewAuthenticateCommand())

	return command
}

func init() {
	// TODO: Add flags
}
