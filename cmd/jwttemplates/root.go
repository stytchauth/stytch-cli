package jwttemplates

import (
	"github.com/spf13/cobra"
)

// NewRootCommand creates the root command for JWT template operations
func NewRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "jwttemplates",
		Short: "Manage project JWT templates",
		Long:  "Manage project JWT templates",
	}

	cmd.AddCommand(NewGetCommand())
	cmd.AddCommand(NewSetCommand())

	return cmd
}
