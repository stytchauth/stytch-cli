package set

import (
	"context"
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/stytchauth/stytch-cli/cmd/internal"
	"github.com/stytchauth/stytch-management-go/v2/pkg/models/sdk"
)

func NewEnableCommand() *cobra.Command {
	var projectID string
	var enabled bool

	cmd := &cobra.Command{
		Use:   "enable",
		Short: "Enable consumer SDK configuration",
		Run: func(cmd *cobra.Command, args []string) {
			// Get current config
			cfgResp, err := internal.MangoClient().SDK.GetConsumerConfig(context.Background(), sdk.GetConsumerConfigRequest{
				ProjectID: projectID,
			})
			if err != nil {
				log.Fatalf("Unable to retrieve SDK config: %v", err)
			}

			// Update config
			updatedCfg := cfgResp.Config
			updatedCfg.Basic.Enabled = enabled

			fmt.Println("SDK configuration updated successfully")
		},
	}

	cmd.Flags().StringVarP(&projectID, "project-id", "p", "", "Project ID")
	cmd.Flags().BoolVarP(&enabled, "enabled", "e", true, "Enable/disable SDK")
	_ = cmd.MarkFlagRequired("project-id")
	_ = cmd.MarkFlagRequired("enabled")

	return cmd
}
