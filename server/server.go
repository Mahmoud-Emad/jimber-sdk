package server

import (
	"fmt"
	"log"

	database "github.com/Mahmoud-Emad/jimber/database"
)

// Server represents the application server.
type Server struct {
	Storage database.Storage // Storage is the database storage for the server.
	Port    string           // Port is the server port number.
	Host    string           // Host is the server host.
}

// Serve starts the server.
func (srv *Server) Serve() {
	srv.Storage = database.Storage{
		DB: nil,
	}

	fmt.Println("|+| Connecting to the database.")
	srv.Storage.Connect()

	// Check if tables have been migrated
	migrated, err := srv.Storage.AreTablesMigrated()
	if err != nil {
		log.Fatalf("Failed to check if tables are migrated: %v", err)
	}

	if !migrated {
		fmt.Println("|+| Migrating the database.")
		srv.Storage.Migrate()
	}

	srv.Start()
}

// Start starts the server.
func (srv Server) Start() {
	fmt.Println("|+| Starting the server.")
}
