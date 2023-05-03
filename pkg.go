package config

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

var defaultConfig = NewWithOptions(ConfigParser{
	config: Config{
		ConfigFileName: ConfigFileNameDefault,
	},
})

func New() *ConfigParser {
	return NewWithOptions(ConfigParser{})
}

func NewWithOptions(config ConfigParser) *ConfigParser {
	c := &ConfigParser{
		config: Config{
			ConfigFileName: ConfigFileNameDefault,
		},
	}

	if c.config.AppDir == "" {
		pwd, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		slice := strings.Split(pwd, string(filepath.Separator))
		c.config.AppDir = slice[len(slice)-1]
	}

	if c.configServiceOptions == nil {
		c.configServiceOptions = func() ConfigOptions {
			return ConfigOptions{}
		}
	}

	return c
}

func SetConfigFileName(configFileName string) {
	defaultConfig.SetConfigFileName(configFileName)
}

func SetAppDir(appDir string) {
	defaultConfig.SetAppDir(appDir)
}

func SetConfigOptions(configServiceOptions ConfigServiceOptions) {
	defaultConfig.SetConfigOptions(configServiceOptions)
}

func ParseConfig() (ConfigOptions, error) {
	return defaultConfig.ParserConfig()
}
