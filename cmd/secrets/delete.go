package secrets

import (
	"context"
	"log"

	"github.com/spf13/cobra"
	"github.com/stytchauth/stytch-cli/cmd/internal"
	"github.com/stytchauth/stytch-management-go/v2/pkg/models/secrets"
)

func NewDeleteCommand() *cobra.Command {
	var projectID, secretID string

	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete a project secret",
		Long:  "Delete a project secret",
		Run: func(c *cobra.Command, args []string) {
			res, err := internal.MangoClient().Secrets.Delete(context.Background(), secrets.DeleteSecretRequest{
				ProjectID: projectID,
				SecretID:  secretID,
			})
			if err != nil {
				log.Fatalf("Delete secret: %s", err)
			}

			internal.PrintJSON(res)
		},
	}

	cmd.Flags().StringVarP(&projectID, "project-id", "p", "", "The project ID")
	cmd.Flags().StringVarP(&secretID, "secret-id", "s", "", "The secret ID")
	_ = cmd.MarkFlagRequired("project-id")
	_ = cmd.MarkFlagRequired("secret-id")
	return cmd
}
