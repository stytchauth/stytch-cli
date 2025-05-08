package consumer

import (
	"github.com/spf13/cobra"
	"github.com/stytchauth/stytch-cli/cmd/sdks/consumer/set"
)

func NewConsumerRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "consumer",
		Short: "Manage consumer SDK configuration",
	}

	cmd.AddCommand(NewGetCommand())
	cmd.AddCommand(set.NewRootCommand())

	return cmd
}
