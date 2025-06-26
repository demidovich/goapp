package profile

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

	if err := u.wantNotExistingEmail(ctx, dto.Email, 0); err != nil {
		return p, failure.Wrapf(err, "")
	}

	if err := u.nextID(ctx, &p.ID); err != nil {
		return p, failure.Wrap(err, "")
	}

	p.SetEmail(dto.Email)
	p.SetPassword(dto.Password)

	_, err := u.db.ExecContext(
		ctx,
		"insert into profile (id, email, password_hash) values ($1, $2, $3)",
		p.ID,
		p.Email,
		p.PasswordHash,
	)

	log.With("email", p.Email).Info("Создан новый профиль")

	return p, err
}

func (u *Usecases) wantNotExistingEmail(ctx context.Context, email string, id int) error {
	existing, err := u.searchByEmail(ctx, email)
	if err == sql.ErrNoRows {
		return nil
	}

	if err != nil {
		return err
	}

	if existing.ID == 0 {
		return nil
	}

	if id == 0 || existing.ID == id {
		return nil
	}

	return errors.New("такой email уже зарегистрирован")
}

func (u *Usecases) nextID(ctx context.Context, id *int) error {
	err := u.db.QueryRowxContext(ctx, "select nextval('profile_id_seq')").Scan(id)
	return err
}
