package grpc_client

import (
	"context"
	"goods/internal/config"
	"goods/internal/model"
	proto "goods/internal/server/grpc/proto/payment"
	logger "goods/pkg/logger/zap"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

type PaymentClient struct {
	conn   *grpc.ClientConn
	client proto.PaymentServiceClient
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
	client := proto.NewPaymentServiceClient(conn)
	return &PaymentClient{conn: conn, client: client}, nil
}

func (c *PaymentClient) Close() {
	if c.conn != nil {
		c.conn.Close()
	}
}

func (c *PaymentClient) ProcessPayment(ctx context.Context, userData model.UserData, orderID int64, amount float64, stripeToken string) (*proto.PaymentResponse, error) {
	paymentIntent := &proto.PaymentIntentData{
		Amount:        int64(amount * 100),
		Currency:      "UAH",
		PaymentMethod: stripeToken,
		Confirm:       true,
		OrderId:       orderID,
	}

	customer := &proto.CustomerData{
		Name:  userData.Fullname,
		Email: userData.Email,
		Phone: userData.Phone,
	}

	req := &proto.CreatePaymentIntentRequest{
		PaymentIntent: paymentIntent,
		Customer:      customer,
	}
	resp, err := c.client.CreatePaymentIntent(ctx, req)
	if err != nil {
		if st, ok := status.FromError(err); ok {
			logger.Error(
				zap.String("server", "grpc"),
				zap.String("action", "createPaymentIntent()"),
				zap.String("response", st.Err().Error()),
			)
			return nil, st.Err()
		}
		logger.Error(
			zap.String("server", "grpc"),
			zap.String("action", "createPaymentIntent()"),
			zap.Error(err),
		)
		return nil, err
	}
	return resp, nil
}
