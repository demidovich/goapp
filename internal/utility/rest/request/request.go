package request

import (
	"goapp/pkg/logger"
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
)

func Logger(log *logger.Logger, r *http.Request) *logger.Logger {
	id, _ := r.Context().Value(middleware.RequestIDKey).(string)
	return log.With("http.request.id", id)
}
