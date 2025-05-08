package set

import (
	"github.com/spf13/cobra"
)

func NewRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set",
		Short: "Set B2B SDK configuration",
	}

	cmd.AddCommand(NewEnableCommand())
	cmd.AddCommand(NewDomainCommand())
	return cmd
}
