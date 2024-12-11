package v1

import (
	"goods/internal/model"
	proto "goods/internal/server/grpc/proto/auth"
	grpc_client "goods/internal/transport/grpc/client"
	logger "goods/pkg/logger/zap"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (h *Handler) AuthCustomerMiddleware(client *grpc_client.AuthClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		grpcResponse, err := h.middleware(c, client)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid session"})
			c.Abort()
			return
		}

		if grpcResponse.Role != model.RoleCustomer {
			logger.Error(
				zap.String("middleware", "customer"),
				zap.Int("status code", http.StatusForbidden),
				zap.Error(model.ErrInvalidRole),
			)
			c.JSON(http.StatusForbidden, gin.H{"message": "Access denied"})
			c.Abort()
			return
		}

		c.Set("userID", grpcResponse.UserId)
		c.Set("role", grpcResponse.Role)

		c.Next()
	}
}

func (h *Handler) AuthShipperMiddleware(client *grpc_client.AuthClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		grpcResponse, err := h.middleware(c, client)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid session"})
			c.Abort()
			return
		}

		if grpcResponse.Role != model.RoleShipper {
			logger.Error(
				zap.String("middleware", "shipper"),
				zap.Int("status code", http.StatusForbidden),
				zap.Error(model.ErrInvalidRole),
			)
			c.JSON(http.StatusForbidden, gin.H{"message": "Access denied"})
			c.Abort()
			return
		}

		c.Set("userID", grpcResponse.UserId)
		c.Set("role", grpcResponse.Role)

		c.Next()
	}
}

func (h *Handler) AuthAdminMiddleware(client *grpc_client.AuthClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		grpcResponse, err := h.middleware(c, client)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid session"})
			c.Abort()
			return
		}

		if grpcResponse.Role != model.RoleAdmin {
			logger.Error(
				zap.String("middleware", "admin"),
				zap.Int("status code", http.StatusForbidden),
				zap.Error(model.ErrInvalidRole),
			)
			c.JSON(http.StatusForbidden, gin.H{"message": "Access denied"})
			c.Abort()
			return
		}

		c.Set("userID", grpcResponse.UserId)
		c.Set("role", grpcResponse.Role)

		c.Next()
	}
}

func (h *Handler) OrderMiddleware(client *grpc_client.AuthClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		grpcResponse, err := h.middleware(c, client)
		if err != nil {
			c.Set("userID", "guest")
			c.Next()
			return
		}
		c.Set("userID", grpcResponse.UserId)
		grpcResponseData, err := client.GetUserInformation(c.Request.Context(), grpcResponse.UserId)
		if err != nil {
			c.Set("userID", "guest")
			c.Next()
			return
		}
		c.Set("userID", grpcResponse.UserId)
		c.Set("fullname", grpcResponseData.Fullname)
		c.Set("email", grpcResponseData.Email)
		c.Set("phone", grpcResponseData.Phone)
		c.Next()
	}
}

func (h *Handler) middleware(c *gin.Context, client *grpc_client.AuthClient) (*proto.UserResponse, error) {
	token, err := c.Cookie("session")
	if err != nil {
		if err == http.ErrNoCookie {
			logger.Error(
				zap.String("middleware", "middleware"),
				zap.Int("status code", http.StatusUnauthorized),
				zap.Error(http.ErrNoCookie),
			)
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Cookie is empty"})
			c.Abort()
			return nil, err
		}
		return nil, err
	}

	grpcResponse, err := client.Verify(c.Request.Context(), token)
	if err != nil {
		logger.Error(
			zap.String("middleware", "middleware"),
			zap.Int("status code", http.StatusUnauthorized),
			zap.Error(err),
		)
		return nil, err
	}
	return grpcResponse, nil
}
