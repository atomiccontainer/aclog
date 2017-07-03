# aclog

App container introspection.

```Go
import "github.com/christianvozar/aclog"
```

aclog is intended for Go applications designed to run specifically within container runtimes. The aim is to provide an inventory of the executing container and its runtime for debugging purposes.

## Usage

```Go
package main

import (
	"log"

	"github.com/christianvozar/aclog"
)

func main() {
	i, err := aclog.New()
	if err != nil {
		log.Fatal(err.Error())
	}

	format := "Container Inventory:\nID: %s\nRuntime: %s\nImage Format: %s\n PID: %s\n"
	_, err := fmt.Printf(format, i.ID, i.Runtime, i.ImageFormat, i.PID)
	fmt.Println("")
}
```
