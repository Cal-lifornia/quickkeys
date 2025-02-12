package config

import (
	"os"

	"github.com/BurntSushi/toml"
	"github.com/Cal-lifornia/quickkeys/types"
	"github.com/spf13/cobra"
)

var config Config

var (
	meta         string = "Meta"
	ctrl         string = "Ctrl"
	shift        string = "Shift"
	altKey       string = "Alt"
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
	LogLevel   string `toml:"log_level" default:"info"`
	Symbols    bool
	AppConfigs []types.AppConfig `toml:"apps"`
}

// Returns a global Config
func C() *Config {
	return &config
}

func (config *Config) Meta() string {
	return meta
}

func (config *Config) Ctrl() string {
	return ctrl
}

func (config *Config) Shift() string {
	return shift
}

func (config *Config) Alt() string {
	return altKey
}

func InitConfig(confPath string) {

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

	initKeys()
}

func initKeys() {
	if config.Symbols == true {
		meta = metaSymbol
		ctrl = ctrlSymbol
		shift = shiftSymbol
		altKey = altKeySymbol
	} else {
		meta = metaText
		ctrl = ctrlText
		shift = shiftText
		altKey = altKeyText
	}
}
