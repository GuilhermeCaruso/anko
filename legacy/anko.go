package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"strings"

	"github.com/fsnotify/fsnotify"
)

const (
	LANG_COMMAND   = "go"
	FILE_EXTENSION = ".go"
	ACT_INIT       = "init"
	ACT_STOP       = "stop"
	ACT_RESET      = "reset"

	ROOT_PATH = "."
	EXEC_PATH = "example/fiber/test.go"
)

var (
	watcher    *fsnotify.Watcher
	isOpen     bool
	done       = make(chan bool)
	dispatcher = make(chan string)
	goExecPath string
	sysOS      = runtime.GOOS
)

func main() {
	fmt.Println("Starting Anko...")
	initialSetup()
	go fileWatcher()
	go appController()
	<-done
}

func initialSetup() {
	resultGoExecPath, err := exec.LookPath(LANG_COMMAND)
	if err != nil {
		log.Fatal("Fail to obtain golang path. Please, verify if you have golang in your PATH")
	}
	goExecPath = resultGoExecPath
}

func appController() {
	for {
		select {
		case action := <-dispatcher:
			switch action {
			case ACT_INIT:
				go initApp()
				isOpen = true
			case ACT_RESET:
				resetApp()
				isOpen = true
			}
		}
	}
}

func initApp() {

	cmd := exec.Command(goExecPath, "run", EXEC_PATH)

	stdout, err := cmd.StdoutPipe()

	cmd.Stderr = cmd.Stdout

	if err != nil {
		done <- true
	}

	if err = cmd.Start(); err != nil {
		done <- true
	}

	for {
		tmp := make([]byte, 1024)
		_, err := stdout.Read(tmp)

		if !strings.Contains(string(tmp), "signal: killed") {
			fmt.Print(string(tmp))
		}

		if err != nil {
			stdout.Close()
			fmt.Println("Restarting...")
			break

		}
	}
}

func resetApp() {
	re := regexp.MustCompile(`(?m)\/([\w_]*).go`)

	match := re.FindStringSubmatch(EXEC_PATH)

	var appName string
	if len(match) > 1 {
		appName = match[1]
	}

	var command string
	if sysOS == "linux" {
		command = "ps -u"
	} else if sysOS == "mac" {
		command = "ps -A"
	}
	b, _ := exec.Command("/bin/sh", "-c", command).Output()
	r := regexp.MustCompile(fmt.Sprintf(`(\d+).*go-build.*/%s`, appName))
	match = r.FindStringSubmatch(string(b))
	if len(match) > 1 {
		i, err := strconv.Atoi(match[1])
		if err != nil {
			done <- true
		}
		p, err := os.FindProcess(i)
		if err != nil {
			done <- true
		}
		p.Kill()
		go initApp()
	}

}

func fileWatcher() {
	fmt.Println("Listening for changes...")

	watcher, _ = fsnotify.NewWatcher()
	defer watcher.Close()

	if err := filepath.Walk(ROOT_PATH, fileListener); err != nil {
		fmt.Println("Root doesn't exists")
		done <- true
	}

	dispatcher <- ACT_INIT

	for {
		select {
		case event := <-watcher.Events:
			if strings.Contains(event.Name, FILE_EXTENSION) && isOpen {
				isOpen = false
				dispatcher <- ACT_RESET
			}
		case err := <-watcher.Errors:
			fmt.Printf("Error. Err: %s\n", err.Error())
			done <- true
		}
	}
}

func fileListener(path string, fi os.FileInfo, err error) error {
	if fi.Mode().IsDir() {
		return watcher.Add(path)
	}
	return nil
}
