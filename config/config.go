package config

import "github.com/Cal-lifornia/quickkeys/reader"

type Config struct {
	LogLevel   string             `toml:"log_level" default:"info"`
	RipGrep    bool               `toml:"use_ripgreg" default:"false"`
	FD         bool               `toml:"use_fd" default:"false"`
	AppConfigs []reader.AppConfig `toml:"apps"`
}
