package cli

import (
	"fmt"
	"log"

	cli "github.com/Mahmoud-Emad/jimber/cli/core"
	utils "github.com/Mahmoud-Emad/jimber/cli/core/utils"
	cobra "github.com/spf13/cobra"
)

func JimberConnect(logger utils.Logger) *cobra.Command {
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
