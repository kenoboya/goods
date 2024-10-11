package rest

import (
	"goods/internal/config"
	"goods/internal/service"
	grpc_client "goods/internal/transport/grpc/client"
	v1 "goods/internal/transport/rest/v1"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services   *service.Services
	authClient *grpc_client.AuthClient
}

func NewHandler(services *service.Services, authClient *grpc_client.AuthClient) *Handler {
	return &Handler{
		services:   services,
		authClient: authClient,
	}
}

func (h *Handler) Init(cfg *config.HttpConfig) *gin.Engine {
	router := gin.Default()
	router.Use(
		gin.Recovery(),
		gin.Logger(),
	)
	h.initAPI(router)
	return router
}

func (h *Handler) initAPI(router *gin.Engine) {
	handlerV1 := v1.NewHandler(h.services, h.authClient)
	api := router.Group("/api")
	{
		handlerV1.Init(api)
	}
}
