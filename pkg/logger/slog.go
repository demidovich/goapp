package logger

import (
	"io"
	"log"
	"log/slog"
	"time"

	"github.com/lmittmann/tint"
)

func newSlog(w io.Writer, config Config) *slog.Logger {
	var logger *slog.Logger

	switch config.Encoding {
	case "json":
		logger = newSlogJSON(w, config)
	case "text":
		logger = newSlogText(w, config)
	case "pretty":
		logger = newSlogPretty(w, config)
	default:
		log.Fatalf("unknow logger encoding \"%s\"", config.Encoding)
	}

	if logger == nil {
		logger = slog.Default()
	} else {
		slog.SetDefault(logger)
	}

	return logger
}

func newSlogJSON(w io.Writer, config Config) *slog.Logger {
	keysToECS := func(groups []string, a slog.Attr) slog.Attr {
		if a.Key == slog.TimeKey {
			a.Key = "@timestamp"
			return a
		}
		if a.Key == slog.LevelKey {
			a.Key = "log.level"
			return a
		}
		if a.Key == slog.MessageKey {
			a.Key = "message"
			return a
		}
		return a
	}

	options := slog.HandlerOptions{
		Level:       slogLevel(config.Level),
		ReplaceAttr: keysToECS,
	}

	return slog.New(
		slog.NewJSONHandler(w, &options),
	)
}

func newSlogText(w io.Writer, config Config) *slog.Logger {
	options := slog.HandlerOptions{
		Level: slogLevel(config.Level),
	}

	return slog.New(
		slog.NewTextHandler(w, &options),
	)
}

func newSlogPretty(w io.Writer, config Config) *slog.Logger {
	hideKeys := map[string]struct{}{
		"@timestamp":                {},
		"log.level":                 {},
		"message":                   {},
		"url.domain":                {},
		"url.path":                  {},
		"http.response.status_code": {},
		"http.request.method":       {},
		"client.ip":                 {},
		"user_agent.original":       {},
		"event.duration":            {},
		// "http.request.id":           {},
		// "http.request.body.bytes":   {},
		// "http.response.body.bytes":  {},
	}

	renameKeys := map[string]string{
		"url.domain":               "host",
		"client.ip":                "ip",
		"user_agent.original":      "agent",
		"http.request.id":          "id",
		"http.request.body.bytes":  "in",
		"http.response.body.bytes": "out",
	}

	replaceAttr := func(groups []string, a slog.Attr) slog.Attr {
		if len(groups) > 0 {
			return a
		}

		if _, ok := hideKeys[a.Key]; ok {
			return slog.Attr{}
		}

		if value, ok := renameKeys[a.Key]; ok {
			a.Key = value
		}

		return a
	}

	options := &tint.Options{
		Level:       slogLevel(config.Level),
		TimeFormat:  time.DateTime,
		ReplaceAttr: replaceAttr,
	}

	return slog.New(
		tint.NewHandler(w, options),
	)
}

func slogLevel(str string) slog.Level {
	var level slog.Level
	var err = level.UnmarshalText([]byte(str))
	if err != nil {
		panic(err)
	}

	return level
}
