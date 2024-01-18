// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.1
// source: proto/todo.proto

package todo

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

// TodoServiceClient is the client API for TodoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TodoServiceClient interface {
	CreateTodo(ctx context.Context, in *NewTodo, opts ...grpc.CallOption) (*Todos, error)
	FetchAllTodos(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*ArrayOfTodos, error)
	EditSingleTodo(ctx context.Context, in *EditOrDeleteRequest, opts ...grpc.CallOption) (*Todos, error)
	DeleteSingleTodo(ctx context.Context, in *EditOrDeleteRequest, opts ...grpc.CallOption) (*Todos, error)
}

type todoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTodoServiceClient(cc grpc.ClientConnInterface) TodoServiceClient {
	return &todoServiceClient{cc}
}

func (c *todoServiceClient) CreateTodo(ctx context.Context, in *NewTodo, opts ...grpc.CallOption) (*Todos, error) {
	out := new(Todos)
	err := c.cc.Invoke(ctx, "/todo.TodoService/CreateTodo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) FetchAllTodos(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*ArrayOfTodos, error) {
	out := new(ArrayOfTodos)
	err := c.cc.Invoke(ctx, "/todo.TodoService/FetchAllTodos", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) EditSingleTodo(ctx context.Context, in *EditOrDeleteRequest, opts ...grpc.CallOption) (*Todos, error) {
	out := new(Todos)
	err := c.cc.Invoke(ctx, "/todo.TodoService/EditSingleTodo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoServiceClient) DeleteSingleTodo(ctx context.Context, in *EditOrDeleteRequest, opts ...grpc.CallOption) (*Todos, error) {
	out := new(Todos)
	err := c.cc.Invoke(ctx, "/todo.TodoService/DeleteSingleTodo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TodoServiceServer is the server API for TodoService service.
// All implementations must embed UnimplementedTodoServiceServer
// for forward compatibility
type TodoServiceServer interface {
	CreateTodo(context.Context, *NewTodo) (*Todos, error)
	FetchAllTodos(context.Context, *EmptyRequest) (*ArrayOfTodos, error)
	EditSingleTodo(context.Context, *EditOrDeleteRequest) (*Todos, error)
	DeleteSingleTodo(context.Context, *EditOrDeleteRequest) (*Todos, error)
	mustEmbedUnimplementedTodoServiceServer()
}

// UnimplementedTodoServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTodoServiceServer struct {
}

func (UnimplementedTodoServiceServer) CreateTodo(context.Context, *NewTodo) (*Todos, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTodo not implemented")
}
func (UnimplementedTodoServiceServer) FetchAllTodos(context.Context, *EmptyRequest) (*ArrayOfTodos, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchAllTodos not implemented")
}
func (UnimplementedTodoServiceServer) EditSingleTodo(context.Context, *EditOrDeleteRequest) (*Todos, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditSingleTodo not implemented")
}
func (UnimplementedTodoServiceServer) DeleteSingleTodo(context.Context, *EditOrDeleteRequest) (*Todos, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSingleTodo not implemented")
}
func (UnimplementedTodoServiceServer) mustEmbedUnimplementedTodoServiceServer() {}

// UnsafeTodoServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TodoServiceServer will
// result in compilation errors.
type UnsafeTodoServiceServer interface {
	mustEmbedUnimplementedTodoServiceServer()
}

func RegisterTodoServiceServer(s grpc.ServiceRegistrar, srv TodoServiceServer) {
	s.RegisterService(&TodoService_ServiceDesc, srv)
}

func _TodoService_CreateTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewTodo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).CreateTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todo.TodoService/CreateTodo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).CreateTodo(ctx, req.(*NewTodo))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoService_FetchAllTodos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).FetchAllTodos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todo.TodoService/FetchAllTodos",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).FetchAllTodos(ctx, req.(*EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoService_EditSingleTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditOrDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).EditSingleTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todo.TodoService/EditSingleTodo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).EditSingleTodo(ctx, req.(*EditOrDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoService_DeleteSingleTodo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditOrDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoServiceServer).DeleteSingleTodo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todo.TodoService/DeleteSingleTodo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoServiceServer).DeleteSingleTodo(ctx, req.(*EditOrDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TodoService_ServiceDesc is the grpc.ServiceDesc for TodoService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TodoService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "todo.TodoService",
	HandlerType: (*TodoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTodo",
			Handler:    _TodoService_CreateTodo_Handler,
		},
		{
			MethodName: "FetchAllTodos",
			Handler:    _TodoService_FetchAllTodos_Handler,
		},
		{
			MethodName: "EditSingleTodo",
			Handler:    _TodoService_EditSingleTodo_Handler,
		},
		{
			MethodName: "DeleteSingleTodo",
			Handler:    _TodoService_DeleteSingleTodo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/todo.proto",
}
