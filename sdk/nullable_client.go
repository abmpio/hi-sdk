package sdk

import (
	"context"

	"github.com/abmpio/abmp/pkg/log"
	pb "github.com/abmpio/hi-sdk/proto"
	"google.golang.org/grpc"
)

type NullableClient struct {
	nullableCodeValueServiceClient
	nullableSmsServiceClient
	nullableSettingsServiceClient
}

var _ IClient = (*NullableClient)(nil)

// #region

func (*NullableClient) HealthCheck(ctx context.Context, in *pb.HealthCheckRequest, opts ...grpc.CallOption) (*pb.HealthCheckResponse, error) {
	return &pb.HealthCheckResponse{
		Status: pb.HealthCheckResponse_NOT_SERVING,
	}, nil
}

// #endregion

type nullableCodeValueServiceClient struct {
}

// #region pb.CodeValueServiceClient members

func (*nullableCodeValueServiceClient) FindListByCodeType(ctx context.Context, in *pb.FindCodeValueListByCodeTypeRequest, opts ...grpc.CallOption) (*pb.FindCodeValueListByCodeTypeResponse, error) {
	log.Logger.Warn("NullableClient.FindListByCodeType method")
	return &pb.FindCodeValueListByCodeTypeResponse{}, nil
}

func (*nullableCodeValueServiceClient) FindOneByCode(ctx context.Context, in *pb.FindOneCodeValueByCodeRequest, opts ...grpc.CallOption) (*pb.FindOneCodeValueByCodeResponse, error) {
	log.Logger.Warn("NullableClient.FindOneByCode method")
	return &pb.FindOneCodeValueByCodeResponse{}, nil
}

// #endregion

type nullableSmsServiceClient struct {
}

// #region pb.SmsServiceClient members

func (*nullableSmsServiceClient) SendSmsMessage(ctx context.Context, in *pb.SendSmsMessageRequest, opts ...grpc.CallOption) (*pb.SendSmsMessageResponse, error) {
	log.Logger.Warn("NullableClient.SendSmsMessage method")
	return &pb.SendSmsMessageResponse{}, nil
}

// #endregion

type nullableSettingsServiceClient struct {
}

// #region pb.SettingsServiceClient members

// 根据 appNameList 获取 Setting 列表
func (*nullableSettingsServiceClient) FindListByAppName(ctx context.Context, in *pb.FindSettingListByAppNameRequest, opts ...grpc.CallOption) (*pb.FindSettingListByAppNameResponse, error) {
	log.Logger.Warn("NullableClient.FindListByAppName method")
	return &pb.FindSettingListByAppNameResponse{}, nil
}

// 根据 appName 和 name 获取单个 Setting
func (*nullableSettingsServiceClient) FindByName(ctx context.Context, in *pb.FindSettingByNameRequest, opts ...grpc.CallOption) (*pb.FindSettingByNameResponse, error) {
	log.Logger.Warn("NullableClient.FindByName method")
	return &pb.FindSettingByNameResponse{}, nil
}

// #endregion
