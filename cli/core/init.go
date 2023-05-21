package cli

import (
	"fmt"

	utils "github.com/Mahmoud-Emad/jimber/cli/core/utils"
)

const jimberDir = ".jimber"

// Initialize the project
// Create .jimber folder inside the project folder.
// This folder should have a file for the configuration.
func Init(logger utils.Logger, projectName string) string {
	if !utils.IsProjectFolderExist(logger, jimberDir) {
		if !utils.IsConfigFileExist(jimberDir, "config") {
			if len(projectName) < 1 {
				defaultProjectName, err := utils.GetGitProjectName()
				if err != nil {
					logger.Error(err.Error())
				}
				projectName = utils.PromptGitProjectName(defaultProjectName)
				utils.CreateFolder(logger, jimberDir)
			}
			SetProjectName(logger, projectName)
		}
		logger.Success("Project Created")
	} else {
		logger.Error(fmt.Sprintf("%s directory found!", jimberDir))
	}
	return projectName
}

func SetProjectName(logger utils.Logger, projectName string) {

	gitName, _ := utils.GetGitProjectName()
	if len(gitName) > 0 {
		projectName = gitName
	}

	repoURL, _ := utils.GetGitRepoURL()

	config := utils.LoadConfigFile(".jimber/config")
	if config.IsError {
		logger.Error(config.ErrorMessage.Error())
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
