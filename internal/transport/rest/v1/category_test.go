package v1

import (
	"encoding/json"
	"fmt"
	"goods/internal/model"
	"goods/internal/service"
	mock_service "goods/internal/service/mocks"
	logger "goods/pkg/logger/zap"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/magiconair/properties/assert"
	gomock "go.uber.org/mock/gomock"
	"golang.org/x/net/context"
)

var (
	category = model.Category{
		CategoryID:   1,
		CategoryName: "burger",
	}
)

func TestHandler_getCategory(t *testing.T) {
	logger.InitLogger()
	type mockBehavior func(mockCategories *mock_service.MockCategories, mockProducts *mock_service.MockProducts, categoryName string)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCategories := mock_service.NewMockCategories(ctrl)
	mockProducts := mock_service.NewMockProducts(ctrl)

	tests := []struct {
		name         string
		categoryName string
		mockBehavior mockBehavior
		statusCode   int
		responseBody string
	}{
		{
			name:         "ok",
			categoryName: "burger",
			mockBehavior: func(mockCategories *mock_service.MockCategories, mockProducts *mock_service.MockProducts, categoryName string) {
				mockCategories.EXPECT().GetCategoryByName(context.Background(), categoryName).Return(category, nil)
				mockProducts.EXPECT().GetProductsByCategoryID(context.Background(), category.CategoryID).Return(products, nil)
			},
			statusCode: http.StatusOK,
			responseBody: func() string {
				goods := model.Goods{
					Category: category,
					Products: products,
				}
				responseBytes, _ := json.Marshal(goods)
				return string(responseBytes)
			}(),
		},
		{
			name:         "empty category",
			categoryName: "",
			mockBehavior: func(mockCategories *mock_service.MockCategories, mockProducts *mock_service.MockProducts, categoryName string) {
			},
			statusCode:   http.StatusNotFound,
			responseBody: `404 page not found`,
		},
		{
			name:         "category not found",
			categoryName: "pizia",
			mockBehavior: func(mockCategories *mock_service.MockCategories, mockProducts *mock_service.MockProducts, categoryName string) {
				mockCategories.EXPECT().GetCategoryByName(context.Background(), categoryName).Return(model.Category{}, model.ErrNotFoundCategory)
			},
			statusCode:   http.StatusBadRequest,
			responseBody: `{"message":"category not found"}`,
		},
		{
			name:         "products not found",
			categoryName: "burger",
			mockBehavior: func(mockCategories *mock_service.MockCategories, mockProducts *mock_service.MockProducts, categoryName string) {
				mockCategories.EXPECT().GetCategoryByName(context.Background(), categoryName).Return(category, nil)
				mockProducts.EXPECT().GetProductsByCategoryID(context.Background(), category.CategoryID).Return([]model.Product{}, model.ErrNotFoundProducts)
			},
			statusCode:   http.StatusInternalServerError,
			responseBody: `{"message":"products not found"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mockBehavior(mockCategories, mockProducts, test.categoryName)

			services := &service.Services{
				Products:   mockProducts,
				Categories: mockCategories,
			}
			handler := NewHandler(services)

			r := gin.New()
			r.GET("/categories/:category", handler.getCategory)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", fmt.Sprintf("/categories/%s", test.categoryName), nil)

			r.ServeHTTP(w, req)

			assert.Equal(t, test.statusCode, w.Code)
			assert.Equal(t, test.responseBody, w.Body.String())
		})
	}
}
