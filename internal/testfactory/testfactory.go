package testfactory

import (
	"fmt"
	"goapp/internal/server"
	"goapp/pkg/logger"
	"math/rand"
	"time"

	"github.com/jmoiron/sqlx"
)

type Factory struct {
	db             *sqlx.DB
	profileFactory *profileFactory
}

func New(s *server.Server, log *logger.Logger) *Factory {
	return &Factory{
		db:             s.DB(),
		profileFactory: newProfileFactory(s, log),
	}
}

func (f *Factory) Profile() *profileFactory {
	return f.profileFactory
}

func (f *Factory) UniqueEmail() string {
	return UniqueEmail()
}

func UniqueEmail() string {
	nano := time.Now().UnixNano()
	nums := rand.Intn(900_000) + 1
	return fmt.Sprintf("fake-%d-%d@fake.com", nano, nums)
}
