package watcher

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/GuilhermeCaruso/anko/internal/banner"
	"github.com/fsnotify/fsnotify"
)

var watcher *fsnotify.Watcher

func (wc *WatcherConfig) WatchForChange() {
	closeDispatcher := false
	watcher, _ = fsnotify.NewWatcher()
	defer watcher.Close()

	if err := filepath.Walk(wc.RootPath, fileListener); err != nil {

		banner.Error(err.Error())
		wc.DoneChan <- true
	}

	normalizedExtensions := normalizeExtensionWithDot(wc.Extensions)

	contentToWatch := []string{}

	contentToWatch = append(contentToWatch, normalizedExtensions...)
	contentToWatch = append(contentToWatch, wc.Files...)

	for {
		select {
		case event := <-watcher.Events:
			if *wc.IsOpen {
				for _, content := range contentToWatch {
					if strings.Contains(event.Name, content) {
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

func normalizeExtensionWithDot(extensions []string) []string {
	normalizedExtensions := []string{}

	for _, extension := range extensions {
		normalizedExtensions = append(normalizedExtensions, fmt.Sprintf(".%s", extension))
	}
	return normalizedExtensions
}
