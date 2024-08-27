// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: app/travel/cmd/rpc/pb/travel.proto

package pb

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

const (
	Travel_HomestayDetail_FullMethodName = "/pb.travel/homestayDetail"
)

// TravelClient is the client API for Travel service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TravelClient interface {
	// homestayDetail
	HomestayDetail(ctx context.Context, in *HomestayDetailReq, opts ...grpc.CallOption) (*HomestayDetailResp, error)
}

type travelClient struct {
	cc grpc.ClientConnInterface
}

func NewTravelClient(cc grpc.ClientConnInterface) TravelClient {
	return &travelClient{cc}
}

func (c *travelClient) HomestayDetail(ctx context.Context, in *HomestayDetailReq, opts ...grpc.CallOption) (*HomestayDetailResp, error) {
	out := new(HomestayDetailResp)
	err := c.cc.Invoke(ctx, Travel_HomestayDetail_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TravelServer is the server API for Travel service.
// All implementations must embed UnimplementedTravelServer
// for forward compatibility
type TravelServer interface {
	// homestayDetail
	HomestayDetail(context.Context, *HomestayDetailReq) (*HomestayDetailResp, error)
	mustEmbedUnimplementedTravelServer()
}

// UnimplementedTravelServer must be embedded to have forward compatible implementations.
type UnimplementedTravelServer struct {
}

func (UnimplementedTravelServer) HomestayDetail(context.Context, *HomestayDetailReq) (*HomestayDetailResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HomestayDetail not implemented")
}
func (UnimplementedTravelServer) mustEmbedUnimplementedTravelServer() {}

// UnsafeTravelServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TravelServer will
// result in compilation errors.
type UnsafeTravelServer interface {
	mustEmbedUnimplementedTravelServer()
}

func RegisterTravelServer(s grpc.ServiceRegistrar, srv TravelServer) {
	s.RegisterService(&Travel_ServiceDesc, srv)
}

func _Travel_HomestayDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HomestayDetailReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TravelServer).HomestayDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Travel_HomestayDetail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TravelServer).HomestayDetail(ctx, req.(*HomestayDetailReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Travel_ServiceDesc is the grpc.ServiceDesc for Travel service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Travel_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.travel",
	HandlerType: (*TravelServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "homestayDetail",
			Handler:    _Travel_HomestayDetail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "app/travel/cmd/rpc/pb/travel.proto",
}
