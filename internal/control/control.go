// Package control provides the external control ability
package control

import (
	"context"

	"github.com/gofiber/fiber/v3"
	"github.com/zeroclash-org/zeroclash/internal/cfg"
	"github.com/zeroclash-org/zeroclash/internal/logger"
	"go.uber.org/zap"
)

const (
	routeLogs    = "/logs"
	routeCfg     = "/cfg"
	routeProxies = "/proxies"
	routeRules   = "/rules"
	routeTraffic = "/traffic"
)

func Run(ctx context.Context) error {
	fbr := fiber.New()

	fbr.Get(routeLogs, handleLogs)
	fbr.Get(routeCfg, handleCfg)
	fbr.Get(routeProxies, handleProxies)
	fbr.Get(routeRules, handleRules)
	fbr.Get(routeTraffic, handleTraffic)

	go func() {
		<-ctx.Done()
		_ = fbr.Shutdown()
	}()

	logger.Get().Info(
		"starting external controller service...",
		zap.String("listen", cfg.Get().ExternalController),
	)

	return fbr.Listen(
		cfg.Get().ExternalController,
		fiber.ListenConfig{DisableStartupMessage: true},
	)
}

func handleLogs(c fiber.Ctx) error {
	_ = c
	return nil
}

func handleCfg(c fiber.Ctx) error {
	return c.JSON(cfg.Get())
}

func handleProxies(c fiber.Ctx) error {
	return c.JSON(cfg.Get().Proxy)
}

func handleRules(c fiber.Ctx) error {
	return c.JSON(cfg.Get().Rule)
}

func handleTraffic(c fiber.Ctx) error {
	_ = c
	return nil
}
