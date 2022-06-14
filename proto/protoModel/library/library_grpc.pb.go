// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package library

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// LibraryServiceClient is the client API for LibraryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LibraryServiceClient interface {
	GetBooks(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetBooksResponse, error)
	FindBook(ctx context.Context, in *FindBookRequest, opts ...grpc.CallOption) (*Book, error)
	AddBook(ctx context.Context, in *AddBookRequest, opts ...grpc.CallOption) (*Book, error)
	UpdateBook(ctx context.Context, in *UpdateBookRequest, opts ...grpc.CallOption) (*Book, error)
	DeleteBook(ctx context.Context, in *DeleteBookRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	ReserveBook(ctx context.Context, in *ReserveBookRequest, opts ...grpc.CallOption) (*Book, error)
}

type libraryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLibraryServiceClient(cc grpc.ClientConnInterface) LibraryServiceClient {
	return &libraryServiceClient{cc}
}

func (c *libraryServiceClient) GetBooks(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetBooksResponse, error) {
	out := new(GetBooksResponse)
	err := c.cc.Invoke(ctx, "/LibraryService/GetBooks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *libraryServiceClient) FindBook(ctx context.Context, in *FindBookRequest, opts ...grpc.CallOption) (*Book, error) {
	out := new(Book)
	err := c.cc.Invoke(ctx, "/LibraryService/FindBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *libraryServiceClient) AddBook(ctx context.Context, in *AddBookRequest, opts ...grpc.CallOption) (*Book, error) {
	out := new(Book)
	err := c.cc.Invoke(ctx, "/LibraryService/AddBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *libraryServiceClient) UpdateBook(ctx context.Context, in *UpdateBookRequest, opts ...grpc.CallOption) (*Book, error) {
	out := new(Book)
	err := c.cc.Invoke(ctx, "/LibraryService/UpdateBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *libraryServiceClient) DeleteBook(ctx context.Context, in *DeleteBookRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/LibraryService/DeleteBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *libraryServiceClient) ReserveBook(ctx context.Context, in *ReserveBookRequest, opts ...grpc.CallOption) (*Book, error) {
	out := new(Book)
	err := c.cc.Invoke(ctx, "/LibraryService/ReserveBook", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LibraryServiceServer is the server API for LibraryService service.
// All implementations must embed UnimplementedLibraryServiceServer
// for forward compatibility
type LibraryServiceServer interface {
	GetBooks(context.Context, *empty.Empty) (*GetBooksResponse, error)
	FindBook(context.Context, *FindBookRequest) (*Book, error)
	AddBook(context.Context, *AddBookRequest) (*Book, error)
	UpdateBook(context.Context, *UpdateBookRequest) (*Book, error)
	DeleteBook(context.Context, *DeleteBookRequest) (*empty.Empty, error)
	ReserveBook(context.Context, *ReserveBookRequest) (*Book, error)
	mustEmbedUnimplementedLibraryServiceServer()
}

// UnimplementedLibraryServiceServer must be embedded to have forward compatible implementations.
type UnimplementedLibraryServiceServer struct {
}

func (UnimplementedLibraryServiceServer) GetBooks(context.Context, *empty.Empty) (*GetBooksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBooks not implemented")
}
func (UnimplementedLibraryServiceServer) FindBook(context.Context, *FindBookRequest) (*Book, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindBook not implemented")
}
func (UnimplementedLibraryServiceServer) AddBook(context.Context, *AddBookRequest) (*Book, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddBook not implemented")
}
func (UnimplementedLibraryServiceServer) UpdateBook(context.Context, *UpdateBookRequest) (*Book, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBook not implemented")
}
func (UnimplementedLibraryServiceServer) DeleteBook(context.Context, *DeleteBookRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBook not implemented")
}
func (UnimplementedLibraryServiceServer) ReserveBook(context.Context, *ReserveBookRequest) (*Book, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReserveBook not implemented")
}
func (UnimplementedLibraryServiceServer) mustEmbedUnimplementedLibraryServiceServer() {}

// UnsafeLibraryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LibraryServiceServer will
// result in compilation errors.
type UnsafeLibraryServiceServer interface {
	mustEmbedUnimplementedLibraryServiceServer()
}

func RegisterLibraryServiceServer(s grpc.ServiceRegistrar, srv LibraryServiceServer) {
	s.RegisterService(&LibraryService_ServiceDesc, srv)
}

func _LibraryService_GetBooks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibraryServiceServer).GetBooks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/LibraryService/GetBooks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibraryServiceServer).GetBooks(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _LibraryService_FindBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibraryServiceServer).FindBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/LibraryService/FindBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibraryServiceServer).FindBook(ctx, req.(*FindBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LibraryService_AddBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibraryServiceServer).AddBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/LibraryService/AddBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibraryServiceServer).AddBook(ctx, req.(*AddBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LibraryService_UpdateBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibraryServiceServer).UpdateBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/LibraryService/UpdateBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibraryServiceServer).UpdateBook(ctx, req.(*UpdateBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LibraryService_DeleteBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibraryServiceServer).DeleteBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/LibraryService/DeleteBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibraryServiceServer).DeleteBook(ctx, req.(*DeleteBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LibraryService_ReserveBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReserveBookRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LibraryServiceServer).ReserveBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/LibraryService/ReserveBook",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LibraryServiceServer).ReserveBook(ctx, req.(*ReserveBookRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LibraryService_ServiceDesc is the grpc.ServiceDesc for LibraryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LibraryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "LibraryService",
	HandlerType: (*LibraryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBooks",
			Handler:    _LibraryService_GetBooks_Handler,
		},
		{
			MethodName: "FindBook",
			Handler:    _LibraryService_FindBook_Handler,
		},
		{
			MethodName: "AddBook",
			Handler:    _LibraryService_AddBook_Handler,
		},
		{
			MethodName: "UpdateBook",
			Handler:    _LibraryService_UpdateBook_Handler,
		},
		{
			MethodName: "DeleteBook",
			Handler:    _LibraryService_DeleteBook_Handler,
		},
		{
			MethodName: "ReserveBook",
			Handler:    _LibraryService_ReserveBook_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "library.proto",
}
