package server

import (
	"github.com/gin-gonic/gin"
)

// Register Project routes inside the jim server
func (s *JimberServer) RegisterAPIRoutes() {
	s.Router = gin.Default()
	api := s.Router.Group("api") // => api/
	s.Api.Projects.RegisterRoutes(s.Storage.DB, api)
}
