// Copyright © 2017 Christian R. Vozar ⚜
// Licensed under the MIT License. All rights reserved.

package main

import (
	"fmt"

	"github.com/atomiccontainer/acprof"
)

func main() {
	i := acprof.New()

	fmt.Println(i.JSON())
}
