package willowlogger

import (
	"context"
	"fmt"

	"go.uber.org/zap"
)

type LOG_LEVEL string

const (
	DEBUG LOG_LEVEL = "debug"
	INFO  LOG_LEVEL = "info"

	UNKNOWN LOG_LEVEL = "unknown"
)

func StringToLogLevel(level string) LOG_LEVEL {
	switch level {
	case "debug":
		return DEBUG
	case "info":
		return INFO
	default:
		return UNKNOWN
	}
}

func NewZapLogger(logLevel LOG_LEVEL) (*zap.Logger, error) {
	zapCfg := zap.NewProductionConfig()
	zapCfg.OutputPaths = []string{"stdout"}
	zapCfg.DisableCaller = true
	zapCfg.DisableStacktrace = true
	zapCfg.Sampling = nil

	switch logLevel {
	case DEBUG:
		zapCfg.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	case INFO:
		zapCfg.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	default:
		return nil, fmt.Errorf("unknown log level received: %s", logLevel)
	}

	logger, err := zapCfg.Build()
	if err != nil {
		return nil, err
	}

	return logger, nil
}

func BaseLogger(logger *zap.Logger) *zap.Logger {
	return zap.New(logger.Core())
}

func StripedContext(logger *zap.Logger) context.Context {
	return context.WithValue(context.Background(), LoggerCtxKey, zap.New(logger.Core()))
}
