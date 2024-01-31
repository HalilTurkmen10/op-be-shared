// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package user

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

// UserSvcClient is the client API for UserSvc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserSvcClient interface {
	GetUsersByFilter(ctx context.Context, in *UserFilter, opts ...grpc.CallOption) (*Users, error)
	CreateUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error)
	DeleteUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error)
	UpdateUserRole(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error)
	UpdateUserBase(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error)
	UpdateUserStatus(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error)
	ChangePassword(ctx context.Context, in *UserPassword, opts ...grpc.CallOption) (*UserPasswordResult, error)
}

type userSvcClient struct {
	cc grpc.ClientConnInterface
}

func NewUserSvcClient(cc grpc.ClientConnInterface) UserSvcClient {
	return &userSvcClient{cc}
}

func (c *userSvcClient) GetUsersByFilter(ctx context.Context, in *UserFilter, opts ...grpc.CallOption) (*Users, error) {
	out := new(Users)
	err := c.cc.Invoke(ctx, "/UserSvc/GetUsersByFilter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userSvcClient) CreateUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/UserSvc/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userSvcClient) DeleteUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/UserSvc/DeleteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userSvcClient) UpdateUserRole(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/UserSvc/UpdateUserRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userSvcClient) UpdateUserBase(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/UserSvc/UpdateUserBase", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userSvcClient) UpdateUserStatus(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/UserSvc/UpdateUserStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userSvcClient) ChangePassword(ctx context.Context, in *UserPassword, opts ...grpc.CallOption) (*UserPasswordResult, error) {
	out := new(UserPasswordResult)
	err := c.cc.Invoke(ctx, "/UserSvc/ChangePassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserSvcServer is the server API for UserSvc service.
// All implementations must embed UnimplementedUserSvcServer
// for forward compatibility
type UserSvcServer interface {
	GetUsersByFilter(context.Context, *UserFilter) (*Users, error)
	CreateUser(context.Context, *User) (*User, error)
	DeleteUser(context.Context, *User) (*User, error)
	UpdateUserRole(context.Context, *User) (*User, error)
	UpdateUserBase(context.Context, *User) (*User, error)
	UpdateUserStatus(context.Context, *User) (*User, error)
	ChangePassword(context.Context, *UserPassword) (*UserPasswordResult, error)
	mustEmbedUnimplementedUserSvcServer()
}

// UnimplementedUserSvcServer must be embedded to have forward compatible implementations.
type UnimplementedUserSvcServer struct {
}

func (UnimplementedUserSvcServer) GetUsersByFilter(context.Context, *UserFilter) (*Users, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsersByFilter not implemented")
}
func (UnimplementedUserSvcServer) CreateUser(context.Context, *User) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedUserSvcServer) DeleteUser(context.Context, *User) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedUserSvcServer) UpdateUserRole(context.Context, *User) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserRole not implemented")
}
func (UnimplementedUserSvcServer) UpdateUserBase(context.Context, *User) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserBase not implemented")
}
func (UnimplementedUserSvcServer) UpdateUserStatus(context.Context, *User) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserStatus not implemented")
}
func (UnimplementedUserSvcServer) ChangePassword(context.Context, *UserPassword) (*UserPasswordResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangePassword not implemented")
}
func (UnimplementedUserSvcServer) mustEmbedUnimplementedUserSvcServer() {}

// UnsafeUserSvcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserSvcServer will
// result in compilation errors.
type UnsafeUserSvcServer interface {
	mustEmbedUnimplementedUserSvcServer()
}

func RegisterUserSvcServer(s grpc.ServiceRegistrar, srv UserSvcServer) {
	s.RegisterService(&UserSvc_ServiceDesc, srv)
}

func _UserSvc_GetUsersByFilter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserFilter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserSvcServer).GetUsersByFilter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserSvc/GetUsersByFilter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserSvcServer).GetUsersByFilter(ctx, req.(*UserFilter))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserSvc_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserSvcServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserSvc/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserSvcServer).CreateUser(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserSvc_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserSvcServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserSvc/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserSvcServer).DeleteUser(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserSvc_UpdateUserRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserSvcServer).UpdateUserRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserSvc/UpdateUserRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserSvcServer).UpdateUserRole(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserSvc_UpdateUserBase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserSvcServer).UpdateUserBase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserSvc/UpdateUserBase",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserSvcServer).UpdateUserBase(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserSvc_UpdateUserStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserSvcServer).UpdateUserStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserSvc/UpdateUserStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserSvcServer).UpdateUserStatus(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserSvc_ChangePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserPassword)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserSvcServer).ChangePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserSvc/ChangePassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserSvcServer).ChangePassword(ctx, req.(*UserPassword))
	}
	return interceptor(ctx, in, info, handler)
}

// UserSvc_ServiceDesc is the grpc.ServiceDesc for UserSvc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserSvc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "UserSvc",
	HandlerType: (*UserSvcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUsersByFilter",
			Handler:    _UserSvc_GetUsersByFilter_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _UserSvc_CreateUser_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _UserSvc_DeleteUser_Handler,
		},
		{
			MethodName: "UpdateUserRole",
			Handler:    _UserSvc_UpdateUserRole_Handler,
		},
		{
			MethodName: "UpdateUserBase",
			Handler:    _UserSvc_UpdateUserBase_Handler,
		},
		{
			MethodName: "UpdateUserStatus",
			Handler:    _UserSvc_UpdateUserStatus_Handler,
		},
		{
			MethodName: "ChangePassword",
			Handler:    _UserSvc_ChangePassword_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
