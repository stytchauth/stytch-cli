package b2b

import (
	"github.com/spf13/cobra"
	"github.com/stytchauth/stytch-cli/cmd/sdks/b2b/set"
)

func NewB2BRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "b2b",
		Short: "Manage B2B SDK configuration",
	}

	cmd.AddCommand(NewGetCommand())
	cmd.AddCommand(set.NewRootCommand())
	return cmd
}

