package logger

import (
	"os"

	"github.com/scholar7r/sugar/singleton"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var global = singleton.New(load)

func Get() *zap.Logger {
	return global.Get()
}

func load() *zap.Logger {
	encoderCfg := zapcore.EncoderConfig{
		MessageKey:    "message",
		LevelKey:      "level",
		TimeKey:       "time",
		NameKey:       "name",
		CallerKey:     "caller",
		StacktraceKey: "stacktrace",
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeLevel:   zapcore.CapitalColorLevelEncoder,
		EncodeTime:    zapcore.ISO8601TimeEncoder,
		EncodeCaller:  zapcore.ShortCallerEncoder,
	}

	encoder := zapcore.NewConsoleEncoder(encoderCfg)

	return zap.New(
		zapcore.NewCore(
			encoder,
			zapcore.AddSync(os.Stdout),
			zapcore.DebugLevel,
		),
		zap.AddCaller(),
		zap.AddStacktrace(zapcore.ErrorLevel),
	)
}
