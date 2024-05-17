// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: protos/blog/blog.proto

package blog

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

// BlogServiceClient is the client API for BlogService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BlogServiceClient interface {
	CreateBlogPost(ctx context.Context, in *Post, opts ...grpc.CallOption) (*Post, error)
	ReadBlogPost(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Post, error)
	UpdateBlogPost(ctx context.Context, in *Post, opts ...grpc.CallOption) (*Post, error)
	DeleteBlogPost(ctx context.Context, in *Id, opts ...grpc.CallOption) (*DeleteResponse, error)
	Read(ctx context.Context, in *Readall, opts ...grpc.CallOption) (*Datas, error)
}

type blogServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBlogServiceClient(cc grpc.ClientConnInterface) BlogServiceClient {
	return &blogServiceClient{cc}
}

func (c *blogServiceClient) CreateBlogPost(ctx context.Context, in *Post, opts ...grpc.CallOption) (*Post, error) {
	out := new(Post)
	err := c.cc.Invoke(ctx, "/blog.BlogService/CreateBlogPost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogServiceClient) ReadBlogPost(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Post, error) {
	out := new(Post)
	err := c.cc.Invoke(ctx, "/blog.BlogService/ReadBlogPost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogServiceClient) UpdateBlogPost(ctx context.Context, in *Post, opts ...grpc.CallOption) (*Post, error) {
	out := new(Post)
	err := c.cc.Invoke(ctx, "/blog.BlogService/UpdateBlogPost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogServiceClient) DeleteBlogPost(ctx context.Context, in *Id, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/blog.BlogService/DeleteBlogPost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogServiceClient) Read(ctx context.Context, in *Readall, opts ...grpc.CallOption) (*Datas, error) {
	out := new(Datas)
	err := c.cc.Invoke(ctx, "/blog.BlogService/Read", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BlogServiceServer is the server API for BlogService service.
// All implementations must embed UnimplementedBlogServiceServer
// for forward compatibility
type BlogServiceServer interface {
	CreateBlogPost(context.Context, *Post) (*Post, error)
	ReadBlogPost(context.Context, *Id) (*Post, error)
	UpdateBlogPost(context.Context, *Post) (*Post, error)
	DeleteBlogPost(context.Context, *Id) (*DeleteResponse, error)
	Read(context.Context, *Readall) (*Datas, error)
	mustEmbedUnimplementedBlogServiceServer()
}

// UnimplementedBlogServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBlogServiceServer struct {
}

func (UnimplementedBlogServiceServer) CreateBlogPost(context.Context, *Post) (*Post, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBlogPost not implemented")
}
func (UnimplementedBlogServiceServer) ReadBlogPost(context.Context, *Id) (*Post, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadBlogPost not implemented")
}
func (UnimplementedBlogServiceServer) UpdateBlogPost(context.Context, *Post) (*Post, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBlogPost not implemented")
}
func (UnimplementedBlogServiceServer) DeleteBlogPost(context.Context, *Id) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBlogPost not implemented")
}
func (UnimplementedBlogServiceServer) Read(context.Context, *Readall) (*Datas, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Read not implemented")
}
func (UnimplementedBlogServiceServer) mustEmbedUnimplementedBlogServiceServer() {}

// UnsafeBlogServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BlogServiceServer will
// result in compilation errors.
type UnsafeBlogServiceServer interface {
	mustEmbedUnimplementedBlogServiceServer()
}

func RegisterBlogServiceServer(s grpc.ServiceRegistrar, srv BlogServiceServer) {
	s.RegisterService(&BlogService_ServiceDesc, srv)
}

func _BlogService_CreateBlogPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Post)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServiceServer).CreateBlogPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blog.BlogService/CreateBlogPost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServiceServer).CreateBlogPost(ctx, req.(*Post))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlogService_ReadBlogPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServiceServer).ReadBlogPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blog.BlogService/ReadBlogPost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServiceServer).ReadBlogPost(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlogService_UpdateBlogPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Post)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServiceServer).UpdateBlogPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blog.BlogService/UpdateBlogPost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServiceServer).UpdateBlogPost(ctx, req.(*Post))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlogService_DeleteBlogPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServiceServer).DeleteBlogPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blog.BlogService/DeleteBlogPost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServiceServer).DeleteBlogPost(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlogService_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Readall)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServiceServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blog.BlogService/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServiceServer).Read(ctx, req.(*Readall))
	}
	return interceptor(ctx, in, info, handler)
}

// BlogService_ServiceDesc is the grpc.ServiceDesc for BlogService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BlogService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "blog.BlogService",
	HandlerType: (*BlogServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBlogPost",
			Handler:    _BlogService_CreateBlogPost_Handler,
		},
		{
			MethodName: "ReadBlogPost",
			Handler:    _BlogService_ReadBlogPost_Handler,
		},
		{
			MethodName: "UpdateBlogPost",
			Handler:    _BlogService_UpdateBlogPost_Handler,
		},
		{
			MethodName: "DeleteBlogPost",
			Handler:    _BlogService_DeleteBlogPost_Handler,
		},
		{
			MethodName: "Read",
			Handler:    _BlogService_Read_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/blog/blog.proto",
}
