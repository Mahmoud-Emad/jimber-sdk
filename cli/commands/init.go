package cli

import (
	cli "github.com/Mahmoud-Emad/jimber/cli/core"
	utils "github.com/Mahmoud-Emad/jimber/cli/core/utils"
	"github.com/spf13/cobra"
)

func JimberInit(logger utils.Logger) *cobra.Command {
	var projectName string

	initCmd := &cobra.Command{
		Use:   "init",
		Short: "initialize the jimber dir",
		Run: func(cmd *cobra.Command, args []string) {
			projectName = cli.Init(logger, projectName)
			if !utils.IsProjectFolderExist(logger, ".jimber") && len(projectName) < 1 {
				cmd.Help()
			}
		},
	}

	initCmd.Flags().StringVarP(&projectName, "name", "n", "", "projectName")
	return initCmd

}
