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

// MetricsClient is the client API for Metrics service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MetricsClient interface {
	Rps(ctx context.Context, opts ...grpc.CallOption) (Metrics_RpsClient, error)
	ResponseTime(ctx context.Context, opts ...grpc.CallOption) (Metrics_ResponseTimeClient, error)
	AddRegister(ctx context.Context, in *AddRegistrationRequest, opts ...grpc.CallOption) (*AddRegistrationResponse, error)
	FilesOnUser(ctx context.Context, opts ...grpc.CallOption) (Metrics_FilesOnUserClient, error)
	ActiveUsers(ctx context.Context, opts ...grpc.CallOption) (Metrics_ActiveUsersClient, error)
	AddInternal(ctx context.Context, in *AddInternalRequest, opts ...grpc.CallOption) (*AddInternalResponse, error)
}

type metricsClient struct {
	cc grpc.ClientConnInterface
}

func NewMetricsClient(cc grpc.ClientConnInterface) MetricsClient {
	return &metricsClient{cc}
}

func (c *metricsClient) Rps(ctx context.Context, opts ...grpc.CallOption) (Metrics_RpsClient, error) {
	stream, err := c.cc.NewStream(ctx, &Metrics_ServiceDesc.Streams[0], "/metrics.Metrics/Rps", opts...)
	if err != nil {
		return nil, err
	}
	x := &metricsRpsClient{stream}
	return x, nil
}

type Metrics_RpsClient interface {
	Send(*RpsRequest) error
	CloseAndRecv() (*RpsResponse, error)
	grpc.ClientStream
}

type metricsRpsClient struct {
	grpc.ClientStream
}

func (x *metricsRpsClient) Send(m *RpsRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *metricsRpsClient) CloseAndRecv() (*RpsResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(RpsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *metricsClient) ResponseTime(ctx context.Context, opts ...grpc.CallOption) (Metrics_ResponseTimeClient, error) {
	stream, err := c.cc.NewStream(ctx, &Metrics_ServiceDesc.Streams[1], "/metrics.Metrics/ResponseTime", opts...)
	if err != nil {
		return nil, err
	}
	x := &metricsResponseTimeClient{stream}
	return x, nil
}

type Metrics_ResponseTimeClient interface {
	Send(*ResponseTimeRequest) error
	Recv() (*ResponseTimeResponse, error)
	grpc.ClientStream
}

type metricsResponseTimeClient struct {
	grpc.ClientStream
}

func (x *metricsResponseTimeClient) Send(m *ResponseTimeRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *metricsResponseTimeClient) Recv() (*ResponseTimeResponse, error) {
	m := new(ResponseTimeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *metricsClient) AddRegister(ctx context.Context, in *AddRegistrationRequest, opts ...grpc.CallOption) (*AddRegistrationResponse, error) {
	out := new(AddRegistrationResponse)
	err := c.cc.Invoke(ctx, "/metrics.Metrics/AddRegister", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metricsClient) FilesOnUser(ctx context.Context, opts ...grpc.CallOption) (Metrics_FilesOnUserClient, error) {
	stream, err := c.cc.NewStream(ctx, &Metrics_ServiceDesc.Streams[2], "/metrics.Metrics/FilesOnUser", opts...)
	if err != nil {
		return nil, err
	}
	x := &metricsFilesOnUserClient{stream}
	return x, nil
}

type Metrics_FilesOnUserClient interface {
	Send(*FilesOnUserRequest) error
	Recv() (*FilesOnUserResponse, error)
	grpc.ClientStream
}

type metricsFilesOnUserClient struct {
	grpc.ClientStream
}

func (x *metricsFilesOnUserClient) Send(m *FilesOnUserRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *metricsFilesOnUserClient) Recv() (*FilesOnUserResponse, error) {
	m := new(FilesOnUserResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *metricsClient) ActiveUsers(ctx context.Context, opts ...grpc.CallOption) (Metrics_ActiveUsersClient, error) {
	stream, err := c.cc.NewStream(ctx, &Metrics_ServiceDesc.Streams[3], "/metrics.Metrics/ActiveUsers", opts...)
	if err != nil {
		return nil, err
	}
	x := &metricsActiveUsersClient{stream}
	return x, nil
}

type Metrics_ActiveUsersClient interface {
	Send(*ActiveUsersRequest) error
	Recv() (*ActiveUsersResponse, error)
	grpc.ClientStream
}

type metricsActiveUsersClient struct {
	grpc.ClientStream
}

func (x *metricsActiveUsersClient) Send(m *ActiveUsersRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *metricsActiveUsersClient) Recv() (*ActiveUsersResponse, error) {
	m := new(ActiveUsersResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *metricsClient) AddInternal(ctx context.Context, in *AddInternalRequest, opts ...grpc.CallOption) (*AddInternalResponse, error) {
	out := new(AddInternalResponse)
	err := c.cc.Invoke(ctx, "/metrics.Metrics/AddInternal", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MetricsServer is the server API for Metrics service.
// All implementations must embed UnimplementedMetricsServer
// for forward compatibility
type MetricsServer interface {
	Rps(Metrics_RpsServer) error
	ResponseTime(Metrics_ResponseTimeServer) error
	AddRegister(context.Context, *AddRegistrationRequest) (*AddRegistrationResponse, error)
	FilesOnUser(Metrics_FilesOnUserServer) error
	ActiveUsers(Metrics_ActiveUsersServer) error
	AddInternal(context.Context, *AddInternalRequest) (*AddInternalResponse, error)
	mustEmbedUnimplementedMetricsServer()
}

// UnimplementedMetricsServer must be embedded to have forward compatible implementations.
type UnimplementedMetricsServer struct {
}

func (UnimplementedMetricsServer) Rps(Metrics_RpsServer) error {
	return status.Errorf(codes.Unimplemented, "method Rps not implemented")
}
func (UnimplementedMetricsServer) ResponseTime(Metrics_ResponseTimeServer) error {
	return status.Errorf(codes.Unimplemented, "method ResponseTime not implemented")
}
func (UnimplementedMetricsServer) AddRegister(context.Context, *AddRegistrationRequest) (*AddRegistrationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddRegister not implemented")
}
func (UnimplementedMetricsServer) FilesOnUser(Metrics_FilesOnUserServer) error {
	return status.Errorf(codes.Unimplemented, "method FilesOnUser not implemented")
}
func (UnimplementedMetricsServer) ActiveUsers(Metrics_ActiveUsersServer) error {
	return status.Errorf(codes.Unimplemented, "method ActiveUsers not implemented")
}
func (UnimplementedMetricsServer) AddInternal(context.Context, *AddInternalRequest) (*AddInternalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddInternal not implemented")
}
func (UnimplementedMetricsServer) mustEmbedUnimplementedMetricsServer() {}

// UnsafeMetricsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MetricsServer will
// result in compilation errors.
type UnsafeMetricsServer interface {
	mustEmbedUnimplementedMetricsServer()
}

func RegisterMetricsServer(s grpc.ServiceRegistrar, srv MetricsServer) {
	s.RegisterService(&Metrics_ServiceDesc, srv)
}

func _Metrics_Rps_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MetricsServer).Rps(&metricsRpsServer{stream})
}

type Metrics_RpsServer interface {
	SendAndClose(*RpsResponse) error
	Recv() (*RpsRequest, error)
	grpc.ServerStream
}

type metricsRpsServer struct {
	grpc.ServerStream
}

func (x *metricsRpsServer) SendAndClose(m *RpsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *metricsRpsServer) Recv() (*RpsRequest, error) {
	m := new(RpsRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Metrics_ResponseTime_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MetricsServer).ResponseTime(&metricsResponseTimeServer{stream})
}

type Metrics_ResponseTimeServer interface {
	Send(*ResponseTimeResponse) error
	Recv() (*ResponseTimeRequest, error)
	grpc.ServerStream
}

type metricsResponseTimeServer struct {
	grpc.ServerStream
}

func (x *metricsResponseTimeServer) Send(m *ResponseTimeResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *metricsResponseTimeServer) Recv() (*ResponseTimeRequest, error) {
	m := new(ResponseTimeRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Metrics_AddRegister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRegistrationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricsServer).AddRegister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/metrics.Metrics/AddRegister",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricsServer).AddRegister(ctx, req.(*AddRegistrationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Metrics_FilesOnUser_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MetricsServer).FilesOnUser(&metricsFilesOnUserServer{stream})
}

type Metrics_FilesOnUserServer interface {
	Send(*FilesOnUserResponse) error
	Recv() (*FilesOnUserRequest, error)
	grpc.ServerStream
}

type metricsFilesOnUserServer struct {
	grpc.ServerStream
}

func (x *metricsFilesOnUserServer) Send(m *FilesOnUserResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *metricsFilesOnUserServer) Recv() (*FilesOnUserRequest, error) {
	m := new(FilesOnUserRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Metrics_ActiveUsers_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MetricsServer).ActiveUsers(&metricsActiveUsersServer{stream})
}

type Metrics_ActiveUsersServer interface {
	Send(*ActiveUsersResponse) error
	Recv() (*ActiveUsersRequest, error)
	grpc.ServerStream
}

type metricsActiveUsersServer struct {
	grpc.ServerStream
}

func (x *metricsActiveUsersServer) Send(m *ActiveUsersResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *metricsActiveUsersServer) Recv() (*ActiveUsersRequest, error) {
	m := new(ActiveUsersRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Metrics_AddInternal_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddInternalRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricsServer).AddInternal(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/metrics.Metrics/AddInternal",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricsServer).AddInternal(ctx, req.(*AddInternalRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Metrics_ServiceDesc is the grpc.ServiceDesc for Metrics service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Metrics_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "metrics.Metrics",
	HandlerType: (*MetricsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddRegister",
			Handler:    _Metrics_AddRegister_Handler,
		},
		{
			MethodName: "AddInternal",
			Handler:    _Metrics_AddInternal_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Rps",
			Handler:       _Metrics_Rps_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "ResponseTime",
			Handler:       _Metrics_ResponseTime_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "FilesOnUser",
			Handler:       _Metrics_FilesOnUser_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "ActiveUsers",
			Handler:       _Metrics_ActiveUsers_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "metrics/metrics.proto",
}