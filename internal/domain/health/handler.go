package health

import (
	"goapp/internal/utility/rest/request"
	"goapp/internal/utility/rest/respond"
	"goapp/pkg/logger"
	"net/http"
)

type Handler struct {
	usecase *Usecase
	log     *logger.Logger
}

func NewHandler(uc *Usecase, log *logger.Logger) *Handler {
	return &Handler{
		usecase: uc,
		log:     log,
	}
}

func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {
	log := request.Logger(h.log, r)
	data, err := h.usecase.Health(r.Context(), log)

	respond.ItemOrFail(w, data, err, log)
}
