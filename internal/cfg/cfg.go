package cfg

import (
	"os"

	"github.com/scholar7r/sugar/singleton"
	"gitlhub.com/scholar7r/zeroclash/internal/logger"
	"go.uber.org/zap"
	"gopkg.in/yaml.v2"
)

var cfgPath string

type ProxyCfg struct {
	Name     string `yaml:"name"     json:"name"`
	Server   string `yaml:"server"   json:"server"`
	Port     uint16 `yaml:"port"     json:"port"`
	Password string `yaml:"password" json:"-"`
}

type RuleCfg struct{}

type BaseCfg struct {
	ControlAddr string     `yaml:"controlAddr" json:"controlAddr"`
	Proxy       []ProxyCfg `yaml:"proxy"       json:"proxy"`
	Rule        []RuleCfg  `yaml:"rule"        json:"rule"`
}

var global = singleton.New(load)

func Cfg(path string) *BaseCfg {
	cfgPath = path
	return global.Get()
}

func Get() *BaseCfg {
	return global.Get()
}

func load() *BaseCfg {
	var (
		b   []byte
		err error
		cfg *BaseCfg
	)

	if cfgPath == "" {
		logger.Get().Fatal("configuration file unspecified")
	}

	logger.Get().Debug("loading configuration from " + cfgPath)

	if _, err = os.Stat(cfgPath); os.IsNotExist(err) {
		logger.Get().Fatal("configuration file does not exists", zap.String("path", cfgPath))
	}

	if b, err = os.ReadFile(cfgPath); err != nil {
		logger.Get().Fatal("failed to read configuration", zap.Any("error", err))
	}

	err = yaml.Unmarshal(b, &cfg)
	if err != nil {
		logger.Get().Fatal("failed to unmarshal configuration", zap.Any("error", err))
	}

	logger.Get().Debug("loaded configuration successfully")
	logger.Get().Debug("counts of configuration", zap.Int("proxy", len(cfg.Proxy)), zap.Int("rule", len(cfg.Rule)))

	return cfg
}
