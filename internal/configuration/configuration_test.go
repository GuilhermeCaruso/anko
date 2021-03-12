package configuration

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func Before(extension string) error {
	content := []byte(`application:
  root_path: "."
  exec_path: "main.go"
  language: go
  watch:
    extensions: 
      - go`)

	err := ioutil.WriteFile(fmt.Sprintf("%s.%s", FILE_NAME, extension), content, 0644)

	if err != nil {
		return err
	}
	return nil
}

func WrongBefore(extension string) error {
	content := []byte(`application:
root_path: "."
 / exec_path: "main.go"`)

	err := ioutil.WriteFile(fmt.Sprintf("%s.%s", FILE_NAME, extension), content, 0644)

	if err != nil {
		return err
	}
	return nil
}

func After(extension string) error {
	globalConfiguration = nil
	return os.Remove(fmt.Sprintf("%s.%s", FILE_NAME, extension))
}

func TestConfigurationFileYml(t *testing.T) {
	if err := Before("yml"); err != nil {
		t.Error("Error to create anko.yml")
	}

	t.Run("Should init the configurationFile correctly", func(it *testing.T) {

		configurationFile, err := Init()

		if err != nil {
			it.Error("Initialization failed")
		}

		expectedLanguage := "go"

		if configurationFile.Application.Language != expectedLanguage {
			it.Errorf("Conversion failed. Expected=%q; Got=%q", expectedLanguage, configurationFile.Application.Language)
		}
	})

	t.Run("Should get the global configurationFile", func(it *testing.T) {
		_, err := Init()

		if err != nil {
			it.Error("Initialization failed")
		}

		configurationFile := Get()
		expectedLanguage := "go"

		if configurationFile.Application.Language != expectedLanguage {
			it.Errorf("Conversion failed. Expected=%q; Got=%q", expectedLanguage, configurationFile.Application.Language)
		}

	})

	if err := After("yml"); err != nil {
		t.Error("Fail to delete anko.yml")
	}

}

func TestNotFoundConfigurationFileYaml(t *testing.T) {
	t.Run("Should'nt return error", func(it *testing.T) {
		_, err := Init()

		if err == nil {
			t.Error("Returned error is invalid, since the file does not exist")
		}
	})

	t.Run("Should return error for invalid formatting", func(it *testing.T) {
		if err := WrongBefore("yaml"); err != nil {
			t.Error("Error to create anko.yaml")
		}

		_, err := Init()

		if err == nil {
			t.Error("Returned error is invalid, since the file formatting is invalid")

		}

		if err := After("yaml"); err != nil {
			t.Error("Fail to delete anko.yaml")
		}
	})

}

func TestConfigurationFileYaml(t *testing.T) {
	if err := Before("yaml"); err != nil {
		t.Error("Error to create anko.yaml")
	}

	t.Run("Should init the configurationFile correctly", func(it *testing.T) {

		configurationFile, err := Init()

		if err != nil {
			it.Error("Initialization failed")
		}

		expectedLanguage := "go"

		if configurationFile.Application.Language != expectedLanguage {
			it.Errorf("Conversion failed. Expected=%q; Got=%q", expectedLanguage, configurationFile.Application.Language)
		}
	})

	t.Run("Should get the global configurationFile", func(it *testing.T) {
		_, err := Init()

		if err != nil {
			it.Error("Initialization failed")
		}

		configurationFile := Get()
		expectedLanguage := "go"

		if configurationFile.Application.Language != expectedLanguage {
			it.Errorf("Conversion failed. Expected=%q; Got=%q", expectedLanguage, configurationFile.Application.Language)
		}

	})

	if err := After("yaml"); err != nil {
		t.Error("Fail to delete anko.yaml")
	}

}
