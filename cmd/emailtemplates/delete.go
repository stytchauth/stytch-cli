package emailtemplates

import (
   "context"
   "log"

   "github.com/spf13/cobra"
   "github.com/stytchauth/stytch-cli/cmd/internal"
   "github.com/stytchauth/stytch-management-go/v2/pkg/models/emailtemplates"
)

// NewDeleteCommand returns a cobra command for deleting an email template
func NewDeleteCommand() *cobra.Command {
   var projectID string
   var templateID string

   cmd := &cobra.Command{
       Use:   "delete",
       Short: "Delete an email template",
       Long:  "Delete an email template",
       Run: func(c *cobra.Command, args []string) {
           if projectID == "" || templateID == "" {
               log.Fatalf("Both --project-id and --template-id must be provided")
           }

           res, err := internal.GetDefaultMangoClient().EmailTemplates.Delete(context.Background(), emailtemplates.DeleteRequest{
               ProjectID:  projectID,
               TemplateID: templateID,
           })
           if err != nil {
               log.Fatalf("Delete email template: %s", err)
           }

           internal.PrintJSON(res)
       },
   }

   cmd.Flags().StringVarP(&projectID, "project-id", "p", "", "The project ID")
   cmd.Flags().StringVarP(&templateID, "template-id", "t", "", "The email template ID")

   return cmd
}