package redirecturls

import (
   "github.com/spf13/cobra"
)

// NewRootCommand creates the root command for redirect URL operations
func NewRootCommand() *cobra.Command {
   cmd := &cobra.Command{
       Use:   "redirecturls",
       Short: "Manage project redirect URLs",
       Long:  "Manage project redirect URLs",
   }

   cmd.AddCommand(NewGetCommand())
   cmd.AddCommand(NewGetAllCommand())
   cmd.AddCommand(NewCreateCommand())
   cmd.AddCommand(NewUpdateCommand())
   cmd.AddCommand(NewDeleteCommand())

   return cmd
}