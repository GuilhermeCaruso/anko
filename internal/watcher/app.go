package watcher

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"

	"github.com/GuilhermeCaruso/anko/internal/banner"
)

func (wc *WatcherConfig) InitApp() {
	var cmd *exec.Cmd

	if wc.Language == GO {
		execPath, err := GetExecPath(wc.Language)

		if err != nil {
			banner.Error(err.Error())
		}

		cmd = exec.Command(execPath, "run", wc.AppPath)
	} else if wc.Language == NODE {
		execPath, err := GetExecPath(wc.Language)

		if err != nil {
			banner.Error(err.Error())
		}

		cmd = exec.Command(execPath, wc.AppPath)
	}

	stdout, err := cmd.StdoutPipe()

	cmd.Stderr = cmd.Stdout

	if err != nil {
		banner.Error(err.Error())
		wc.DoneChan <- true
	}

	if err = cmd.Start(); err != nil {

		banner.Error(err.Error())
		wc.DoneChan <- true
	}

	for {
		tmp := make([]byte, 1024)
		_, err := stdout.Read(tmp)

		if !strings.Contains(string(tmp), "signal: killed") {
			fmt.Print(string(tmp))
		}

		if err != nil {
			stdout.Close()
			banner.Restarting()
			break

		}
	}
}

func (wc *WatcherConfig) resetApp() {
	re := regexp.MustCompile(`(?m)\/([\w_]*).go`)

	match := re.FindStringSubmatch(wc.AppPath)

	var appName string
	if len(match) > 1 {
		appName = match[1]
	}

	var command string
	if wc.SysOS == "linux" {
		command = "ps -u"
	} else if wc.SysOS == "mac" {
		command = "ps -A"
	}
	b, _ := exec.Command("/bin/sh", "-c", command).Output()
	r := regexp.MustCompile(fmt.Sprintf(`(\d+).*go-build.*/%s`, appName))
	match = r.FindStringSubmatch(string(b))
	if len(match) > 1 {
		i, err := strconv.Atoi(match[1])
		if err != nil {

			banner.Error(err.Error())
			wc.DoneChan <- true
		}
		p, err := os.FindProcess(i)
		if err != nil {

			banner.Error(err.Error())
			wc.DoneChan <- true
		}
		p.Kill()
		go wc.InitApp()
	}
}

func (wc *WatcherConfig) AppController() {
	openDispacher := true
	go wc.InitApp()
	for {
		select {
		case action := <-wc.DispatcherChan:
			switch action {
			case ACT_INIT:
				go wc.InitApp()
				wc.IsOpen = &openDispacher
			case ACT_RESET:
				wc.resetApp()
				wc.IsOpen = &openDispacher
			}
		}
	}
}
