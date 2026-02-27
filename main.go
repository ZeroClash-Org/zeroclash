// SPDX-License-Identifier: MIT
// Copyright (c) 2026 scholar7r.

package main

import (
	"os"
	"os/signal"
	"syscall"

	"gitlhub.com/scholar7r/zeroclash/internal/cfg"
	"gitlhub.com/scholar7r/zeroclash/internal/logger"
)

func main() {
	FromArgs()

	cfg := cfg.Cfg(cfgPath)
	_ = cfg

	logger.Get().Info("zeroclash starting...")

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	<-c

	logger.Get().Info("zeroclash shutdown")
}
