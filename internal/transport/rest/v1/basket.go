package v1

import (
	"goods/internal/model"
	logger "goods/pkg/logger/zap"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (h *Handler) InitBasketsRoutes(v1 *gin.RouterGroup) {
	baskets := v1.Group("/customers")
	{
		baskets.GET("", h.getProducts, h.AuthCustomerMiddleware(h.authClient))
		baskets.POST("", h.addProduct, h.AuthCustomerMiddleware(h.authClient))
		baskets.PATCH("", h.updateProduct, h.AuthCustomerMiddleware(h.authClient))
		baskets.DELETE("", h.deleteProduct, h.AuthCustomerMiddleware(h.authClient))
	}
}

func (h *Handler) getProducts(c *gin.Context) {
	customerID, exists := c.Get("userID")
	if !exists {
		logger.Error(
			zap.String("action", "getValueFromContext"),
			zap.Int("status code", http.StatusUnauthorized),
			zap.Error(model.ErrContextIsEmpty),
		)
		c.JSON(http.StatusUnauthorized, gin.H{"message": "User not authorized"})
		return
	}
	customerIDStr, ok := customerID.(string)
	if !ok {
		logger.Error(
			zap.String("action", "convertValueToString"),
			zap.Int("status code", http.StatusUnauthorized),
			zap.Error(model.ErrContextIsEmpty),
		)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to get user ID"})
		return
	}

	products, err := h.services.Baskets.GetProducts(c.Request.Context(), customerIDStr)
	if err != nil {
		logger.Error(
			zap.String("action", "getProducts()"),
			zap.Int("status code", http.StatusBadRequest),
			zap.Error(err),
		)
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, products)
}

func (h *Handler) addProduct(c *gin.Context) {
	customerID, exists := c.Get("userID")
	if !exists {
		logger.Error(
			zap.String("action", "getValueFromContext"),
			zap.Int("status code", http.StatusUnauthorized),
			zap.Error(model.ErrContextIsEmpty),
		)
		c.JSON(http.StatusUnauthorized, gin.H{"message": "User not authorized"})
		return
	}
	customerIDStr, ok := customerID.(string)
	if !ok {
		logger.Error(
			zap.String("action", "convertValueToString"),
			zap.Int("status code", http.StatusUnauthorized),
			zap.Error(model.ErrContextIsEmpty),
		)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to get user ID"})
		return
	}

	var productRequest model.ProductRequest
	if err := c.BindJSON(&productRequest); err != nil {
		logger.Error(
			zap.String("action", "addProduct()"),
			zap.Int("status code", http.StatusBadRequest),
			zap.Error(err),
		)
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Baskets.AddProduct(c.Request.Context(), customerIDStr, productRequest); err != nil {
		logger.Error("Failed to add product to basket",
			zap.String("action", "AddProduct()"),
			zap.Int("status code", http.StatusBadRequest),
			zap.Error(err),
		)
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Product was added successfully"})
}

func (h *Handler) updateProduct(c *gin.Context) {
	customerID, exists := c.Get("userID")
	if !exists {
		logger.Error(
			zap.String("action", "getValueFromContext"),
			zap.Int("status code", http.StatusUnauthorized),
			zap.Error(model.ErrContextIsEmpty),
		)
		c.JSON(http.StatusUnauthorized, gin.H{"message": "User not authorized"})
		return
	}
	customerIDStr, ok := customerID.(string)
	if !ok {
		logger.Error(
			zap.String("action", "convertValueToString"),
			zap.Int("status code", http.StatusUnauthorized),
			zap.Error(model.ErrContextIsEmpty),
		)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to get user ID"})
		return
	}

	var productRequest model.ProductRequest
	if err := c.BindJSON(&productRequest); err != nil {
		logger.Error(
			zap.String("action", "updateProduct()"),
			zap.Int("status code", http.StatusBadRequest),
			zap.Error(err),
		)
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Baskets.UpdateProduct(c.Request.Context(), customerIDStr, productRequest); err != nil {
		logger.Error("Failed to update product to basket",
			zap.String("action", "UpdateProduct()"),
			zap.Int("status code", http.StatusBadRequest),
			zap.Error(err),
		)
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Product was updated successfully"})
}

func (h *Handler) deleteProduct(c *gin.Context) {
	customerID, exists := c.Get("userID")
	if !exists {
		logger.Error(
			zap.String("action", "getValueFromContext"),
			zap.Int("status code", http.StatusUnauthorized),
			zap.Error(model.ErrContextIsEmpty),
		)
		c.JSON(http.StatusUnauthorized, gin.H{"message": "User not authorized"})
		return
	}
	customerIDStr, ok := customerID.(string)
	if !ok {
		logger.Error(
			zap.String("action", "convertValueToString"),
			zap.Int("status code", http.StatusUnauthorized),
			zap.Error(model.ErrContextIsEmpty),
		)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to get user ID"})
		return
	}

	var productID int
	if err := c.BindJSON(&productID); err != nil {
		logger.Error(
			zap.String("action", "deleteProduct()"),
			zap.Int("status code", http.StatusBadRequest),
			zap.Error(err),
		)
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Baskets.DeleteProduct(c.Request.Context(), customerIDStr, productID); err != nil {
		logger.Error("Failed to add product to basket",
			zap.String("action", "AddProduct()"),
			zap.Int("status code", http.StatusBadRequest),
			zap.Error(err),
		)
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Product was deleted successfully"})
}
