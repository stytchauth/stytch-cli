package sdks

import (
	"github.com/spf13/cobra"
	"github.com/stytchauth/stytch-cli/cmd/sdks/b2b"
	"github.com/stytchauth/stytch-cli/cmd/sdks/consumer"
)

// NewRootCommand creates the root command for public tokens operations
func NewRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sdks",
		Short: "SDKs",
		Long:  "SDKs",
	}

	cmd.AddCommand(b2b.NewB2BRootCommand())
	cmd.AddCommand(consumer.NewConsumerRootCommand())

	return cmd
}
