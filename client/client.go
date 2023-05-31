package client

import cli "jimber.com/sdk/cmd/cli/commands"

func NewJimberClient() *JimberClient {
	return &JimberClient{
		Commands: cli.InitCommands(),
	}
}
