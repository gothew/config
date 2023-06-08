package config

// ConfigService returns the default credentials for the application.
type ConfigServiceOptions func() ConfigOptions

type Config struct {
	AppDir         string
	ConfigFileName string
}

// Services represents the config for the credentials.
type Services map[string]interface{}

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

// ConfigOptions is the config options for file config.
type ConfigOptions struct {
	Services interface{}
}
