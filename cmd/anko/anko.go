package main

import "github.com/GuilhermeCaruso/anko/internal/banner"

type Anko struct {
	ExecPath string
}

func main() {
	banner.Intro()
	banner.SettingUp()
	banner.Listening()
	// config := configuration.Init()

	// execPath, err := watcher.GetExecPath(config.Application.Language)

	// if err != nil {
	// log.Fatal(err)
	// }

	// mainAnko := Anko{
	// 	ExecPath: execPath,
	// }
}
