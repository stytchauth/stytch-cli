package set

import (
	"github.com/spf13/cobra"
)

func NewEnableCommand() *cobra.Command {
	var projectID string
	var enabled bool
	cmd := &cobra.Command{
		Use:   "enable",
		Short: "Enable B2B SDK configuration",
	}

	cmd.Flags().StringVarP(&projectID, "project-id", "p", "", "Project ID")
	cmd.Flags().BoolVarP(&enabled, "enabled", "e", true, "Enable/disable SDK")
	_ = cmd.MarkFlagRequired("project-id")
	_ = cmd.MarkFlagRequired("enabled")
	return cmd
}
