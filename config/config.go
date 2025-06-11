package config

import (
	"fmt"
	"goapp/internal/server"
	"goapp/pkg/logger"

	"github.com/spf13/viper"
)

type Config struct {
	App    AppConfig
	Server server.Config
	Logger logger.Config
}

type AppConfig struct {
	Mode string
}

func New(filename string) (*Config, error) {
	v := viper.New()
	v.SetConfigFile(filename)
	v.AddConfigPath(".")
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	var c Config
	if err := v.Unmarshal(&c); err != nil {
		return nil, fmt.Errorf("unable to decode config into struct, %v", err)
	}

	return &c, nil
}
