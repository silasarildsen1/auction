// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.7
// source: grpc/interface.proto

package skrr

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

// AuctionClient is the client API for Auction service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuctionClient interface {
	Result(ctx context.Context, in *Void, opts ...grpc.CallOption) (*Outcome, error)
	Bid(ctx context.Context, in *BidMessage, opts ...grpc.CallOption) (*Ack, error)
	BackUp(ctx context.Context, in *BackUpMessage, opts ...grpc.CallOption) (*Ack, error)
}

type auctionClient struct {
	cc grpc.ClientConnInterface
}

func NewAuctionClient(cc grpc.ClientConnInterface) AuctionClient {
	return &auctionClient{cc}
}

func (c *auctionClient) Result(ctx context.Context, in *Void, opts ...grpc.CallOption) (*Outcome, error) {
	out := new(Outcome)
	err := c.cc.Invoke(ctx, "/auction.Auction/Result", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *auctionClient) Bid(ctx context.Context, in *BidMessage, opts ...grpc.CallOption) (*Ack, error) {
	out := new(Ack)
	err := c.cc.Invoke(ctx, "/auction.Auction/Bid", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *auctionClient) BackUp(ctx context.Context, in *BackUpMessage, opts ...grpc.CallOption) (*Ack, error) {
	out := new(Ack)
	err := c.cc.Invoke(ctx, "/auction.Auction/BackUp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuctionServer is the server API for Auction service.
// All implementations must embed UnimplementedAuctionServer
// for forward compatibility
type AuctionServer interface {
	Result(context.Context, *Void) (*Outcome, error)
	Bid(context.Context, *BidMessage) (*Ack, error)
	BackUp(context.Context, *BackUpMessage) (*Ack, error)
	mustEmbedUnimplementedAuctionServer()
}

// UnimplementedAuctionServer must be embedded to have forward compatible implementations.
type UnimplementedAuctionServer struct {
}

func (UnimplementedAuctionServer) Result(context.Context, *Void) (*Outcome, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Result not implemented")
}
func (UnimplementedAuctionServer) Bid(context.Context, *BidMessage) (*Ack, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Bid not implemented")
}
func (UnimplementedAuctionServer) BackUp(context.Context, *BackUpMessage) (*Ack, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BackUp not implemented")
}
func (UnimplementedAuctionServer) mustEmbedUnimplementedAuctionServer() {}

// UnsafeAuctionServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuctionServer will
// result in compilation errors.
type UnsafeAuctionServer interface {
	mustEmbedUnimplementedAuctionServer()
}

func RegisterAuctionServer(s grpc.ServiceRegistrar, srv AuctionServer) {
	s.RegisterService(&Auction_ServiceDesc, srv)
}

func _Auction_Result_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuctionServer).Result(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auction.Auction/Result",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuctionServer).Result(ctx, req.(*Void))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auction_Bid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BidMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuctionServer).Bid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auction.Auction/Bid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuctionServer).Bid(ctx, req.(*BidMessage))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auction_BackUp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BackUpMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuctionServer).BackUp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auction.Auction/BackUp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuctionServer).BackUp(ctx, req.(*BackUpMessage))
	}
	return interceptor(ctx, in, info, handler)
}

// Auction_ServiceDesc is the grpc.ServiceDesc for Auction service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Auction_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auction.Auction",
	HandlerType: (*AuctionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Result",
			Handler:    _Auction_Result_Handler,
		},
		{
			MethodName: "Bid",
			Handler:    _Auction_Bid_Handler,
		},
		{
			MethodName: "BackUp",
			Handler:    _Auction_BackUp_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc/interface.proto",
}
