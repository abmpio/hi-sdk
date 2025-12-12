package sdk

import (
	"context"

	"github.com/abmpio/abmp/pkg/log"
	pb "github.com/abmpio/hi-sdk/proto"
	"google.golang.org/grpc"
)

type nullableCodeValueServiceClient struct {
}

var _ pb.CodeValueServiceClient = (*nullableCodeValueServiceClient)(nil)

// #region pb.CodeValueServiceClient members

func (*nullableCodeValueServiceClient) FindListByCodeType(ctx context.Context, in *pb.FindCodeValueListByCodeTypeRequest, opts ...grpc.CallOption) (*pb.FindCodeValueListByCodeTypeResponse, error) {
	log.Logger.Warn("nullableCodeValueServiceClient.FindListByCodeType method")
	return &pb.FindCodeValueListByCodeTypeResponse{}, nil
}

func (*nullableCodeValueServiceClient) FindOneByCode(ctx context.Context, in *pb.FindOneCodeValueByCodeRequest, opts ...grpc.CallOption) (*pb.FindOneCodeValueByCodeResponse, error) {
	log.Logger.Warn("nullableCodeValueServiceClient.FindOneByCode method")
	return &pb.FindOneCodeValueByCodeResponse{}, nil
}

// #endregion
