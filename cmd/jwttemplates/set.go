package jwttemplates

import (
	"context"
	"log"

	"github.com/spf13/cobra"
	"github.com/stytchauth/stytch-cli/cmd/internal"
	"github.com/stytchauth/stytch-management-go/v2/pkg/models/jwttemplates"
)

// NewSetCommand returns a cobra command for setting a JWT template
func NewSetCommand() *cobra.Command {
	var projectID string
	var templateType string
	var templateContent string
	var customAudience string

	cmd := &cobra.Command{
		Use:   "set",
		Short: "Set a JWT template",
		Long:  "Set a JWT template",
		Run: func(c *cobra.Command, args []string) {
			req := &jwttemplates.SetRequest{
				ProjectID: projectID,
				JWTTemplate: jwttemplates.JWTTemplate{
					TemplateType:    jwttemplates.TemplateType(templateType),
					TemplateContent: templateContent,
					CustomAudience:  customAudience,
				},
			}
			res, err := internal.MangoClient().JWTTemplates.Set(context.Background(), req)
			if err != nil {
				log.Fatalf("Set JWT template: %s", err)
			}

			internal.PrintJSON(res)
		},
	}

	cmd.Flags().StringVarP(&projectID, "project-id", "p", "", "The project ID")
	cmd.Flags().StringVarP(&templateType, "template-type", "t", "", "The JWT template type (e.g., SESSION or M2M)")
	cmd.Flags().StringVarP(&templateContent, "template-content", "c", "", "The JWT template content")
	cmd.Flags().StringVarP(&customAudience, "custom-audience", "a", "", "The custom audience for the JWT template (optional)")

	_ = cmd.MarkFlagRequired("project-id")
	_ = cmd.MarkFlagRequired("template-type")
	_ = cmd.MarkFlagRequired("template-content")
	return cmd
}
