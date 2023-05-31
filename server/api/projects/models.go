package server

import user "jimber.com/sdk/server/api/users"

// / Project represents a project in the database.
type Projects struct {
	ID          int         `json:"id"`                                 // ID is the unique identifier of the project.
	Name        string      `json:"name"`                               // Name is the name of the project.
	Token       string      `json:"token"`                              // Token is the token associated with the project.
	CreatedDate string      `json:"created_date"`                       // CreatedDate is the timestamp when the project was created.
	UpdatedDate string      `json:"updated_date"`                       // UpdatedDate is the timestamp when the project was last updated.
	OwnerID     int         `json:"owner_id"`                           // OwnerID is the ID of the project owner.
	Owner       user.User   `json:"owner"`                              // Owner is the user who owns the project.
	Team        []user.User `json:"team" gorm:"many2many:project_team"` // Team is the list of users who work on the project.
}
