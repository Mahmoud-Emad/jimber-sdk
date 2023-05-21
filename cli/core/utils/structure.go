package cli

import (
	"fmt"
	"os"
	"path/filepath"
)

// Check if the .jimber directory exists
func IsProjectFolderExist(logger Logger, jimberDir string) bool {
	dirInfo, err := os.Stat(jimberDir)
	if err != nil {
		if !os.IsNotExist(err) {
			logger.Error(fmt.Sprintf("Error checking %s directory: %s", jimberDir, err))
		}
		return false
	}

	// Check if it's a directory
	if !dirInfo.IsDir() {
		return false
	}
	return true
}

func IsConfigFileExist(jimberDir string, configFile string) bool {
	configPath := filepath.Join(jimberDir, configFile)
	_, err := os.Stat(configPath)
	if err != nil {
		return false
	}
	return true
}

// Creating the .jimber folder if it doesn't exist
func CreateFolder(logger Logger, folderName string) {
	_, err := os.Stat(folderName)

	if err != nil {
		if os.IsNotExist(err) {
			err = os.Mkdir(folderName, 0777)
			if err != nil {
				logger.Error(fmt.Sprintf("Error creating %s directory: %s", folderName, err))
			}
		} else {
			logger.Error(fmt.Sprintf("failed to check %s directory: %v", folderName, err))
		}
	}
}
