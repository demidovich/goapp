package respond

import (
	"encoding/json"
	"goapp/pkg/logger"
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

var (
	prettyJSONEnabled = true
	errorStackEnabled = true
)

func SetPrettyJSONEnabled(v bool) {
	prettyJSONEnabled = v
}

func SetErrorStackEnabled(v bool) {
	errorStackEnabled = v
}

func Success(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
}

func SuccessOrFail(w http.ResponseWriter, err error, log *logger.Logger) {
	if err != nil {
		Error(w, err, log)
	} else {
		Success(w)
	}
}

func Message(w http.ResponseWriter, message string) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_ = jsonEncoder(w).Encode(messageBody{Message: message})
}

func MessageOrFail(w http.ResponseWriter, message string, err error, log *logger.Logger) {
	if err != nil {
		Error(w, err, log)
	} else {
		Message(w, message)
	}
}

func Created(w http.ResponseWriter, item any) {
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	_ = jsonEncoder(w).Encode(item)
}

func CreatedOrFail(w http.ResponseWriter, item any, err error, log *logger.Logger) {
	if err != nil {
		Error(w, err, log)
	} else {
		Created(w, item)
	}
}

func Item(w http.ResponseWriter, item any) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_ = jsonEncoder(w).Encode(item)
}

func ItemOrFail(w http.ResponseWriter, item any, err error, log *logger.Logger) {
	if err != nil {
		Error(w, err, log)
	} else {
		Item(w, item)
	}
}

func Collection(w http.ResponseWriter, collection []any) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_ = jsonEncoder(w).Encode(collectionBody{Resources: collection})
}

func CollectionOrFail(w http.ResponseWriter, collection []any, err error, log *logger.Logger) {
	if err != nil {
		Error(w, err, log)
	} else {
		Collection(w, collection)
	}
}

func Error(w http.ResponseWriter, err error, log *logger.Logger) {
	log.WithError(err).Error("")

	body := errorBody{}
	body.Error = err.Error()

	if !errorStackEnabled {
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
	_ = jsonEncoder(w).Encode(body)
}

func jsonEncoder(w http.ResponseWriter) *json.Encoder {
	e := json.NewEncoder(w)
	if prettyJSONEnabled {
		e.SetIndent("", "    ")
	}
	return e
}
