package database

import (
	"database/sql"
	"fmt"
)

// Storage represents the database storage and provides methods to interact with the database.
type Storage struct {
	DB *sql.DB
}

// Project represents a project in the database.
type Project struct {
	ID          int    `json:"id"`           // ID is the unique identifier of the project.
	Name        string `json:"name"`         // Name is the name of the project.
	Token       string `json:"token"`        // Token is the token associated with the project.
	CreatedDate string `json:"created_date"` // CreatedDate is the timestamp when the project was created.
	UpdatedDate string `json:"updated_date"` // UpdatedDate is the timestamp when the project was last updated.
	OwnerID     int    `json:"owner_id"`     // OwnerID is the ID of the project owner.
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
)

var connStr = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
	host, port, user, password, dbname)
