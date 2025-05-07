package project

import (
	"log"
	"strings"

	"github.com/spf13/cobra"
	"github.com/stytchauth/stytch-management-go/v2/pkg/models/projects"

	"github.com/stytchauth/stytch-cli/cmd/internal"
)

var (
	vertical    string // for the --vertical flag
	projectName string // for the --name flag
)

func NewCreateCommand() *cobra.Command {
	createCommand := &cobra.Command{
		Use:   "create",
		Short: "Create a new project",
		Run: func(c *cobra.Command, args []string) {
			// Send the request
			var verticalType projects.Vertical
			switch strings.ToLower(vertical) {
			case "b2b":
				verticalType = projects.VerticalB2B
			case "consumer":
				verticalType = projects.VerticalConsumer
			default:
				log.Fatalf("Invalid vertical: %s", vertical)
			}

			res, err := internal.MangoClient().Projects.Create(c.Context(), projects.CreateRequest{
				ProjectName: projectName,
				Vertical:    verticalType,
			})
			if err != nil {
				log.Fatalf("Error creating B2B project: %v", err)
			}
			internal.PrintJSON(res)
		},
	}
	createCommand.Flags().StringVarP(&vertical, "vertical", "v", "", "The vertical of the project")
	createCommand.Flags().StringVarP(&projectName, "name", "n", "", "The name of the project")
	var errors []error
	errors = append(errors, createCommand.MarkFlagRequired("vertical"))
	errors = append(errors, createCommand.MarkFlagRequired("name"))
	if len(errors) > 0 {
		for _, err := range errors {
			log.Fatalf("Error marking flag required: %v", err)
		}
	}
	return createCommand
}
