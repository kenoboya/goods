package v1

import (
	"goods/internal/model"
	logger "goods/pkg/logger/zap"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (h *Handler) InitCustomersRoutes(v1 *gin.RouterGroup) {
	customers := v1.Group("/customers")
	{
		customers.GET("", h.getCustomers, h.AuthAdminMiddleware(h.authClient))
		customers.GET("/:customerID", h.getCustomer, h.AuthAdminMiddleware(h.authClient))
	}
}
func (h *Handler) getCustomers(c *gin.Context) {
	customers, err := h.services.Customers.GetCustomers(c.Request.Context())
	if err != nil {
		logger.Error(
			zap.String("action", "getCustomers()"),
			zap.Int("status code", http.StatusBadRequest),
			zap.Error(err),
		)
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, customers)
}

func (h *Handler) getCustomer(c *gin.Context) {
	title := getPrefixFromParam(c, "customerID")
	if title == "" {
		logger.Error(
			zap.String("action", "getCustomerByID"),
			zap.String("prefix", title),
			zap.Int("status code", http.StatusBadRequest),
			zap.Error(model.ErrEmptyParam),
		)
		newResponse(c, http.StatusBadRequest, model.ErrEmptyParam.Error())
		return
	}
	customerID, err := strconv.ParseInt(title, 10, 64)
	if err != nil {
		logger.Error(
			zap.String("prefix", title),
			zap.Int("status code", http.StatusBadRequest),
			zap.Error(err),
		)
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	customer, err := h.services.Customers.GetCustomerByID(c.Request.Context(), customerID)
	if err != nil {
		logger.Error(
			zap.String("action", "getCustomerByID()"),
			zap.String("prefix", title),
			zap.Int("status code", http.StatusBadRequest),
			zap.Error(err),
		)
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, customer)
}
