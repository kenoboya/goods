package v1

import (
	"goods/internal/service"
	grpc_client "goods/internal/transport/grpc/client"

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

func (h *Handler) Init(router *gin.RouterGroup) {
	v1 := router.Group("/v1")
	{
		h.InitCategoriesRoutes(v1)
		h.InitProductsRoutes(v1)
	}
}
