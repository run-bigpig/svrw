package config

import "sync"

var (
	once       sync.Once
	globalConf *Config
)

type Config struct {
	Address        string
	BiliBiliCookie string
}

func Init(c *Config) {
	once.Do(func() {
		globalConf = c
	})
}

func Get() *Config {
	return globalConf
}
