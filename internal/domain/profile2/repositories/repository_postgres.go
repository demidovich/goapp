package repositories

import (
	"context"
	"goapp/internal/domain/profile2"

	"github.com/demidovich/failure"
	"github.com/jmoiron/sqlx"
)

type repositoryPostgres struct {
	db *sqlx.DB
}

func NewPostgresRepository(db *sqlx.DB) *repositoryPostgres {
	return &repositoryPostgres{db: db}
}

func (r *repositoryPostgres) NextID(ctx context.Context) (int, error) {
	var val int
	err := r.db.QueryRowxContext(ctx, "select nextval('profile_id_seq') as val").Scan(&val)

	return val, err
}

func (r *repositoryPostgres) Create(ctx context.Context, p *profile2.Profile) error {
	_, err := r.db.ExecContext(
		ctx,
		"insert into profile (id, email, password_hash) values ($1, $2, $3)",
		p.ID,
		p.Email,
		p.PasswordHash,
	)

	return err
}

func (r *repositoryPostgres) SearchByEmail(ctx context.Context, email string) (*profile2.Profile, error) {
	item := &profile2.Profile{}

	if email == "" {
		return item, failure.New("передан пустой email")
	}

	err := r.db.GetContext(ctx, item, "select * from profile where email = lower($1) limit 1", email)

	return item, err
}
