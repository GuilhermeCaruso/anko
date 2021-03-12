package banner

import (
	"fmt"
	"strings"
)

// PrintWithColor is responsible for fomatting the message
func PrintWithColor(base string, args ...string) {
	finalString := base
	for _, arg := range args {
		finalString = strings.Replace(finalString, "?", arg, 1)
	}
	fmt.Printf("> %s %s\n", finalString, reset)
}
