# aclog

Container profile & introspection.

```Go
import "github.com/christianvozar/aclog"
```

aclog detects information about the running container and its runtime. The aim is to provide an inventory of the executing container and its runtime for debugging purposes.

## Usage

```Go
package main

import (
	"fmt"

	"github.com/christianvozar/aclog"
)

func main() {
	i := aclog.New()

	fmt.Println(i.JSON())
}
```
