package v1

import (
	"goods/internal/service"
	grpc_client "goods/internal/transport/grpc/client"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services   *service.Services
	authClient *grpc_client.AuthClient
	paymentClient *grpc_client.PaymentClient
}

func NewHandler(services *service.Services, authClient *grpc_client.AuthClient, paymentClient *grpc_client.PaymentClient) *Handler {
	return &Handler{
		services:      services,
		authClient:    authClient,
		paymentClient: paymentClient,
	}
}

func (h *Handler) Init(router *gin.RouterGroup) {
	v1 := router.Group("/v1")
	{
		h.InitCategoriesRoutes(v1)
		h.InitProductsRoutes(v1)
		h.InitCustomersRoutes(v1)
		h.InitBasketsRoutes(v1)
		h.InitOrdersRoutes(v1)
	}
}
