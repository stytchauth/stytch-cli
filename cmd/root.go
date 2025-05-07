package cmd

import (
	"github.com/spf13/cobra"

	"github.com/stytchauth/stytch-cli/cmd/demoapps"
	"github.com/stytchauth/stytch-cli/cmd/emailtemplates"
	"github.com/stytchauth/stytch-cli/cmd/jwttemplates"
	"github.com/stytchauth/stytch-cli/cmd/passwordstrengthconfig"
	project "github.com/stytchauth/stytch-cli/cmd/projects"
	"github.com/stytchauth/stytch-cli/cmd/publictokens"
	"github.com/stytchauth/stytch-cli/cmd/redirecturls"
	"github.com/stytchauth/stytch-cli/cmd/secrets"
)

func NewRootCommand() *cobra.Command {
	// rootCmd represents the base command when called without any subcommands
	var command = &cobra.Command{
		Use:   "stytch-cli",
		Short: "A brief description of your application",
	}

	command.AddCommand(project.NewRootCommand())
	command.AddCommand(NewVersionCommand())
	command.AddCommand(NewAuthenticateCommand())
	command.AddCommand(jwttemplates.NewRootCommand())
	command.AddCommand(passwordstrengthconfig.NewRootCommand())
	command.AddCommand(redirecturls.NewRootCommand())
	command.AddCommand(emailtemplates.NewRootCommand())
	command.AddCommand(publictokens.NewRootCommand())
	command.AddCommand(secrets.NewRootCommand())
	command.AddCommand(demoapps.NewReactB2CSetup())

	return command
}
