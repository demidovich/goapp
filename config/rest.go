package config

import "time"

type Rest struct {
	Listen                string
	MaxBodySize           string
	GzipLevel             int
	ReadTimeout           time.Duration
	WriteTimeout          time.Duration
	ResponsePrettyEnabled bool
	ResponseStackEnabled  bool
}
