package sdk

import (
	"context"

	"github.com/abmpio/abmp/pkg/log"
	pb "github.com/abmpio/hi-sdk/proto"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type nullableSettingsServiceClient struct {
}

var _ pb.SettingsServiceClient = (*nullableSettingsServiceClient)(nil)

// #region pb.SettingsServiceClient members

// 根据 appNameList 获取 Setting 列表
func (*nullableSettingsServiceClient) FindListByAppName(ctx context.Context, in *pb.FindSettingListByAppNameRequest, opts ...grpc.CallOption) (*pb.FindSettingListByAppNameResponse, error) {
	log.Logger.Warn("nullableSettingsServiceClient.FindListByAppName method")
	return &pb.FindSettingListByAppNameResponse{}, nil
}

// 根据 appName 和 name 获取单个 Setting
func (*nullableSettingsServiceClient) FindByName(ctx context.Context, in *pb.FindSettingByNameRequest, opts ...grpc.CallOption) (*pb.FindSettingByNameResponse, error) {
	log.Logger.Warn("nullableSettingsServiceClient.FindByName method")
	return &pb.FindSettingByNameResponse{}, nil
}

// 更新 Setting 的值
func (*nullableSettingsServiceClient) UpdateSettingValue(ctx context.Context, in *pb.UpdateSettingValueRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	log.Logger.Warn("nullableSettingsServiceClient.UpdateSettingValue method")
	return &emptypb.Empty{}, nil
}

// #endregion
