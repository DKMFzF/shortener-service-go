package clientConfig

type ClientConfig struct {
	Addr string
}

func New() *ClientConfig {
	return &ClientConfig{
		Addr: "5050",
	}
}
