package banner

import (
	"fmt"
	"strings"

	"github.com/GuilhermeCaruso/anko/internal/configuration"
)

func PrintWithColor(base string, args ...string) {
	if !configuration.Get().Application.ShowBanner {
		return
	}

	finalString := base
	for _, arg := range args {
		finalString = strings.Replace(finalString, "?", arg, 1)
	}
	fmt.Printf("> %s %s\n", finalString, Reset)
}
