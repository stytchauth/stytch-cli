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
	var slugPattern string

	cmd := &cobra.Command{
		Use:   "domain",
		Short: "Set B2B SDK domains",
		Run: func(cmd *cobra.Command, args []string) {
			// Get current config
			cfgResp, err := internal.MangoClient().SDK.GetB2BConfig(context.Background(), sdk.GetB2BConfigRequest{
				ProjectID: projectID,
			})
			if err != nil {
				log.Fatalf("Unable to retrieve B2B SDK config: %v", err)
			}

			// Update config
			updatedCfg := cfgResp.Config
			updatedCfg.Basic.Domains = append(updatedCfg.Basic.Domains, sdk.AuthorizedB2BDomain{
				Domain:      domain,
				SlugPattern: slugPattern,
			})

			// Set new config
			_, err = internal.MangoClient().SDK.SetB2BConfig(context.Background(), sdk.SetB2BConfigRequest{
				ProjectID: projectID,
				Config:    updatedCfg,
			})
			if err != nil {
				log.Fatalf("Unable to update B2B SDK config: %v", err)
			}
		},
	}

	cmd.Flags().StringVarP(&projectID, "project-id", "p", "", "Project ID")
	cmd.Flags().StringVarP(&domain, "domain", "d", "", "Domain")
	cmd.Flags().StringVarP(&slugPattern, "slug-pattern", "s", "", "Slug pattern")
	_ = cmd.MarkFlagRequired("project-id")
	_ = cmd.MarkFlagRequired("domain")
	_ = cmd.MarkFlagRequired("slug-pattern")

	return cmd
}
