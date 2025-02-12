package config

import (
	"os"

	"github.com/BurntSushi/toml"
	"github.com/spf13/cobra"
)

var globalConfig *Config

var (
	metaText     string = "Meta"
	ctrlText     string = "Ctrl"
	shiftText    string = "Shift"
	altKeyText   string = "Alt"
	metaSymbol   string = "Meta"
	ctrlSymbol   string = "Ctrl"
	shiftSymbol  string = "Shift"
	altKeySymbol string = "Alt"
)

type Config struct {
	LogLevel string `toml:"log_level" default:"info"`
	Symbols  bool   `toml:"symbols" default:"false"`
	meta     string
	ctrl     string
	shift    string
	altKey   string
	// AppConfigs []types.AppConfig `toml:"apps"`
}

// Returns a global Config
func C() *Config {
	return globalConfig
}

func (config *Config) Meta() string {
	return config.meta
}

func (config *Config) Ctrl() string {
	return config.ctrl
}

func (config *Config) Shift() string {
	return config.shift
}

func (config *Config) Alt() string {
	return config.altKey
}

func InitConfig(confPath string) {
	var config *Config
	if confPath != "" {

		// Read file
		configFile, err := os.ReadFile(confPath)
		if err != nil {
			cobra.CheckErr(err)
		}

		// Decode file to toml config
		_, err = toml.Decode(string(configFile), &config)
		if err != nil {
			cobra.CheckErr(err)
		}

	} else {
		config = &Config{
			LogLevel: "debug",
			Symbols:  false,
		}
	}
	config.InitKeys()
	SetGlobalConfig(config)
}

func SetGlobalConfig(config *Config) {
	globalConfig = config
}

func (config *Config) InitKeys() {
	if config.Symbols == true {
		config.meta = metaSymbol
		config.ctrl = ctrlSymbol
		config.shift = shiftSymbol
		config.altKey = altKeySymbol
	} else {
		config.meta = metaText
		config.ctrl = ctrlText
		config.shift = shiftText
		config.altKey = altKeyText
	}
}
