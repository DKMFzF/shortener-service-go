package configs

import (
	"shortener/internal/configs/flags"
)

type Config struct {
	Addr string
}

func New() *Config {
	addr := flags.ParseFlags()

	if addr == "" {
		// TODO: сделать подкачку .env файла
		addr = "8080"
	}

	return &Config{
		Addr: addr,
	}
}
