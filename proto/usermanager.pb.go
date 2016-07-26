// Code generated by protoc-gen-go.
// source: usermanager.proto
// DO NOT EDIT!

/*
Package usermanager is a generated protocol buffer package.

It is generated from these files:
	usermanager.proto

It has these top-level messages:
	LoginRequest
	LoginReply
*/
package usermanager

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type LoginRequest struct {
	Userid string `protobuf:"bytes,1,opt,name=userid" json:"userid,omitempty"`
	Prodid string `protobuf:"bytes,2,opt,name=prodid" json:"prodid,omitempty"`
	Task   string `protobuf:"bytes,3,opt,name=task" json:"task,omitempty"`
}

func (m *LoginRequest) Reset()                    { *m = LoginRequest{} }
func (m *LoginRequest) String() string            { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()               {}
func (*LoginRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type LoginReply struct {
	Token string `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
	Err   string `protobuf:"bytes,2,opt,name=err" json:"err,omitempty"`
}

func (m *LoginReply) Reset()                    { *m = LoginReply{} }
func (m *LoginReply) String() string            { return proto.CompactTextString(m) }
func (*LoginReply) ProtoMessage()               {}
func (*LoginReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*LoginRequest)(nil), "usermanager.LoginRequest")
	proto.RegisterType((*LoginReply)(nil), "usermanager.LoginReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for UserManager service

type UserManagerClient interface {
	// Obtains the feature at a given position.
	GetToken(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginReply, error)
}

type userManagerClient struct {
	cc *grpc.ClientConn
}

func NewUserManagerClient(cc *grpc.ClientConn) UserManagerClient {
	return &userManagerClient{cc}
}

func (c *userManagerClient) GetToken(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginReply, error) {
	out := new(LoginReply)
	err := grpc.Invoke(ctx, "/usermanager.UserManager/GetToken", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserManager service

type UserManagerServer interface {
	// Obtains the feature at a given position.
	GetToken(context.Context, *LoginRequest) (*LoginReply, error)
}

func RegisterUserManagerServer(s *grpc.Server, srv UserManagerServer) {
	s.RegisterService(&_UserManager_serviceDesc, srv)
}

func _UserManager_GetToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserManagerServer).GetToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/usermanager.UserManager/GetToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserManagerServer).GetToken(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserManager_serviceDesc = grpc.ServiceDesc{
	ServiceName: "usermanager.UserManager",
	HandlerType: (*UserManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetToken",
			Handler:    _UserManager_GetToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("usermanager.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 218 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x2c, 0x2d, 0x4e, 0x2d,
	0xca, 0x4d, 0xcc, 0x4b, 0x4c, 0x4f, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x46,
	0x12, 0x52, 0x0a, 0xe2, 0xe2, 0xf1, 0xc9, 0x4f, 0xcf, 0xcc, 0x0b, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d,
	0x2e, 0x11, 0x12, 0xe3, 0x62, 0x03, 0x49, 0x67, 0xa6, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06,
	0x41, 0x79, 0x20, 0x71, 0xa0, 0xee, 0x14, 0xa0, 0x38, 0x13, 0x44, 0x1c, 0xc2, 0x13, 0x12, 0xe2,
	0x62, 0x29, 0x49, 0x2c, 0xce, 0x96, 0x60, 0x06, 0x8b, 0x82, 0xd9, 0x4a, 0x26, 0x5c, 0x5c, 0x50,
	0x33, 0x0b, 0x72, 0x2a, 0x85, 0x44, 0xb8, 0x58, 0x4b, 0xf2, 0xb3, 0x53, 0xf3, 0xa0, 0x06, 0x42,
	0x38, 0x42, 0x02, 0x5c, 0xcc, 0xa9, 0x45, 0x45, 0x50, 0xc3, 0x40, 0x4c, 0x23, 0x7f, 0x2e, 0xee,
	0x50, 0xa0, 0x5d, 0xbe, 0x10, 0x87, 0x09, 0x39, 0x70, 0x71, 0xb8, 0xa7, 0x96, 0x84, 0x80, 0x15,
	0x4b, 0xea, 0x21, 0xfb, 0x02, 0xd9, 0xbd, 0x52, 0xe2, 0xd8, 0xa4, 0x80, 0xd6, 0x2a, 0x31, 0x38,
	0x59, 0x71, 0xa9, 0x24, 0x66, 0xea, 0x25, 0x96, 0x96, 0xe4, 0xe7, 0xe5, 0xe7, 0xe6, 0x97, 0x16,
	0xeb, 0x15, 0x17, 0xa4, 0xa6, 0x26, 0x67, 0x24, 0x16, 0x14, 0x14, 0x23, 0x6b, 0x72, 0x12, 0x40,
	0xb2, 0x36, 0x00, 0x14, 0x42, 0x01, 0x8c, 0x49, 0x6c, 0xe0, 0xa0, 0x32, 0x06, 0x04, 0x00, 0x00,
	0xff, 0xff, 0xe0, 0x5f, 0x17, 0x7d, 0x3f, 0x01, 0x00, 0x00,
}
