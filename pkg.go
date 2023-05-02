package config

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

var defaultConfig = NewWithOptions()

func New() *ConfigParser {
	return NewWithOptions()
}

func NewWithOptions() *ConfigParser {
	c := &ConfigParser{
		Config{
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
	return c
}

func SetConfigFileName(configFileName string) {
	defaultConfig.SetConfigFileName(configFileName)
}

func ParseConfig() {
	defaultConfig.ParserConfig()
}
