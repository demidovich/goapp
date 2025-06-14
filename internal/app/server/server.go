package server

import (
	"goapp/pkg/logger"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Config struct {
	Listen                string
	MaxBodySize           string
	GzipLevel             int
	ReadTimeout           time.Duration
	WriteTimeout          time.Duration
	ResponsePrettyEnabled bool
	ResponseStackEnabled  bool
}

type Server struct {
	config   Config
	logger   *logger.Logger
	router   *chi.Mux
	handlers handlers
}

func New(c Config, l *logger.Logger) *Server {
	response := newResponse(c)

	s := &Server{
		config:   c,
		logger:   l,
		router:   chi.NewRouter(),
		handlers: newHandlers(l, response),
	}

	s.setupMiddlewares()
	s.setupLogger()
	s.setupRoutes()

	return s
}

func (s *Server) Run() error {
	s.logger.Info("HTTP server starting on " + s.config.Listen)

	return http.ListenAndServe(s.config.Listen, s.router)
}

func (s *Server) setupMiddlewares() {
	s.logger.Info("HTTP server setup middlewares")

	s.router.Use(middleware.Recoverer)
	s.router.Use(middleware.RequestID)
	s.router.Use(middleware.RealIP)
	s.router.Use(middleware.Timeout(s.config.ReadTimeout * time.Second))

	// router.Use(middleware.Compress(s.config.GzipLevel))

	// s.router.Use(cors.Handler(cors.Options{
	// 	AllowedOrigins: []string{"https://*", "http://*"},
	// 	// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
	// 	AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	// 	AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
	// 	ExposedHeaders:   []string{"Link"},
	// 	AllowCredentials: false,
	// 	MaxAge:           300, // Maximum value not ignored by any of major browsers
	// }))
}
