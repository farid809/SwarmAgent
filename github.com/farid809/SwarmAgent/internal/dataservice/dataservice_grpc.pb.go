// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.0
// source: internal/dataservice/dataservice.proto

package dataservice

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
	DataService_CreateOrUpdate_FullMethodName = "/dataservice.DataService/CreateOrUpdate"
	DataService_Read_FullMethodName           = "/dataservice.DataService/Read"
	DataService_Delete_FullMethodName         = "/dataservice.DataService/Delete"
)

// DataServiceClient is the client API for DataService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DataServiceClient interface {
	CreateOrUpdate(ctx context.Context, in *KeyValuePair, opts ...grpc.CallOption) (*OperationResult, error)
	Read(ctx context.Context, in *Key, opts ...grpc.CallOption) (*KeyValuePair, error)
	Delete(ctx context.Context, in *Key, opts ...grpc.CallOption) (*OperationResult, error)
}

type dataServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDataServiceClient(cc grpc.ClientConnInterface) DataServiceClient {
	return &dataServiceClient{cc}
}

func (c *dataServiceClient) CreateOrUpdate(ctx context.Context, in *KeyValuePair, opts ...grpc.CallOption) (*OperationResult, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OperationResult)
	err := c.cc.Invoke(ctx, DataService_CreateOrUpdate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataServiceClient) Read(ctx context.Context, in *Key, opts ...grpc.CallOption) (*KeyValuePair, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(KeyValuePair)
	err := c.cc.Invoke(ctx, DataService_Read_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataServiceClient) Delete(ctx context.Context, in *Key, opts ...grpc.CallOption) (*OperationResult, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OperationResult)
	err := c.cc.Invoke(ctx, DataService_Delete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DataServiceServer is the server API for DataService service.
// All implementations must embed UnimplementedDataServiceServer
// for forward compatibility.
type DataServiceServer interface {
	CreateOrUpdate(context.Context, *KeyValuePair) (*OperationResult, error)
	Read(context.Context, *Key) (*KeyValuePair, error)
	Delete(context.Context, *Key) (*OperationResult, error)
	mustEmbedUnimplementedDataServiceServer()
}

// UnimplementedDataServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedDataServiceServer struct{}

func (UnimplementedDataServiceServer) CreateOrUpdate(context.Context, *KeyValuePair) (*OperationResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrUpdate not implemented")
}
func (UnimplementedDataServiceServer) Read(context.Context, *Key) (*KeyValuePair, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Read not implemented")
}
func (UnimplementedDataServiceServer) Delete(context.Context, *Key) (*OperationResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedDataServiceServer) mustEmbedUnimplementedDataServiceServer() {}
func (UnimplementedDataServiceServer) testEmbeddedByValue()                     {}

// UnsafeDataServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DataServiceServer will
// result in compilation errors.
type UnsafeDataServiceServer interface {
	mustEmbedUnimplementedDataServiceServer()
}

func RegisterDataServiceServer(s grpc.ServiceRegistrar, srv DataServiceServer) {
	// If the following call pancis, it indicates UnimplementedDataServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&DataService_ServiceDesc, srv)
}

func _DataService_CreateOrUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyValuePair)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataServiceServer).CreateOrUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DataService_CreateOrUpdate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataServiceServer).CreateOrUpdate(ctx, req.(*KeyValuePair))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataService_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Key)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataServiceServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DataService_Read_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataServiceServer).Read(ctx, req.(*Key))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Key)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DataService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataServiceServer).Delete(ctx, req.(*Key))
	}
	return interceptor(ctx, in, info, handler)
}

// DataService_ServiceDesc is the grpc.ServiceDesc for DataService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DataService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dataservice.DataService",
	HandlerType: (*DataServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateOrUpdate",
			Handler:    _DataService_CreateOrUpdate_Handler,
		},
		{
			MethodName: "Read",
			Handler:    _DataService_Read_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _DataService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/dataservice/dataservice.proto",
}