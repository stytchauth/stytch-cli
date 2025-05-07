package passwordstrengthconfig

import (
	"context"
	"log"

	"github.com/spf13/cobra"
	"github.com/stytchauth/stytch-cli/cmd/internal"
	"github.com/stytchauth/stytch-management-go/v2/pkg/models/passwordstrengthconfig"
)

// NewSetCommand returns a cobra command for updating password strength configuration
func NewSetCommand() *cobra.Command {
	var projectID string
	var checkBreachOnCreation bool
	var checkBreachOnAuthentication bool
	var validateOnAuthentication bool
	var validationPolicy string
	var ludsMinPasswordLength int
	var ludsMinPasswordComplexity int

	cmd := &cobra.Command{
		Use:   "set",
		Short: "Update password strength configuration",
		Long:  "Update password strength configuration for a project",
		Run: func(c *cobra.Command, args []string) {
			if projectID == "" {
				log.Fatalf("Missing --project-id")
			}
			if validationPolicy == "" {
				log.Fatalf("Missing --validation-policy")
			}

			req := passwordstrengthconfig.SetRequest{
				ProjectID: projectID,
				PasswordStrengthConfig: passwordstrengthconfig.PasswordStrengthConfig{
					CheckBreachOnCreation:       checkBreachOnCreation,
					CheckBreachOnAuthentication: checkBreachOnAuthentication,
					ValidateOnAuthentication:    validateOnAuthentication,
					ValidationPolicy:            passwordstrengthconfig.ValidationPolicy(validationPolicy),
					LudsMinPasswordLength:       ludsMinPasswordLength,
					LudsMinPasswordComplexity:   ludsMinPasswordComplexity,
				},
			}
			res, err := internal.GetDefaultMangoClient().PasswordStrengthConfig.Set(
				context.Background(), req,
			)
			if err != nil {
				log.Fatalf("Set password strength config: %s", err)
			}

			internal.PrintJSON(res)
		},
	}

	cmd.Flags().StringVarP(&projectID, "project-id", "p", "", "The project ID")
	cmd.Flags().BoolVarP(&checkBreachOnCreation, "check-breach-on-creation", "c", false, "Check breach on creation")
	cmd.Flags().BoolVarP(&checkBreachOnAuthentication, "check-breach-on-authentication", "b", false, "Check breach on authentication")
	cmd.Flags().BoolVarP(&validateOnAuthentication, "validate-on-authentication", "v", false, "Validate on authentication")
	cmd.Flags().StringVarP(&validationPolicy, "validation-policy", "y", "", "The validation policy (LUDS or ZXCVBN)")
	cmd.Flags().IntVarP(&ludsMinPasswordLength, "luds-min-password-length", "m", 0, "Minimum password length for LUDS policy")
	cmd.Flags().IntVarP(&ludsMinPasswordComplexity, "luds-min-password-complexity", "x", 0, "Minimum password complexity for LUDS policy")

	return cmd
}
