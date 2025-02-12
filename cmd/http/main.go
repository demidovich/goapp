package main

import (
	"fmt"
	"goapp-boilerplate/config"
	"net/http"
)

func main() {
	fmt.Println("Starting http server")

	cfg := configOrFail("./config/config.yml")

	fmt.Printf("Listen %s\n", cfg.Server.Listen)
	http.HandleFunc("/", homeHandler)
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

func homeHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "<h1 style='margin: 50px'>goapp boilerplate</h1>")
}
