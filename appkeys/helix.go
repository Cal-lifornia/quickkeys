package appkeys

import (
	"os"
	"strings"

	"github.com/Cal-lifornia/quickkeys/appkeys/parsers"
	"go.uber.org/zap"
)

var helixConfig AppConfig = AppConfig{
	Name:       "Helix",
	Alias:      []string{"hx"},
	ConfigPath: "HOME/.config/helix/config.toml",
}

func getHelixKeys(conf *AppConfig) (*[]parsers.Entry, error) {
	localLogger := logger.With(
		zap.String("file", conf.ConfigPath),
	)

	localLogger.Debug("Starting parse of Helix config")

	file, err := os.Open(conf.ConfigPath)
	if err != nil {
		localLogger.Error("failed to open Helix config file")
		return nil, err
	}
	parsedFile, err := parsers.TomlParser.Parse(conf.ConfigPath, file)

	if err != nil {
		localLogger.Error("failed to parse Helix config file")
		return nil, err

	}

	var keysEntries []parsers.Entry = []parsers.Entry{}

	for _, entry := range parsedFile.Entries {
		if strings.Contains(entry.Section.Name, "keys") {
			keysEntries = append(keysEntries, *entry)
		}
	}

	return &keysEntries, nil
}
