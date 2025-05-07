package emailtemplates

import (
	"context"
	"log"

	"github.com/spf13/cobra"
	"github.com/stytchauth/stytch-cli/cmd/internal"
	"github.com/stytchauth/stytch-management-go/v2/pkg/models/emailtemplates"
)

// NewGetAllCommand returns a cobra command for listing email templates
func NewGetAllCommand() *cobra.Command {
	var projectID string

	cmd := &cobra.Command{
		Use:   "get-all",
		Short: "Retrieve a list of email templates",
		Long:  "Retrieve a list of email templates",
		Run: func(c *cobra.Command, args []string) {
			if projectID == "" {
				log.Fatalf("Missing --project-id")
			}

			res, err := internal.GetDefaultMangoClient().EmailTemplates.GetAll(context.Background(), emailtemplates.GetAllRequest{
				ProjectID: projectID,
			})
			if err != nil {
				log.Fatalf("Get all email templates: %s", err)
			}

			internal.PrintJSON(res)
		},
	}

	cmd.Flags().StringVarP(&projectID, "project-id", "p", "", "The project ID")

	return cmd
}
