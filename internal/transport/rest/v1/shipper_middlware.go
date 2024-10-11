package v1

import (
	"goods/internal/model"
	grpc_client "goods/internal/transport/grpc/client"
	logger "goods/pkg/logger/zap"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (h *Handler) AuthShipperMiddleware(client *grpc_client.AuthClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("session")
		if err != nil {
			if err == http.ErrNoCookie {
				logger.Error(
					zap.String("middleware", "shipper"),
					zap.Int("status code", http.StatusUnauthorized),
					zap.Error(http.ErrNoCookie),
				)
				c.JSON(http.StatusUnauthorized, gin.H{"message": "Cookie is empty"})
				c.Abort()
				return
			}
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			c.Abort()
			return
		}

		grpcResponse, err := client.Verify(token)
		if err != nil {
			logger.Error(
				zap.String("middleware", "shipper"),
				zap.Int("status code", http.StatusUnauthorized),
				zap.Error(err),
			)
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid session"})
			c.Abort()
			return
		}

		if grpcResponse.Role != model.RoleShipper {
			logger.Error(
				zap.String("middleware", "shipper"),
				zap.Int("status code", http.StatusForbidden),
				zap.Error(err),
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
