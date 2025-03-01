package logger

import (
	"context"
	"fmt"

	"go.uber.org/zap"
)

const (
	Key = "logger"

	RequestId = "request_id"
)

type Logger struct {
	l *zap.Logger
}

func New(ctx context.Context) (context.Context, error) {
	logger, err := zap.NewProduction()
	if err != nil {
		return nil, fmt.Errorf("register logger error: %w", err)
	}

	ctx = context.WithValue(ctx, Key, &Logger{logger})
	return ctx, err
}

func GetLoggerFromCtx(ctx context.Context) *Logger {
	return ctx.Value(Key).(*Logger)
}

func (l *Logger) Info(ctx context.Context, msg string, fields ...zap.Field) {
	if ctx.Value(RequestId) != nil {
		fields = append(fields, zap.String(RequestId, ctx.Value(RequestId).(string)))
	}

	l.l.Info(msg, fields...)
}

func (l *Logger) Fatal(ctx context.Context, msg string, fields ...zap.Field) {
	if ctx.Value(RequestId) != nil {
		fields = append(fields, zap.String(RequestId, ctx.Value(RequestId).(string)))
	}

	l.l.Fatal(msg, fields...)
}
