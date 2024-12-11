package grpc_client

import (
	"context"
	"goods/internal/config"
	proto "goods/internal/server/grpc/proto/auth"
	logger "goods/pkg/logger/zap"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

type AuthClient struct {
	conn   *grpc.ClientConn
	client proto.AuthServiceClient
}

func NewAuthClient(cfg config.GrpcConfig) (*AuthClient, error) {
	var opts []grpc.DialOption
	// There is currently no certificate available.
	// creds, err := credentials.NewClientTLSFromFile(cfg.CertFile, "")
	// if err != nil {
	// 	return nil, err
	// }
	// opts = append(opts, grpc.WithTransportCredentials(creds))
	conn, err := grpc.NewClient(cfg.AuthAddr, opts...)
	if err != nil {
		logger.Error("Failed to create auth client",
			zap.String("server", "grpc"),
			zap.String("address", cfg.AuthAddr),
			zap.Error(err),
		)
		return nil, err
	}
	client := proto.NewAuthServiceClient(conn)
	return &AuthClient{conn: conn, client: client}, nil
}

func (c *AuthClient) GetUserInformation(ctx context.Context, userID string) (*proto.UserInformationResponse, error) {
	req := &proto.UserInformationRequest{UserId: userID}
	resp, err := c.client.GetUserInformation(ctx, req)
	if err != nil {
		if st, ok := status.FromError(err); ok {
			logger.Error(
				zap.String("server", "grpc"),
				zap.String("action", "GetUserInformation"),
				zap.String("response", st.Err().Error()),
			)
			return nil, st.Err()
		}
		logger.Error(
			zap.String("server", "grpc"),
			zap.String("action", "GetUserInformation"),
			zap.Error(err),
		)
		return nil, err
	}
	return resp, nil
}
func (c *AuthClient) Verify(ctx context.Context, sessionToken string) (*proto.UserResponse, error) {
	req := &proto.TokenRequest{SessionToken: sessionToken}
	resp, err := c.client.Verify(ctx, req)
	if err != nil {
		if st, ok := status.FromError(err); ok {
			logger.Error(
				zap.String("server", "grpc"),
				zap.String("action", "verify()"),
				zap.String("response", st.Err().Error()),
			)
			return nil, st.Err()
		}
		logger.Error(
			zap.String("server", "grpc"),
			zap.String("action", "verify()"),
			zap.Error(err),
		)
		return nil, err
	}
	return resp, nil
}

func (c *AuthClient) Close() {
	if c.conn != nil {
		c.conn.Close()
	}
}
