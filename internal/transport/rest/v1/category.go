package v1

import (
	"goods/internal/model"
	logger "goods/pkg/logger/zap"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (h *Handler) InitCategoriesRoutes(v1 *gin.RouterGroup) {
	categories := v1.Group("/categories")
	{
		categories.GET("", h.GetCategories)
		categories.GET("/:category", h.GetCategory)
	}
}

func (h *Handler) GetCategories(c *gin.Context) {
	categories, err := h.services.Categories.GetCategories(c.Request.Context())
	if err != nil {
		logger.Error("Failed to get categories",
			zap.String("action", "GetCategories()"),
			zap.Int("status code", http.StatusBadRequest),
			zap.Error(err),
		)
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, categories)
}

func (h *Handler) GetCategory(c *gin.Context) {
	title := getPrefixFromParam(c, "category")
	category, err := h.services.Categories.GetCategoryByName(c.Request.Context(), title)
	if err != nil {
		logger.Error("Failed to get category",
			zap.String("action", "GetCategoryByName()"),
			zap.String("prefix", title),
			zap.Int("status code", http.StatusBadRequest),
			zap.Error(err),
		)
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	products, err := h.services.Products.GetProductsByCategoryID(c.Request.Context(), category.CategoryID)
	if err != nil {
		logger.Error("Failed to get category",
			zap.String("action", "GetCategoryByCategoryName()"),
			zap.Int("category_id", int(category.CategoryID)),
			zap.Int("status code", http.StatusInternalServerError),
			zap.Error(err),
		)
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, model.Goods{
		Category: category,
		Products: products,
	})
}
