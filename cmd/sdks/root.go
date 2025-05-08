package sdks

import (
	"github.com/spf13/cobra"
)

// NewRootCommand creates the root command for public tokens operations
func NewRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sdks",
		Short: "SDKs",
		Long:  "SDKs",
	}

	cmd.AddCommand(NewGetCommand())
	cmd.AddCommand(NewSetCommand())
	
	return cmd
}
