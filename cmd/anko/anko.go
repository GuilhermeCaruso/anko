package main

import (
	"runtime"

	"github.com/GuilhermeCaruso/anko/internal/banner"
	"github.com/GuilhermeCaruso/anko/internal/configuration"
	"github.com/GuilhermeCaruso/anko/internal/watcher"
)

var (
	done       = make(chan bool)
	dispatcher = make(chan string)
	isOpen     = true
)

func main() {

	config, err := configuration.Init()

	if err != nil {
		banner.Error(err.Error())
		return
	}

	banner.Intro()
	banner.SettingUp()

	w := watcher.New(watcher.WatcherConfig{
		Files:          config.Application.Watch.Files,
		Extensions:     config.Application.Watch.Extensions,
		RootPath:       config.Application.RootPath,
		DispatcherChan: dispatcher,
		DoneChan:       done,
		IsOpen:         &isOpen,
		AppPath:        config.Application.ExecPath,
		Language:       config.Application.Language,
		SysOS:          runtime.GOOS,
	})

	banner.Listening()

	go w.WatchForChange()
	go w.AppController()

	<-done
}
