package internal

import (
	"io/fs"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	fileContent = `
[database]
host = "localhost"
user = "postgres"
password = "postgres"
port = 5432
name = "postgres"
[server]
port = 8080
host = "localhost"
`
	fileName = "/config.toml"
)

// Test create and write the config content inside a config file.
func TestWriteConfigFile(t *testing.T) {
	t.Run("read config file ", func(t *testing.T) {
		dir := t.TempDir()
		configPath := filepath.Join(dir, fileName)

		err := os.WriteFile(configPath, []byte(fileContent), 0644)
		assert.NoError(t, err)

		data, err := ReadConfFile(configPath)
		assert.NoError(t, err)
		assert.NotEmpty(t, data)
	})

	t.Run("change permissions of file", func(t *testing.T) {
		dir := t.TempDir()
		configPath := filepath.Join(dir, fileName)

		err := os.WriteFile(configPath, []byte(fileContent), fs.FileMode(os.O_RDONLY))
		assert.NoError(t, err)

		data, err := ReadConfFile(configPath)
		assert.Error(t, err)
		assert.Empty(t, data)
	})
}

func TestParseConf(t *testing.T) {
	t.Run("parse config file", func(t *testing.T) {
		dir := t.TempDir()
		configPath := filepath.Join(dir, fileContent)

		err := os.WriteFile(configPath, []byte(fileContent), 0644)
		assert.NoError(t, err)

		got, err := ReadConfFile(configPath)
		assert.NoError(t, err)

		expected := Configuration{
			Database: databaseExpectation(),
			Server:   serverExpectation(),
		}

		assert.NoError(t, err)
		assert.Equal(t, got.Database, expected.Database)
	})

	t.Run("no file found", func(t *testing.T) {
		_, err := ReadConfFile("config.toml")
		assert.Error(t, err)
	})
}

// The expected database struct, used for testing.
func databaseExpectation() DatabaseConfiguration {
	return DatabaseConfiguration{
		Port:     5432,
		Host:     "localhost",
		Name:     "postgres",
		Password: "postgres",
		User:     "postgres",
	}
}

// The expected database struct, used for testing.
func serverExpectation() ServerConfiguration {
	return ServerConfiguration{
		Port: 8080,
		Host: "localhost",
	}
}
