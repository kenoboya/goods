package api

import (
	"context"
	"goods/internal/config"
	repo "goods/internal/repository"
	grpc_server "goods/internal/server/grpc"
	http_server "goods/internal/server/http"
	"goods/internal/service"
	grpc_client "goods/internal/transport/grpc/client"
	"goods/internal/transport/rest"
	"goods/pkg/database/psql"
	logger "goods/pkg/logger/zap"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"
)

func Run(configDIR string, envDIR string) {
	logger.InitLogger()
	cfg, err := config.Init(configDIR, envDIR)
	if err != nil {
		logger.Fatal("Failed to initialize config",
			zap.Error(err),
			zap.String("context", "Initializing application"),
			zap.String("version", "1.0.0"),
			zap.String("environment", "development"),
		)
	}
	db, err := psql.NewPostgresConnection(cfg.PSQL)
	if err != nil {
		logger.Fatal("Failed to connect to database",
			zap.Error(err),
			zap.String("context", "Initializing application"),
			zap.String("version", "1.0.0"),
			zap.String("environment", "development"),
		)
	}

	authClient, err := grpc_client.NewAuthClient(cfg.GRPC)
	if err != nil {
		logger.Fatal("Failed to create auth client",
			zap.Error(err),
			zap.String("context", "Initializing application"),
			zap.String("version", "1.0.0"),
			zap.String("environment", "development"),
		)
	}

	paymentClient, err := grpc_client.NewPaymentClient(cfg.GRPC)
	if err != nil {
		logger.Fatal("Failed to create payment client",
			zap.Error(err),
			zap.String("context", "Initializing application"),
			zap.String("version", "1.0.0"),
			zap.String("environment", "development"),
		)
	}

	repositories := repo.NewRepositories(db)
	services := service.NewServices(repositories)
	handler := rest.NewHandler(services, authClient, paymentClient)

	// HTTP
	httpServer := http_server.NewServer(cfg.HTTP, handler)

	go func() {
		if err := httpServer.Run(); err != nil {
			logger.Fatalf("The http server didn't start: %s\n", err)
		}
	}()

	logger.Info("Http server started")

	// GRPC
	grpcServer := grpc_server.NewServer(cfg.GRPC)

	go func() {
		if err := grpcServer.Run(); err != nil {
			logger.Fatalf("The grpc server didn't start: %s\n", err)
		}
	}()

	// EXIT
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	const timeout = 5 * time.Second

	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	if err := httpServer.Shutdown(ctx); err != nil {
		logger.Errorf("failed to stop http server: %v", err)
	}

	grpcServer.Stop()

	if err := db.Close(); err != nil {
		logger.Errorf("failed to stop postgres database: %v", err)
	}
}
