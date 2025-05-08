package demoapps

import (
	"context"
	"errors"
	"fmt"
	"io/fs"
	"log"
	"net/url"
	"os"
	"regexp"
	"slices"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/stytchauth/stytch-management-go/v2/pkg/models/projects"
	"github.com/stytchauth/stytch-management-go/v2/pkg/models/publictokens"
	"github.com/stytchauth/stytch-management-go/v2/pkg/models/redirecturls"
	"github.com/stytchauth/stytch-management-go/v2/pkg/models/sdk"

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

			// Assert that FE SDKs are enabled and, if not, enable them.
			checkSDKActive(c.Context(), projectID)

			// Assert that the project has valid redirect URLs.
			checkRedirectURLs(c.Context(), projectID)

			// Grab public token.
			projectPublicToken := projectToken(c.Context(), projectID)
			writeEnvFile(projectPublicToken)
		},
	}
}

func writeEnvFile(projectPublicToken string) {
	fmt.Println("✍️ Writing public token to .env.local")
	// hardcode the path to the example app for now, remove
	envFile := "../stytch-react-example/.env.local"

	// read in env file if it exists, otherwise create it
	content, err := os.ReadFile(envFile)
	// Convert content to string and check for existing token
	fileContent := string(content)
	tokenLine := "REACT_APP_STYTCH_PUBLIC_TOKEN=" + projectPublicToken + "\n"

	if os.IsNotExist(err) {
		// Create new file if it doesn't exist
		err = os.WriteFile(envFile, []byte(tokenLine), fs.FileMode(0644))
		if err != nil {
			log.Fatalf("Failed to create %s file: %v", envFile, err)
		}
	} else {
		// Replace existing token or append if not found
		if strings.Contains(fileContent, "REACT_APP_STYTCH_PUBLIC_TOKEN=") {
			fileContent = regexp.MustCompile(`REACT_APP_STYTCH_PUBLIC_TOKEN=.*\n`).ReplaceAllString(fileContent, tokenLine)
		} else {
			fileContent += tokenLine
		}

		err = os.WriteFile(envFile, []byte(fileContent), fs.FileMode(0644))
		if err != nil {
			log.Fatalf("Failed to write to %s file: %v", envFile, err)
		}
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

func checkSDKActive(ctx context.Context, projectID string) {
	cfgResp, err := internal.MangoClient().SDK.GetConsumerConfig(ctx, sdk.GetConsumerConfigRequest{
		ProjectID: projectID,
	})
	if err != nil {
		log.Fatalf("Unable to retrieve SDK config: %v", err)
	}
	updatedCfg := cfgResp.Config
	fmt.Println("Enabling usage of Frontend SDKs in your project...")
	updatedCfg.Basic.Enabled = true

	if len(updatedCfg.Basic.Domains) == 0 {
		fmt.Println("Frontend SDKs does not have domains set, setting to localhost:3000")
		updatedCfg.Basic.Domains = []string{"http://localhost:3000"}
	}
	// these are very specific to the example app
	updatedCfg.MagicLinks.LoginOrCreateEnabled = true
	updatedCfg.MagicLinks.SendEnabled = true
	updatedCfg.Basic.CreateNewUsers = true

	_, err = internal.MangoClient().SDK.SetConsumerConfig(ctx, sdk.SetConsumerConfigRequest{
		ProjectID: projectID,
		Config:    updatedCfg,
	})
	if err != nil {
		log.Fatalf("Unable to update SDK config: %v", err)
	}
}

func checkRedirectURLs(ctx context.Context, projectID string) {
	resp, err := internal.MangoClient().RedirectURLs.GetAll(ctx, redirecturls.GetAllRequest{
		ProjectID: projectID,
	})
	if err != nil {
		log.Fatalf("Unable to retrieve redirect URLs: %v", err)
	}

	// A valid redirect URL has to be enabled for login AND signup.
	var validRedirectURLs []string
	for _, u := range resp.RedirectURLs {
		var tpes []redirecturls.RedirectType
		for _, tpe := range u.ValidTypes {
			tpes = append(tpes, tpe.Type)
		}
		if slices.Contains(tpes, redirecturls.RedirectTypeLogin) && slices.Contains(tpes, redirecturls.RedirectTypeSignup) {
			validRedirectURLs = append(validRedirectURLs, u.URL)
		}
	}
	if len(validRedirectURLs) > 0 {
		fmt.Println("Valid Redirect URLs already configured for this project.")
		return
	}

	fmt.Println("No Redirect URLs configured for this project. Please add one to continue.")
	redirectPrompt := promptui.Prompt{
		Label: "Enter a redirect URL:",
		Validate: func(input string) error {
			_, err := url.Parse(input)
			return err
		},
	}
	urlInput, err := redirectPrompt.Run()
	if err != nil {
		log.Fatalf("Unable to parse redirect URL: %v", err)
	}

	_, err = internal.MangoClient().RedirectURLs.Create(ctx, redirecturls.CreateRequest{
		ProjectID: projectID,
		RedirectURL: redirecturls.RedirectURL{
			URL: urlInput,
			ValidTypes: []redirecturls.URLRedirectType{
				{
					Type:      redirecturls.RedirectTypeLogin,
					IsDefault: false,
				},
				{
					Type:      redirecturls.RedirectTypeSignup,
					IsDefault: false,
				},
			},
		},
	})
	if err != nil {
		log.Fatalf("Unable to create redirect URL: %v", err)
	}
	fmt.Println("Redirect URL created successfully.")
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
			return token.PublicToken
		}
	}

	log.Fatalf("Unable to find project token")
	return ""
}
