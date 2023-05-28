package main

import (
	"github.com/Mahmoud-Emad/jimber/server"
)

func main() {
	server := server.Server{
		Port: "8080",
		Host: "localhost",
	}
	server.Serve()
}
