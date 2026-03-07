package tunnel_test

import (
	"testing"

	"github.com/zeroclash-org/zeroclash/adaptor"
	"github.com/zeroclash-org/zeroclash/rule"
	"github.com/zeroclash-org/zeroclash/tunnel"
)

var _ tunnel.Tunnel = (*mockTunnel)(nil)

type mockTunnel struct {
	proxies map[string]adaptor.Adaptor
	rules   []rule.Rule
}

func newMockTunnel() tunnel.Tunnel {
	return &mockTunnel{
		proxies: make(map[string]adaptor.Adaptor),
		rules:   []rule.Rule{},
	}
}

func (x *mockTunnel) Update() error {
	return nil
}

func Test_tunnel_Update(t *testing.T) {
	t.Run("update tunnel", func(t *testing.T) {
		v := newMockTunnel()
		if err := v.Update(); err != nil {
			t.Fatalf("expected nil, got error")
		}
	})
}
