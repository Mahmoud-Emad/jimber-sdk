// Package internal for internal details
package internal

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

type Configuration struct {
	Database DatabaseConfiguration `toml:"database"`
	Server   ServerConfiguration   `toml:"server"`
}

type ServerConfiguration struct {
	Host string `toml:"host"`
	Port int64  `toml:"port"`
}

type DatabaseConfiguration struct {
	Host     string `toml:"host"`
	Port     int64  `toml:"port"`
	User     string `toml:"user"`
	Password string `toml:"password"`
	Name     string `toml:"name"`
}

// Read the config file.
func ReadConfFile(path string) (Configuration, error) {
	config := Configuration{}
	_, err := toml.DecodeFile(path, &config)
	if err != nil {
		return Configuration{}, fmt.Errorf("failed to open config file: %w", err)
	}
	return config, nil
}

// Parse the config file.
func ParseConfigFile(path string) error {
	config, err := ReadConfFile(path)
	if err != nil {
		return err
	}
	fmt.Println(config)
	return nil
}
