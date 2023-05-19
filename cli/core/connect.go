package cli

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Connect to jimber server
func Connect(username string, password string) (string, error) {
	log.Println("Connecting to the server")
	postBody, err := json.Marshal(map[string]string{
		"username": username,
		"password": password,
	})
	if err != nil {
		return "", err
	}

	responseBody := bytes.NewBuffer(postBody)
	log.Println("Request body:", responseBody.String())

	response, err := http.Post("http://localhost:8080/login", "application/json", responseBody)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return "", fmt.Errorf("login failed with status: %d", response.StatusCode)
	}

	var responseJSON struct {
		Token string `json:"token"`
	}

	err = json.NewDecoder(response.Body).Decode(&responseJSON)
	if err != nil {
		return "", err
	}

	return responseJSON.Token, nil
}
