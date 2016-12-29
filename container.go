// Copyright © 2017 Christian R. Vozar ⚜
// Licensed under the MIT License. All rights reserved.

package aclog

import (
	"io/ioutil"
	"os"
	"strings"
)

// IsContainer returns true if the application is running within a container
// runtime/engine.
func IsContainer() bool {
	if _, err := os.Stat("/.dockerinit"); err == nil {
		return true
	}

	if _, err := os.Stat("/.dockerenv"); err == nil {
		return true
	}

	if c := getContainerID(); c != "" {
		return true
	}

	return false
}

func getContainerID() string {
	cgroup, err := ioutil.ReadFile("/proc/self/cgroup")
	if err != nil {
		return ""
	}

	strCgroup := string(cgroup)

	b := strings.Index(strCgroup, "cpu:/docker/")
	e := strings.LastIndex(strCgroup, "1:cpuset")

	if b != -1 {
		return strCgroup[b+12 : e-2]
	}

	// Not empty and not native Docker, attempt CoreOS
	if b = strings.Index(strCgroup, "cpuset:/system.slice/docker-"); b != -1 {
		return strCgroup[b:]
	}

	return ""
}
