package tunnel

import (
	"strings"

	"gitlhub.com/scholar7r/zeroclash/internal/cfg"
)

type Tunnel struct {
	proxies map[string]any
	rules   []any
}

func (x *Tunnel) Update() error {
	x.proxies = x.updateProxy()
	x.rules = x.updateRule()

	return nil
}

func (x *Tunnel) updateProxy() map[string]any {
	proxies := make(map[string]any)

	for _, v := range cfg.Get().Proxy {
		switch v.Type {
		// TODO: Create adapter instance for each proxy protocol
		// Maybe we can do some design here for protocols will be added future?
		case "shadowsocks":
			proxies[v.Name] = nil
		case "hysteria":
			proxies[v.Name] = nil
		case "hysteria2":
			proxies[v.Name] = nil
		}
	}

	return proxies
}

func (x *Tunnel) updateRule() []any {
	rules := []any{}

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
		}
	}

	return rules
}

func trimAround(v []string) []string {
	var trimed []string

	for _, w := range v {
		trimed = append(trimed, strings.Trim(w, " "))
	}

	return trimed
}
