package project

import (
	"context"
	"log"

	"github.com/spf13/cobra"
	"github.com/stytchauth/stytch-cli/cmd/internal"
	"github.com/stytchauth/stytch-management-go/v2/pkg/models/projects"
)

func NewGetAllCommand() *cobra.Command {
	getAllCommand := &cobra.Command{
		Use:   "getall",
		Short: "Get all all projects",
		Run: func(c *cobra.Command, args []string) {
			client := internal.MangoClient()
			ctx := context.Background()

			// Call the list endpoint
			res, err := client.Projects.GetAll(ctx, projects.GetAllRequest{})
			if err != nil {
				log.Fatalf("Error getting all projects: %v", err)
			}

			internal.PrintJSON(res)
		},
	}
	return getAllCommand
}
