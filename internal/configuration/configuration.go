package configuration

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

const (
	FILE_NAME = "anko"
)

type ConfigurationFile struct {
	Version     string      `yaml:"version"`
	Application Application `yaml:"application"`
}

type Application struct {
	RootPath   string `yaml:"root_path"`
	ExecPath   string `yaml:"exec_path"`
	Language   string `yaml:"language"`
	ShowOutput bool   `yaml:"show_output"`
	Watch      Watch  `yaml:"watch"`
}

type Watch struct {
	Extensions []string `yaml:"extensions"`
	Files      []string `yaml:"files"`
}

var globalConfiguration *ConfigurationFile

func readFile() *ConfigurationFile {
	fileName, err := getFileNameIfExists()

	if err != nil {
		log.Fatal(err.Error())
	}

	byteContent, _ := ioutil.ReadFile(fileName)

	var configuration ConfigurationFile

	err = yaml.Unmarshal(byteContent, &configuration)

	if err != nil {
		log.Fatal("Fail to read anko configuration file")
	}

	globalConfiguration = &configuration

	return globalConfiguration
}

func Init() *ConfigurationFile {
	if globalConfiguration == nil {
		return readFile()
	}
	return globalConfiguration
}

func Get() *ConfigurationFile {
	return Init()
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

	return "", errors.New("File doesn`t exists")
}
