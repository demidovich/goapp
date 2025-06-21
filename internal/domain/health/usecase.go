package health

import (
	"context"
	"goapp/pkg/logger"
	"time"

	"github.com/jmoiron/sqlx"
)

type Response struct {
	Status  string `json:"status"`
	Details struct {
		Database struct {
			Status string `json:"status"`
			Error  string `json:"error,omitempty"`
		} `json:"database"`
	} `json:"details"`
}

type Usecase struct {
	db *sqlx.DB
}

func NewUsecase(db *sqlx.DB) *Usecase {
	return &Usecase{
		db: db,
	}
}

func (u *Usecase) Health(ctx context.Context, log *logger.Logger) (Response, error) {
	r := Response{}
	r.Status = "UP"
	r.Details.Database.Status = "UP"

	u.checkDatabase(ctx, &r)

	return r, nil
}

func (u *Usecase) checkDatabase(ctx context.Context, response *Response) {
	timeoutCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	_, err := u.db.QueryContext(timeoutCtx, "select now()")
	if err != nil {
		response.Status = "DOWN"
		response.Details.Database.Status = "DOWN"
		response.Details.Database.Error = err.Error()
	}
}
