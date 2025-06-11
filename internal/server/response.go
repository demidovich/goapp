package server

import "net/http"

type ResponseSuccess func(w http.ResponseWriter, r *http.Request) error

func (h ResponseSuccess) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h(w, r); err != nil {
		// handle returned error here.
		w.WriteHeader(503)
		w.Write([]byte(err.Error()))
	}
}
