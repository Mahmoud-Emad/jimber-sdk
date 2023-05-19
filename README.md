# Jimber

Jimber is a web-based and CLI tool for managing environment variables for users and teams.

## Overview

Jimber allows users to create and manage their own environment variables as well as share variables within their teams. It provides a server backend deployed online and a simple web-based user interface for managing variables. Additionally, it offers a command-line interface (CLI) tool that can be installed via `go get` to interact with the server and manage variables.

## Features

- User Authentication: Users can register, log in, and manage their accounts.
- Role-based Permissions: Users have different roles and permissions, allowing for fine-grained access control.
- Environment Variable Management: Users can create, update, delete, and retrieve their own environment variables.
- Team Collaboration: Users can share variables within teams and manage team-specific settings.
- Web-based UI: A user-friendly web-based interface for interacting with the server and managing variables.
- CLI Tool: A command-line interface (CLI) tool for accessing and managing variables from the command line.

## Technologies Used

- Golang: Backend server and CLI tool development.
- [Your chosen web framework]: Web-based UI development.
- [Your chosen database]: Data storage and management.
- [Other libraries or dependencies you use]: [List any additional libraries or dependencies used in the project.]

## Installation and Usage

### Server Setup

1. Clone the Jimber repository.
2. Install the necessary dependencies using `go mod download`.
3. Configure the database connection settings in the server configuration file.
4. Build and run the server using `go run server.go`.

### Web-based UI

1. Install the required dependencies for the web-based UI (e.g., Node.js, npm or yarn).
2. Navigate to the `web-ui` directory.
3. Install the necessary packages using `npm install` or `yarn install`.
4. Configure the API endpoint in the web UI codebase.
5. Start the web-based UI development server using `npm start` or `yarn start`.

### CLI Tool

1. Install the CLI tool by running `go get github.com/your-username/jimber/cli`.
2. Authenticate with the server using the CLI tool (`jimber login`) and follow the prompts.
3. Use the CLI tool commands to manage environment variables (`jimber add`, `jimber get`, etc.).

## Contributing

Contributions are welcome! If you'd like to contribute to Jimber, please follow the guidelines outlined in [CONTRIBUTING.md](CONTRIBUTING.md).

## License

[MIT License](LICENSE)
