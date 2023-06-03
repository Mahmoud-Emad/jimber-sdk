package server

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Register the projects routes under the api group.
// This function takes two arguments:
// 1. db: a pointer to the gorm.DB instance.
// 2. api: the api group to register the routes under.
func (p *Project) RegisterRoutes(db *gorm.DB, api *gin.RouterGroup) {
	api.GET("/projects", p.GetProjects(db))
	api.POST("/projects", p.CreateProject(db))
}
