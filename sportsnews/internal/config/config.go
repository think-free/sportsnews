package config

type Config struct {
	PageSize int64
}

func New() *Config {
	return &Config{
		PageSize: 10,
	}
}
