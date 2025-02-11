package appkeys

import (
	"os"
	"strings"

	"github.com/Cal-lifornia/quickkeys/appkeys/parsers"
	"github.com/Cal-lifornia/quickkeys/types"
	"go.uber.org/zap"
)

var helixConfig types.AppConfig = types.AppConfig{
	Name:       "Helix",
	Alias:      []string{"hx"},
	ConfigPath: "HOME/.config/helix/config.toml",
}

func getHelixKeysEntries(confPath string) ([]parsers.Entry, error) {
	localLogger := zap.L().With(
		zap.String("file", confPath),
	)

	localLogger.Debug("Starting parse of Helix config")

	file, err := os.Open(confPath)
	defer func() {
		err = file.Close()
		if err != nil {
			localLogger.Fatal("failed to close files")
		}
	}()
	if err != nil {
		localLogger.Error("failed to open Helix config file")
		return nil, err
	}
	localLogger.Debug("opened helix file successfully")
	parsedFile, err := parsers.TomlParser.Parse(confPath, file)

	if err != nil {
		localLogger.Error("failed to parse Helix config file")
		return nil, err
	}
	localLogger.Debug("parsed helix config file")

	keysEntries := []parsers.Entry{}

	for _, entry := range parsedFile.Entries {
		if entry.Section != nil {
			if strings.Contains(entry.Section.Name, "keys") {
				keysEntries = append(keysEntries, *entry)
			}
		}
	}

	return keysEntries, nil
}

// func parseHelixKeys(conf *AppConfig) error {

// 	return nil
// }
