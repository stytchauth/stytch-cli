package project

import (
	"github.com/spf13/cobra"
)

func NewRootCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "project",
		Short: "Manage projects",
	}

	command.AddCommand(NewCreateCommand())

	return command
}
