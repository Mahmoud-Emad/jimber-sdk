package cli

import (
	// "bufio"
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"syscall"

	terminal "golang.org/x/crypto/ssh/terminal"
	logger "jimber.com/sdk/cmd/logger"
)

type Credentials struct {
	username string
	password string
}

func PromptConnect(logger *logger.Logger) Credentials {
	crds := Credentials{
		username: "",
		password: "",
	}

	fmt.Println("use `jimber logout` to remove the saved credentials.")

	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Failed to get current directory:", err)
		os.Exit(1)
	}
	configPath := filepath.Join(currentDir, ".jimber", "config")
	fmt.Printf("Your password will be stored in %s.\n", configPath)

	if crds.GetCredentials() {
		fmt.Println("Authenticating with existing credentials...")
		fmt.Println(crds)
	} else {
		fmt.Println("Authenticating with username/password...")
		crds.SetCredentials(logger)
	}

	return crds
}

func (crds *Credentials) GetCredentials() bool {
	config := LoadConfigFile(".jimber/config")
	section, err := config.Config.GetSection("credentials")
	if err != nil {
		return false
	}
	fmt.Println(section)
	fmt.Println(config.Config.ChildSections("credentials"))
	// crds.username = config.Config.ChildSections("credentials")
	return true
}

func (crds *Credentials) SetCredentials(logger *logger.Logger) (*Credentials, error) {

	config := LoadConfigFile(".jimber/config")
	_, err := config.Config.GetSection("project")

	if err != nil {
		config.IsError = true
		config.ErrorMessage = fmt.Errorf("There is no config file %s", err)
		logger.Error(config.ErrorMessage.Error(), nil)
		return crds, err
	}

	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Username: ")
	username, _ := reader.ReadString('\n')
	username = strings.ToLower(strings.TrimSpace(username))

	fmt.Printf("Password: ")
	password, err := terminal.ReadPassword(int(syscall.Stdin))
	fmt.Println()

	if err != nil {
		config.IsError = true
		config.ErrorMessage = fmt.Errorf("Error reading password: %s", err)
		logger.Error(config.ErrorMessage.Error(), nil)
		os.Exit(1)
	}

	if err != nil {
		fmt.Println("Error reading password:", err)
		os.Exit(1)
	}

	crds.username = username
	crds.password, err = hashPassword(string(password))
	if err != nil {
		fmt.Println("Failed to hash password:", err)
		os.Exit(1)
	}

	credentialsSection, err := config.Config.NewSection("credentials")
	if err != nil {
		config.IsError = true
		config.ErrorMessage = fmt.Errorf("Can not create credentials section due error: %s", err)
		logger.Error(config.ErrorMessage.Error(), nil)
		return crds, err
	}

	config.SetValue(credentialsSection, "username", crds.username)
	config.SetValue(credentialsSection, "password", crds.password)
	config.Save()

	return crds, nil
}
