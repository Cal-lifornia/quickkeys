package reader

var helixConfig AppConfig = AppConfig{
	Name:         "Helix",
	Alias:        []string{"hx"},
	ConfigPath:   ".config/helix/config.toml",
	commentToken: "#",
}

type helixToml struct {
	Keys struct {
		Normal []KeyGroup `toml:"normal"`
		Insert []KeyGroup `toml:"insert"`
	} `toml:"keys"`
}
