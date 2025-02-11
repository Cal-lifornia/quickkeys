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

type KeyBind struct {
	Keys    string `json:"keys"`
	Command string `json:"cmd"`
	Desc    string `json:"desc,omitempty"`
}

type KeyGroup struct {
	Name string    `json:"name"`
	Keys []KeyBind `json:"keys"`
}
