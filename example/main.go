package main

import (
	"fmt"
	"log"

	"github.com/gothew/config"
)

type ConfigExample struct {
  Command string `yaml:"command"`
}

func fnConfigExample() config.ConfigOptions {
	m := make(map[string]interface{})
	m["ftp"] = ConfigExample{Command: "ls"}

	return config.ConfigOptions{
		Services: m,
	}
}

func main() {
	config.SetAppDir("linuxtest")
	config.SetConfigOptions(fnConfigExample)
	cfg, err := config.ParseConfig()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("s1: %s", cfg.Services["ftp"])
}
