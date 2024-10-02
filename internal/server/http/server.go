package http

import (
	"context"
	"goods/internal/config"
	"goods/internal/transport/rest"
	logger "goods/pkg/logger/zap"
	"net/http"

	"go.uber.org/zap"
)

type Server struct {
	srv *http.Server
}

func NewServer(config config.HttpConfig, handler *rest.Handler) *Server {
	return &Server{
		srv: &http.Server{
			Addr:           config.Addr,
			ReadTimeout:    config.ReadTimeout,
			WriteTimeout:   config.WriteTimeout,
			MaxHeaderBytes: config.MaxHeaderBytes,
			Handler:        handler.Init(&config),
		},
	}
}

func (s *Server) Run() error {
	if err := s.srv.ListenAndServe(); err != nil {
		logger.Error("Failed to run server",
			zap.String("server", "http"),
			zap.Error(err),
		)
		return err
	}
	return nil
}

func (s *Server) Shutdown(ctx context.Context) error {
	if err := s.srv.Shutdown(ctx); err != nil {
		logger.Error("Failed to shutdown server",
			zap.String("server", "http"),
			zap.Error(err),
		)
		return err
	}
	return nil
}
