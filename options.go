package config

// ConfigService returns the default credentials for the application.
type ConfigServiceOptions = func() ConfigOptions

type Config struct {
	AppDir         string
	ConfigFileName string
}

// SettingsConfig struct represents the config for the credentials.
type SettingsConfig struct {
	Prefix string `yaml:"prefix"`
	Querys string `yaml:"querys"`
	Path   string `yaml:"path"`
}

type configError struct {
	configDir string
	parser    *ConfigParser
	err       error
	config    Config
}

// ConfigParser is the parser for the config file.
type ConfigParser struct {
	config               Config
	configServiceOptions ConfigServiceOptions
}

type ConfigOptions struct {
	Services map[string]SettingsConfig `yaml:"services"`
}
