package app

import (
	"fmt"
	"net/http"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type Server struct {
	Host string
	Port string
}

// NewServer creates a new server with host and port attributes.
func NewServer(host string, port string) *Server {
	return &Server{
		Port: port,
		Host: host,
	}
}

// Start the server
func (s *Server) Serve() error {
	initZerolog()
	log.Info().Msgf("Server is listening on http://%s:%s", s.Host, s.Port)

	// Define a handler function for the root route ("/")
	rootHandler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, World!")
	}

	// Define a handler function for the root route ("/hello")
	helloHandler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Just Hello!")
	}

	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/hello", helloHandler)

	err := http.ListenAndServe(fmt.Sprintf("%s:%s", s.Host, s.Port), nil)

	if err != nil {
		return err
	}

	return nil
}

func initZerolog() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
}
