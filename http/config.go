package http

const (
	DefaultHTTPPort = 8090
	DefaultHTTPHost = "127.0.0.1"
)

type (
	Config struct {
		Host string `json:"host"`
		Port int32  `json:"port"`
	}
)

func NewDefaultConfig() *Config {
	return &Config{
		Host: DefaultHTTPHost,
		Port: DefaultHTTPPort,
	}
}
