package secrets

import (
	"github.com/spf13/cobra"
)

func NewRootCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "secrets",
		Short: "Manage project secrets",
		Long:  "Manage project secrets",
	}

	command.AddCommand(NewGetCommand())
	command.AddCommand(NewGetAllCommand())
	command.AddCommand(NewCreateCommand())
	command.AddCommand(NewDeleteCommand())

	return command
}
