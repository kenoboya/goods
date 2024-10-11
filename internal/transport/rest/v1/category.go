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
		categories.POST("/create", h.createCategory, h.AuthAdminMiddleware(h.authClient))
		categories.GET("", h.getCategories)
		categories.GET("/:category", h.getCategory)
	}
}

func (h Handler) createCategory(c *gin.Context) {
	var categoryRequest model.CreateCategoryRequest
	if err := c.BindJSON(&categoryRequest); err != nil {
		logger.Error(
			zap.String("action", "createCategory()"),
			zap.Int("status code", http.StatusBadRequest),
			zap.Error(err),
		)
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Categories.CreateCategories(c.Request.Context(), categoryRequest); err != nil {
		logger.Error("Failed to create category",
			zap.String("action", "createCategory()"),
			zap.Int("status code", http.StatusBadRequest),
			zap.Error(err),
		)
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Category was created successfully"})
}
func (h *Handler) getCategories(c *gin.Context) {
	categories, err := h.services.Categories.GetCategories(c.Request.Context())
	if err != nil {
		logger.Error(
			zap.String("action", "getCategories()"),
			zap.Int("status code", http.StatusBadRequest),
			zap.Error(err),
		)
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, categories)
}

func (h *Handler) getCategory(c *gin.Context) {
	title := getPrefixFromParam(c, "category")
	if title == "" {
		logger.Error(
			zap.String("action", "getCategoryByName()"),
			zap.String("prefix", title),
			zap.Int("status code", http.StatusBadRequest),
			zap.Error(model.ErrEmptyParam),
		)
		newResponse(c, http.StatusBadRequest, model.ErrEmptyParam.Error())
		return
	}
	category, err := h.services.Categories.GetCategoryByName(c.Request.Context(), title)
	if err != nil {
		logger.Error(
			zap.String("action", "getCategoryByName()"),
			zap.String("prefix", title),
			zap.Int("status code", http.StatusBadRequest),
			zap.Error(err),
		)
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	products, err := h.services.Products.GetProductsByCategoryID(c.Request.Context(), category.CategoryID)
	if err != nil {
		logger.Error(
			zap.String("action", "getCategoryByCategoryName()"),
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
