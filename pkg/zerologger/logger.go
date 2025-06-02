package zerologger

import (
	"os"

	"github.com/rs/zerolog"
)

type Config struct {
	Encoding          string
	Level             string
	CallerEnabled     bool
	StacktraceEnabled bool
}

type Logger struct {
	encoding          string
	level             string
	callerEnabled     bool
	stacktraceEnabled bool
	zlog              *zerolog.Logger
}

func New(config Config) (*Logger, error) {
	zlog, err := newZerolog(config)
	if err != nil {
		return nil, err
	}

	return &Logger{
		encoding:          config.Encoding,
		level:             config.Level,
		callerEnabled:     config.CallerEnabled,
		stacktraceEnabled: config.StacktraceEnabled,
		zlog:              zlog,
	}, nil
}

func newZerolog(config Config) (*zerolog.Logger, error) {
	level, err := zerolog.ParseLevel(config.Level)
	if err != nil {
		return nil, err
	}

	zerolog.SetGlobalLevel(level)
	log := zerolog.New(os.Stdout)

	return &log, nil
}
