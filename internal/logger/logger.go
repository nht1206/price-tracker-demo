package logger

import (
	"fmt"
	"path"

	"github.com/nht1206/pricetracker/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.SugaredLogger

func InitLogger(cfg *config.LogConfig) error {
	if cfg == nil {
		return fmt.Errorf("cfg is nil")
	}
	config := zap.NewProductionConfig()
	level, err := getZapLevel(cfg.Level)
	if err != nil {
		return err
	}
	config.Level.SetLevel(level)
	config.OutputPaths = []string{
		path.Join(cfg.OutputPath, cfg.FileName),
	}
	config.DisableStacktrace = true
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	l, err := config.Build()
	if err != nil {
		return err
	}
	Logger = l.Sugar()
	return nil
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
		return -2, fmt.Errorf("undefined log level")
	}
}
