package project

import (
	"context"
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/stytchauth/stytch-cli/cmd/internal"
	"github.com/stytchauth/stytch-management-go/v2/pkg/models/projects"
)

func NewDeleteCommand() *cobra.Command {
	var projectID string

	deleteCommand := &cobra.Command{
		Use:   "delete",
		Short: "Delete a project by ID",
		Run: func(c *cobra.Command, args []string) {
			client := internal.MangoClient()
			ctx := context.Background()

			// Call the delete endpoint
			_, err := client.Projects.Delete(ctx, projects.DeleteRequest{
				ProjectID: projectID,
			})
			if err != nil {
				log.Fatalf("Error deleting project: %v", err)
			}

			fmt.Printf("Project %s deleted successfully.\n", projectID)
		},
	}
	deleteCommand.Flags().StringVarP(&projectID, "id", "i", "", "The ID of the project to delete")
	err := deleteCommand.MarkFlagRequired("id")
	if err != nil {
		log.Fatalf("Error marking project-id flag required: %v", err)
	}
	return deleteCommand
}
