package log

import (
	"context"

	"go.uber.org/zap"
)

var contextKey int

func WithLogger(ctx context.Context, logger *zap.SugaredLogger) context.Context {
	return context.WithValue(ctx, contextKey, *logger)
}

func FromContext(ctx context.Context) *zap.SugaredLogger {
	if logger, ok := ctx.Value(contextKey).(*zap.SugaredLogger); ok {
		return logger
	}
	return nil
}
