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
		orders.POST("", h.createOrder, h.OrderMiddleware(h.authClient))
	}
}

func (h *Handler) createOrder(c *gin.Context) {
	var orderRequest model.OrderRequest
	if err := c.BindJSON(&orderRequest); err != nil {
		logger.Error("Invalid JSON request",
			zap.String("action", "createOrder()"),
			zap.Int("status code", http.StatusBadRequest),
			zap.Error(err),
		)
		newResponse(c, http.StatusBadRequest, "Invalid request format")
		return
	}

	// Создание заказа
	orderID, err := h.services.Orders.CreateOrder(c.Request.Context(), orderRequest)
	if err != nil {
		logger.Error("Failed to create order",
			zap.String("action", "createOrder()"),
			zap.Int("status code", http.StatusInternalServerError),
			zap.Error(err),
		)
		newResponse(c, http.StatusInternalServerError, "Failed to create order")
		return
	}

	// Получение суммы заказа
	amount, err := h.services.Orders.GetTotalSumOrder(c.Request.Context(), orderID)
	if err != nil || amount < 0 {
		logger.Error("Failed to get total order sum",
			zap.String("action", "GetTotalSumOrder()"),
			zap.Int("status code", http.StatusInternalServerError),
			zap.Error(err),
		)
		newResponse(c, http.StatusInternalServerError, "Invalid order total sum")
		return
	}

	// Получение данных пользователя
	userID, fullname, email, phone := getUserDataFromParam(c)
	if userID == "guest" || userID == "" {
		fullname = orderRequest.Customer.CustomerFullName
		email = ""
		phone = orderRequest.Customer.CustomerPhone
	}

	// Обработка способа оплаты
	switch orderRequest.PaymentMethod {
	case model.PaymentMethodCard:
		h.handleCardPayment(c, orderID, amount, fullname, email, phone, orderRequest.PaymentToken)
	case model.PaymentMethodCash:
		c.JSON(http.StatusOK, gin.H{"message": "Order was created successfully"})
	default:
		newResponse(c, http.StatusBadRequest, "Unknown payment method")
	}
}

// Обработка платежа картой
func (h *Handler) handleCardPayment(c *gin.Context, orderID int64, amount float64, fullname, email, phone string, paymentToken *string) {
	resp, err := h.paymentClient.ProcessPayment(c.Request.Context(), model.UserData{
		Fullname: fullname,
		Email:    email,
		Phone:    phone,
	}, orderID, amount, *paymentToken)
	if err != nil {
		logger.Error("Payment failed",
			zap.String("action", "ProcessPayment"),
			zap.Int("status code", http.StatusInternalServerError),
			zap.Error(err),
		)
		newResponse(c, http.StatusInternalServerError, "Payment failed")
		return
	}

	if err = h.services.Orders.LinkTransactionToOrder(c.Request.Context(), orderID, resp.TransactionId); err != nil {
		logger.Error("Failed to link transaction to order",
			zap.String("action", "LinkTransactionToOrder"),
			zap.Int("status code", http.StatusInternalServerError),
			zap.Error(err),
		)
		newResponse(c, http.StatusInternalServerError, "Failed to link transaction to order")
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"message": "Order created and payment processed successfully"})
}
