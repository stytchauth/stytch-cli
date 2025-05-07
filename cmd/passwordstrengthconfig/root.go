package passwordstrengthconfig

import (
	"github.com/spf13/cobra"
)

// NewRootCommand creates the root command for password strength configuration operations
func NewRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "passwordstrengthconfig",
		Short: "Manage project password strength configuration",
		Long:  "Manage project password strength configuration",
	}

	cmd.AddCommand(NewGetCommand())
	cmd.AddCommand(NewSetCommand())

	return cmd
}
