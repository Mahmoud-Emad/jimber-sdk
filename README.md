# Jimber SDK

The Jimber SDK is a comprehensive software development kit that enables seamless interaction with the Jimber system. It provides functionalities for managing environment variables, setting development and production modes, and handling various configurations using a command-line interface (CLI) and server API endpoints. The SDK offers modular components for enhanced code organization and ease of use.

## Features

- **CLI Client**: Interact with the Jimber CLI using commands such as `init`, `connect`, `add`, `push`, and `pull`. The CLI client generates the `jimber` binary for easy command execution.

- **Client Module**: Initialize the CLI module within the Jimber struct for seamless integration with the CLI client.

- **CMD Module**: Includes the CLI module and a logger module for error logging, success logging, and stage tracking, allowing users to track their journey within the Jimber system.

- **Server Module**: Core module responsible for running the HTTP server. It provides API endpoints for storing flags, managing variables, and handling project configurations. The server module includes sub-modules for each database table (e.g., Projects, Users), along with a URL file for registering API endpoints.

- **Database Module**: Handles database operations within the Jimber system. It includes a Storage struct with an instance of `gorm.DB` and implements all necessary `gorm.DB` methods.

- **Testing Module**: Provides a testing framework for writing unit tests for the server endpoints. Ensures comprehensive test coverage for the Jimber system.

## Installation

To use the Jimber SDK, follow these steps:

1. Clone the Jimber SDK repository:

```bash
   git clone https://github.com/Mahmoud-Emad/jimber-sdk.git
```

2. Install the necessary dependencies:

```bash
go mod tidy
```

3. Build the Jimber binary:

```bash
go build -o jimber cmd/cli-client/main.go
```

## Usage

1. Run the Jimber Server:
for name the serving helper implemnted inside the `testing` folder, will move it inside a folder called scripts.

```bash
go run testing/main.go
```

2. Run the Jimber CLI client:

```bash
./jimber <command> [flags]
```

Available commands:

- init: Initialize a new Jimber project.
- connect: Connect to an existing Jimber project.
- add: Add flags or variables to the Jimber project.
- push: Push the changes to the Jimber server.
- pull: Pull the latest changes from the Jimber server.

3. Access the Jimber server API endpoints:

```go
package main
// Import the required packages
import (
    jimberSDK "jimber.com/sdk"
)

func main() {
  jimber := jimberSDK.NewJimber("localhost", "8080")
  jimber.RunServer()
}

```

The server API endpoints are now accessible at http://localhost:8080.

## Contributions

to the Jimber SDK are welcome! If you find any issues or have suggestions for improvements, please submit a pull request or open an issue in the GitHub repository

## License

The Jimber SDK is released under the MIT License