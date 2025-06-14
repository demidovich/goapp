package server

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httplog/v3"
)

// var loggerSchemaECS = httplog.Schema{
// 	Timestamp:          "@timestamp",
// 	Level:              "log.level",
// 	Message:            "message",
// 	Error:              "error.message",
// 	ErrorStackTrace:    "error.stack_trace",
// 	RequestURL:         "url.full",
// 	RequestScheme:      "url.scheme",
// 	RequestHost:        "url.domain",
// 	RequestPath:        "url.path",
// 	ResponseStatus:     "http.response.status_code",
// 	RequestProto:       "http.version",
// 	RequestMethod:      "http.request.method",
// 	RequestHeaders:     "http.request.headers",
// 	RequestBody:        "http.request.body.content",
// 	RequestBytes:       "http.request.body.bytes",
// 	RequestReferer:     "http.request.referrer",
// 	RequestBytesUnread: "http.request.body.unread.bytes",
// 	ResponseBody:       "http.response.body.content",
// 	ResponseHeaders:    "http.response.headers",
// 	ResponseBytes:      "http.response.body.bytes",
// 	RequestRemoteIP:    "client.ip",
// 	RequestUserAgent:   "user_agent.original",
// 	ResponseDuration:   "event.duration",
// }

var loggerSchemaECS = httplog.Schema{
	Timestamp:        "@timestamp",
	Level:            "log.level",
	Message:          "message",
	Error:            "error.message",
	ErrorStackTrace:  "error.stack_trace",
	RequestHost:      "url.domain",
	RequestPath:      "url.path",
	RequestMethod:    "http.request.method",
	ResponseStatus:   "http.response.status_code",
	RequestBytes:     "http.request.body.bytes",
	ResponseBytes:    "http.response.body.bytes",
	RequestRemoteIP:  "client.ip",
	RequestUserAgent: "user_agent.original",
	ResponseDuration: "event.duration",
}

func (s *Server) setupLogger() {
	s.logger.Info("HTTP server setup request logger")

	options := &httplog.Options{
		Schema:        loggerSchemaECS,
		RecoverPanics: true,
		Skip: func(req *http.Request, respStatus int) bool {
			return req.URL.String() == "/favicon.ico"
		},
		// // Log all requests with invalid payload as curl command.
		// LogExtraAttrs: func(req *http.Request, reqBody string, respStatus int) []slog.Attr {
		// 	if respStatus == 400 || respStatus == 422 {
		// 		req.Header.Del("Authorization")
		// 		return []slog.Attr{slog.String("curl", httplog.CURL(req, reqBody))}
		// 	}
		// 	return nil
		// },
	}
	s.router.Use(httplog.RequestLogger(s.logger.Slog(), options))

	logRequestID := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			if id, ok := ctx.Value(middleware.RequestIDKey).(string); ok {
				httplog.SetAttrs(ctx, slog.String("http.request.id", id))
			}

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
	s.router.Use(logRequestID)
}
