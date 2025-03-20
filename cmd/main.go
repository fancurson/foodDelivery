package main

import (
	"context"
	"delivery/internal/config"
	test "delivery/pkg/api/test/api"
	"delivery/pkg/logger"
	"delivery/pkg/postgres"
	"delivery/pkg/service"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {

	ctx := context.Background()
	ctx, stop := signal.NotifyContext(ctx, os.Interrupt)
	defer stop()
	ctx, err := logger.New(ctx)
	if err != nil {
		log.Fatalf("failed to initialize logger: %v", err)
	}

	cfg, err := config.New()
	if err != nil {
		logger.GetLoggerFromCtx(ctx).Fatal(ctx, "failed to create configs", zap.Error(err))
	}

	db, err := postgres.NewDB(cfg.Postgres)
	if err != nil {
		logger.GetLoggerFromCtx(ctx).Fatal(ctx, "failed to connect to db", zap.Error(err))
	}
	fmt.Println(db.Ping(ctx))

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", cfg.GRPCPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	srv := service.New()
	server := grpc.NewServer(
		grpc.UnaryInterceptor(logger.Interceptor),
	)
	test.RegisterOrderServiceServer(server, srv)

	go func() {
		if err := server.Serve(lis); err != nil {
			logger.GetLoggerFromCtx(ctx).Info(ctx, "failed to serve", zap.Error(err))
		}
	}()

	select {
	case <-ctx.Done():
		server.Stop()

	}
}
