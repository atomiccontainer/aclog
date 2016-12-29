// Copyright © 2017 Christian R. Vozar ⚜
// Licensed under the MIT License. All rights reserved.

package aclog

import (
	"os"

	"github.com/uber-go/zap"
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
	zap.New(
		zap.NewJSONEncoder(zap.RFC3339Formatter("timestamp")),
		zap.Fields(
			zap.Int("pid", acinv.PID),
			zap.String("container_id", acinv.ID),
			zap.String("container_runtime", acinv.Runtime),
			zap.String("container_image_format", acinv.ImageFormat),
		),
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
