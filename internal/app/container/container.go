package container

import (
	"goapp/config"
	"goapp/pkg/logger"
)

type Container struct {
	config config.Config
	logger *logger.Logger
}

func New(c config.Config, l *logger.Logger) *Container {
	return &Container{
		config: c,
		logger: l,
	}
}
