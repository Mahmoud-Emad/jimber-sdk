package jimber

import (
	client "jimber.com/sdk/client"
	cmd "jimber.com/sdk/cmd"
	server "jimber.com/sdk/server"
	storage "jimber.com/sdk/server/database"
)

type Jimber struct {
	Server *server.JimberServer
	CLI    *cmd.JimberCLI
	Client *client.JimberClient
}

func NewJimber(host string, port string) *Jimber {
	// Initialize instances of each component
	jimber := &Jimber{
		Server: server.NewJimberServer(host, port),
		CLI:    cmd.NewJimberCLI(),
		Client: client.NewJimberClient(),
	}
	jimber.Server.Storage = storage.NewStorage()
	jimber.Server.Storage.DB = jimber.Server.Storage.Connect()
	jimber.Server.RegisterAPIRoutes()
	return jimber
}

func (jim *Jimber) RunServer() {
	jim.Server.Start()
}
