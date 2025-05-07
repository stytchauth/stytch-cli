package redirecturls

import (
	"context"
	"log"

	"github.com/spf13/cobra"
	"github.com/stytchauth/stytch-cli/cmd/internal"
	"github.com/stytchauth/stytch-management-go/v2/pkg/models/redirecturls"
)

// NewDeleteCommand returns a cobra command for deleting a redirect URL
func NewDeleteCommand() *cobra.Command {
	var projectID string
	var url string

	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete a redirect URL",
		Long:  "Delete a redirect URL for a project",
		Run: func(c *cobra.Command, args []string) {
			res, err := internal.GetDefaultMangoClient().RedirectURLs.Delete(
				context.Background(), redirecturls.DeleteRequest{ProjectID: projectID, URL: url},
			)
			if err != nil {
				log.Fatalf("Delete redirect URL: %s", err)
			}

			internal.PrintJSON(res)
		},
	}

	cmd.Flags().StringVarP(&projectID, "project-id", "p", "", "The project ID")
	cmd.Flags().StringVarP(&url, "url", "u", "", "The redirect URL to delete")
	cmd.MarkFlagRequired("project-id")
	cmd.MarkFlagRequired("url")
	return cmd
}
