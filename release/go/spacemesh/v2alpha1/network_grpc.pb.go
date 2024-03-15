// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: spacemesh/v2alpha1/network.proto

package spacemeshv2alpha1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	NetworkService_Info_FullMethodName = "/spacemesh.v2alpha1.NetworkService/Info"
)

// NetworkServiceClient is the client API for NetworkService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NetworkServiceClient interface {
	Info(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*NetworkInfoResponse, error)
}

type networkServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNetworkServiceClient(cc grpc.ClientConnInterface) NetworkServiceClient {
	return &networkServiceClient{cc}
}

func (c *networkServiceClient) Info(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*NetworkInfoResponse, error) {
	out := new(NetworkInfoResponse)
	err := c.cc.Invoke(ctx, NetworkService_Info_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NetworkServiceServer is the server API for NetworkService service.
// All implementations should embed UnimplementedNetworkServiceServer
// for forward compatibility
type NetworkServiceServer interface {
	Info(context.Context, *emptypb.Empty) (*NetworkInfoResponse, error)
}

// UnimplementedNetworkServiceServer should be embedded to have forward compatible implementations.
type UnimplementedNetworkServiceServer struct {
}

func (UnimplementedNetworkServiceServer) Info(context.Context, *emptypb.Empty) (*NetworkInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Info not implemented")
}

// UnsafeNetworkServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NetworkServiceServer will
// result in compilation errors.
type UnsafeNetworkServiceServer interface {
	mustEmbedUnimplementedNetworkServiceServer()
}

func RegisterNetworkServiceServer(s grpc.ServiceRegistrar, srv NetworkServiceServer) {
	s.RegisterService(&NetworkService_ServiceDesc, srv)
}

func _NetworkService_Info_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServiceServer).Info(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NetworkService_Info_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServiceServer).Info(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// NetworkService_ServiceDesc is the grpc.ServiceDesc for NetworkService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NetworkService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "spacemesh.v2alpha1.NetworkService",
	HandlerType: (*NetworkServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Info",
			Handler:    _NetworkService_Info_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "spacemesh/v2alpha1/network.proto",
}
