package config

import (
	"goods/internal/model"
	"goods/pkg/database/psql"
	logger "goods/pkg/logger/zap"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	logger.InitLogger()
	tests := []struct {
		configDIR string
		envDIR    string
		expected  *Config
		err       error
	}{
		{
			configDIR: "./fixturess",
			envDIR:    "./fixtures/.env",
			expected:  &Config{},
			err:       model.ErrNotFoundConfigFile,
		},
		{
			configDIR: "./fixtures",
			envDIR:    "./fixturess/.env",
			expected:  &Config{},
			err:       model.ErrNotFoundEnvFile,
		},
		{
			configDIR: "./fixtures",
			envDIR:    "./fixtures/.env",
			expected: &Config{
				HTTP: HttpConfig{
					Addr:           ":8080",
					ReadTimeout:    10 * time.Second,
					WriteTimeout:   10 * time.Second,
					MaxHeaderBytes: 1,
				},
				GRPC: GrpcConfig{},
				PSQL: psql.PSQlConfig{
					Host:     "test",
					Port:     9999,
					Username: "test",
					Name:     "test",
					SSLmode:  "disable",
					Password: "test",
				},
			},
			err: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.configDIR, func(t *testing.T) {
			actual, err := Init(test.configDIR, test.envDIR)
			assert.Equal(t, test.expected, actual)
			assert.Equal(t, test.err, err)
		})
	}
}
