# Config
Configuration based on yaml files

## Getting Started
The best way to install the config is through `go get`.

```sh
go get github.com/gothew/config
```

## Documentation
working...

## Quickstart
For more examples of using ngrok-go, check out the [/examples](/examples) folder.

The following example uses config with a custom configuration interface to the yaml file and as an end result shows the configuration read through the `Services` property.

```go
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
```

See [LICENSE](./LICENSE) for details.
