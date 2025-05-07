package redirecturls

import (
	"context"
	"log"

	"github.com/spf13/cobra"
	"github.com/stytchauth/stytch-cli/cmd/internal"
	"github.com/stytchauth/stytch-management-go/v2/pkg/models/redirecturls"
)

// NewCreateCommand returns a cobra command for creating a redirect URL
func NewCreateCommand() *cobra.Command {
	var projectID string
	var url string
	var redirectType string
	var isDefault bool

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a new redirect URL",
		Long:  "Create a new redirect URL for a project",
		Run: func(c *cobra.Command, args []string) {
			// Build the request
			req := redirecturls.CreateRequest{
				ProjectID: projectID,
				RedirectURL: redirecturls.RedirectURL{
					URL: url,
					ValidTypes: []redirecturls.URLRedirectType{
						{Type: redirecturls.RedirectType(redirectType), IsDefault: isDefault},
					},
				},
			}
			res, err := internal.GetDefaultMangoClient().RedirectURLs.Create(
				context.Background(), req,
			)
			if err != nil {
				log.Fatalf("Create redirect URL: %s", err)
			}

			internal.PrintJSON(res)
		},
	}

	cmd.Flags().StringVarP(&projectID, "project-id", "p", "", "The project ID")
	cmd.Flags().StringVarP(&url, "url", "u", "", "The redirect URL to create")
	cmd.Flags().StringVarP(&redirectType, "redirect-type", "t", "", "The redirect type (e.g., LOGIN, SIGNUP)")
	cmd.Flags().BoolVarP(&isDefault, "is-default", "d", false, "Whether to set this URL as the default for the given redirect type")
	cmd.MarkFlagRequired("project-id")
	cmd.MarkFlagRequired("url")
	cmd.MarkFlagRequired("redirect-type")

	return cmd
}
