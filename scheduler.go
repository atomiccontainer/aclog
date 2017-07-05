// Copyright © 2017 Christian R. Vozar ⚜
// Licensed under the MIT License. All rights reserved.

package acprof

const (
	schedulerKubernetes   = "kubernetes"
	schedulerNomad        = "nomad"
	schedulerUndetermined = "undetermined"
)

// getScheduler returns the scheduler if one is detected.
func getScheduler() string {
	if _, ok := EnvironmentVariables["KUBERNETES_SERVICE_HOST"]; ok {
		return schedulerKubernetes
	}

	return schedulerUndetermined
}
