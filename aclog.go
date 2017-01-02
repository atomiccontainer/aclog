// Copyright © 2017 Christian R. Vozar ⚜
// Licensed under the MIT License. All rights reserved.

package aclog

import (
	"os"

	log "github.com/uber-go/zap"
)

// Inventory holds an application's container and runtime information.
type Inventory struct {
	ID          string
	Runtime     string
	ImageFormat string
	PID         int
}

func init() {
	acinv := NewInventory()

	acLogger := log.New(
		log.NewJSONEncoder(
			log.RFC3339Formatter("timestamp"),
			log.MessageKey("message"),
			log.LevelString("level"),
		),
		log.Fields(
			log.String("dissembler_version", Version),
		),
	)

	acLogger.Info("appc_inventory",
		log.Int("pid", acinv.PID),
		log.String("container_id", acinv.ID),
		log.String("container_runtime", acinv.Runtime),
		log.String("container_image_format", acinv.ImageFormat),
	)
}

// NewInventory returns a new Inventory with populated values.
func NewInventory() *Inventory {
	return &Inventory{
		ID:          getContainerID(),
		Runtime:     getRuntime(),
		ImageFormat: getImageFormat(),
		PID:         os.Getpid(),
	}
}
