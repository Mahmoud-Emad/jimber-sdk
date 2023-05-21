package cli

import (
	"log"

	utils "github.com/Mahmoud-Emad/jimber/cli/core/utils"
	"github.com/spf13/cobra"
)

type JimberCommands struct {
	command *cobra.Command
}

func InitCommands() JimberCommands {
	logger := utils.NewLogger()
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
