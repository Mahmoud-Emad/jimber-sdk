package cli

import (
	"fmt"
	"log"

	cli "github.com/Mahmoud-Emad/jimber/cli/core"
	"github.com/spf13/cobra"
)

func JimberInit() *cobra.Command {
	var projectName string

	initCmd := &cobra.Command{
		Use:   "init",
		Short: "initialize the jimber dir",
		Run: func(cmd *cobra.Command, args []string) {
			initialized, err := cli.Init(projectName)
			if err != nil {
				log.Fatal("Login failed:", err)
			}
			fmt.Println("Project initialzed", initialized)
		},
	}

	initCmd.Flags().StringVarP(&projectName, "name", "n", "", "projectName")
	// initCmd.MarkFlagRequired("name")

	return initCmd

}
