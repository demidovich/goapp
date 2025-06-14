package logger

import (
	"fmt"
	"io"
	"log/slog"
)

type Config struct {
	Encoding string
	Level    string
}

type Logger struct {
	config Config
	slog   *slog.Logger
}

func New(w io.Writer, config Config) *Logger {
	return &Logger{
		slog: newSlog(w, config),
	}
}

func (l *Logger) Slog() *slog.Logger {
	return l.slog
}

func (l *Logger) Debug(msg string, args ...any) {
	l.slog.Debug(msg, args...)
}

func (l *Logger) Debugf(format string, args ...any) {
	l.slog.Debug(fmt.Sprintf(format, args...))
}

func (l *Logger) Info(msg string, args ...any) {
	l.slog.Info(msg, args...)
}

func (l *Logger) Infof(format string, args ...any) {
	l.slog.Info(fmt.Sprintf(format, args...))
}

func (l *Logger) Warn(msg string, args ...any) {
	l.slog.Warn(msg, args...)
}

func (l *Logger) Warnf(format string, args ...any) {
	l.slog.Warn(fmt.Sprintf(format, args...))
}

func (l *Logger) Error(msg string, args ...any) {
	l.slog.Error(msg, args...)
}

func (l *Logger) Errorf(format string, args ...any) {
	l.slog.Error(fmt.Sprintf(format, args...))
}

func (l *Logger) With(args ...any) *Logger {
	return &Logger{
		slog: l.slog.With(args...),
	}
}

func (l *Logger) WithGroup(name string, args ...any) *Logger {
	return &Logger{
		slog: l.slog.WithGroup(name).With(args...),
	}
}

func (l *Logger) WithError(err error) *Logger {
	args := []any{
		"message",
		err.Error(),
	}

	type stackable interface {
		Stack() []string
	}

	if e, ok := err.(stackable); ok {
		stack := e.Stack()
		switch len(stack) {
		case 0:
		case 1:
			args = append(args, "caller", stack[0])
		default:
			args = append(args, "caller", stack[0], "stack", stack)
		}
	}

	return &Logger{
		slog: l.slog.WithGroup("error").With(args...),
	}
}
