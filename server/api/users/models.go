package server

// Project represents a project in the database.
type User struct {
	ID          int    `json:"id"`           // ID is the unique identifier of the user.
	FirstName   string `json:"first_name"`   // FirstName is the first name of the user.
	LastName    string `json:"last_name"`    // LastName is the last name of the user.
	CreatedDate string `json:"created_date"` // CreatedDate is the timestamp when the user was created.
	UpdatedDate string `json:"updated_date"` // UpdatedDate is the timestamp when the user was last updated.
	Email       string `json:"email"`        // Email is the email address of the user.
}
