package server

import (
	"goapp/internal/user"
	"goapp/pkg/logger"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5/middleware"
)

type handlers struct {
	logger   *logger.Logger
	response response
}

func newHandlers(l *logger.Logger, r response) handlers {
	return handlers{
		logger:   l,
		response: r,
	}
}

func (h *handlers) requestLogger(r *http.Request) *logger.Logger {
	id, _ := r.Context().Value(middleware.RequestIDKey).(string)
	return h.logger.With("http.request.id", id)
}

func (h *handlers) errorHandler(log *logger.Logger, err error) http.Handler {
	log.WithError(err).Error("")

	return h.response.Error(err)
}

func (h *handlers) Home(w http.ResponseWriter, r *http.Request) http.Handler {
	return h.response.Message("Hello world")
}

func (h *handlers) User(w http.ResponseWriter, r *http.Request) http.Handler {
	id, _ := strconv.Atoi(r.PathValue("id"))
	log := h.requestLogger(r)

	item, err := user.Find(log, id)
	if err != nil {
		return h.errorHandler(log, err)
	}

	return h.response.Item(item)
}
