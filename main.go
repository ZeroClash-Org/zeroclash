// SPDX-License-Identifier: MIT
// Copyright (c) 2026 scholar7r.

package main

import (
	"context"
	"os/signal"
	"syscall"

	"github.com/zeroclash-org/zeroclash/internal/cfg"
	"github.com/zeroclash-org/zeroclash/internal/control"
	"github.com/zeroclash-org/zeroclash/internal/logger"
	"go.uber.org/zap"
)

func main() {
	// FromArgs initialize global variable cfgPath, externalControl via program arguments
	FromArgs()

	cfg.FromFile(cfgPath) // Load configuration from the specified configuration file

	if externalController != "" {
		cfg.Get().ExternalController = externalController
	}

	logger.Get().Info("zeroclash starting...")

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	if err := control.Run(ctx); err != nil {
		logger.Get().Fatal("failed to run control service", zap.Error(err))
	}

	<-ctx.Done()
	logger.Get().Info("zeroclash shutdown")
}
