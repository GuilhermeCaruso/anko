package watcher

import (
	"errors"
	"fmt"
	"os/exec"
)

// Language is a structure responsible for the control language information such as Name,
// CommandPath, process expression and other details used for the control processes.
type Language struct {
	ExecName      string
	ExecCmd       string
	ProcessRegexp string
	ExecPath      string
}

var supportedLanguages = map[string]*Language{
	"go": {
		ExecName:      "golang",
		ExecCmd:       "run",
		ProcessRegexp: `(\d+).*go-build.*/%s`,
	},
	"node": {
		ExecName:      "node",
		ProcessRegexp: `(\d+).* %s`,
	},
}

// GetLanguage is a simple selector using supportedLanguages
func GetLanguage(language string) (*Language, error) {
	var err error
	var execPath string

	selectedLanguage := supportedLanguages[language]

	if selectedLanguage == nil {
		return nil, errors.New("Language not implemented")
	}

	execPath, err = exec.LookPath(language)
	if err != nil {
		err = errors.New(buildMsgError(selectedLanguage.ExecName))
	}

	selectedLanguage.ExecPath = execPath

	return selectedLanguage, err
}

func buildMsgError(name string) string {
	return fmt.Sprintf("Fail to obtain %s path. Please, verify if you have %s in your PATH", name, name)
}
