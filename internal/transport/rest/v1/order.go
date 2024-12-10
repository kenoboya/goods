package v1

import (
	"goods/internal/model"
	logger "goods/pkg/logger/zap"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (h *Handler) InitOrdersRoutes(v1 *gin.RouterGroup) {
	orders := v1.Group("/checkout")
	{
		orders.POST("", h.createOrder)
	}
}

func (h *Handler) createOrder(c *gin.Context) {
	var orderID int64
	var amount float64
	var err error
	var orderRequest model.OrderRequest

	if err = c.BindJSON(&orderRequest); err != nil {
		logger.Error(
			zap.String("action", "createOrder()"),
			zap.Int("status code", http.StatusBadRequest),
			zap.Error(err),
		)
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if orderID, err = h.services.Orders.CreateOrder(c.Request.Context(), orderRequest); err != nil {
		logger.Error("Failed to create order",
			zap.String("action", "createOrder()"),
			zap.Int("status code", http.StatusBadRequest),
			zap.Error(err),
		)
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	switch orderRequest.PaymentMethod {
	case model.PaymentMethodCard:
		if amount, err = h.services.Orders.GetTotalSumOrder(c.Request.Context(), orderID); err != nil {
			logger.Error("Failed to get total sum of order",
				zap.String("action", "GetTotalSumOrder()"),
				zap.Int("status code", http.StatusBadRequest),
				zap.Error(err),
			)
			newResponse(c, http.StatusBadRequest, err.Error())
			return
		}

	}

	c.JSON(http.StatusOK, gin.H{"message": "Order was created successfully"})
}

func (h *Handler) processPayment(orderID int64, amount float64) {

}
