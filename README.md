# Envserver

## Project Description

envserver is a server application that allows users to store and manage environment keys for their projects. It provides a centralized platform for securely storing and accessing environment variables, similar to the functionality provided by tools like `flagsmeth`. With envserver, users can easily store and retrieve environment keys for their projects, enhancing their development workflow.

The server application includes a database with two main tables: "User" and "Project". The "User" table handles user registration and login functionality, with each user being assigned a unique token for authentication. The "Project" table is used to manage projects, their teams, and their associated environment variables.

## CLI Tool

envserver provides a command-line interface (CLI) tool to facilitate key management for users. The CLI tool offers several commands to interact with the server and manage environment keys effectively. The available commands are:

- pull: Pulls the latest changes from the server and creates or updates the local configuration file.
- push: Pushes the local changes to the server, updating the environment keys.
- add: Adds new environment keys to the local configuration file.
- commit: Commits the changes to the local configuration file, providing a commit message. The commit message can be customized and will be updated if conflicts occur.

### Please note that, all of these commands are still under implementation.

## Project Configuration

For detailed information on configuring the envserver project, refer to the [Project Configuration](./docs/configuration.md) document. This document provides instructions on setting up the config.toml configuration file, which includes important settings such as database connection details and server port.

## Makefile Commands

- `build`: This command builds the project by compiling the `cmd/server.go` file.
- `run`: This command first builds the project by invoking the build command, and then it runs the built executable file `./server`.
- `test`: This command first builds the project by invoking the build command, and then it runs all the tests in the project using the go test command.
- `clean`: This command will remove the executable file `./server`.

## Contributing

If you would like to contribute to the envserver project, please refer to the [Contributing Guidelines](./docs/contributing.md) document. It outlines the steps to contribute, including guidelines for reporting issues, suggesting improvements, and submitting pull requests.
