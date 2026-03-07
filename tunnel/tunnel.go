// Package tunnel provides tunnel ability to forward data package to upstream
// proxy server, like shadowsocks, hysteria and others.
package tunnel

import (
	"strings"

	"github.com/zeroclash-org/zeroclash/adaptor"
	"github.com/zeroclash-org/zeroclash/internal/cfg"
	"github.com/zeroclash-org/zeroclash/internal/logger"
	"github.com/zeroclash-org/zeroclash/rule"
	"go.uber.org/zap"
)

type Tunnel struct {
	proxies map[string]adaptor.Adaptor
	rules   []rule.Rule
}

func (x *Tunnel) Update() error {
	x.proxies = x.updateProxy()
	x.rules = x.updateRule()

	return nil
}

func (x *Tunnel) updateProxy() map[string]adaptor.Adaptor {
	proxies := make(map[string]adaptor.Adaptor)

	for _, v := range cfg.Get().Proxy {
		switch strings.ToLower(v.Type) {
		case "shadowsocks":
			proxies[v.Name] = nil
		case "hysteria":
			proxies[v.Name] = nil
		case "hysteria2":
			proxies[v.Name] = nil
		default:
			logger.Get().Warn("unknown protocol", zap.String("protocol", proxies[v.Name].Name()))
		}
	}

	return proxies
}

func (x *Tunnel) updateRule() []rule.Rule {
	rules := []rule.Rule{}

	for _, v := range cfg.Get().Rule {
		parts := trimAround(strings.Split(v, ","))
		if len(parts) == 0 {
			continue
		}

		switch parts[0] {
		case "DOMAIN-SUFFIX":
		case "DOMAIN-KEYWORD":
		case "GEOIP":
		case "IP-CIDR", "IP-CIDR6":
		case "FINAL":
		default:
			logger.Get().Warn("unknown rule type", zap.String("ruleType", parts[0]))
		}
	}

	return rules
}

func trimAround(v []string) []string {
	var trimmed []string

	for _, w := range v {
		trimmed = append(trimmed, strings.Trim(w, " "))
	}

	return trimmed
}
