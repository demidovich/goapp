package logger

import (
	"fmt"
	"goapp/pkg/errors"
	"strings"

	"go.uber.org/zap"
)

type Config struct {
	Encoding        string
	Level           string
	DevMode         bool
	WithErrorCaller bool
	WithErrorTrace  bool
}

type Log struct {
	encoding        string
	level           string
	devMode         bool
	withErrorCaller bool
	withErrorTrace  bool
	zaplog          *zap.SugaredLogger
}

func New(config Config) (*Log, error) {
	zaplog, err := newZaplog(config)
	if err != nil {
		return nil, err
	}

	instance := &Log{
		encoding:        config.Encoding,
		level:           config.Level,
		devMode:         config.DevMode,
		withErrorCaller: config.WithErrorCaller,
		withErrorTrace:  config.WithErrorTrace,
		zaplog:          zaplog.Sugar(),
	}

	return instance, nil
}

func (l *Log) Debug(args ...interface{}) {
	l.zaplog.Debug(args...)
}

func (l *Log) Debugf(template string, args ...interface{}) {
	l.zaplog.Debugf(template, args...)
}

func (l *Log) Info(args ...interface{}) {
	l.zaplog.Info(args...)
}

func (l *Log) Infof(template string, args ...interface{}) {
	l.zaplog.Infof(template, args...)
}

func (l *Log) Warn(args ...interface{}) {
	l.zaplog.Warn(args...)
}

func (l *Log) Warnf(template string, args ...interface{}) {
	l.zaplog.Warnf(template, args...)
}

func (l *Log) Error(args ...interface{}) {
	l.zaplog.Error(
		l.argsMessage(args...),
		l.argsDetailsString(args...),
	)
}

func (l *Log) Errorf(template string, args ...interface{}) {
	l.zaplog.Errorf(template, args...)
}

func (l *Log) Panic(args ...interface{}) {
	l.zaplog.DPanic(args...)
}

func (l *Log) Panicf(template string, args ...interface{}) {
	l.zaplog.DPanicf(template, args...)
}

func (l *Log) Fatal(args ...interface{}) {
	l.zaplog.Fatal(args...)
}

func (l *Log) Fatalf(template string, args ...interface{}) {
	l.zaplog.Fatalf(template, args...)
}

func (l *Log) argsMessage(args ...interface{}) (results string) {
	if len(args) < 1 {
		return
	}

	switch t := args[0].(type) {
	case string:
		results = t
	case error:
		results = t.Error()
	}

	return
}

func (l *Log) argsDetailsJson(args ...interface{}) (results []zap.Field) {
	if len(args) < 1 {
		return
	}

	if traceErr, ok := args[0].(errors.Stacktracer); ok {
		trace := strings.Builder{}
		for _, frame := range traceErr.Stacktrace().Frames() {
			trace.WriteString(
				fmt.Sprintf("%s:%d (%s)", frame.File, frame.Line, frame.Function),
			)
		}
		results = append(results, zap.String("trace", trace.String()))
	}

	return
}

func (l *Log) argsDetailsString(args ...interface{}) (results []string) {
	if len(args) < 1 {
		return
	}

	results = append(results, fmt.Sprintf("%v+", args[0]))
	// if traceErr, ok := args[0].(errors.Stacktracer); ok {
	// 	trace := strings.Builder{}
	// 	trace.WriteString("\ntrace:")
	// 	for _, frame := range traceErr.Stacktrace().ToSlice() {
	// 		trace.WriteString(
	// 			fmt.Sprintf("\n-> %s:%d (%s)", frame.File, frame.Line, frame.Function),
	// 		)
	// 	}
	// 	results = append(results, trace.String())
	// }

	return
}
