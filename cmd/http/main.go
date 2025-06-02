package main

import (
	"fmt"
	"goapp/config"
	"goapp/pkg/logger"
	"net/http"

	"github.com/demidovich/failure"
)

func main() {
	fmt.Println("Starting http server")

	cfg := configOrFail("./config/config.yml")
	log := logOrFail(cfg.Logger)

	failure.SetStackMode(failure.StackModeRoot)
	failure.SetStackRootDir("../../")
	failure.SetStackPrefix(" --- ")

	log.Infof("Listen %s\n", cfg.Server.Listen)
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/example-error", exampleErrorHandler)
	http.ListenAndServe(cfg.Server.Listen, nil)
}

func configOrFail(file string) config.Config {
	fmt.Printf("Init configuration from %s\n", file)

	instance, err := config.New(file)
	if err != nil {
		panic(err)
	}

	return *instance
}

func logOrFail(cfg logger.Config) *logger.Log {
	fmt.Println("Init logger")

	instance, err := logger.New(cfg)
	if err != nil {
		panic(err)
	}

	return instance
}

func homeHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "<h1 style='margin: 50px'>goapp</h1>")
}

func exampleErrorHandler(w http.ResponseWriter, req *http.Request) {
	err := failure.New("example error")
	fmt.Fprintf(w, "%+v", err)
}
