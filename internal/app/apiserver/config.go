package apiserver

import "github.com/dvd-denis/legko-server/internal/app/store"

// Config ...
type Config struct {
	BindAddr string
	LogLevel string `toml:"log_level"`
	GinDebug bool   `toml:"gin_debug"`
	Store    *store.Config
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
		GinDebug: true,
		Store:    store.NewConfig(),
	}
}
