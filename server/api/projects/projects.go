package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (p *Projects) GetProjects(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Retrieving projects
		var projects []Projects
		result := db.Find(&projects)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}

		// Return the projects as a response
		c.JSON(http.StatusOK, gin.H{"projects": projects})
	}
}

func (p *Projects) CreateProject(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var project Projects
		if err := c.ShouldBindJSON(&project); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project data"})
			return
		}

		fmt.Println(&project.Name)

		result := db.Create(&project)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create project"})
			return
		}

		c.JSON(http.StatusCreated, project)
	}
}

// Initialize a new projects struct
func NewProjects() *Projects {
	return &Projects{}
}