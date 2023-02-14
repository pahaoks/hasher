package config

import (
	"time"

	"example.com/hasher/pkg/envconfig"
)

// Application specific config structure
type Config struct {
	envconfig.DefaultConfig

	HashTTL  time.Duration `envconfig:"HASH_TTL" default:"5m"`
	GrpcPort string        `envconfig:"GRPC_PORT" default:"8081"`
	HttpPort string        `envconfig:"HTTP_PORT" default:"8080"`
}

// Init config or panic, if something is missing
func Init() *Config {
	cfg := &Config{}

	envconfig.Init(cfg)

	return cfg
}
