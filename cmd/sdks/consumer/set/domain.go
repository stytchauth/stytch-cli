package set

import (
	"context"
	"log"

	"github.com/spf13/cobra"
	"github.com/stytchauth/stytch-cli/cmd/internal"
	"github.com/stytchauth/stytch-management-go/v2/pkg/models/sdk"
)

func NewDomainCommand() *cobra.Command {
	var projectID string
	var domain string

	cmd := &cobra.Command{
		Use:   "domain",
		Short: "Set consumer SDK domain",
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
			updatedCfg.Basic.Domains = append(updatedCfg.Basic.Domains, domain)

			// Set new config
			_, err = internal.MangoClient().SDK.SetConsumerConfig(context.Background(), sdk.SetConsumerConfigRequest{
				ProjectID: projectID,
				Config:    updatedCfg,
			})
			if err != nil {
				log.Fatalf("Unable to update SDK config: %v", err)
			}
		},
	}

	cmd.Flags().StringVarP(&projectID, "project-id", "p", "", "Project ID")
	cmd.Flags().StringVarP(&domain, "domain", "d", "", "Domain")
	_ = cmd.MarkFlagRequired("project-id")
	_ = cmd.MarkFlagRequired("domain")

	return cmd
}
