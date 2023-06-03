package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (p *User) GetUsers(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Retrieving users
		var users []User
		result := db.Find(&users)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}

		// Return the users as a response
		c.JSON(http.StatusOK, gin.H{"users": users})
	}
}

// Initialize a new User struct
func NewUser() *User {
	return &User{}
}
