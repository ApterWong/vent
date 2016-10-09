// Code generated by protoc-gen-go.
// source: rpc.proto
// DO NOT EDIT!

package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for Register service

type RegisterClient interface {
	Register(ctx context.Context, in *RegisterC2S, opts ...grpc.CallOption) (*RegisterS2C, error)
}

type registerClient struct {
	cc *grpc.ClientConn
}

func NewRegisterClient(cc *grpc.ClientConn) RegisterClient {
	return &registerClient{cc}
}

func (c *registerClient) Register(ctx context.Context, in *RegisterC2S, opts ...grpc.CallOption) (*RegisterS2C, error) {
	out := new(RegisterS2C)
	err := grpc.Invoke(ctx, "/proto.Register/Register", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Register service

type RegisterServer interface {
	Register(context.Context, *RegisterC2S) (*RegisterS2C, error)
}

func RegisterRegisterServer(s *grpc.Server, srv RegisterServer) {
	s.RegisterService(&_Register_serviceDesc, srv)
}

func _Register_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterC2S)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegisterServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Register/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegisterServer).Register(ctx, req.(*RegisterC2S))
	}
	return interceptor(ctx, in, info, handler)
}

var _Register_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Register",
	HandlerType: (*RegisterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _Register_Register_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor1,
}

// Client API for Login service

type LoginClient interface {
	Login(ctx context.Context, in *LoginC2S, opts ...grpc.CallOption) (*LoginS2C, error)
}

type loginClient struct {
	cc *grpc.ClientConn
}

func NewLoginClient(cc *grpc.ClientConn) LoginClient {
	return &loginClient{cc}
}

func (c *loginClient) Login(ctx context.Context, in *LoginC2S, opts ...grpc.CallOption) (*LoginS2C, error) {
	out := new(LoginS2C)
	err := grpc.Invoke(ctx, "/proto.Login/Login", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Login service

type LoginServer interface {
	Login(context.Context, *LoginC2S) (*LoginS2C, error)
}

func RegisterLoginServer(s *grpc.Server, srv LoginServer) {
	s.RegisterService(&_Login_serviceDesc, srv)
}

func _Login_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginC2S)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoginServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Login/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoginServer).Login(ctx, req.(*LoginC2S))
	}
	return interceptor(ctx, in, info, handler)
}

var _Login_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Login",
	HandlerType: (*LoginServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _Login_Login_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor1,
}

// Client API for UserInfoManager service

type UserInfoManagerClient interface {
	UserInfoModify(ctx context.Context, in *UserInfoModifyC2S, opts ...grpc.CallOption) (*UserInfoModifyS2C, error)
	UserInfoGet(ctx context.Context, in *UserInfoGetC2S, opts ...grpc.CallOption) (*UserInfoGetS2C, error)
}

type userInfoManagerClient struct {
	cc *grpc.ClientConn
}

func NewUserInfoManagerClient(cc *grpc.ClientConn) UserInfoManagerClient {
	return &userInfoManagerClient{cc}
}

func (c *userInfoManagerClient) UserInfoModify(ctx context.Context, in *UserInfoModifyC2S, opts ...grpc.CallOption) (*UserInfoModifyS2C, error) {
	out := new(UserInfoModifyS2C)
	err := grpc.Invoke(ctx, "/proto.UserInfoManager/UserInfoModify", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userInfoManagerClient) UserInfoGet(ctx context.Context, in *UserInfoGetC2S, opts ...grpc.CallOption) (*UserInfoGetS2C, error) {
	out := new(UserInfoGetS2C)
	err := grpc.Invoke(ctx, "/proto.UserInfoManager/UserInfoGet", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserInfoManager service

type UserInfoManagerServer interface {
	UserInfoModify(context.Context, *UserInfoModifyC2S) (*UserInfoModifyS2C, error)
	UserInfoGet(context.Context, *UserInfoGetC2S) (*UserInfoGetS2C, error)
}

func RegisterUserInfoManagerServer(s *grpc.Server, srv UserInfoManagerServer) {
	s.RegisterService(&_UserInfoManager_serviceDesc, srv)
}

func _UserInfoManager_UserInfoModify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserInfoModifyC2S)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserInfoManagerServer).UserInfoModify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UserInfoManager/UserInfoModify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserInfoManagerServer).UserInfoModify(ctx, req.(*UserInfoModifyC2S))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserInfoManager_UserInfoGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserInfoGetC2S)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserInfoManagerServer).UserInfoGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.UserInfoManager/UserInfoGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserInfoManagerServer).UserInfoGet(ctx, req.(*UserInfoGetC2S))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserInfoManager_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.UserInfoManager",
	HandlerType: (*UserInfoManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UserInfoModify",
			Handler:    _UserInfoManager_UserInfoModify_Handler,
		},
		{
			MethodName: "UserInfoGet",
			Handler:    _UserInfoManager_UserInfoGet_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor1,
}

// Client API for GeoManager service

type GeoManagerClient interface {
	UserGeoUpload(ctx context.Context, in *GeoUploadC2S, opts ...grpc.CallOption) (*CommonS2C, error)
}

type geoManagerClient struct {
	cc *grpc.ClientConn
}

func NewGeoManagerClient(cc *grpc.ClientConn) GeoManagerClient {
	return &geoManagerClient{cc}
}

func (c *geoManagerClient) UserGeoUpload(ctx context.Context, in *GeoUploadC2S, opts ...grpc.CallOption) (*CommonS2C, error) {
	out := new(CommonS2C)
	err := grpc.Invoke(ctx, "/proto.GeoManager/UserGeoUpload", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for GeoManager service

type GeoManagerServer interface {
	UserGeoUpload(context.Context, *GeoUploadC2S) (*CommonS2C, error)
}

func RegisterGeoManagerServer(s *grpc.Server, srv GeoManagerServer) {
	s.RegisterService(&_GeoManager_serviceDesc, srv)
}

func _GeoManager_UserGeoUpload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GeoUploadC2S)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GeoManagerServer).UserGeoUpload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.GeoManager/UserGeoUpload",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GeoManagerServer).UserGeoUpload(ctx, req.(*GeoUploadC2S))
	}
	return interceptor(ctx, in, info, handler)
}

var _GeoManager_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.GeoManager",
	HandlerType: (*GeoManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UserGeoUpload",
			Handler:    _GeoManager_UserGeoUpload_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor1,
}

// Client API for SessionManager service

type SessionManagerClient interface {
	GetSession(ctx context.Context, in *GetSessionReq, opts ...grpc.CallOption) (*GetSessionRes, error)
	AuthSession(ctx context.Context, in *AuthReq, opts ...grpc.CallOption) (*CommonS2C, error)
}

type sessionManagerClient struct {
	cc *grpc.ClientConn
}

func NewSessionManagerClient(cc *grpc.ClientConn) SessionManagerClient {
	return &sessionManagerClient{cc}
}

func (c *sessionManagerClient) GetSession(ctx context.Context, in *GetSessionReq, opts ...grpc.CallOption) (*GetSessionRes, error) {
	out := new(GetSessionRes)
	err := grpc.Invoke(ctx, "/proto.SessionManager/GetSession", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sessionManagerClient) AuthSession(ctx context.Context, in *AuthReq, opts ...grpc.CallOption) (*CommonS2C, error) {
	out := new(CommonS2C)
	err := grpc.Invoke(ctx, "/proto.SessionManager/AuthSession", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SessionManager service

type SessionManagerServer interface {
	GetSession(context.Context, *GetSessionReq) (*GetSessionRes, error)
	AuthSession(context.Context, *AuthReq) (*CommonS2C, error)
}

func RegisterSessionManagerServer(s *grpc.Server, srv SessionManagerServer) {
	s.RegisterService(&_SessionManager_serviceDesc, srv)
}

func _SessionManager_GetSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSessionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionManagerServer).GetSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.SessionManager/GetSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionManagerServer).GetSession(ctx, req.(*GetSessionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SessionManager_AuthSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SessionManagerServer).AuthSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.SessionManager/AuthSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SessionManagerServer).AuthSession(ctx, req.(*AuthReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _SessionManager_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.SessionManager",
	HandlerType: (*SessionManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSession",
			Handler:    _SessionManager_GetSession_Handler,
		},
		{
			MethodName: "AuthSession",
			Handler:    _SessionManager_AuthSession_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor1,
}

func init() { proto1.RegisterFile("rpc.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 272 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x8f, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0xeb, 0xa1, 0xa2, 0x53, 0x4d, 0x65, 0x54, 0x90, 0x1c, 0x7b, 0x15, 0x02, 0xae, 0x3d,
	0x88, 0x20, 0x28, 0x81, 0x16, 0x41, 0x2f, 0x0d, 0x7d, 0x80, 0x98, 0x4e, 0x63, 0xa0, 0xdd, 0x89,
	0xbb, 0xab, 0xe0, 0xc9, 0xd7, 0xf0, 0x71, 0xdd, 0x4d, 0x76, 0x83, 0x2d, 0x39, 0x65, 0xe7, 0xdb,
	0xff, 0xfb, 0x33, 0x0b, 0xc7, 0xaa, 0x2e, 0x92, 0x5a, 0xb1, 0x61, 0x1c, 0x36, 0x9f, 0xf8, 0xa4,
	0xd8, 0x54, 0x24, 0x4d, 0xe2, 0x27, 0x4d, 0xea, 0x8b, 0x54, 0x3b, 0x89, 0x47, 0x38, 0x5a, 0x50,
	0x59, 0x69, 0x43, 0x0a, 0xa7, 0xff, 0xce, 0xd8, 0xde, 0x27, 0x01, 0xa4, 0x42, 0xc7, 0xfb, 0x2c,
	0x13, 0xc5, 0x64, 0x20, 0xa6, 0x30, 0x7c, 0xe1, 0xb2, 0x92, 0x78, 0x1d, 0x0e, 0x63, 0x9f, 0x6b,
	0x26, 0x27, 0xee, 0x80, 0xd6, 0xfa, 0x3d, 0x80, 0xf1, 0xd2, 0x6e, 0xf2, 0x2c, 0xd7, 0xfc, 0x9a,
	0xcb, 0xbc, 0xb4, 0xff, 0x9c, 0x41, 0xd4, 0x21, 0x5e, 0x55, 0xeb, 0x6f, 0xbc, 0xf2, 0xe2, 0x2e,
	0x4e, 0x45, 0x16, 0xf7, 0xdf, 0x64, 0x22, 0x9d, 0x0c, 0xf0, 0x01, 0x46, 0x01, 0xcf, 0xc9, 0xe0,
	0xe5, 0x5e, 0xd4, 0x32, 0xd7, 0xd0, 0x83, 0x1b, 0x5d, 0xcc, 0x00, 0xe6, 0xd4, 0x2d, 0x75, 0x07,
	0xa7, 0x2e, 0x61, 0xc9, 0xb2, 0xde, 0x70, 0xbe, 0xc2, 0x73, 0xef, 0x75, 0xc4, 0x95, 0x9d, 0x79,
	0x98, 0xf2, 0x76, 0xcb, 0xb2, 0xed, 0xf9, 0x81, 0x28, 0x23, 0xad, 0x2b, 0x96, 0xa1, 0xeb, 0xde,
	0x35, 0x1b, 0x0f, 0xf1, 0xa2, 0x2b, 0x0a, 0x68, 0x41, 0x1f, 0x71, 0x1f, 0xd5, 0xf6, 0x51, 0x37,
	0x30, 0x7a, 0xfa, 0x34, 0xef, 0x41, 0x8e, 0x7c, 0xcc, 0x31, 0xa7, 0xf5, 0x2c, 0xf0, 0x76, 0xd8,
	0xa0, 0xdb, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe5, 0x45, 0x94, 0x72, 0x12, 0x02, 0x00, 0x00,
}
