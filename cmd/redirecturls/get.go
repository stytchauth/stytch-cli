package redirecturls

import (
	"context"
	"log"

	"github.com/spf13/cobra"
	"github.com/stytchauth/stytch-cli/cmd/internal"
	"github.com/stytchauth/stytch-management-go/v2/pkg/models/redirecturls"
)

// NewGetCommand returns a cobra command for retrieving a redirect URL
func NewGetCommand() *cobra.Command {
	var projectID string
	var url string

	cmd := &cobra.Command{
		Use:   "get",
		Short: "Retrieve details about a redirect URL",
		Long:  "Retrieve details about a redirect URL for a project",
		Run: func(c *cobra.Command, args []string) {
			res, err := internal.MangoClient().RedirectURLs.Get(
				context.Background(), redirecturls.GetRequest{ProjectID: projectID, URL: url},
			)
			if err != nil {
				log.Fatalf("Get redirect URL: %s", err)
			}

			internal.PrintJSON(res)
		},
	}

	cmd.Flags().StringVarP(&projectID, "project-id", "p", "", "The project ID")
	cmd.Flags().StringVarP(&url, "url", "u", "", "The redirect URL")
	_ = cmd.MarkFlagRequired("project-id")
	_ = cmd.MarkFlagRequired("url")
	return cmd
}
