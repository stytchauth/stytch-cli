package publictokens

import (
   "context"
   "log"

   "github.com/spf13/cobra"
   "github.com/stytchauth/stytch-cli/cmd/internal"
   "github.com/stytchauth/stytch-management-go/v2/pkg/models/publictokens"
)

// NewCreateCommand returns a cobra command for creating a public token
func NewCreateCommand() *cobra.Command {
   var projectID string

   cmd := &cobra.Command{
       Use:   "create",
       Short: "Create a new public token",
       Long:  "Create a new public token for a project",
       Run: func(c *cobra.Command, args []string) {
           if projectID == "" {
               log.Fatalf("Missing --project-id")
           }

           res, err := internal.GetDefaultMangoClient().PublicTokens.Create(
               context.Background(), publictokens.CreateRequest{ProjectID: projectID},
           )
           if err != nil {
               log.Fatalf("Create public token: %s", err)
           }

           internal.PrintJSON(res)
       },
   }

   cmd.Flags().StringVarP(&projectID, "project-id", "p", "", "The project ID")

   return cmd
}