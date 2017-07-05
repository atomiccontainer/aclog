package aclog

import (
	"os"
	"strings"
)

// environMap returns the results of os.Environ as a map.
func environMap() map[string]string {
	env := os.Environ()
	vars := make(map[string]string)
	for _, pair := range env {
		e := strings.Split(pair, "=")
		vars[e[0]] = e[1]
	}
	return vars
}
