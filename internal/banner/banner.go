package banner

import "fmt"

// Intro prints the initial message from the application
func Intro() {
	PrintWithColor("?Starting Anko!", bCyan)
}

// SettingUp prints the setting up message from the application
func SettingUp() {
	PrintWithColor("?Configuring watcher.", bCyan)
}

// Listening prints message after application is listening for changes
func Listening() {
	PrintWithColor("?Listening for changes...", bCyan)
}

// Restarting is called when application requires changes
func Restarting() {
	PrintWithColor("?Restarting...", bCyan)
}

// Error is called to format the error message in the application
func Error(message string) {
	PrintWithColor(fmt.Sprintf("?[ERROR] ?%s", message), bRed, bCyan)
}
