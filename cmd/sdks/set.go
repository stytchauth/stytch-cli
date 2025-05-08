package sdks

import (
	"context"
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/stytchauth/stytch-cli/cmd/internal"
	"github.com/stytchauth/stytch-management-go/v2/pkg/models/sdk"
)

func NewSetCommand() *cobra.Command {
	var (
		projectID string
		enabled   bool
		domains   []string
		createNewUsersEnabled bool
		magicLinksEnabled bool
	)

	cmd := &cobra.Command{
		Use:   "set",
		Short: "Set SDK configuration for a project",
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
			if len(domains) > 0 {
				updatedCfg.Basic.Domains = domains
			}
			updatedCfg.Basic.CreateNewUsers = createNewUsersEnabled
			updatedCfg.MagicLinks.LoginOrCreateEnabled = magicLinksEnabled
			// Set new config
			_, err = internal.MangoClient().SDK.SetConsumerConfig(context.Background(), sdk.SetConsumerConfigRequest{
				ProjectID: projectID,
				Config:    updatedCfg,
			})
			if err != nil {
				log.Fatalf("Unable to update SDK config: %v", err)
			}

			fmt.Println("SDK configuration updated successfully")
		},
	}

	cmd.Flags().StringVarP(&projectID, "project-id", "i", "", "Project ID")
	cmd.Flags().BoolVarP(&enabled, "enabled", "e", true, "Enable/disable SDK")
	cmd.Flags().StringSliceVarP(&domains, "domains", "d", []string{}, "Allowed domains (comma-separated)")
	cmd.Flags().BoolVarP(&createNewUsersEnabled, "create-new-users", "c", true, "Enable/disable create new users")
	cmd.Flags().BoolVarP(&magicLinksEnabled, "magic-links", "l", false, "Enable/disable Magic Links login or create and send")
	
	cmd.MarkFlagRequired("project-id")
	cmd.MarkFlagRequired("enabled")
	return cmd
}