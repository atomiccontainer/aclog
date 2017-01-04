// Copyright © 2017 Christian R. Vozar ⚜
// Licensed under the MIT License. All rights reserved.

package aclog

import (
	"io/ioutil"
	"os"
	"regexp"
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
	dockerIDMatch := regexp.MustCompile(`cpu\:\/docker\/([0-9a-z]+)`)
	coreOSIDMatch := regexp.MustCompile(`cpuset\:\/system.slice\/docker-([0-9a-z]+)`)

	cgroup, err := ioutil.ReadFile("/proc/self/cgroup")
	if err != nil {
		return ""
	}

	strCgroup := string(cgroup)

	loc := dockerIDMatch.FindStringIndex(strCgroup)

	if loc != nil {
		return strCgroup[loc[0]+12 : loc[1]-2]
	}

	// cgroup not nil, not Docker. Check for CoreOS.
	loc = coreOSIDMatch.FindStringIndex(strCgroup)

	if loc != nil {
		return strCgroup[loc[0]+27:]
	}

	return ""
}
