package main

import (
	"goapp/config"
	"goapp/internal/app/server"
	"goapp/pkg/logger"
	"os"

	"github.com/demidovich/failure"
)

func main() {
	cfg := newConfig("./config/config.yml")
	log := logger.New(os.Stdout, cfg.Logger)

	log.Info("Init configuration ./config/config.yml")

	failure.SetStackMode(failure.StackModeRoot)
	failure.SetStackRootDir("./")

	srv := server.New(cfg.Server, log)
	err := srv.Run()
	if err != nil {
		log.Error(err.Error())
	}
}

func newConfig(file string) config.Config {
	instance, err := config.New(file)
	if err != nil {
		panic(err)
	}

	return *instance
}
