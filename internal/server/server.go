package server

import (
	"goapp/config"
	"goapp/internal/domain/health"
	"goapp/internal/utility/rest/respond"
	"goapp/pkg/logger"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	Version    string
	config     *config.Config
	logger     *logger.Logger
	router     *chi.Mux
	httpServer *http.Server
	health     *health.Handler
}

func New(cfg *config.Config, log *logger.Logger) *Server {
	s := Server{
		config: cfg,
		logger: log,
		router: chi.NewRouter(),
	}

	return &s
}

func (s *Server) Init() {
	respond.SetPrettyJSONEnabled(s.config.Rest.ResponsePrettyEnabled)
	respond.SetErrorStackEnabled(s.config.Rest.ResponseStackEnabled)

	s.initDomain()
	s.initRoutes()
}

func (s *Server) Run() {
	s.logger.Info("REST server starting on " + s.config.Rest.Listen)

	s.httpServer = &http.Server{
		Addr:    s.config.Rest.Listen,
		Handler: s.router,
		// ReadHeaderTimeout: s.config.Rest.ReadTimeout,
	}

	err := s.httpServer.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
