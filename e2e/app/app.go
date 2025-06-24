package app

import (
	"goapp/config"
	"goapp/internal/server"
	"goapp/pkg/logger"
	"io"

	"github.com/arthurkushman/buildsqlx"
	"github.com/jmoiron/sqlx"
	"github.com/steinfletcher/apitest"
)

var srv *server.Server

func serverInstance() *server.Server {
	if srv == nil {
		srv = newServerInstance()
	}

	return srv
}

func newServerInstance() *server.Server {
	cfg := newConfig("../config/config.yml")
	log := logger.New(io.Discard, cfg.Logger)

	srv := server.New(cfg, log)
	srv.Init()

	return srv
}

func newConfig(file string) *config.Config {
	cfg, err := config.New(file)
	if err != nil {
		panic(err)
	}

	cfg.Postgres.Dbname = "test_db"

	return cfg
}

func DB() *sqlx.DB {
	return serverInstance().DB()
}

func DBQueryBuilder() *buildsqlx.DB {
	return buildsqlx.NewDb(
		buildsqlx.NewConnectionFromDb(
			serverInstance().DB().DB,
		),
	)
}

func API() *apitest.APITest {
	router := serverInstance().Router()

	return apitest.New().Handler(router)
}
