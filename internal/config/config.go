package config

import (
	"goods/pkg/database/psql"
	"log"

	"github.com/kelseyhightower/envconfig"
	"github.com/spf13/viper"
	"github.com/subosito/gotenv"
)

const configDIR = "api/configs"
const envDIR = "api/.env"

type Config struct {
	HTTP HttpConfig
	GRPC GrpcConfig
	PSQL psql.PSQlConfig
}

type HttpConfig struct {
}

type GrpcConfig struct {
}

func Init() (*Config, error){
	if err:= loadViperConfig(configDIR); err != nil{
		return &Config{}, err
	}

	var cfg Config
	if err:= unmarshal(&cfg); err != nil{
		return &Config{}, err
	}

	if err:= loadFromEnv(&cfg); err != nil{
		return &Config{}, err
	}
	
	return &cfg, nil
}

func unmarshal(config *Config) error{
	if err:= viper.UnmarshalKey("http", &config.HTTP); err != nil{
		return err
	}

	if err:= viper.UnmarshalKey("grpc", &config.GRPC); err != nil{
		return err
	}

	return nil
}

func loadFromEnv(cfg *Config) error{
	if err:= gotenv.Load(envDIR); err != nil{
		return err
	}
	
	if err:= envconfig.Process("DB", &cfg.PSQL); err != nil{
		return err
	}

	return nil
}

func loadViperConfig(path string) error {
	viper.SetConfigFile("server")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path)
	
	if err:= viper.ReadInConfig(); err!= nil{
		if _, ok:= err.(viper.ConfigFileNotFoundError); ok{
			log.Fatalf("Viper didn't find config file: server.yaml: %v", err)
			return err
		}else{
			return err
		}
	}
	return nil
}