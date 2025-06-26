package request

import (
	"encoding/json"
	"fmt"
	"goapp/pkg/logger"
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
)

func Logger(log *logger.Logger, request *http.Request) *logger.Logger {
	id, _ := request.Context().Value(middleware.RequestIDKey).(string)

	return log.With("http.request.id", id)
}

func DTOFromJSON(dto any, request *http.Request) error {
	decoder := json.NewDecoder(request.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&dto)
	if err != nil {
		return fmt.Errorf("invalid JSON request: %w", err)
	}

	return nil
}
