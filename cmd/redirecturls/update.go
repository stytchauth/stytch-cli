package redirecturls

import (
	"context"
	"log"

	"github.com/spf13/cobra"
	"github.com/stytchauth/stytch-cli/cmd/internal"
	"github.com/stytchauth/stytch-management-go/v2/pkg/models/redirecturls"
)

// NewUpdateCommand returns a cobra command for updating a redirect URL
func NewUpdateCommand() *cobra.Command {
	var projectID string
	var url string
	var redirectType string

	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update a redirect URL",
		Long:  "Update a redirect URL for a project",
		Run: func(c *cobra.Command, args []string) {
			if projectID == "" || url == "" {
				log.Fatalf("Both --project-id and --url must be provided")
			}
			if redirectType == "" {
				log.Fatalf("Missing --redirect-type for update")
			}

			// Build the request
			req := redirecturls.UpdateRequest{
				ProjectID: projectID,
				RedirectURL: redirecturls.RedirectURL{
					URL: url,
					ValidTypes: []redirecturls.URLRedirectType{
						{Type: redirecturls.RedirectType(redirectType), IsDefault: true},
					},
				},
			}
			res, err := internal.GetDefaultMangoClient().RedirectURLs.Update(
				context.Background(), req,
			)
			if err != nil {
				log.Fatalf("Update redirect URL: %s", err)
			}

			internal.PrintJSON(res)
		},
	}

	cmd.Flags().StringVarP(&projectID, "project-id", "p", "", "The project ID")
	cmd.Flags().StringVarP(&url, "url", "u", "", "The redirect URL to update")
	cmd.Flags().StringVarP(&redirectType, "redirect-type", "t", "", "The new redirect type (e.g., LOGIN, SIGNUP)")

	return cmd
}
