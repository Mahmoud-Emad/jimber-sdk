package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// User struct holds data of users
type User struct {
	gorm.Model
	ID             uuid.UUID  `gorm:"primary_key; unique; type:uuid; column:id"`
	Name           string     `json:"name" binding:"required"`
	Email          string     `json:"email" gorm:"unique" binding:"required"`
	HashedPassword []byte     `json:"hashed_password" binding:"required"`
	UpdatedAt      time.Time  `json:"updated_at"`
	ISOwner        bool       `json:"is_owner"`
	Projects       []*Project `gorm:"many2many:user_projects;"`
}

// BeforeCreate generates a new uuid.
func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return err
	}

	user.ID = id
	return
}
