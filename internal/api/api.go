package api

import (
	"goods/internal/config"
	"goods/pkg/database/psql"
	logger "goods/pkg/logger/zap"

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
}
