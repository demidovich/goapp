package server

import (
	"goapp/pkg/logger"
	"log/slog"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httplog/v3"
)

type Config struct {
	Listen       string
	MaxBodySize  string
	GzipLevel    int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type Server struct {
	config Config
	logger *logger.Logger
	router *chi.Mux
}

func New(config Config, logger *logger.Logger) *Server {
	return &Server{
		config: config,
		logger: logger,
		router: chi.NewRouter(),
	}
}

func (s *Server) Run() {
	s.setupMiddlewares()
	s.setupLogger()
	s.setupHandlers()

	s.logger.Info("HTTP server starting on " + s.config.Listen)
	http.ListenAndServe(s.config.Listen, s.router)
}

func (s *Server) setupMiddlewares() {
	s.logger.Info("HTTP server setup middlewares")

	router := s.router
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(s.config.ReadTimeout * time.Second))
	// router.Use(middleware.Compress(s.config.GzipLevel))
}

func (s *Server) setupLogger() {
	s.logger.Info("HTTP server setup request logger")

	s.router.Use(httplog.RequestLogger(s.logger.Slog(), &httplog.Options{
		Schema:        logSchemaECS,
		RecoverPanics: true,
		Skip: func(req *http.Request, respStatus int) bool {
			return req.URL.String() == "/favicon.ico"
		},
	}))

	// Добавление RequestID в контекст логгера запроса
	s.router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			if id, ok := ctx.Value(middleware.RequestIDKey).(string); ok {
				httplog.SetAttrs(ctx, slog.String("http.request.id", id))
			}

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	})
}

// Роутинг и вызов usecases
// Создание логгера HTTP запроса c информацией о HTTP запросе
// Передача логгера в usecases в виде явной зависимости
func (s *Server) setupHandlers() {
	s.logger.Info("HTTP server setup routes")

	requestLogger := func(r *http.Request) *logger.Logger {
		requestID, _ := r.Context().Value(middleware.RequestIDKey).(string)
		return s.logger.With("http.request.id", requestID)
	}

	router := s.router
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	router.Get("/example", func(w http.ResponseWriter, r *http.Request) {
		log := requestLogger(r)
		log.Warn("example warning")

		w.Write([]byte("welcome"))
	})
}
