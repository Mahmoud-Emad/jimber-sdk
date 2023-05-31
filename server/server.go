package server

import (
	"fmt"

	logger "jimber.com/sdk/cmd/logger"
	server "jimber.com/sdk/server/api/projects"
)

func NewJimberServer(host string, port string) *JimberServer {
	s := JimberServer{}
	s.Host = host
	s.Port = port
	s.Api = NewAPIRequest()
	s.logger = logger.NewLogger()
	return &s
}

func NewAPIRequest() *APIRequest {
	return &APIRequest{
		Projects: server.NewProjects(),
	}
}

// Start the server
func (s *JimberServer) Start() {
	// Create the router

	addr := fmt.Sprintf("%s:%s", s.Host, s.Port)
	err := s.Router.Run(addr)
	if err != nil {
		s.logger.Error("Failed to start the server: ", err)
	}
}
