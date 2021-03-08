package banner

import "fmt"

func Intro() {
	PrintWithColor("?Starting Anko!", BCyan)
}

func SettingUp() {
	PrintWithColor("?Configuring watcher.", BCyan)
}

func Listening() {
	PrintWithColor("?Listening for changes...", BCyan)
}

func Restarting() {
	PrintWithColor("?Restarting...", BCyan)
}

func Error(message string) {
	PrintWithColor(fmt.Sprintf("?[ERROR] ?%s", message), BRed, BCyan)
}
