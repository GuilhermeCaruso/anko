package configuration

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

const (
	FILE_NAME = "anko"
)

// ConfigurationFile is the main structure of the anko project.
// All the information contained in anko.yaml is converted using
// the ConfigurationFile structure.
type ConfigurationFile struct {
	Application Application `yaml:"application"`
}

// Application is responsible to mapping the datails of the
// root application.
type Application struct {
	RootPath   string `yaml:"root_path"`
	ExecPath   string `yaml:"exec_path"`
	Language   string `yaml:"language"`
	ShowBanner bool   `yaml:"show_banner"`
	Watch      Watch  `yaml:"watch"`
}

// Watch contains all files and extensions to be observed
// during the application cycle.
type Watch struct {
	Extensions []string `yaml:"extensions"`
	Files      []string `yaml:"files"`
}

var globalConfiguration *ConfigurationFile

func readFile() (*ConfigurationFile, error) {
	fileName, err := getFileNameIfExists()

	if err != nil {
		return nil, err
	}

	byteContent, _ := ioutil.ReadFile(fileName)

	var configuration ConfigurationFile

	err = yaml.Unmarshal(byteContent, &configuration)

	if err != nil {
		return nil, errors.New("Fail to read anko configuration file")
	}

	globalConfiguration = &configuration

	return globalConfiguration, nil
}

// Init is the basic configuration function, responsible to
// read and defining the configuration
func Init() (*ConfigurationFile, error) {
	if globalConfiguration == nil {
		return readFile()
	}
	return globalConfiguration, nil
}

// Get is the getter of the configurationFile
func Get() *ConfigurationFile {
	return globalConfiguration
}

func getFileNameIfExists() (string, error) {
	yamlName := fmt.Sprintf("%s.yaml", FILE_NAME)
	ymlName := fmt.Sprintf("%s.yml", FILE_NAME)

	if _, err := os.Stat(yamlName); !os.IsNotExist(err) {
		return yamlName, nil
	}

	if _, err := os.Stat(ymlName); !os.IsNotExist(err) {
		return ymlName, nil
	}

	return "", errors.New("Anko file not found")
}
