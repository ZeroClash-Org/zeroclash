package cfg

import (
	"os"

	"github.com/scholar7r/sugar/singleton"
	"gitlhub.com/scholar7r/zeroclash/internal/logger"
	"go.uber.org/zap"
	"gopkg.in/yaml.v2"
)

var cfgPath = "./cfg.yaml"

type BaseCfg struct{}

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

	logger.Get().Debug("loading configuration from " + cfgPath)

	if _, err := os.Stat(cfgPath); os.IsNotExist(err) {
		logger.Get().Fatal("configuration file do not exists")
	}

	if b, err = os.ReadFile(cfgPath); err != nil {
		logger.Get().Fatal("failed to read configuration", zap.Any("error", err))
	}

	err = yaml.Unmarshal(b, &cfg)
	if err != nil {
		logger.Get().Fatal("failed to unmarshal configuration", zap.Any("error", err))
	}

	logger.Get().Debug("loaded configuration successfully")

	return cfg
}
