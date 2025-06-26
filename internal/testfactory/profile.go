package testfactory

import (
	"goapp/internal/domain/profile"
	"goapp/internal/server"
	"goapp/pkg/logger"
	"testing"

	"github.com/jackc/fake"
	"github.com/jmoiron/sqlx"
)

type profileFactory struct {
	log      *logger.Logger
	db       *sqlx.DB
	usecases *profile.Usecases
}

func newProfileFactory(s *server.Server, log *logger.Logger) *profileFactory {
	return &profileFactory{
		log:      log,
		db:       s.DB(),
		usecases: s.ProfileUsecases(),
	}
}

func (f *profileFactory) ClearAll(t *testing.T) {
	_, err := f.db.Exec("delete from profile")
	if err != nil {
		t.Errorf("profile factory: errorexecution ClearAll(): %v", err)
	}
}

func (f *profileFactory) New(t *testing.T) *profile.Profile {
	dto := profile.CreateDTO{
		Email:    UniqueEmail(),
		Password: fake.SimplePassword(),
	}

	item, err := f.usecases.Create(t.Context(), f.log, dto)
	if err != nil {
		t.Errorf("profile factory: errorexecution New(): %v", err)
	}

	return &item
}
