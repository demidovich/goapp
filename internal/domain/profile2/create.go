package profile2

import (
	"context"
	"database/sql"
	"errors"
	"goapp/pkg/logger"
	"goapp/pkg/validation"

	"github.com/demidovich/failure"
)

type CreateDTO struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (c *CreateDTO) Validate() error {
	return validation.ValidateStruct(c)
}

func (u *Usecases) Create(ctx context.Context, log *logger.Logger, dto CreateDTO) (Profile, error) {
	p := Profile{}

	if err := dto.Validate(); err != nil {
		return p, failure.Wrapf(err, "")
	}

	if err := u.wantUniqueEmail(ctx, dto.Email, 0); err != nil {
		return p, failure.Wrapf(err, "")
	}

	if newID, err := u.repo.NextID(ctx); err != nil {
		return p, failure.Wrap(err, "")
	} else {
		p.ID = newID
	}

	p.SetEmail(dto.Email)
	p.SetPassword(dto.Password)

	if err := u.repo.Create(ctx, &p); err != nil {
		return p, failure.Wrap(err, "ошибка создания профиля")
	}

	log.With("email", p.Email).Info("создан новый профиль")

	return p, nil
}

func (u *Usecases) wantUniqueEmail(ctx context.Context, email string, id int) error {
	found, err := u.repo.SearchByEmail(ctx, email)

	if err == sql.ErrNoRows {
		return nil
	}

	if err != nil {
		return err
	}

	if found.ID == 0 {
		return nil
	}

	if id == 0 || found.ID == id {
		return nil
	}

	return errors.New("такой email уже зарегистрирован")
}
