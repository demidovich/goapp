package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", homeHandler)
	http.ListenAndServe(":7100", nil)
}

func homeHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "<h1 style='margin: 50px'>goapp boilerplate</h1>")
}
