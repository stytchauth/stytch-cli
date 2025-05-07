package emailtemplates

import (
	"context"
	"log"

	"github.com/spf13/cobra"
	"github.com/stytchauth/stytch-cli/cmd/internal"
	"github.com/stytchauth/stytch-management-go/v2/pkg/models/emailtemplates"
)

// NewCreateCommand returns a cobra command for creating an email template
func NewCreateCommand() *cobra.Command {
	var projectID string
	var templateID string
	var name string

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a new email template",
		Long:  "Create a new email template",
		Run: func(c *cobra.Command, args []string) {
			res, err := internal.MangoClient().EmailTemplates.Create(context.Background(), emailtemplates.CreateRequest{
				ProjectID: projectID,
				EmailTemplate: emailtemplates.EmailTemplate{
					TemplateID: templateID,
					Name:       &name,
				},
			})
			if err != nil {
				log.Fatalf("Create email template: %s", err)
			}

			internal.PrintJSON(res)
		},
	}

	cmd.Flags().StringVarP(&projectID, "project-id", "p", "", "The project ID")
	cmd.Flags().StringVarP(&templateID, "template-id", "t", "", "The email template ID")
	cmd.Flags().StringVarP(&name, "name", "n", "", "The name of the email template")
	var errors []error
	errors = append(errors, cmd.MarkFlagRequired("project-id"))
	errors = append(errors, cmd.MarkFlagRequired("template-id"))
	errors = append(errors, cmd.MarkFlagRequired("name"))
	if len(errors) > 0 {
		for _, err := range errors {
			log.Fatalf("Error marking flag required: %v", err)
		}
	}
	return cmd
}
