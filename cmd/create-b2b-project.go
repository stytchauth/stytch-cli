package cmd

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

var createB2BProjectCmd = &cobra.Command{
	Use:   "create-b2b-project",
	Short: "Create a new B2B project in Stytch",
	Run: func(cmd *cobra.Command, args []string) {
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
		 res, err := client.Projects.Create(ctx, projects.CreateRequest{
			ProjectName: "My new project",
			Vertical: projects.VerticalB2B,
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

func init() {
	rootCmd.AddCommand(createB2BProjectCmd)
}