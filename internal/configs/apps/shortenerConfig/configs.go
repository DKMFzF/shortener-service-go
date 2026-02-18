package shortenerConfig

import (
	"shortener/internal/configs/flags"
	"time"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	Addr                 string        `env:"PORT" envDefault:"8080"`
	Version              string        `env:"VERSION" envDefault:"1"`
	GracefulShutdownTime time.Duration `env:"GRACEFUL_SHUTDOWN_TIME" envDefault:"5s"`
}

func New() *Config {
	_ = godotenv.Load()

	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		return nil
	}

	addr := flags.ParsePort()
	if addr != "" {
		cfg.Addr = addr
	}

	return &cfg
}
