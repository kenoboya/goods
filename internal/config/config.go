package config

import (
	"goods/pkg/database/psql"
	logger "goods/pkg/logger/zap"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/spf13/viper"
	"github.com/subosito/gotenv"
	"go.uber.org/zap"
)

type Config struct {
	HTTP HttpConfig
	GRPC GrpcConfig
	PSQL psql.PSQlConfig
}

type HttpConfig struct {
	Addr           string        `mapstructure:"port"`
	ReadTimeout    time.Duration `mapstructure:"readTimeout"`
	WriteTimeout   time.Duration `mapstructure:"writeTimeout"`
	MaxHeaderBytes int           `mapstructure:"maxHeaderBytes"`
}

type GrpcConfig struct {
}

func Init(configDIR string, envDIR string) (*Config, error) {
	if err := loadViperConfig(configDIR); err != nil {
		return &Config{}, err
	}

	var cfg Config
	if err := unmarshal(&cfg); err != nil {
		return &Config{}, err
	}

	if err := loadFromEnv(&cfg, envDIR); err != nil {
		return &Config{}, err
	}

	return &cfg, nil
}

func unmarshal(config *Config) error {
	if err := viper.UnmarshalKey("http", &config.HTTP); err != nil {
		logger.Error("Failed to unmarshal config file",
			zap.String("prefix", "http"),
			zap.Error(err),
		)
		return err
	}

	if err := viper.UnmarshalKey("grpc", &config.GRPC); err != nil {
		logger.Error("Failed to unmarshal config file",
			zap.String("prefix", "grpc"),
			zap.Error(err),
		)
		return err
	}

	return nil
}

func loadFromEnv(cfg *Config, envDIR string) error {
	if err := gotenv.Load(envDIR); err != nil {
		logger.Error("Failed to load environment file",
			zap.String("file", ".env"),
			zap.Error(err),
		)
		return err
	}

	if err := envconfig.Process("DB", &cfg.PSQL); err != nil {
		logger.Error("Failed to unmarshal environment file",
			zap.String("prefix", "DB"),
			zap.String("file", ".env"),
			zap.Error(err),
		)
		return err
	}

	return nil
}

func loadViperConfig(path string) error {
	viper.SetConfigFile("server")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path)

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			logger.Error("Failed to find config file",
				zap.String("file", "server.yaml"),
				zap.String("path", path),
				zap.Error(err),
			)
			return err
		} else {
			return err
		}
	}
	return nil
}
