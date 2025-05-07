package secrets

import (
	"context"
	"log"

	"github.com/spf13/cobra"
	"github.com/stytchauth/stytch-management-go/v2/pkg/models/secrets"

	"github.com/stytchauth/stytch-cli/cmd/internal"
)

func NewCreateCommand() *cobra.Command {
	var projectID string

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a new project secret",
		Long:  "Create a new project secret",
		Run: func(c *cobra.Command, args []string) {
			res, err := internal.GetDefaultMangoClient().Secrets.Create(context.Background(), secrets.CreateSecretRequest{
				ProjectID: projectID,
			})
			if err != nil {
				log.Fatalf("Create secret: %s", err)
			}

			internal.PrintJSON(res)
		},
	}

	cmd.Flags().StringVarP(&projectID, "project-id", "p", "", "The project ID")
	cmd.MarkFlagRequired("project-id")
	return cmd
}
