// Copyright © 2017 Christian R. Vozar ⚜
// Licensed under the MIT License. All rights reserved.

package aclog

import (
	"io/ioutil"
	"os"
	"strings"
)

const (
	formatDocker       = "docker"
	formatACI          = "aci"
	formatOCF          = "ocf"
	formatUndetermined = "undetermined"
)

func getImageFormat() string {
	if _, err := os.Stat("/.dockerinit"); err == nil {
		return formatDocker
	}

	if _, err := os.Stat("/.dockerenv"); err == nil {
		return formatDocker
	}

	cgroup, _ := ioutil.ReadFile("/proc/self/cgroup")
	if strings.Contains(string(cgroup), "docker") {
		return formatDocker
	}

	if ac := os.Getenv("AC_METADATA_URL"); ac != "" {
		return formatACI
	}

	if ac := os.Getenv("AC_APP_NAME"); ac != "" {
		return formatACI
	}

	return formatUndetermined
}
