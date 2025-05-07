package publictokens

import (
	"context"
	"log"

	"github.com/spf13/cobra"
	"github.com/stytchauth/stytch-cli/cmd/internal"
	"github.com/stytchauth/stytch-management-go/v2/pkg/models/publictokens"
)

// NewGetAllCommand returns a cobra command for listing public tokens
func NewGetAllCommand() *cobra.Command {
	var projectID string

	cmd := &cobra.Command{
		Use:   "get-all",
		Short: "Retrieve a list of public tokens",
		Long:  "Retrieve a list of public tokens for a project",
		Run: func(c *cobra.Command, args []string) {
			res, err := internal.MangoClient().PublicTokens.GetAll(
				context.Background(), publictokens.GetAllRequest{ProjectID: projectID},
			)
			if err != nil {
				log.Fatalf("Get all public tokens: %s", err)
			}

			internal.PrintJSON(res)
		},
	}

	cmd.Flags().StringVarP(&projectID, "project-id", "p", "", "The project ID")
	_ = cmd.MarkFlagRequired("project-id")
	return cmd
}
