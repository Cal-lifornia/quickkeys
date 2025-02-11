package config

import "github.com/Cal-lifornia/quickkeys/types"

type Config struct {
	LogLevel   string            `toml:"log_level" default:"info"`
	AppConfigs []types.AppConfig `toml:"apps"`
}
