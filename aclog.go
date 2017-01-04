// Copyright © 2017 Christian R. Vozar ⚜
// Licensed under the MIT License. All rights reserved.

package aclog

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
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

	acLogger, _ := zap.Config{
		Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
		Development: false,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding: "json",
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "timestamp",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			MessageKey:     "message",
			StacktraceKey:  "stacktrace",
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}.Build(zap.Fields(
		zap.String("aclog_version", Version),
	))

	acLogger.Info("appc_inventory",
		zap.Int("pid", acinv.PID),
		zap.String("container_id", acinv.ID),
		zap.String("container_runtime", acinv.Runtime),
		zap.String("container_image_format", acinv.ImageFormat),
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
