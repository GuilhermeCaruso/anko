package watcher

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/GuilhermeCaruso/anko/internal/banner"
	"github.com/fsnotify/fsnotify"
)

type WatcherConfig struct {
	RootPath       string
	FileExtensions []string
	IsOpen         *bool
	Language       string
	AppPath        string
	DoneChan       chan bool
	SysOS          string
	DispatcherChan chan string
}

const (
	ACT_INIT  = "init"
	ACT_STOP  = "stop"
	ACT_RESET = "reset"
)

func New(args WatcherConfig) *WatcherConfig {
	return &args
}

var watcher *fsnotify.Watcher

func (wc *WatcherConfig) WatchForChange() {
	closeDispatcher := false
	watcher, _ = fsnotify.NewWatcher()
	defer watcher.Close()

	if err := filepath.Walk(wc.RootPath, fileListener); err != nil {
		wc.DoneChan <- true
	}

	for {
		select {
		case event := <-watcher.Events:
			if *wc.IsOpen {
				for _, extension := range wc.FileExtensions {
					if strings.Contains(event.Name, fmt.Sprintf(".%s", extension)) {
						wc.IsOpen = &closeDispatcher
						wc.DispatcherChan <- ACT_RESET
					}
				}
			}
		case err := <-watcher.Errors:
			banner.Error(err.Error())
			wc.DoneChan <- true
		}
	}
}

func fileListener(path string, fi os.FileInfo, err error) error {
	if fi.Mode().IsDir() {
		return watcher.Add(path)
	}
	return nil
}
