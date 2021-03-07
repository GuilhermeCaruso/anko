package watcher

import (
	"errors"
	"fmt"
	"os/exec"
)

const (
	GO   = "go"
	MAKE = "make"
	SH   = "sh"
)

var supportedLanguage = map[string]string{
	GO:   "golang",
	MAKE: "makefile",
	SH:   "shellscript",
}

func GetExecPath(language string) (string, error) {
	var err error
	var execPath string

	switch language {
	case GO:
		execPath, err = exec.LookPath(GO)
		if err != nil {
			err = errors.New(buildMsgError(supportedLanguage[GO]))
		}
	case MAKE:
		execPath, err = exec.LookPath(MAKE)
		if err != nil {
			err = errors.New(buildMsgError(supportedLanguage[MAKE]))
		}
	case SH:
		execPath, err = exec.LookPath(SH)
		if err != nil {
			err = errors.New(buildMsgError(supportedLanguage[SH]))
		}
	default:
		err = errors.New("Language not implemented")
	}
	return execPath, err
}

func buildMsgError(name string) string {
	return fmt.Sprintf("Fail to obtain %s path. Please, verify if you have %s in your PATH", name, name)
}
