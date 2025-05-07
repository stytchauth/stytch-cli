package jwttemplates

import (
	"context"
	"log"

	"github.com/spf13/cobra"
	"github.com/stytchauth/stytch-cli/cmd/internal"
	"github.com/stytchauth/stytch-management-go/v2/pkg/models/jwttemplates"
)

// NewGetCommand returns a cobra command for retrieving a JWT template
func NewGetCommand() *cobra.Command {
	var projectID string
	var templateType string

	cmd := &cobra.Command{
		Use:   "get",
		Short: "Retrieve details about a JWT template",
		Long:  "Retrieve details about a JWT template",
		Run: func(c *cobra.Command, args []string) {
			if projectID == "" || templateType == "" {
				log.Fatalf("Both --project-id and --template-type must be provided")
			}

			req := &jwttemplates.GetRequest{
				ProjectID:    projectID,
				TemplateType: jwttemplates.TemplateType(templateType),
			}
			res, err := internal.GetDefaultMangoClient().JWTTemplates.Get(context.Background(), req)
			if err != nil {
				log.Fatalf("Get JWT template: %s", err)
			}

			internal.PrintJSON(res)
		},
	}

	cmd.Flags().StringVarP(&projectID, "project-id", "p", "", "The project ID")
	cmd.Flags().StringVarP(&templateType, "template-type", "t", "", "The JWT template type (e.g., SESSION or M2M)")

	return cmd
}
