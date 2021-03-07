package main

import (
	"fmt"

	"github.com/GuilhermeCaruso/anko/internal/configuration"
)

func main() {
	config := configuration.Init()

	fmt.Println(config)
}
