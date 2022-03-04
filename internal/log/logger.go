package log

import (
	"errors"
	"path"

	"github.com/nht1206/pricetracker/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func InitLogger(cfg *config.LogConfig) (*zap.SugaredLogger, error) {
	if cfg == nil {
		return nil, errors.New("cfg is nil")
	}
	config := zap.NewProductionConfig()
	level, err := getZapLevel(cfg.Level)
	if err != nil {
		return nil, err
	}
	config.Level.SetLevel(level)
	if cfg.OutputPath != "" {
		config.OutputPaths = []string{
			path.Join(cfg.OutputPath, cfg.FileName),
		}
	}
	config.DisableStacktrace = true
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	l, err := config.Build()
	if err != nil {
		return nil, err
	}
	return l.Sugar(), nil
}

func getZapLevel(level string) (zapcore.Level, error) {
	switch level {
	case "debug":
		return zap.DebugLevel, nil
	case "info":
		return zap.InfoLevel, nil
	case "warn":
		return zap.WarnLevel, nil
	case "error":
		return zap.ErrorLevel, nil
	case "dpanic":
		return zap.DPanicLevel, nil
	case "panic":
		return zap.PanicLevel, nil
	case "fatal":
		return zap.FatalLevel, nil
	default:
		return -2, errors.New("undefined log level")
	}
}
