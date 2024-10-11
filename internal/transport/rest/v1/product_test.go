package v1

import (
	"context"
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
	"github.com/stretchr/testify/assert"
	gomock "go.uber.org/mock/gomock"
)

var (
	products = []model.Product{
		{
			ProductID:       1,
			ProductName:     "Cheeseburger",
			SupplierID:      5,
			CategoryID:      1,
			UnitPrice:       5.99,
			Stock:           true,
			Discount:        0.1,
			QuantityPerUnit: func(s string) *string { return &s }("1 burger"),
			Weight:          func(s string) *string { return &s }("200g"),
			Image:           func(s string) *string { return &s }("cheeseburger.png"),
		},
		{
			ProductID:       2,
			ProductName:     "Chicken Sandwich",
			SupplierID:      3,
			CategoryID:      1,
			UnitPrice:       6.49,
			Stock:           false,
			Discount:        0.15,
			QuantityPerUnit: func(s string) *string { return &s }("1 sandwich"),
			Weight:          func(s string) *string { return &s }("250g"),
			Image:           nil,
		},
		{
			ProductID:       3,
			ProductName:     "Veggie Wrap",
			SupplierID:      2,
			CategoryID:      1,
			UnitPrice:       4.99,
			Stock:           true,
			Discount:        0.05,
			QuantityPerUnit: nil,
			Weight:          func(s string) *string { return &s }("180g"),
			Image:           nil,
		},
	}
)

func TestHandler_getProduct(t *testing.T) {
	logger.InitLogger()
	type mockBehavior func(mock *mock_service.MockProducts, productName string)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock_service.NewMockProducts(ctrl)

	tests := []struct {
		name         string
		productName  string
		mockBehavior mockBehavior
		statusCode   int
		responseBody string
	}{
		{
			name:        "ok",
			productName: "Cheeseburger",
			mockBehavior: func(mock *mock_service.MockProducts, productName string) {
				mock.EXPECT().GetProductByName(context.Background(), productName).Return(products[0], nil)
			},
			statusCode: http.StatusOK,
			responseBody: func() string {
				responseBytes, _ := json.Marshal(products[0])
				return string(responseBytes)
			}(),
		},
		{
			name:        "empty product",
			productName: "",
			mockBehavior: func(mock *mock_service.MockProducts, productName string) {
			},
			statusCode:   http.StatusNotFound,
			responseBody: `404 page not found`,
		},
		{
			name:        "product not found",
			productName: "bred",
			mockBehavior: func(mock *mock_service.MockProducts, productName string) {
				mock.EXPECT().GetProductByName(context.Background(), productName).Return(model.Product{}, model.ErrNotFoundProduct)
			},
			statusCode:   http.StatusBadRequest,
			responseBody: `{"message":"product not found"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mockBehavior(mock, test.productName)

			services := &service.Services{Products: mock}
			handler := NewHandler(services, nil)

			r := gin.New()

			r.GET("/products/:product", handler.getProduct)

			w := httptest.NewRecorder()

			req := httptest.NewRequest("GET", fmt.Sprintf("/products/%s", test.productName), nil)

			r.ServeHTTP(w, req)

			assert.Equal(t, test.statusCode, w.Code)
			assert.Equal(t, test.responseBody, w.Body.String())
		})
	}
}
