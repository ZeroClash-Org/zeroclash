// SPDX-License-Identifier: MIT
// Copyright (c) 2026 scholar7r.

package main

import (
	"context"
	"os/signal"
	"syscall"

	"gitlhub.com/scholar7r/zeroclash/internal/cfg"
	"gitlhub.com/scholar7r/zeroclash/internal/control"
	"gitlhub.com/scholar7r/zeroclash/internal/logger"
	"go.uber.org/zap"
)

func main() {
	FromArgs()

	cfg := cfg.Cfg(cfgPath)
	_ = cfg

	logger.Get().Info("zeroclash starting...")

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	if err := control.Run(ctx); err != nil {
		logger.Get().Fatal("failed to run control service", zap.Error(err))
	}

	<-ctx.Done()
	logger.Get().Info("zeroclash shutdown")
}
