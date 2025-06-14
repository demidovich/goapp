package server

import (
	"encoding/json"
	"net/http"

	"github.com/demidovich/failure"
)

type messageBody struct {
	Message string `json:"message,omitempty"`
}

type collectionBody struct {
	Meta struct {
		Page    int `json:"page,omitempty"`
		PerPage int `json:"per_page,omitempty"`
	} `json:"meta,omitempty"`
	Resources []any `json:"resources,omitempty"`
}

type errorBody struct {
	Error  string   `json:"error,omitempty"`
	Token  string   `json:"token,omitempty"`
	Caller string   `json:"caller,omitempty"`
	Stack  []string `json:"stack,omitempty"`
}

type response struct {
	prettyEnabled bool
	stackEnabled  bool
}

func newResponse(c Config) response {
	return response{
		prettyEnabled: c.ResponsePrettyEnabled,
		stackEnabled:  c.ResponseStackEnabled,
	}
}

func (r *response) Success() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, hr *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
	})
}

func (r *response) Message(message string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, hr *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		_ = r.jsonEncoder(w).Encode(messageBody{Message: message})
	})
}

func (r *response) Created(item any) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, hr *http.Request) {
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		_ = r.jsonEncoder(w).Encode(item)
	})
}

func (r *response) Item(item any) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, hr *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		_ = r.jsonEncoder(w).Encode(item)
	})
}

func (r *response) Collection(collection []any) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, hr *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		_ = r.jsonEncoder(w).Encode(collectionBody{Resources: collection})
	})
}

func (r *response) Error(err error) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, hr *http.Request) {
		body := errorBody{}
		body.Error = err.Error()

		if !r.stackEnabled {
		} else if f, ok := err.(failure.Error); ok {
			s := f.Stack()
			switch len(s) {
			case 0:
			case 1:
				body.Caller = s[0]
			default:
				body.Caller = s[0]
				body.Stack = s
			}
		}

		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		_ = r.jsonEncoder(w).Encode(body)
	})
}

func (r *response) jsonEncoder(w http.ResponseWriter) *json.Encoder {
	e := json.NewEncoder(w)
	if r.prettyEnabled {
		e.SetIndent("", "    ")
	}
	return e
}
