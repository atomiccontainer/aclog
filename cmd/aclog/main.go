// Copyright © 2017 Christian R. Vozar ⚜
// Licensed under the MIT License. All rights reserved.

package main

import (
	"fmt"

	"github.com/christianvozar/aclog"
)

func main() {
	i := aclog.New()

	fmt.Println(i.JSON())
}
