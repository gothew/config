package config

var defaultConfig = NewWithOptions()

func New() *ConfigParser {
	return NewWithOptions()
}

func NewWithOptions() *ConfigParser {
	return &ConfigParser{
		Config{
			AppDir:         AppDirDefault,
			ConfigFileName: ConfigFileNameDefault,
		},
	}
}

func SetConfigFileName(configFileName string) {
	defaultConfig.SetConfigFileName(configFileName)
}

func ParseConfig() {
	defaultConfig.ParserConfig()
}
