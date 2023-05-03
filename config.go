package config

import (
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// getDefaultConfigYamlContents returns the default config credentials.
func (parser ConfigParser) getDefaultConfigYamlContents() string {
	defaultConfig := parser.configServiceOptions()
	yaml, _ := yaml.Marshal(defaultConfig)

	return string(yaml)
}

// Error returns the error message for when a config file is not found.
func (e configError) Error() string {
	return fmt.Sprintf(
		`Couldn't find a config.yml configuration file.
Create one under: %s
Example of a config.yml file:
%s
For more info, go to https://github.com/gothew/config
press q to exit.
Original error: %v`,
		path.Join(e.configDir, e.config.AppDir, e.config.ConfigFileName),
		e.parser.getDefaultConfigYamlContents(),
		e.err,
	)
}

// writeDefaultConfingContents writes the default config file contents.
func (parser ConfigParser) writeDefaultConfingContents(newConfigFile *os.File) error {
	_, err := newConfigFile.WriteString(parser.getDefaultConfigYamlContents())

	if err != nil {
		return err
	}

	return nil
}

// createConfigFileIfMissing creates the config file if it doesn't exist.
func (parser ConfigParser) createConfigFileIfMissing(configFilePath string) error {
	if _, err := os.Stat(configFilePath); errors.Is(err, os.ErrNotExist) {
		newConfigFile, err := os.OpenFile(configFilePath, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0666)
		if err != nil {
			return err
		}

		defer newConfigFile.Close()
		return parser.writeDefaultConfingContents(newConfigFile)
	}
	return nil
}

// getConfigFileOrCreateIfMissing returns the config file path or creates the config file if it doesn't exist.
func (parser *ConfigParser) getConfigFileOrCreateIfMissing() (*string, error) {
	var err error

	configDir := os.Getenv("XDG_CONFIG_HOME")
	if configDir == "" {
		configDir, err = os.UserConfigDir()
		if err != nil {
			return nil, configError{parser: parser, configDir: configDir, err: err}
		}
	}

	prsConfigDir := filepath.Join(configDir, parser.config.AppDir)
	err = os.MkdirAll(prsConfigDir, os.ModePerm)
	if err != nil {
		return nil, configError{parser: parser, configDir: configDir, err: err}
	}

	configFilePath := filepath.Join(prsConfigDir, parser.config.ConfigFileName)
	err = parser.createConfigFileIfMissing(configFilePath)
	if err != nil {
		return nil, configError{parser: parser, configDir: configDir, err: err}
	}

	return &configFilePath, nil
}

// parsingError represents an error that ocurred while parsing the config file.
type parsingError struct {
	err error
}

// Error represents an error that ocurred while parsing the config file
func (e parsingError) Error() string {
	return fmt.Sprintf("failed parsing config.yml: %v", e.err)
}

// readConfigFile reads the config file and return config
func (parser *ConfigParser) readConfigFile(path string) (ConfigOptions, error) {
	config := parser.configServiceOptions()
	data, err := os.ReadFile(path)
	if err != nil {
		return config, configError{parser: parser, configDir: path, err: err}
	}

	err = yaml.Unmarshal((data), &config)
	return config, err
}

// ParserConfig parse the config file and returns config
func (configParser *ConfigParser) ParserConfig() (ConfigOptions, error) {
	var config ConfigOptions
	var err error

	configFilePath, err := configParser.getConfigFileOrCreateIfMissing()
	if err != nil {
		return config, parsingError{err: err}
	}

	config, err = configParser.readConfigFile(*configFilePath)
	if err != nil {
		return config, parsingError{err: err}
	}

	return config, nil
}

// GetConfigDir return config
func (configParser ConfigParser) GetConfigDir() string {
	dir, _ := os.UserConfigDir()
	return filepath.Join(dir, configParser.config.AppDir, configParser.config.ConfigFileName)
}

// SetConfigFileName sets the config file.
func (configParser *ConfigParser) SetConfigFileName(configFileName string) {
	configParser.config.ConfigFileName = configFileName
}

// SetAppDir sets the dir config.
func (configParser *ConfigParser) SetAppDir(appDir string) {
	configParser.config.AppDir = appDir
}

// SetConfigOptions sets fn for config interface yaml.
func (configParser *ConfigParser) SetConfigOptions(configServiceOptions ConfigServiceOptions) {
	configParser.configServiceOptions = configServiceOptions
}
