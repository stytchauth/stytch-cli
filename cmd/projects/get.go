package project

import (
	"context"
	"log"

	"github.com/spf13/cobra"
	"github.com/stytchauth/stytch-cli/cmd/internal"
	"github.com/stytchauth/stytch-management-go/v2/pkg/models/projects"
)

var getProjectID string // for the --id flag

func NewGetCommand() *cobra.Command {
	getCommand := &cobra.Command{
		Use:   "get",
		Short: "Get a project by ID",
		Run: func(c *cobra.Command, args []string) {
			client := internal.MangoClient()
			ctx := context.Background()

			// Call the get endpoint
			res, err := client.Projects.Get(ctx, projects.GetRequest{
				ProjectID: getProjectID,
			})
			if err != nil {
				log.Fatalf("Error getting project: %v", err)
			}

			internal.PrintJSON(res)
		},
	}
	getCommand.Flags().StringVarP(&getProjectID, "id", "i", "", "The ID of the project to get")
	err := getCommand.MarkFlagRequired("id")
	if err != nil {
		log.Fatalf("Error marking project-id flag required: %v", err)
	}
	return getCommand
}