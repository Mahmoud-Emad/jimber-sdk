package cmd

import (
	"fmt"
	"log"

	cobra "github.com/spf13/cobra"
	cli "jimber.com/sdk/cmd/cli"
	utils "jimber.com/sdk/cmd/cli/core/utils"
	logger "jimber.com/sdk/cmd/logger"
)

func InitCommands() JimberCommands {
	logger := logger.NewLogger()
	rootCmd := &cobra.Command{
		Use:   "jimber",
		Short: "Jimber CLI tool",
	}

	initCmd := JimberInit(logger)
	connectCmd := JimberConnect(logger)
	rootCmd.AddCommand(connectCmd)
	rootCmd.AddCommand(initCmd)

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}

	return JimberCommands{
		command: rootCmd,
	}
}

func JimberInit(logger *logger.Logger) *cobra.Command {
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

func JimberConnect(logger *logger.Logger) *cobra.Command {
	var username, password string

	connectCmd := &cobra.Command{
		Use:   "connect",
		Short: "Connect to the server",
		Run: func(cmd *cobra.Command, args []string) {
			if utils.IsProjectInitialized(logger) {
				utils.PromptConnect(logger)
				token, err := cli.Connect(logger, username, password)
				if err != nil {
					log.Fatal("Login failed:", err)
				}
				fmt.Println("Authentication token:", token)
			}
		},
	}

	connectCmd.Flags().StringVarP(&username, "username", "u", "", "Username")
	connectCmd.Flags().StringVarP(&password, "password", "p", "", "Password")

	return connectCmd
}
