package emailtemplates

import (
   "github.com/spf13/cobra"
)

// NewRootCommand creates the root command for emailtemplates
func NewRootCommand() *cobra.Command {
   cmd := &cobra.Command{
       Use:   "emailtemplates",
       Short: "Manage project email templates",
       Long:  "Manage project email templates",
   }

   cmd.AddCommand(NewGetCommand())
   cmd.AddCommand(NewGetAllCommand())
   cmd.AddCommand(NewCreateCommand())
   cmd.AddCommand(NewUpdateCommand())
   cmd.AddCommand(NewDeleteCommand())

   return cmd
}