package passwordstrengthconfig

import (
	"context"
	"log"

	"github.com/spf13/cobra"
	"github.com/stytchauth/stytch-cli/cmd/internal"
	"github.com/stytchauth/stytch-management-go/v2/pkg/models/passwordstrengthconfig"
)

// NewGetCommand returns a cobra command for retrieving password strength configuration
func NewGetCommand() *cobra.Command {
	var projectID string

	cmd := &cobra.Command{
		Use:   "get",
		Short: "Retrieve password strength configuration",
		Long:  "Retrieve password strength configuration for a project",
		Run: func(c *cobra.Command, args []string) {
			if projectID == "" {
				log.Fatalf("Missing --project-id")
			}

			res, err := internal.GetDefaultMangoClient().PasswordStrengthConfig.Get(
				context.Background(),
				passwordstrengthconfig.GetRequest{ProjectID: projectID},
			)
			if err != nil {
				log.Fatalf("Get password strength config: %s", err)
			}

			internal.PrintJSON(res)
		},
	}

	cmd.Flags().StringVarP(&projectID, "project-id", "p", "", "The project ID")

	return cmd
}
