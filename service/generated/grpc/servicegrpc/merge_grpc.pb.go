// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.1
// source: servicegrpc/merge.proto

package servicegrpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	MergeBlobService_SendMergeBlob_FullMethodName = "/app.merge_blob.MergeBlobService/SendMergeBlob"
)

// MergeBlobServiceClient is the client API for MergeBlobService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MergeBlobServiceClient interface {
	SendMergeBlob(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[SendMergeBlobRequest, SendMergeBlobResponse], error)
}

type mergeBlobServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMergeBlobServiceClient(cc grpc.ClientConnInterface) MergeBlobServiceClient {
	return &mergeBlobServiceClient{cc}
}

func (c *mergeBlobServiceClient) SendMergeBlob(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[SendMergeBlobRequest, SendMergeBlobResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &MergeBlobService_ServiceDesc.Streams[0], MergeBlobService_SendMergeBlob_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[SendMergeBlobRequest, SendMergeBlobResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type MergeBlobService_SendMergeBlobClient = grpc.ClientStreamingClient[SendMergeBlobRequest, SendMergeBlobResponse]

// MergeBlobServiceServer is the server API for MergeBlobService service.
// All implementations must embed UnimplementedMergeBlobServiceServer
// for forward compatibility.
type MergeBlobServiceServer interface {
	SendMergeBlob(grpc.ClientStreamingServer[SendMergeBlobRequest, SendMergeBlobResponse]) error
	mustEmbedUnimplementedMergeBlobServiceServer()
}

// UnimplementedMergeBlobServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedMergeBlobServiceServer struct{}

func (UnimplementedMergeBlobServiceServer) SendMergeBlob(grpc.ClientStreamingServer[SendMergeBlobRequest, SendMergeBlobResponse]) error {
	return status.Errorf(codes.Unimplemented, "method SendMergeBlob not implemented")
}
func (UnimplementedMergeBlobServiceServer) mustEmbedUnimplementedMergeBlobServiceServer() {}
func (UnimplementedMergeBlobServiceServer) testEmbeddedByValue()                          {}

// UnsafeMergeBlobServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MergeBlobServiceServer will
// result in compilation errors.
type UnsafeMergeBlobServiceServer interface {
	mustEmbedUnimplementedMergeBlobServiceServer()
}

func RegisterMergeBlobServiceServer(s grpc.ServiceRegistrar, srv MergeBlobServiceServer) {
	// If the following call pancis, it indicates UnimplementedMergeBlobServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&MergeBlobService_ServiceDesc, srv)
}

func _MergeBlobService_SendMergeBlob_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MergeBlobServiceServer).SendMergeBlob(&grpc.GenericServerStream[SendMergeBlobRequest, SendMergeBlobResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type MergeBlobService_SendMergeBlobServer = grpc.ClientStreamingServer[SendMergeBlobRequest, SendMergeBlobResponse]

// MergeBlobService_ServiceDesc is the grpc.ServiceDesc for MergeBlobService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MergeBlobService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "app.merge_blob.MergeBlobService",
	HandlerType: (*MergeBlobServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SendMergeBlob",
			Handler:       _MergeBlobService_SendMergeBlob_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "servicegrpc/merge.proto",
}