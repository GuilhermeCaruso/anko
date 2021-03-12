package watcher

import (
	"errors"
	"fmt"
	"os/exec"
)

type Language struct {
	ExecName      string
	ExecCmd       string
	ProcessRegexp string
	ExecPath      string
}

var support = map[string]*Language{
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

func GetLanguage(language string) (*Language, error) {
	var err error
	var execPath string

	selectedLanguage := support[language]

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
