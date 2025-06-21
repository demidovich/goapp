package server

import (
	"goapp/config"
	"goapp/internal/domain/health"
	"goapp/internal/utility/rest/respond"
	"goapp/pkg/logger"
	"goapp/pkg/postgres"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
)

type Server struct {
	Version    string
	config     *config.Config
	logger     *logger.Logger
	router     *chi.Mux
	httpServer *http.Server
	db         *sqlx.DB
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

	s.initDB()
	s.initDomain()
	s.initRoutes()
}

func (s *Server) initDB() {
	s.logger.Info("DB connecting...")

	db, err := postgres.NewConnection(s.config.Postgres)
	if err != nil {
		panic(err)
	}
	s.db = db

	s.logger.Infof("DB connected on %s:%d", s.config.Postgres.Host, s.config.Postgres.Port)
}

func (s *Server) Run() {
	s.logger.Infof("REST server starting on %s", s.config.Rest.Listen)

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
