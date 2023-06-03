package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Get all registered users from the database
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

// Method to create user record into the database.
func (p *User) CreateUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user data"})
			return
		}

		result := db.Create(&user)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
			return
		}

		c.JSON(http.StatusCreated, user)
	}
}

// Initialize a new User struct
func NewUser() *User {
	return &User{}
}
