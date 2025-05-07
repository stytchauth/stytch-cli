package demoapps

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/stytchauth/stytch-management-go/v2/pkg/models/projects"
	"github.com/stytchauth/stytch-management-go/v2/pkg/models/publictokens"

	"github.com/stytchauth/stytch-cli/cmd/internal"
)

func NewReactB2CSetup() *cobra.Command {
	return &cobra.Command{
		Use:   "setup-react-app",
		Short: "Setup a React B2C demo app",
		Run: func(c *cobra.Command, _ []string) {
			projectSelection := promptui.Select{
				Label: "Create a new project or use an existing one",
				Items: []string{"Create a new project", "Use an existing project"},
			}
			_, choice, err := projectSelection.Run()
			if err != nil {
				log.Fatalf("Error creating B2B project: %v", err)
			}

			var projectID string
			switch choice {
			case "Create a new project":
				project := createB2CProject(c.Context())
				projectID = project.TestProjectID
			case "Use an existing project":
				projectID = chooseExistingProject(c.Context())
			default:
				log.Fatalf("Invalid choice: %s", choice)
			}

			// Grab public token.
			projectPublicToken := projectToken(c.Context(), projectID)
			fmt.Printf("Public token: %s\n", projectPublicToken)
		},
	}
}

func createB2CProject(ctx context.Context) projects.Project {
	projectNamePrompt := promptui.Prompt{
		Label: "Choose a project name:",
		Validate: func(input string) error {
			if len(input) == 0 {
				return errors.New("project name cannot be empty")
			}
			return nil
		},
	}
	projectName, err := projectNamePrompt.Run()
	if err != nil {
		log.Fatalf("Unable to parse project name: %v", err)
	}

	// Create a project.
	createResp, err := internal.MangoClient().Projects.Create(ctx, projects.CreateRequest{
		ProjectName: projectName,
		Vertical:    projects.VerticalConsumer,
	})
	if err != nil {
		log.Fatalf("Unable to create project: %v", err)
	}
	fmt.Println("Project created successfully.")

	return createResp.Project
}

func chooseExistingProject(ctx context.Context) string {
	resp, err := internal.MangoClient().Projects.GetAll(ctx, projects.GetAllRequest{})
	if err != nil {
		log.Fatalf("Unable to retrieve projects: %v", err)
	}

	var projectNames []string
	for _, project := range resp.Projects {
		projectNames = append(projectNames, project.Name)
	}

	projectSelection := promptui.Select{
		Label: "Choose an existing project",
		Items: projectNames,
	}
	_, choice, err := projectSelection.Run()
	if err != nil {
		log.Fatalf("Unable to parse project choice: %v", err)
	}

	for _, project := range resp.Projects {
		if project.Name == choice {
			return project.TestProjectID
		}
	}
	log.Fatalf("Unable to find project")
	return ""
}

func projectToken(ctx context.Context, projectID string) string {
	getResp, err := internal.MangoClient().PublicTokens.GetAll(ctx, publictokens.GetAllRequest{
		ProjectID: projectID,
	})
	if err != nil {
		log.Fatalf("Unable to retrieve public tokens: %v", err)
	}

	for _, token := range getResp.PublicTokens {
		if token.ProjectID == projectID {
			return token.ProjectID
		}
	}

	log.Fatalf("Unable to find project token")
	return ""
}
