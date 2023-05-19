package cli

import (
	"fmt"
	"os"

	utils "github.com/Mahmoud-Emad/jimber/cli/core/utils"
	ini "github.com/go-ini/ini"
)

// Initialize the project
// Create .jimber folder inside the project folder.
// This folder should have a file for the configuration.
func Init(projectName string) (string, error) {
	if len(projectName) < 1 {
		defaultProjectName, err := utils.GetGitProjectName()
		if err != nil {
			fmt.Printf("Failed to get Git project name: %v\n", err)

			return "Failed to get Git project name: %v\n", err
		}
		projectName = utils.PromptGitProjectName(defaultProjectName)
	}

	result, err := createProject(projectName)
	if err != nil {
		fmt.Printf("Failed to initialize project: %v\n", err)
		return "Failed to get Git project name: %v\n", err
	}

	fmt.Println(result)
	return "Initialized..", nil
}

func createProject(projectName string) (string, error) {
	// Creating the .jimber folder if it doesn't exist
	_, err := os.Stat(".jimber")
	gitName, _ := utils.GetGitProjectName()
	fmt.Println("Project name founded, ", gitName)
	if err != nil {
		if os.IsNotExist(err) {
			err = os.Mkdir(".jimber", 0777)
			if err != nil {
				return "", fmt.Errorf("failed to create .jimber directory: %v", err)
			}
		} else {
			return "", fmt.Errorf("failed to check .jimber directory: %v", err)
		}
	}

	// Creating or updating the config file
	cfg, err := ini.Load(".jimber/config")
	if err != nil {
		if os.IsNotExist(err) {
			// If the config file doesn't exist, create a new one
			cfg = ini.Empty()
		} else {
			return "", fmt.Errorf("failed to load config: %v", err)
		}
	}

	// Create or get the 'project' section
	projectSection, err := cfg.GetSection("project")
	gitSection, err := cfg.GetSection("git")
	if err != nil {
		projectSection, err = cfg.NewSection("project")
		gitSection, err = cfg.NewSection("git")
		if err != nil {
			return "", fmt.Errorf("failed to create 'project' section: %v", err)
		}
	}

	// Set the project details
	projectSection.Key("name").SetValue(projectName)
	projectSection.Key("url").SetValue("localhost:8080")
	repoURL, err := utils.GetGitRepoURL()
	gitSection.Key("url").SetValue(repoURL)

	// Save the config file
	err = cfg.SaveTo(".jimber/config")
	if err != nil {
		return "", fmt.Errorf("failed to save config: %v", err)
	}
	return "Initialized..", nil
}
