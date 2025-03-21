package logger

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"
	"google.golang.org/grpc"
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
	l, ok := ctx.Value(Key).(*Logger)
	if !ok || l == nil {
		return &Logger{zap.NewNop()} // Возвращаем "пустой" логгер, чтобы избежать паники
	}
	return l
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

func InterceptorWithLogger(logger *Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context,
		req any,
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (any, error) {

		guid := uuid.New().String()
		ctx = context.WithValue(ctx, RequestId, guid)

		logger.Info(ctx,
			"request",
			zap.String("method", info.FullMethod),
			zap.Time("request time", time.Now()),
		)

		return handler(ctx, req)
	}
}
