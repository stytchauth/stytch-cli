package publictokens

import (
	"context"
	"log"

	"github.com/spf13/cobra"
	"github.com/stytchauth/stytch-cli/cmd/internal"
	"github.com/stytchauth/stytch-management-go/v2/pkg/models/publictokens"
)

// NewDeleteCommand returns a cobra command for deleting a public token
func NewDeleteCommand() *cobra.Command {
	var projectID string
	var token string

	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete a public token",
		Long:  "Delete a public token for a project",
		Run: func(c *cobra.Command, args []string) {
			res, err := internal.MangoClient().PublicTokens.Delete(
				context.Background(), publictokens.DeleteRequest{ProjectID: projectID, PublicToken: token},
			)
			if err != nil {
				log.Fatalf("Delete public token: %s", err)
			}

			internal.PrintJSON(res)
		},
	}

	cmd.Flags().StringVarP(&projectID, "project-id", "p", "", "The project ID")
	cmd.Flags().StringVarP(&token, "public-token", "t", "", "The public token to delete")
	var errors []error
	errors = append(errors, cmd.MarkFlagRequired("project-id"))
	errors = append(errors, cmd.MarkFlagRequired("public-token"))
	if len(errors) > 0 {
		for _, err := range errors {
			log.Fatalf("Error marking flag required: %v", err)
		}
	}

	return cmd
}
