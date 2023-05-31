package cli

import (
	"fmt"

	utils "jimber.com/sdk/cmd/cli/core/utils"
	logger "jimber.com/sdk/cmd/logger"
)

const jimberDir = ".jimber"

// Initialize the project
// Create .jimber folder inside the project folder.
// This folder should have a file for the configuration.
func Init(logger *logger.Logger, projectName string) string {
	if !utils.IsProjectFolderExist(logger, jimberDir) {
		if !utils.IsConfigFileExist(jimberDir, "config") {
			if len(projectName) < 1 {
				defaultProjectName, err := utils.GetGitProjectName()
				if err != nil {
					logger.Error(err.Error(), nil)
					// os.Exit(0)
				}
				projectName = utils.PromptGitProjectName(defaultProjectName)
				utils.CreateFolder(logger, jimberDir)
			}
			SetProjectName(logger, projectName)
		}
		logger.Success("Project Created")
	} else {
		logger.Error(fmt.Sprintf("%s directory found!", jimberDir), nil)
	}
	return projectName
}

func SetProjectName(logger *logger.Logger, projectName string) {

	gitName, _ := utils.GetGitProjectName()
	if len(gitName) > 0 {
		projectName = gitName
	}

	repoURL, _ := utils.GetGitRepoURL()

	config := utils.LoadConfigFile(".jimber/config")
	if config.IsError {
		logger.Error(config.ErrorMessage.Error(), nil)
	}

	logger.State("Handle config file.")

	projectSection := config.CreateSection("project")
	config.SetValue(projectSection, "name", projectName)
	config.SetValue(projectSection, "url", "localhost:8080")

	if utils.IsGitRepo() {
		gitSection := config.CreateSection("git")
		config.SetValue(gitSection, "remote", repoURL)
	}

	config.Save()
	logger.State("Config file has been created")
}
