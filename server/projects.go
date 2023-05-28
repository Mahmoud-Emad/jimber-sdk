package server

import (
	"net/http"
	"sync"

	database "github.com/Mahmoud-Emad/jimber/database"
	"github.com/gin-gonic/gin"
)

type App struct {
	storage *database.Storage
	mutex   sync.Mutex
}

func (app *App) CreateProject(c *gin.Context) {
	var project database.Project

	// Bind JSON request body to the project struct
	if err := c.ShouldBindJSON(&project); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse request body"})
		return
	}

	// Validate the project data
	// (Add your validation logic here)

	// Acquire the lock
	app.mutex.Lock()
	defer app.mutex.Unlock()

	// Save the project to the database
	err := app.storage.CreateProject(&project)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create project"})
		return
	}

	// Return the created project as the response
	c.JSON(http.StatusOK, gin.H{"project": project})
}
