package v1

import (
	"goods/internal/model"
	logger "goods/pkg/logger/zap"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (h *Handler) InitProductsRoutes(v1 *gin.RouterGroup) {
	products := v1.Group("/products")
	{
		products.GET("/:product", h.getProduct)
	}
}

func (h *Handler) getProduct(c *gin.Context) {
	title := getPrefixFromParam(c, "product")
	if title == "" {
		logger.Error(
			zap.String("action", "GetCategoryByName()"),
			zap.String("prefix", title),
			zap.Int("status code", http.StatusBadRequest),
			zap.Error(model.ErrEmptyParam),
		)
		newResponse(c, http.StatusBadRequest, model.ErrEmptyParam.Error())
		return
	}
	product, err := h.services.Products.GetProductByName(c.Request.Context(), title)
	if err != nil {
		logger.Error(
			zap.String("action", "GetProductByName()"),
			zap.String("prefix", title),
			zap.Int("status code", http.StatusBadRequest),
			zap.Error(err),
		)
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, product)
}
