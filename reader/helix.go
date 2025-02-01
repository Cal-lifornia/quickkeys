package reader

var helixConfig AppConfig = AppConfig{
	Name:         "Helix",
	Alias:        []string{"hx"},
	ConfigPath:   ".config/helix/config.toml",
	commentToken: "#",
}

type helixToml struct {
	Keys struct {
		Normal map[string]string `toml:"normal"`
		Insert map[string]string `toml:"insert"`
	}
}
