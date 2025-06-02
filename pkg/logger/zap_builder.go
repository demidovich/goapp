package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func newZaplog(config Config) (*zap.Logger, error) {
	var encoderConfig zapcore.EncoderConfig
	if config.DevMode {
		encoderConfig = zap.NewDevelopmentEncoderConfig()
	} else {
		encoderConfig = zap.NewProductionEncoderConfig()
	}

	encoderConfig.LevelKey = "level"
	encoderConfig.CallerKey = "caller"
	encoderConfig.TimeKey = "time"
	encoderConfig.NameKey = "name"
	encoderConfig.MessageKey = "message"
	encoderConfig.StacktraceKey = "trace"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	var encoder zapcore.Encoder
	if config.Encoding == "console" {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	} else {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	}

	level, err := zapcore.ParseLevel(config.Level)
	if err != nil {
		return nil, err
	}

	writer := zapcore.AddSync(os.Stdout)
	core := zapcore.NewCore(
		encoder,
		writer,
		zap.NewAtomicLevelAt(level),
	)

	options := []zap.Option{}
	// if config.WithCaller {
	// 	options = append(options, zap.AddCaller())
	// 	options = append(options, zap.AddCallerSkip(1))
	// }

	// if config.WithStacktrace {
	// 	options = append(options, zap.AddStacktrace(level))
	// }

	return zap.New(core, options...), nil
}
