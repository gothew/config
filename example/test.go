package main

import (
	"fmt"
	"log"

	"github.com/gothew/config"
)

func main() {
	config, err := config.New().ParserConfig()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v", config)
}
