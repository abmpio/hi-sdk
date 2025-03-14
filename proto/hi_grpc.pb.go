// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: hi.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Hi_HealthCheck_FullMethodName = "/proto.Hi/HealthCheck"
)

// HiClient is the client API for Hi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HiClient interface {
	HealthCheck(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error)
}

type hiClient struct {
	cc grpc.ClientConnInterface
}

func NewHiClient(cc grpc.ClientConnInterface) HiClient {
	return &hiClient{cc}
}

func (c *hiClient) HealthCheck(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(HealthCheckResponse)
	err := c.cc.Invoke(ctx, Hi_HealthCheck_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HiServer is the server API for Hi service.
// All implementations must embed UnimplementedHiServer
// for forward compatibility.
type HiServer interface {
	HealthCheck(context.Context, *HealthCheckRequest) (*HealthCheckResponse, error)
	mustEmbedUnimplementedHiServer()
}

// UnimplementedHiServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedHiServer struct{}

func (UnimplementedHiServer) HealthCheck(context.Context, *HealthCheckRequest) (*HealthCheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HealthCheck not implemented")
}
func (UnimplementedHiServer) mustEmbedUnimplementedHiServer() {}
func (UnimplementedHiServer) testEmbeddedByValue()            {}

// UnsafeHiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HiServer will
// result in compilation errors.
type UnsafeHiServer interface {
	mustEmbedUnimplementedHiServer()
}

func RegisterHiServer(s grpc.ServiceRegistrar, srv HiServer) {
	// If the following call pancis, it indicates UnimplementedHiServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Hi_ServiceDesc, srv)
}

func _Hi_HealthCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HiServer).HealthCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Hi_HealthCheck_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HiServer).HealthCheck(ctx, req.(*HealthCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Hi_ServiceDesc is the grpc.ServiceDesc for Hi service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Hi_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Hi",
	HandlerType: (*HiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HealthCheck",
			Handler:    _Hi_HealthCheck_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hi.proto",
}

const (
	CodeValueService_FindListByCodeType_FullMethodName = "/proto.CodeValueService/FindListByCodeType"
	CodeValueService_FindOneByCode_FullMethodName      = "/proto.CodeValueService/FindOneByCode"
)

// CodeValueServiceClient is the client API for CodeValueService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// code value service
type CodeValueServiceClient interface {
	FindListByCodeType(ctx context.Context, in *FindCodeValueListByCodeTypeRequest, opts ...grpc.CallOption) (*FindCodeValueListByCodeTypeResponse, error)
	FindOneByCode(ctx context.Context, in *FindOneCodeValueByCodeRequest, opts ...grpc.CallOption) (*FindOneCodeValueByCodeResponse, error)
}

type codeValueServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCodeValueServiceClient(cc grpc.ClientConnInterface) CodeValueServiceClient {
	return &codeValueServiceClient{cc}
}

func (c *codeValueServiceClient) FindListByCodeType(ctx context.Context, in *FindCodeValueListByCodeTypeRequest, opts ...grpc.CallOption) (*FindCodeValueListByCodeTypeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FindCodeValueListByCodeTypeResponse)
	err := c.cc.Invoke(ctx, CodeValueService_FindListByCodeType_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *codeValueServiceClient) FindOneByCode(ctx context.Context, in *FindOneCodeValueByCodeRequest, opts ...grpc.CallOption) (*FindOneCodeValueByCodeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FindOneCodeValueByCodeResponse)
	err := c.cc.Invoke(ctx, CodeValueService_FindOneByCode_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CodeValueServiceServer is the server API for CodeValueService service.
// All implementations must embed UnimplementedCodeValueServiceServer
// for forward compatibility.
//
// code value service
type CodeValueServiceServer interface {
	FindListByCodeType(context.Context, *FindCodeValueListByCodeTypeRequest) (*FindCodeValueListByCodeTypeResponse, error)
	FindOneByCode(context.Context, *FindOneCodeValueByCodeRequest) (*FindOneCodeValueByCodeResponse, error)
	mustEmbedUnimplementedCodeValueServiceServer()
}

// UnimplementedCodeValueServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCodeValueServiceServer struct{}

func (UnimplementedCodeValueServiceServer) FindListByCodeType(context.Context, *FindCodeValueListByCodeTypeRequest) (*FindCodeValueListByCodeTypeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindListByCodeType not implemented")
}
func (UnimplementedCodeValueServiceServer) FindOneByCode(context.Context, *FindOneCodeValueByCodeRequest) (*FindOneCodeValueByCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindOneByCode not implemented")
}
func (UnimplementedCodeValueServiceServer) mustEmbedUnimplementedCodeValueServiceServer() {}
func (UnimplementedCodeValueServiceServer) testEmbeddedByValue()                          {}

// UnsafeCodeValueServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CodeValueServiceServer will
// result in compilation errors.
type UnsafeCodeValueServiceServer interface {
	mustEmbedUnimplementedCodeValueServiceServer()
}

func RegisterCodeValueServiceServer(s grpc.ServiceRegistrar, srv CodeValueServiceServer) {
	// If the following call pancis, it indicates UnimplementedCodeValueServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&CodeValueService_ServiceDesc, srv)
}

func _CodeValueService_FindListByCodeType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindCodeValueListByCodeTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CodeValueServiceServer).FindListByCodeType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CodeValueService_FindListByCodeType_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CodeValueServiceServer).FindListByCodeType(ctx, req.(*FindCodeValueListByCodeTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CodeValueService_FindOneByCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindOneCodeValueByCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CodeValueServiceServer).FindOneByCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CodeValueService_FindOneByCode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CodeValueServiceServer).FindOneByCode(ctx, req.(*FindOneCodeValueByCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CodeValueService_ServiceDesc is the grpc.ServiceDesc for CodeValueService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CodeValueService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.CodeValueService",
	HandlerType: (*CodeValueServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindListByCodeType",
			Handler:    _CodeValueService_FindListByCodeType_Handler,
		},
		{
			MethodName: "FindOneByCode",
			Handler:    _CodeValueService_FindOneByCode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hi.proto",
}

const (
	SettingsService_FindListByAppName_FullMethodName = "/proto.SettingsService/FindListByAppName"
	SettingsService_FindByName_FullMethodName        = "/proto.SettingsService/FindByName"
)

// SettingsServiceClient is the client API for SettingsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// 定义 Settings 服务
type SettingsServiceClient interface {
	// 根据 appNameList 获取 Setting 列表
	FindListByAppName(ctx context.Context, in *FindSettingListByAppNameRequest, opts ...grpc.CallOption) (*FindSettingListByAppNameResponse, error)
	// 根据 appName 和 name 获取单个 Setting
	FindByName(ctx context.Context, in *FindSettingByNameRequest, opts ...grpc.CallOption) (*FindSettingByNameResponse, error)
}

type settingsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSettingsServiceClient(cc grpc.ClientConnInterface) SettingsServiceClient {
	return &settingsServiceClient{cc}
}

func (c *settingsServiceClient) FindListByAppName(ctx context.Context, in *FindSettingListByAppNameRequest, opts ...grpc.CallOption) (*FindSettingListByAppNameResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FindSettingListByAppNameResponse)
	err := c.cc.Invoke(ctx, SettingsService_FindListByAppName_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *settingsServiceClient) FindByName(ctx context.Context, in *FindSettingByNameRequest, opts ...grpc.CallOption) (*FindSettingByNameResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FindSettingByNameResponse)
	err := c.cc.Invoke(ctx, SettingsService_FindByName_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SettingsServiceServer is the server API for SettingsService service.
// All implementations must embed UnimplementedSettingsServiceServer
// for forward compatibility.
//
// 定义 Settings 服务
type SettingsServiceServer interface {
	// 根据 appNameList 获取 Setting 列表
	FindListByAppName(context.Context, *FindSettingListByAppNameRequest) (*FindSettingListByAppNameResponse, error)
	// 根据 appName 和 name 获取单个 Setting
	FindByName(context.Context, *FindSettingByNameRequest) (*FindSettingByNameResponse, error)
	mustEmbedUnimplementedSettingsServiceServer()
}

// UnimplementedSettingsServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedSettingsServiceServer struct{}

func (UnimplementedSettingsServiceServer) FindListByAppName(context.Context, *FindSettingListByAppNameRequest) (*FindSettingListByAppNameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindListByAppName not implemented")
}
func (UnimplementedSettingsServiceServer) FindByName(context.Context, *FindSettingByNameRequest) (*FindSettingByNameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindByName not implemented")
}
func (UnimplementedSettingsServiceServer) mustEmbedUnimplementedSettingsServiceServer() {}
func (UnimplementedSettingsServiceServer) testEmbeddedByValue()                         {}

// UnsafeSettingsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SettingsServiceServer will
// result in compilation errors.
type UnsafeSettingsServiceServer interface {
	mustEmbedUnimplementedSettingsServiceServer()
}

func RegisterSettingsServiceServer(s grpc.ServiceRegistrar, srv SettingsServiceServer) {
	// If the following call pancis, it indicates UnimplementedSettingsServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&SettingsService_ServiceDesc, srv)
}

func _SettingsService_FindListByAppName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindSettingListByAppNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingsServiceServer).FindListByAppName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SettingsService_FindListByAppName_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingsServiceServer).FindListByAppName(ctx, req.(*FindSettingListByAppNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SettingsService_FindByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindSettingByNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SettingsServiceServer).FindByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SettingsService_FindByName_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SettingsServiceServer).FindByName(ctx, req.(*FindSettingByNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SettingsService_ServiceDesc is the grpc.ServiceDesc for SettingsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SettingsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.SettingsService",
	HandlerType: (*SettingsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindListByAppName",
			Handler:    _SettingsService_FindListByAppName_Handler,
		},
		{
			MethodName: "FindByName",
			Handler:    _SettingsService_FindByName_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hi.proto",
}

const (
	AppService_EnsureAppBuiltinRoleListExist_FullMethodName = "/proto.AppService/EnsureAppBuiltinRoleListExist"
	AppService_FindAppBuiltinRoleListByApp_FullMethodName   = "/proto.AppService/FindAppBuiltinRoleListByApp"
)

// AppServiceClient is the client API for AppService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AppServiceClient interface {
	// Ensure that a list of app builtin role items exist for a specific app
	EnsureAppBuiltinRoleListExist(ctx context.Context, in *EnsureAppBuiltinRoleListExistRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// find list by app
	FindAppBuiltinRoleListByApp(ctx context.Context, in *FindAppBuiltinRoleListByAppRequest, opts ...grpc.CallOption) (*FindAppBuiltinRoleListByAppResponse, error)
}

type appServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAppServiceClient(cc grpc.ClientConnInterface) AppServiceClient {
	return &appServiceClient{cc}
}

func (c *appServiceClient) EnsureAppBuiltinRoleListExist(ctx context.Context, in *EnsureAppBuiltinRoleListExistRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, AppService_EnsureAppBuiltinRoleListExist_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appServiceClient) FindAppBuiltinRoleListByApp(ctx context.Context, in *FindAppBuiltinRoleListByAppRequest, opts ...grpc.CallOption) (*FindAppBuiltinRoleListByAppResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FindAppBuiltinRoleListByAppResponse)
	err := c.cc.Invoke(ctx, AppService_FindAppBuiltinRoleListByApp_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AppServiceServer is the server API for AppService service.
// All implementations must embed UnimplementedAppServiceServer
// for forward compatibility.
type AppServiceServer interface {
	// Ensure that a list of app builtin role items exist for a specific app
	EnsureAppBuiltinRoleListExist(context.Context, *EnsureAppBuiltinRoleListExistRequest) (*emptypb.Empty, error)
	// find list by app
	FindAppBuiltinRoleListByApp(context.Context, *FindAppBuiltinRoleListByAppRequest) (*FindAppBuiltinRoleListByAppResponse, error)
	mustEmbedUnimplementedAppServiceServer()
}

// UnimplementedAppServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAppServiceServer struct{}

func (UnimplementedAppServiceServer) EnsureAppBuiltinRoleListExist(context.Context, *EnsureAppBuiltinRoleListExistRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EnsureAppBuiltinRoleListExist not implemented")
}
func (UnimplementedAppServiceServer) FindAppBuiltinRoleListByApp(context.Context, *FindAppBuiltinRoleListByAppRequest) (*FindAppBuiltinRoleListByAppResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAppBuiltinRoleListByApp not implemented")
}
func (UnimplementedAppServiceServer) mustEmbedUnimplementedAppServiceServer() {}
func (UnimplementedAppServiceServer) testEmbeddedByValue()                    {}

// UnsafeAppServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AppServiceServer will
// result in compilation errors.
type UnsafeAppServiceServer interface {
	mustEmbedUnimplementedAppServiceServer()
}

func RegisterAppServiceServer(s grpc.ServiceRegistrar, srv AppServiceServer) {
	// If the following call pancis, it indicates UnimplementedAppServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&AppService_ServiceDesc, srv)
}

func _AppService_EnsureAppBuiltinRoleListExist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EnsureAppBuiltinRoleListExistRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServiceServer).EnsureAppBuiltinRoleListExist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AppService_EnsureAppBuiltinRoleListExist_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServiceServer).EnsureAppBuiltinRoleListExist(ctx, req.(*EnsureAppBuiltinRoleListExistRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppService_FindAppBuiltinRoleListByApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindAppBuiltinRoleListByAppRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServiceServer).FindAppBuiltinRoleListByApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AppService_FindAppBuiltinRoleListByApp_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServiceServer).FindAppBuiltinRoleListByApp(ctx, req.(*FindAppBuiltinRoleListByAppRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AppService_ServiceDesc is the grpc.ServiceDesc for AppService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AppService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.AppService",
	HandlerType: (*AppServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EnsureAppBuiltinRoleListExist",
			Handler:    _AppService_EnsureAppBuiltinRoleListExist_Handler,
		},
		{
			MethodName: "FindAppBuiltinRoleListByApp",
			Handler:    _AppService_FindAppBuiltinRoleListByApp_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hi.proto",
}

const (
	SmsService_SendSmsMessage_FullMethodName = "/proto.SmsService/SendSmsMessage"
)

// SmsServiceClient is the client API for SmsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// sms service
type SmsServiceClient interface {
	// send sms message
	SendSmsMessage(ctx context.Context, in *SendSmsMessageRequest, opts ...grpc.CallOption) (*SendSmsMessageResponse, error)
}

type smsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSmsServiceClient(cc grpc.ClientConnInterface) SmsServiceClient {
	return &smsServiceClient{cc}
}

func (c *smsServiceClient) SendSmsMessage(ctx context.Context, in *SendSmsMessageRequest, opts ...grpc.CallOption) (*SendSmsMessageResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SendSmsMessageResponse)
	err := c.cc.Invoke(ctx, SmsService_SendSmsMessage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SmsServiceServer is the server API for SmsService service.
// All implementations must embed UnimplementedSmsServiceServer
// for forward compatibility.
//
// sms service
type SmsServiceServer interface {
	// send sms message
	SendSmsMessage(context.Context, *SendSmsMessageRequest) (*SendSmsMessageResponse, error)
	mustEmbedUnimplementedSmsServiceServer()
}

// UnimplementedSmsServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedSmsServiceServer struct{}

func (UnimplementedSmsServiceServer) SendSmsMessage(context.Context, *SendSmsMessageRequest) (*SendSmsMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendSmsMessage not implemented")
}
func (UnimplementedSmsServiceServer) mustEmbedUnimplementedSmsServiceServer() {}
func (UnimplementedSmsServiceServer) testEmbeddedByValue()                    {}

// UnsafeSmsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SmsServiceServer will
// result in compilation errors.
type UnsafeSmsServiceServer interface {
	mustEmbedUnimplementedSmsServiceServer()
}

func RegisterSmsServiceServer(s grpc.ServiceRegistrar, srv SmsServiceServer) {
	// If the following call pancis, it indicates UnimplementedSmsServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&SmsService_ServiceDesc, srv)
}

func _SmsService_SendSmsMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendSmsMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SmsServiceServer).SendSmsMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SmsService_SendSmsMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SmsServiceServer).SendSmsMessage(ctx, req.(*SendSmsMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SmsService_ServiceDesc is the grpc.ServiceDesc for SmsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SmsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.SmsService",
	HandlerType: (*SmsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendSmsMessage",
			Handler:    _SmsService_SendSmsMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hi.proto",
}
