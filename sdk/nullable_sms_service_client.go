package sdk

import (
	"context"

	"github.com/abmpio/abmp/pkg/log"
	pb "github.com/abmpio/hi-sdk/proto"
	"google.golang.org/grpc"
)

type nullableSmsServiceClient struct {
}

var _ pb.SmsServiceClient = (*nullableSmsServiceClient)(nil)

// #region pb.SmsServiceClient members

func (*nullableSmsServiceClient) SendSmsMessage(ctx context.Context, in *pb.SendSmsMessageRequest, opts ...grpc.CallOption) (*pb.SendSmsMessageResponse, error) {
	log.Logger.Warn("nullableSmsServiceClient.SendSmsMessage method")
	return &pb.SendSmsMessageResponse{}, nil
}

// #endregion
