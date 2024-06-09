// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: api/proto/group.proto

package group_v1

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
	Group_Register_FullMethodName        = "/group.Group/Register"
	Group_GetCode_FullMethodName         = "/group.Group/GetCode"
	Group_AddToGroup_FullMethodName      = "/group.Group/AddToGroup"
	Group_RemoveFromGroup_FullMethodName = "/group.Group/RemoveFromGroup"
	Group_GetMembers_FullMethodName      = "/group.Group/GetMembers"
)

// GroupClient is the client API for Group service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GroupClient interface {
	Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*BoolResp, error)
	GetCode(ctx context.Context, in *GetReq, opts ...grpc.CallOption) (*GetCodeResp, error)
	AddToGroup(ctx context.Context, in *AddReq, opts ...grpc.CallOption) (*BoolResp, error)
	RemoveFromGroup(ctx context.Context, in *RemoveReq, opts ...grpc.CallOption) (*BoolResp, error)
	GetMembers(ctx context.Context, in *GetReq, opts ...grpc.CallOption) (*GetMembersResp, error)
}

type groupClient struct {
	cc grpc.ClientConnInterface
}

func NewGroupClient(cc grpc.ClientConnInterface) GroupClient {
	return &groupClient{cc}
}

func (c *groupClient) Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*BoolResp, error) {
	out := new(BoolResp)
	err := c.cc.Invoke(ctx, Group_Register_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupClient) GetCode(ctx context.Context, in *GetReq, opts ...grpc.CallOption) (*GetCodeResp, error) {
	out := new(GetCodeResp)
	err := c.cc.Invoke(ctx, Group_GetCode_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupClient) AddToGroup(ctx context.Context, in *AddReq, opts ...grpc.CallOption) (*BoolResp, error) {
	out := new(BoolResp)
	err := c.cc.Invoke(ctx, Group_AddToGroup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupClient) RemoveFromGroup(ctx context.Context, in *RemoveReq, opts ...grpc.CallOption) (*BoolResp, error) {
	out := new(BoolResp)
	err := c.cc.Invoke(ctx, Group_RemoveFromGroup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupClient) GetMembers(ctx context.Context, in *GetReq, opts ...grpc.CallOption) (*GetMembersResp, error) {
	out := new(GetMembersResp)
	err := c.cc.Invoke(ctx, Group_GetMembers_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GroupServer is the server API for Group service.
// All implementations must embed UnimplementedGroupServer
// for forward compatibility
type GroupServer interface {
	Register(context.Context, *RegisterReq) (*BoolResp, error)
	GetCode(context.Context, *GetReq) (*GetCodeResp, error)
	AddToGroup(context.Context, *AddReq) (*BoolResp, error)
	RemoveFromGroup(context.Context, *RemoveReq) (*BoolResp, error)
	GetMembers(context.Context, *GetReq) (*GetMembersResp, error)
	mustEmbedUnimplementedGroupServer()
}

// UnimplementedGroupServer must be embedded to have forward compatible implementations.
type UnimplementedGroupServer struct {
}

func (UnimplementedGroupServer) Register(context.Context, *RegisterReq) (*BoolResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedGroupServer) GetCode(context.Context, *GetReq) (*GetCodeResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCode not implemented")
}
func (UnimplementedGroupServer) AddToGroup(context.Context, *AddReq) (*BoolResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddToGroup not implemented")
}
func (UnimplementedGroupServer) RemoveFromGroup(context.Context, *RemoveReq) (*BoolResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveFromGroup not implemented")
}
func (UnimplementedGroupServer) GetMembers(context.Context, *GetReq) (*GetMembersResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMembers not implemented")
}
func (UnimplementedGroupServer) mustEmbedUnimplementedGroupServer() {}

// UnsafeGroupServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GroupServer will
// result in compilation errors.
type UnsafeGroupServer interface {
	mustEmbedUnimplementedGroupServer()
}

func RegisterGroupServer(s grpc.ServiceRegistrar, srv GroupServer) {
	s.RegisterService(&Group_ServiceDesc, srv)
}

func _Group_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Group_Register_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServer).Register(ctx, req.(*RegisterReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Group_GetCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServer).GetCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Group_GetCode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServer).GetCode(ctx, req.(*GetReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Group_AddToGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServer).AddToGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Group_AddToGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServer).AddToGroup(ctx, req.(*AddReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Group_RemoveFromGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServer).RemoveFromGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Group_RemoveFromGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServer).RemoveFromGroup(ctx, req.(*RemoveReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Group_GetMembers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServer).GetMembers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Group_GetMembers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServer).GetMembers(ctx, req.(*GetReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Group_ServiceDesc is the grpc.ServiceDesc for Group service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Group_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "group.Group",
	HandlerType: (*GroupServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _Group_Register_Handler,
		},
		{
			MethodName: "GetCode",
			Handler:    _Group_GetCode_Handler,
		},
		{
			MethodName: "AddToGroup",
			Handler:    _Group_AddToGroup_Handler,
		},
		{
			MethodName: "RemoveFromGroup",
			Handler:    _Group_RemoveFromGroup_Handler,
		},
		{
			MethodName: "GetMembers",
			Handler:    _Group_GetMembers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/proto/group.proto",
}