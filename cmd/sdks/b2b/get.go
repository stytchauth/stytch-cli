package b2b

import (
	"context"
	"log"

	"github.com/spf13/cobra"
	"github.com/stytchauth/stytch-cli/cmd/internal"
	"github.com/stytchauth/stytch-management-go/v2/pkg/models/sdk"
)

func NewGetCommand() *cobra.Command {
	var projectID string

	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get B2B SDK configuration for a project",
		Run: func(cmd *cobra.Command, args []string) {
			cfgResp, err := internal.MangoClient().SDK.GetB2BConfig(context.Background(), sdk.GetB2BConfigRequest{
				ProjectID: projectID,
			})
			if err != nil {
				log.Fatalf("Unable to retrieve B2B SDK config: %v", err)
			}

			internal.PrintJSON(cfgResp)
		},
	}

	cmd.Flags().StringVarP(&projectID, "project-id", "p", "", "Project ID")
	_ = cmd.MarkFlagRequired("project-id")

	return cmd
}
