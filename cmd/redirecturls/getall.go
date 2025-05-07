package redirecturls

import (
   "context"
   "log"

   "github.com/spf13/cobra"
   "github.com/stytchauth/stytch-cli/cmd/internal"
   "github.com/stytchauth/stytch-management-go/v2/pkg/models/redirecturls"
)

// NewGetAllCommand returns a cobra command for listing redirect URLs
func NewGetAllCommand() *cobra.Command {
   var projectID string

   cmd := &cobra.Command{
       Use:   "get-all",
       Short: "Retrieve a list of redirect URLs",
       Long:  "Retrieve a list of redirect URLs for a project",
       Run: func(c *cobra.Command, args []string) {
           if projectID == "" {
               log.Fatalf("Missing --project-id")
           }

           res, err := internal.GetDefaultMangoClient().RedirectURLs.GetAll(
               context.Background(), redirecturls.GetAllRequest{ProjectID: projectID},
           )
           if err != nil {
               log.Fatalf("Get all redirect URLs: %s", err)
           }

           internal.PrintJSON(res)
       },
   }

   cmd.Flags().StringVarP(&projectID, "project-id", "p", "", "The project ID")

   return cmd
}