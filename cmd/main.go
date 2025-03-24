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
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
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
		grpc.UnaryInterceptor(logger.InterceptorWithLogger(logger.GetLoggerFromCtx(ctx))),
	)
	test.RegisterOrderServiceServer(server, srv)

	rt := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err = test.RegisterOrderServiceHandlerFromEndpoint(ctx, rt, "localhost:"+strconv.Itoa(cfg.GRPCPort), opts)
	if err != nil {
		logger.GetLoggerFromCtx(ctx).Fatal(ctx, "failed to register handler server", zap.Error(err))
	}

	go func() {
		if err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.RestPort), rt); err != nil {
			logger.GetLoggerFromCtx(ctx).Fatal(ctx, "failed to server", zap.Error(err))
		}
	}()

	go func() {
		if err := server.Serve(lis); err != nil {
			logger.GetLoggerFromCtx(ctx).Info(ctx, "failed to serve", zap.Error(err))
		}
	}()

	<-ctx.Done()
	logger.GetLoggerFromCtx(ctx).Info(ctx, "Shutting down server...")

	server.GracefulStop()
	logger.GetLoggerFromCtx(ctx).Info(ctx, "Server Stopped")

}
