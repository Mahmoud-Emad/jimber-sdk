package cli

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// GetGitProjectName retrieves the project name from the Git configuration.
func GetGitProjectName() (string, error) {
	cmd := exec.Command("git", "config", "--get", "remote.origin.url")
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to execute git command: %v", err)
	}

	remoteURL := strings.TrimSpace(string(output))
	parts := strings.Split(remoteURL, "/")
	if len(parts) < 2 {
		return "", fmt.Errorf("invalid remote URL: %s", remoteURL)
	}

	repoName := strings.TrimSuffix(parts[len(parts)-1], ".git")
	return repoName, nil
}

func PromptGitProjectName(defaultName string) string {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Use default project name '%s'? (y/n): ", defaultName)
	answer, _ := reader.ReadString('\n')
	answer = strings.ToLower(strings.TrimSpace(answer))

	if answer == "n" || answer == "no" {
		fmt.Print("Enter new project name: ")
		projectName, _ := reader.ReadString('\n')
		return strings.TrimSpace(projectName)
	}

	return defaultName
}

func GetGitRepoURL() (string, error) {
	cmd := exec.Command("git", "config", "--get", "remote.origin.url")
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to execute 'git' command: %v", err)
	}

	repoURL := strings.TrimSpace(string(output))
	return repoURL, nil
}

// Func to check if there is a git repo inside the project
func IsGitRepo() bool {
	cmd := exec.Command("git", "rev-parse", "--is-inside-work-tree")
	_, err := cmd.Output()
	if err != nil {
		return false
	}
	return true
}
