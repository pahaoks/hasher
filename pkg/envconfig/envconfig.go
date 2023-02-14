package envconfig

import (
	"example.com/hasher/pkg/dotenv"
	config "github.com/kelseyhightower/envconfig"
)

// Default config, suitable for any application
type DefaultConfig struct {
	App         string `envconfig:"APP" required:"true"`
	Environment string `envconfig:"ENVIRONMENT" required:"true"`
	LogLevel    string `envconfig:"LOG_LEVEL" default:"debug"`
}

// Init new config
// pass cfg as a pointer to structure
func Init(cfg any) {
	// if there is .env file in root, use it first
	dotenv.Load()

	if err := config.Process("", cfg); err != nil {
		panic(err)
	}
}
