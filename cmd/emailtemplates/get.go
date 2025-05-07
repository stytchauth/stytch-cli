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
			res, err := internal.MangoClient().EmailTemplates.Get(context.Background(), emailtemplates.GetRequest{
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
	var errors []error
	errors = append(errors, cmd.MarkFlagRequired("project-id"))
	errors = append(errors, cmd.MarkFlagRequired("template-id"))
	if len(errors) > 0 {
		for _, err := range errors {
			log.Fatalf("Error marking flag required: %v", err)
		}
	}
	return cmd
}
