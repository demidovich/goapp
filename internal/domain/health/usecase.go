package health

import (
	"goapp/pkg/logger"
)

type Response struct {
	Status  string `json:"status"`
	Details struct {
		Database struct {
			Status  string `json:"status"`
			Message string `json:"message"`
		} `json:"database"`
	} `json:"details,omitempty"`
}

type Usecase struct {
}

func NewUsecase() *Usecase {
	return &Usecase{}
}

func (s *Usecase) Health(log *logger.Logger) (Response, error) {
	r := Response{
		Status: "UP",
	}

	return r, nil
}
