package server

import (
	"github.com/gin-gonic/gin"
	logger "jimber.com/sdk/cmd/logger"
	projects "jimber.com/sdk/server/api/projects"
	users "jimber.com/sdk/server/api/users"
	storage "jimber.com/sdk/server/database"
)

type APIRequest struct {
	Projects *projects.Project
	Users    *users.User
}

type JimberServer struct {
	Api     *APIRequest
	Storage *storage.Storage
	Port    string
	Router  *gin.Engine // Router is the HTTP router.
	Host    string
	logger  *logger.Logger
}
