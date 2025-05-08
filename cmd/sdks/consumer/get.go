package consumer

import (
	"context"
	"log"

	"github.com/spf13/cobra"
	"github.com/stytchauth/stytch-cli/cmd/internal"
	"github.com/stytchauth/stytch-management-go/v2/pkg/models/sdk"
)

// NewCreateCommand returns a cobra command for creating a public token
func NewGetCommand() *cobra.Command {
	var projectID string

	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get SDK config for project",
		Run: func(c *cobra.Command, args []string) {
			res, err := internal.MangoClient().SDK.GetConsumerConfig(
				context.Background(), sdk.GetConsumerConfigRequest{ProjectID: projectID},
			)
			if err != nil {
				log.Fatalf("Enable sdks: %s", err)
			}

			internal.PrintJSON(res)
		},
	}

	cmd.Flags().StringVarP(&projectID, "project-id", "p", "", "The project ID")
	_ = cmd.MarkFlagRequired("project-id")
	return cmd
}
