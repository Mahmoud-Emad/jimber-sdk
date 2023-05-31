package jimber

import (
	"fmt"

	"github.com/gin-gonic/gin"
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
	jimber.Server.Router = gin.Default()
	jimber.RegisterAPIRoutes()
	return jimber
}

// Register Project routes inside the jim server
func (jim *Jimber) RegisterAPIRoutes() {
	api := jim.Server.Router.Group("api") // => api/
	fmt.Println(api)
	jim.Server.Api.Projects.RegisterRoutes(jim.Server.Storage.DB, api)
}

func (jim *Jimber) RunServer() {
	jim.Server.Start()
}
