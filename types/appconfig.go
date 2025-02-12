package types

type AppConfig struct {
	Name string `toml:"name"`
	// Alias for the app name to make searching easier
	Alias []string `toml:"aliases"`
	// Leading slash means folder otherwise it's a file
	ConfigPath string `toml:"path"`
	findKey    func(args args) []KeyBind
}

// Arguments for the findKey function in an AppConfig
type args struct {
	Key     string `default:".*"`
	Command string `default:".*"`
	Desc    string `default:".*"`
}
