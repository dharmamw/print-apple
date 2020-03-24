package http

import (
	"net/http"

	"github.com/rs/cors"
	"github.com/vilbert/go-skeleton/pkg/grace"
)

// AppleHandler ...
type AppleHandler interface {
	AppleHandler(w http.ResponseWriter, r *http.Request)
}

// Server ...
type Server struct {
	server *http.Server
	Apple  AppleHandler
}

// Serve is serving HTTP gracefully on port x ...
func (s *Server) Serve(port string) error {
	handler := cors.AllowAll().Handler(s.Handler())
	return grace.Serve(port, handler)
}
