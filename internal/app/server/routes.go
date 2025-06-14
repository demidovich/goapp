package server

import (
	"net/http"
)

type appHandler func(w http.ResponseWriter, r *http.Request) http.Handler

func (fn appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if handler := fn(w, r); handler != nil {
		handler.ServeHTTP(w, r)
	}
}

func (s *Server) setupRoutes() {
	s.logger.Info("HTTP server setup routes")

	r := s.router
	r.Handle("GET /", appHandler(s.handlers.Home))
	r.Handle("GET /user/{id}", appHandler(s.handlers.User))
}
