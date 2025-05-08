package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/stytchauth/stytch-cli/utils"
)

func NewLogoutCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "logout",
		Short: "Logout and revoke your access token",
		Run: func(cmd *cobra.Command, args []string) {
			if err := utils.DeleteToken(utils.AccessToken); err != nil {
				log.Fatalf("Failed to logout: %v", err)
			}
			fmt.Println("Successfully logged out")
		},
	}
}
