package grpc_client

package grpc_client

import (
	"context"
	"goods/internal/config"
	"goods/internal/server/grpc/proto"
	logger "goods/pkg/logger/zap"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

type PaymentClient struct {
	conn   *grpc.ClientConn
	client proto.AuthServiceClient
}

func NewPaymentClient(cfg config.GrpcConfig) (*PaymentClient, error) {
	var opts []grpc.DialOption
	// There is currently no certificate available.
	// creds, err := credentials.NewClientTLSFromFile(cfg.CertFile, "")
	// if err != nil {
	// 	return nil, err
	// }
	// opts = append(opts, grpc.WithTransportCredentials(creds))

	conn, err := grpc.NewClient(cfg.AuthAddr, opts...)
	if err != nil {
		logger.Error("Failed to create payment client",
			zap.String("server", "grpc"),
			zap.String("address", cfg.AuthAddr),
			zap.Error(err),
		)
		return nil, err
	}
	// client := proto.NewAuthServiceClient(conn)
	return &PaymentClient{conn: conn, client: client}, nil
}

func (c *PaymentClient) Close() {
	if c.conn != nil {
		c.conn.Close()
	}
}


func (c *PaymentClient) ProcessPayment(orderID int64, amount float64) (error){

}