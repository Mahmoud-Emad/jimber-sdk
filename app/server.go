package app

import (
	"errors"
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
	log.Info().Msgf("Server is listening on %s:%s", s.Host, s.Port)

	// Define a handler function for the root route ("/")
	rootHandler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, World!")
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", s.Host, s.Port),
		Handler: http.HandlerFunc(rootHandler),
	}

	go func() {
		log.Info().Msg("From here")
		if err := srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			log.Fatal().Err(err).Msg("HTTP server error")
		}
		log.Info().Msg("Stopped serving new connections")
	}()

	return nil
}

func initZerolog() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
}
