package sdk

import (
	"context"

	pb "github.com/abmpio/hi-sdk/proto"
	"google.golang.org/grpc"
)

type NullableClient struct {
	nullableCodeValueServiceClient
	nullableSmsServiceClient
	nullableSettingsServiceClient
}

var _ IClient = (*NullableClient)(nil)
var _ pb.HiClient = (*NullableClient)(nil)

// #region

func (*NullableClient) HealthCheck(ctx context.Context, in *pb.HealthCheckRequest, opts ...grpc.CallOption) (*pb.HealthCheckResponse, error) {
	return &pb.HealthCheckResponse{
		Status: pb.HealthCheckResponse_NOT_SERVING,
	}, nil
}

// #endregion
