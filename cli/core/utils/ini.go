package cli

import (
	"fmt"
	"os"

	ini "github.com/go-ini/ini"
)

type IniParser struct {
	Config       *ini.File
	IsError      bool
	ErrorMessage error
	location     string
}

// Load the config file from an exact location
func LoadConfigFile(location string) IniParser {
	config, err := ini.Load(location)

	if err != nil {
		if os.IsNotExist(err) {
			// If the config file doesn't exist, create a new one
			config = ini.Empty()
		} else {
			return IniParser{
				IsError:      true,
				ErrorMessage: fmt.Errorf("failed to load config: %v", err),
			}
		}
	}
	return IniParser{
		Config:       config,
		IsError:      false,
		ErrorMessage: nil,
		location:     location,
	}
}

// Create a new section ito the config file
func (psr IniParser) CreateSection(sectionName string) *ini.Section {
	section, err := psr.Config.GetSection("git")
	if err != nil {
		section, err = psr.Config.NewSection(sectionName)
		if err != nil {
			psr.IsError = true
			psr.ErrorMessage = fmt.Errorf("failed to create %s section: %v", sectionName, err)
		}
	}
	return section
}

// Set new key and value inside an exact section
func (psr IniParser) SetValue(section *ini.Section, key string, value string) {
	section.Key(key).SetValue(value)
}

// Save the config file
func (psr IniParser) Save() *ini.File {
	err := psr.Config.SaveTo(psr.location)
	if err != nil {
		psr.IsError = true
		psr.ErrorMessage = fmt.Errorf("failed to save config: %v", err)
	}
	return psr.Config
}
