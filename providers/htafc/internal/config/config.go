package config

import "time"

type Config struct {
	TeamID          string
	FeedURL         string
	PoolingInterval time.Duration
	ArticlesCount   int
}

func New() *Config {
	return &Config{
		TeamID:          "Huddersfield Town",
		FeedURL:         "https://www.htafc.com/api/incrowd",
		PoolingInterval: time.Second * 30,
		ArticlesCount:   50, // We request 50 articles, we should do less than that
	}
}
