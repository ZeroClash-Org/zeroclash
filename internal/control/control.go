// Package control provides the control ability
package control

import (
	"context"

	"github.com/gofiber/fiber/v3"
	"gitlhub.com/scholar7r/zeroclash/internal/cfg"
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

	return fbr.Listen(
		cfg.Get().ControlAddr,
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
