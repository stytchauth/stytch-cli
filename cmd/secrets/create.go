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
			if projectID == "" {
				log.Fatalf("Missing --project-id")
			}

			res, err := internal.GetDefaultMangoClient().Secrets.Create(context.Background(), secrets.CreateSecretRequest{
				ProjectID: projectID,
			})
			if err != nil {
				log.Fatalf("Create secret: %s", err)
			}

			internal.PrintJSON(res)
		},
	}

	cmd.Flags().StringVar(&projectID, "project-id", "", "The project ID")

	return cmd
}
