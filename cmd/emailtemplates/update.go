package emailtemplates

import (
	"context"
	"log"

	"github.com/spf13/cobra"
	"github.com/stytchauth/stytch-cli/cmd/internal"
	"github.com/stytchauth/stytch-management-go/v2/pkg/models/emailtemplates"
)

// NewUpdateCommand returns a cobra command for updating an email template
func NewUpdateCommand() *cobra.Command {
	var projectID string
	var templateID string
	var name string

	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update an email template",
		Long:  "Update an email template",
		Run: func(c *cobra.Command, args []string) {
			res, err := internal.MangoClient().EmailTemplates.Update(context.Background(), emailtemplates.UpdateRequest{
				ProjectID: projectID,
				EmailTemplate: emailtemplates.EmailTemplate{
					TemplateID: templateID,
					Name:       &name,
				},
			})
			if err != nil {
				log.Fatalf("Update email template: %s", err)
			}

			internal.PrintJSON(res)
		},
	}

	cmd.Flags().StringVarP(&projectID, "project-id", "p", "", "The project ID")
	cmd.Flags().StringVarP(&templateID, "template-id", "t", "", "The email template ID")
	cmd.Flags().StringVarP(&name, "name", "n", "", "The new name of the email template")
	_ = cmd.MarkFlagRequired("project-id")
	_ = cmd.MarkFlagRequired("template-id")
	_ = cmd.MarkFlagRequired("name")
	return cmd
}
