package cli

import (
	"log"

	"github.com/spf13/cobra"
)

type JimberCommands struct {
	command *cobra.Command
}

func InitCommands() JimberCommands {
	rootCmd := &cobra.Command{
		Use:   "jimber",
		Short: "Jimber CLI tool",
	}

	initCmd := JimberInit()
	connectCmd := JimberConnect()
	rootCmd.AddCommand(connectCmd)
	rootCmd.AddCommand(initCmd)

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}

	return JimberCommands{
		command: rootCmd,
	}
}
