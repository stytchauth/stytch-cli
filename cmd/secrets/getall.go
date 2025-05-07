package secrets

import (
	"context"
	"log"

	"github.com/spf13/cobra"
	"github.com/stytchauth/stytch-cli/cmd/internal"
	"github.com/stytchauth/stytch-management-go/v2/pkg/models/secrets"
)

func NewGetAllCommand() *cobra.Command {
	var projectID string

	cmd := &cobra.Command{
		Use:   "get-all",
		Short: "Retrieve a list of project secrets",
		Long:  "Retrieve a list of project secrets",
		Run: func(c *cobra.Command, args []string) {
			if projectID == "" {
				log.Fatalf("Missing --project-id")
			}

			res, err := internal.MangoClient().Secrets.GetAll(context.Background(), secrets.GetAllSecretsRequest{
				ProjectID: projectID,
			})
			if err != nil {
				log.Fatalf("Get all secrets: %s", err)
			}

			internal.PrintJSON(res)
		},
	}

	cmd.Flags().StringVar(&projectID, "project-id", "", "The project ID")

	return cmd
}
