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
	var enableAll bool
	cmd := &cobra.Command{
		Use:   "enable",
		Short: "Enable B2B SDK configuration",
		Run: func(cmd *cobra.Command, args []string) {
			// Get current config
			cfgResp, err := internal.MangoClient().SDK.GetB2BConfig(context.Background(), sdk.GetB2BConfigRequest{
				ProjectID: projectID,
			})
			if err != nil {
				log.Fatalf("Unable to retrieve SDK config: %v", err)
			}

			// Update config
			updatedCfg := cfgResp.Config
			updatedCfg.Basic.Enabled = enabled
			if enableAll {
				updatedCfg.Basic.CreateNewMembers = true
				updatedCfg.Basic.AllowSelfOnboarding = true
				updatedCfg.Basic.EnableMemberPermissions = true
				updatedCfg.MagicLinks.Enabled = true
			}

			_, err = internal.MangoClient().SDK.SetB2BConfig(context.Background(), sdk.SetB2BConfigRequest{
				ProjectID: projectID,
				Config:    updatedCfg,
			})
			if err != nil {
				log.Fatalf("Unable to update SDK config: %v", err)
			}

			fmt.Println("SDK configuration updated successfully")
		},
	}

	cmd.Flags().StringVarP(&projectID, "project-id", "p", "", "Project ID")
	cmd.Flags().BoolVarP(&enabled, "enabled", "e", true, "Enable/disable SDK")
	cmd.Flags().BoolVarP(&enableAll, "all", "a", false, "Enables standard settings that get enabled from dashboard when you turn on SDKs")
	_ = cmd.MarkFlagRequired("project-id")
	_ = cmd.MarkFlagRequired("enabled")
	return cmd
}
