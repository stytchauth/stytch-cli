package emailtemplates

import (
	"context"
	"log"

	"github.com/spf13/cobra"
	"github.com/stytchauth/stytch-cli/cmd/internal"
	"github.com/stytchauth/stytch-management-go/v2/pkg/models/emailtemplates"
)

// NewGetCommand returns a cobra command for retrieving an email template
func NewGetCommand() *cobra.Command {
	var projectID string
	var templateID string

	cmd := &cobra.Command{
		Use:   "get",
		Short: "Retrieve details about an email template",
		Long:  "Retrieve details about an email template",
		Run: func(c *cobra.Command, args []string) {
			if projectID == "" || templateID == "" {
				log.Fatalf("Both --project-id and --template-id must be provided")
			}

			res, err := internal.GetDefaultMangoClient().EmailTemplates.Get(context.Background(), emailtemplates.GetRequest{
				ProjectID:  projectID,
				TemplateID: templateID,
			})
			if err != nil {
				log.Fatalf("Get email template: %s", err)
			}

			internal.PrintJSON(res)
		},
	}

	cmd.Flags().StringVarP(&projectID, "project-id", "p", "", "The project ID")
	cmd.Flags().StringVarP(&templateID, "template-id", "t", "", "The email template ID")

	return cmd
}
