package publictokens

import (
   "github.com/spf13/cobra"
)

// NewRootCommand creates the root command for public tokens operations
func NewRootCommand() *cobra.Command {
   cmd := &cobra.Command{
       Use:   "publictokens",
       Short: "Manage project public tokens",
       Long:  "Manage project public tokens",
   }

   cmd.AddCommand(NewGetAllCommand())
   cmd.AddCommand(NewCreateCommand())
   cmd.AddCommand(NewDeleteCommand())

   return cmd
}