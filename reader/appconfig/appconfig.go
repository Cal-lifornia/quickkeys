package appconfig

// For limiting the kind of file type that can be read
type fileType string

type AppConfig struct {
	Name     string
	fileType fileType
	FilePath string
	// TODO: Add func handler
}

const (
	toml fileType = "toml"
	conf fileType = "conf"
	yaml fileType = "yaml"
	kdl  fileType = "kdl"
)
