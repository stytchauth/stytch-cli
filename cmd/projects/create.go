package project

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"github.com/stytchauth/stytch-management-go/pkg/api"
	"github.com/stytchauth/stytch-management-go/pkg/models/projects"
)

var vertical string // for the --vertical flag
var projectName string // for the --name flag

func NewCreateCommand() *cobra.Command {
	createCommand := &cobra.Command{
		Use:   "create",
		Short: "Create a new project",
		Run: func(c *cobra.Command, args []string) {
			// Load environment variables from .env file
			err := godotenv.Load()
			if err != nil {
				log.Fatal("Error loading .env file")
			}
			// Set your Stytch Management API credentials as env variables
			keyID := os.Getenv("STYTCH_WORKSPACE_KEY_ID")
			keySecret := os.Getenv("STYTCH_WORKSPACE_KEY_SECRET")

			if keyID == "" || keySecret == "" {
				log.Fatal("STYTCH_WORKSPACE_KEY_ID and STYTCH_WORKSPACE_KEY_SECRET must be set")
			}

			client := api.NewClient(keyID, keySecret)
			ctx := context.Background()

			// Send the request
			var verticalType projects.Vertical
			if vertical == "b2b" {
				verticalType = projects.VerticalB2B
			} else if vertical == "consumer" {
				verticalType = projects.VerticalConsumer
			} else {
				log.Fatalf("Invalid vertical: %s", vertical)
			}
			res, err := client.Projects.Create(ctx, projects.CreateRequest{
				ProjectName: projectName,
				Vertical:  verticalType  ,
			})
			if err != nil {
				log.Fatalf("Error creating B2B project: %v", err)
			}

			// Get the new project information
			// This is used in examples below
			newProject := res.Project
			fmt.Printf("New project created: %+v\n", newProject.Name)
		},
	}
	createCommand.Flags().StringVarP(&vertical, "vertical", "v", "", "The vertical of the project")
	createCommand.Flags().StringVarP(&projectName, "name", "n", "", "The name of the project")
	return createCommand
}
