package models

import (
	"testing"

	internal "github.com/Mahmoud-Emad/envserver/internal"
	"github.com/stretchr/testify/assert"
)

// Setup database helper, created to be used inside test case functions.
func setupDB(t *testing.T) (Database, internal.Configuration) {
	db := NewDatabase()
	// Read the database config.
	dbConfig, err := internal.ReadConfFile("../config.toml")
	assert.NoError(t, err)
	err = db.Connect(dbConfig.Database)
	assert.NoError(t, err)
	err = db.Migrate()
	assert.NoError(t, err)
	return db, dbConfig
}

// Test connect to database.
// This function expect to initialize a database with a specific configuration.
// The expected behavior is
// 1 - Raising an error in the first block based on the wrong configuration.
// 2 - Passing the second scenario because it's a valid database configuration.
func TestDatabaseConnect(t *testing.T) {
	db, conf := setupDB(t)

	t.Run("invalid database", func(t *testing.T) {
		err := db.Connect(internal.DatabaseConfiguration{})
		assert.Error(t, err)
	})
	t.Run("valid database", func(t *testing.T) {
		err := db.Connect(conf.Database)
		assert.NoError(t, err)
	})
}

func TestUser(t *testing.T) {
	t.Run("create new user object", func(t *testing.T) {
		// Test create new user record into the database.
		db, _ := setupDB(t)
		err := db.CreateUser(&User{
			Name:     "test",
			Email:    "test@test.test",
			Projects: []*Project{},
		})

		assert.NoError(t, err)

		user, err := db.GetUserByEmail("test@test.test")

		assert.Equal(t, user.Name, "test")
		assert.NoError(t, err)
	})

	t.Run("delete created user", func(t *testing.T) {
		// Test delete user record from the database by it's email.
		db, _ := setupDB(t)
		var user User

		user, err := db.GetUserByEmail("test@test.test")

		assert.NoError(t, err)
		assert.Equal(t, user.Name, "test")

		err = db.DeleteUserByEmail("test@test.test")

		assert.NoError(t, err)

		user, err = db.GetUserByEmail("test@test.test")
		assert.Error(t, err)
	})
}

func TestProject(t *testing.T) {
	t.Run("create new project object", func(t *testing.T) {
		// Test create new project record into the database.
		db, _ := setupDB(t)
		err := db.CreateProject(&Project{
			Name: "test",
		})

		assert.NoError(t, err)

		p, err := db.GetProjectByName("test")

		assert.Equal(t, p.Name, "test")
		assert.NoError(t, err)
	})

	t.Run("delete created project", func(t *testing.T) {
		// Test delete project record from the database by its name.
		db, _ := setupDB(t)

		p, err := db.GetProjectByName("test")

		assert.NoError(t, err)
		assert.Equal(t, p.Name, "test")

		err = db.DeleteProjectByName("test")

		assert.NoError(t, err)

		_, err = db.GetProjectByName("test")
		assert.Error(t, err)
	})
}

func TestEnvironmentKey(t *testing.T) {
	t.Run("create new environment Key object", func(t *testing.T) {
		// Test create new env key|value record into the database.
		db, _ := setupDB(t)

		err := db.CreateProject(&Project{
			Name: "test",
		})

		assert.NoError(t, err)

		p, err := db.GetProjectByName("test")

		err = db.CreateEnvKey(&EnvironmentKey{
			Key:       "test",
			Value:     "test",
			ProjectID: p.ID,
		})

		assert.NoError(t, err)

		env, err := db.GetEnvKeyByKeyName("test")

		assert.Equal(t, env.Key, "test")
		assert.NoError(t, err)
	})

	t.Run("delete created environment key", func(t *testing.T) {
		// Test delete key|value record from the database by its key name.
		db, _ := setupDB(t)

		env, err := db.GetEnvKeyByKeyName("test")

		assert.NoError(t, err)
		assert.Equal(t, env.Key, "test")

		err = db.DeleteEnvKeyByKeyName("test")

		assert.NoError(t, err)

		_, err = db.GetEnvKeyByKeyName("test")
		assert.Error(t, err)

		err = db.DeleteProjectByName("test")

		assert.NoError(t, err)

		_, err = db.GetProjectByName("test")
		assert.Error(t, err)
	})
}
