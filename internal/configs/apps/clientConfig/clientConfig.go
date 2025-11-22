package clientConfig

import "shortener/internal/configs/flags"

type ClientConfig struct {
	Url string
}

func New() *ClientConfig {
	endpoint := flags.ParseEndpoint()
	if endpoint == "" {
		endpoint = "ping"
	}

	return &ClientConfig{
		Url: "http://localhost:" + "8080" + "/api" + endpoint,
	}
}
