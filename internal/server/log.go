package server

import "github.com/go-chi/httplog/v3"

// var SchemaECS = httplog.Schema{
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

var logSchemaECS = httplog.Schema{
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
