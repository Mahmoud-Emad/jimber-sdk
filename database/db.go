package database

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	_ "github.com/lib/pq"
)

// getDBFilePath returns the full path to the specified database file.
func (stg *Storage) getDBFilePath(filename string) string {
	// Get the current working directory
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Printf("Failed to get current working directory: %v", err)
		os.Exit(1)
	}

	// Create the full path to the file
	return filepath.Join(currentDir, "database/queries", filename)
}

// Migrate performs database migration by executing the SQL script in create_tables.sql.
// It creates the necessary tables in the database.
func (stg *Storage) Migrate() error {
	err := stg.executeFile("create_tables.sql")
	if err != nil {
		return fmt.Errorf("failed to migrate: %w", err)
	}

	fmt.Println("|+| Tables created successfully!")
	return nil
}

// Connect establishes a connection to the database.
func (stg *Storage) Connect() error {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return fmt.Errorf("failed to establish connection: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return fmt.Errorf("failed to ping the database: %w", err)
	}

	stg.DB = db
	fmt.Println("|+| Connected to the database!")
	return nil
}

// AreTablesMigrated checks if the necessary tables are already migrated in the database.
func (stg *Storage) AreTablesMigrated() (bool, error) {
	query := `SELECT EXISTS (
		SELECT 1
		FROM   information_schema.tables
		WHERE  table_schema = 'public'
		AND    table_name = 'projects'
	)`

	var exists bool
	err := stg.DB.QueryRow(query).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("failed to check table migration: %w", err)
	}

	return exists, nil
}

// CreateProject creates a new project in the database.
func (stg *Storage) CreateProject(project *Project) error {
	// Implement the logic to create a project in the database
	return nil
}

// executeFile executes the SQL script in the specified file.
func (stg *Storage) executeFile(filename string) error {
	sqlFile := stg.getDBFilePath(filename)

	// Read the SQL file
	fileContents, err := ioutil.ReadFile(sqlFile)
	if err != nil {
		return fmt.Errorf("failed to read SQL file: %w", err)
	}

	// Execute the SQL script
	_, err = stg.DB.Exec(string(fileContents))
	if err != nil {
		return fmt.Errorf("failed to execute SQL file %s: %w", filename, err)
	}

	return nil
}
