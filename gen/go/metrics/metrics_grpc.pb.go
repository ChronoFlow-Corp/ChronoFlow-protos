// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.2
// source: metrics/metrics.proto

package metricsv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// MetricsServiceClient is the client API for MetricsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MetricsServiceClient interface {
	Rps(ctx context.Context, opts ...grpc.CallOption) (MetricsService_RpsClient, error)
	ResponseTime(ctx context.Context, opts ...grpc.CallOption) (MetricsService_ResponseTimeClient, error)
	AddRegister(ctx context.Context, in *AddRegistrationRequest, opts ...grpc.CallOption) (*AddRegistrationResponse, error)
	FilesOnUser(ctx context.Context, opts ...grpc.CallOption) (MetricsService_FilesOnUserClient, error)
	ActiveUsers(ctx context.Context, opts ...grpc.CallOption) (MetricsService_ActiveUsersClient, error)
	AddInternal(ctx context.Context, in *AddInternalRequest, opts ...grpc.CallOption) (*AddInternalResponse, error)
}

type metricsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMetricsServiceClient(cc grpc.ClientConnInterface) MetricsServiceClient {
	return &metricsServiceClient{cc}
}

func (c *metricsServiceClient) Rps(ctx context.Context, opts ...grpc.CallOption) (MetricsService_RpsClient, error) {
	stream, err := c.cc.NewStream(ctx, &MetricsService_ServiceDesc.Streams[0], "/metrics.MetricsService/Rps", opts...)
	if err != nil {
		return nil, err
	}
	x := &metricsServiceRpsClient{stream}
	return x, nil
}

type MetricsService_RpsClient interface {
	Send(*RpsRequest) error
	CloseAndRecv() (*RpsResponse, error)
	grpc.ClientStream
}

type metricsServiceRpsClient struct {
	grpc.ClientStream
}

func (x *metricsServiceRpsClient) Send(m *RpsRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *metricsServiceRpsClient) CloseAndRecv() (*RpsResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(RpsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *metricsServiceClient) ResponseTime(ctx context.Context, opts ...grpc.CallOption) (MetricsService_ResponseTimeClient, error) {
	stream, err := c.cc.NewStream(ctx, &MetricsService_ServiceDesc.Streams[1], "/metrics.MetricsService/ResponseTime", opts...)
	if err != nil {
		return nil, err
	}
	x := &metricsServiceResponseTimeClient{stream}
	return x, nil
}

type MetricsService_ResponseTimeClient interface {
	Send(*ResponseTimeRequest) error
	Recv() (*ResponseTimeResponse, error)
	grpc.ClientStream
}

type metricsServiceResponseTimeClient struct {
	grpc.ClientStream
}

func (x *metricsServiceResponseTimeClient) Send(m *ResponseTimeRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *metricsServiceResponseTimeClient) Recv() (*ResponseTimeResponse, error) {
	m := new(ResponseTimeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *metricsServiceClient) AddRegister(ctx context.Context, in *AddRegistrationRequest, opts ...grpc.CallOption) (*AddRegistrationResponse, error) {
	out := new(AddRegistrationResponse)
	err := c.cc.Invoke(ctx, "/metrics.MetricsService/AddRegister", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metricsServiceClient) FilesOnUser(ctx context.Context, opts ...grpc.CallOption) (MetricsService_FilesOnUserClient, error) {
	stream, err := c.cc.NewStream(ctx, &MetricsService_ServiceDesc.Streams[2], "/metrics.MetricsService/FilesOnUser", opts...)
	if err != nil {
		return nil, err
	}
	x := &metricsServiceFilesOnUserClient{stream}
	return x, nil
}

type MetricsService_FilesOnUserClient interface {
	Send(*FilesOnUserRequest) error
	Recv() (*FilesOnUserResponse, error)
	grpc.ClientStream
}

type metricsServiceFilesOnUserClient struct {
	grpc.ClientStream
}

func (x *metricsServiceFilesOnUserClient) Send(m *FilesOnUserRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *metricsServiceFilesOnUserClient) Recv() (*FilesOnUserResponse, error) {
	m := new(FilesOnUserResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *metricsServiceClient) ActiveUsers(ctx context.Context, opts ...grpc.CallOption) (MetricsService_ActiveUsersClient, error) {
	stream, err := c.cc.NewStream(ctx, &MetricsService_ServiceDesc.Streams[3], "/metrics.MetricsService/ActiveUsers", opts...)
	if err != nil {
		return nil, err
	}
	x := &metricsServiceActiveUsersClient{stream}
	return x, nil
}

type MetricsService_ActiveUsersClient interface {
	Send(*ActiveUsersRequest) error
	Recv() (*ActiveUsersResponse, error)
	grpc.ClientStream
}

type metricsServiceActiveUsersClient struct {
	grpc.ClientStream
}

func (x *metricsServiceActiveUsersClient) Send(m *ActiveUsersRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *metricsServiceActiveUsersClient) Recv() (*ActiveUsersResponse, error) {
	m := new(ActiveUsersResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *metricsServiceClient) AddInternal(ctx context.Context, in *AddInternalRequest, opts ...grpc.CallOption) (*AddInternalResponse, error) {
	out := new(AddInternalResponse)
	err := c.cc.Invoke(ctx, "/metrics.MetricsService/AddInternal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MetricsServiceServer is the server API for MetricsService service.
// All implementations must embed UnimplementedMetricsServiceServer
// for forward compatibility
type MetricsServiceServer interface {
	Rps(MetricsService_RpsServer) error
	ResponseTime(MetricsService_ResponseTimeServer) error
	AddRegister(context.Context, *AddRegistrationRequest) (*AddRegistrationResponse, error)
	FilesOnUser(MetricsService_FilesOnUserServer) error
	ActiveUsers(MetricsService_ActiveUsersServer) error
	AddInternal(context.Context, *AddInternalRequest) (*AddInternalResponse, error)
	mustEmbedUnimplementedMetricsServiceServer()
}

// UnimplementedMetricsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMetricsServiceServer struct {
}

func (UnimplementedMetricsServiceServer) Rps(MetricsService_RpsServer) error {
	return status.Errorf(codes.Unimplemented, "method Rps not implemented")
}
func (UnimplementedMetricsServiceServer) ResponseTime(MetricsService_ResponseTimeServer) error {
	return status.Errorf(codes.Unimplemented, "method ResponseTime not implemented")
}
func (UnimplementedMetricsServiceServer) AddRegister(context.Context, *AddRegistrationRequest) (*AddRegistrationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddRegister not implemented")
}
func (UnimplementedMetricsServiceServer) FilesOnUser(MetricsService_FilesOnUserServer) error {
	return status.Errorf(codes.Unimplemented, "method FilesOnUser not implemented")
}
func (UnimplementedMetricsServiceServer) ActiveUsers(MetricsService_ActiveUsersServer) error {
	return status.Errorf(codes.Unimplemented, "method ActiveUsers not implemented")
}
func (UnimplementedMetricsServiceServer) AddInternal(context.Context, *AddInternalRequest) (*AddInternalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddInternal not implemented")
}
func (UnimplementedMetricsServiceServer) mustEmbedUnimplementedMetricsServiceServer() {}

// UnsafeMetricsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MetricsServiceServer will
// result in compilation errors.
type UnsafeMetricsServiceServer interface {
	mustEmbedUnimplementedMetricsServiceServer()
}

func RegisterMetricsServiceServer(s grpc.ServiceRegistrar, srv MetricsServiceServer) {
	s.RegisterService(&MetricsService_ServiceDesc, srv)
}

func _MetricsService_Rps_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MetricsServiceServer).Rps(&metricsServiceRpsServer{stream})
}

type MetricsService_RpsServer interface {
	SendAndClose(*RpsResponse) error
	Recv() (*RpsRequest, error)
	grpc.ServerStream
}

type metricsServiceRpsServer struct {
	grpc.ServerStream
}

func (x *metricsServiceRpsServer) SendAndClose(m *RpsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *metricsServiceRpsServer) Recv() (*RpsRequest, error) {
	m := new(RpsRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _MetricsService_ResponseTime_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MetricsServiceServer).ResponseTime(&metricsServiceResponseTimeServer{stream})
}

type MetricsService_ResponseTimeServer interface {
	Send(*ResponseTimeResponse) error
	Recv() (*ResponseTimeRequest, error)
	grpc.ServerStream
}

type metricsServiceResponseTimeServer struct {
	grpc.ServerStream
}

func (x *metricsServiceResponseTimeServer) Send(m *ResponseTimeResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *metricsServiceResponseTimeServer) Recv() (*ResponseTimeRequest, error) {
	m := new(ResponseTimeRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _MetricsService_AddRegister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRegistrationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricsServiceServer).AddRegister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/metrics.MetricsService/AddRegister",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricsServiceServer).AddRegister(ctx, req.(*AddRegistrationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetricsService_FilesOnUser_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MetricsServiceServer).FilesOnUser(&metricsServiceFilesOnUserServer{stream})
}

type MetricsService_FilesOnUserServer interface {
	Send(*FilesOnUserResponse) error
	Recv() (*FilesOnUserRequest, error)
	grpc.ServerStream
}

type metricsServiceFilesOnUserServer struct {
	grpc.ServerStream
}

func (x *metricsServiceFilesOnUserServer) Send(m *FilesOnUserResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *metricsServiceFilesOnUserServer) Recv() (*FilesOnUserRequest, error) {
	m := new(FilesOnUserRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _MetricsService_ActiveUsers_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MetricsServiceServer).ActiveUsers(&metricsServiceActiveUsersServer{stream})
}

type MetricsService_ActiveUsersServer interface {
	Send(*ActiveUsersResponse) error
	Recv() (*ActiveUsersRequest, error)
	grpc.ServerStream
}

type metricsServiceActiveUsersServer struct {
	grpc.ServerStream
}

func (x *metricsServiceActiveUsersServer) Send(m *ActiveUsersResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *metricsServiceActiveUsersServer) Recv() (*ActiveUsersRequest, error) {
	m := new(ActiveUsersRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _MetricsService_AddInternal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddInternalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricsServiceServer).AddInternal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/metrics.MetricsService/AddInternal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricsServiceServer).AddInternal(ctx, req.(*AddInternalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MetricsService_ServiceDesc is the grpc.ServiceDesc for MetricsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MetricsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "metrics.MetricsService",
	HandlerType: (*MetricsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddRegister",
			Handler:    _MetricsService_AddRegister_Handler,
		},
		{
			MethodName: "AddInternal",
			Handler:    _MetricsService_AddInternal_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Rps",
			Handler:       _MetricsService_Rps_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "ResponseTime",
			Handler:       _MetricsService_ResponseTime_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "FilesOnUser",
			Handler:       _MetricsService_FilesOnUser_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "ActiveUsers",
			Handler:       _MetricsService_ActiveUsers_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "metrics/metrics.proto",
}
