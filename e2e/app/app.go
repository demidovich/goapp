package app

import (
	"goapp/config"
	"goapp/internal/server"
	"goapp/internal/testfactory"
	"goapp/pkg/logger"
	"io"

	"github.com/arthurkushman/buildsqlx"
	"github.com/jmoiron/sqlx"
	"github.com/steinfletcher/apitest"
)

var cfg *config.Config
var log *logger.Logger
var srv *server.Server
var factory *testfactory.Factory

func configInstance() *config.Config {
	if cfg == nil {
		var err error
		cfg, err = config.New("../config/config.yml")
		if err != nil {
			panic(err)
		}

		cfg.Postgres.Dbname = "test_db"
	}

	return cfg
}

func loggerInstance() *logger.Logger {
	if log == nil {
		cfg := configInstance()
		log = logger.New(io.Discard, cfg.Logger)
	}

	return log
}

func serverInstance() *server.Server {
	if srv == nil {
		srv = server.New(
			configInstance(),
			loggerInstance(),
		)
		srv.Init()
	}

	return srv
}

func Factory() *testfactory.Factory {
	if factory == nil {
		factory = testfactory.New(
			serverInstance(),
			loggerInstance(),
		)
	}

	return factory
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
