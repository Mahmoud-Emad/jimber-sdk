package cli

import (
	"fmt"
	"log"

	cli "github.com/Mahmoud-Emad/jimber/cli/core"
	"github.com/spf13/cobra"
)

func JimberConnect() *cobra.Command {
	var username, password string

	connectCmd := &cobra.Command{
		Use:   "connect",
		Short: "Connect to the server",
		Run: func(cmd *cobra.Command, args []string) {
			token, err := cli.Connect(username, password)
			if err != nil {
				log.Fatal("Login failed:", err)
			}
			fmt.Println("Authentication token:", token)
		},
	}

	connectCmd.Flags().StringVarP(&username, "username", "u", "", "Username")
	connectCmd.Flags().StringVarP(&password, "password", "p", "", "Password")
	connectCmd.MarkFlagRequired("username")
	connectCmd.MarkFlagRequired("password")
	return connectCmd
}
