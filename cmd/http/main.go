package main

import (
	"goapp/config"
	"goapp/internal/server"
	"goapp/pkg/logger"
	"os"

	"github.com/demidovich/failure"
)

func main() {
	cfg := newConfig("./config/config.yml")
	log := logger.New(os.Stdout, cfg.Logger)

	log.Info("Init configuration ./config/config.yml")

	failure.SetStackMode(failure.StackModeRoot)
	failure.SetStackRootDir("../../")

	srv := server.New(cfg.Server, log)
	srv.Run()
}

func newConfig(file string) config.Config {
	instance, err := config.New(file)
	if err != nil {
		panic(err)
	}

	return *instance
}
