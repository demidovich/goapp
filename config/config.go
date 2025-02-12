package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	App    AppConfig
	Server ServerConfig
}

type AppConfig struct {
	Mode string
}

type ServerConfig struct {
	Listen       string
	MaxBodySize  string
	GzipLevel    int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
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
