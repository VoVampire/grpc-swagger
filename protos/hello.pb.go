// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protos/hello.proto

package hello

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// SayHello请求
type SayHelloRequest struct {
	// 姓名
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SayHelloRequest) Reset()         { *m = SayHelloRequest{} }
func (m *SayHelloRequest) String() string { return proto.CompactTextString(m) }
func (*SayHelloRequest) ProtoMessage()    {}
func (*SayHelloRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_525aa0d637d17553, []int{0}
}

func (m *SayHelloRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SayHelloRequest.Unmarshal(m, b)
}
func (m *SayHelloRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SayHelloRequest.Marshal(b, m, deterministic)
}
func (m *SayHelloRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SayHelloRequest.Merge(m, src)
}
func (m *SayHelloRequest) XXX_Size() int {
	return xxx_messageInfo_SayHelloRequest.Size(m)
}
func (m *SayHelloRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SayHelloRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SayHelloRequest proto.InternalMessageInfo

func (m *SayHelloRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// SayHello响应
type SayHelloResponse struct {
	// 返回信息
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SayHelloResponse) Reset()         { *m = SayHelloResponse{} }
func (m *SayHelloResponse) String() string { return proto.CompactTextString(m) }
func (*SayHelloResponse) ProtoMessage()    {}
func (*SayHelloResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_525aa0d637d17553, []int{1}
}

func (m *SayHelloResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SayHelloResponse.Unmarshal(m, b)
}
func (m *SayHelloResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SayHelloResponse.Marshal(b, m, deterministic)
}
func (m *SayHelloResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SayHelloResponse.Merge(m, src)
}
func (m *SayHelloResponse) XXX_Size() int {
	return xxx_messageInfo_SayHelloResponse.Size(m)
}
func (m *SayHelloResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SayHelloResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SayHelloResponse proto.InternalMessageInfo

func (m *SayHelloResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

// SayGoodbye请求
type SayGoodbyeRequest struct {
	// 姓名
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SayGoodbyeRequest) Reset()         { *m = SayGoodbyeRequest{} }
func (m *SayGoodbyeRequest) String() string { return proto.CompactTextString(m) }
func (*SayGoodbyeRequest) ProtoMessage()    {}
func (*SayGoodbyeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_525aa0d637d17553, []int{2}
}

func (m *SayGoodbyeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SayGoodbyeRequest.Unmarshal(m, b)
}
func (m *SayGoodbyeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SayGoodbyeRequest.Marshal(b, m, deterministic)
}
func (m *SayGoodbyeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SayGoodbyeRequest.Merge(m, src)
}
func (m *SayGoodbyeRequest) XXX_Size() int {
	return xxx_messageInfo_SayGoodbyeRequest.Size(m)
}
func (m *SayGoodbyeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SayGoodbyeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SayGoodbyeRequest proto.InternalMessageInfo

func (m *SayGoodbyeRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// SayGoodbye响应
type SayGoodbyeResponse struct {
	// 返回信息
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SayGoodbyeResponse) Reset()         { *m = SayGoodbyeResponse{} }
func (m *SayGoodbyeResponse) String() string { return proto.CompactTextString(m) }
func (*SayGoodbyeResponse) ProtoMessage()    {}
func (*SayGoodbyeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_525aa0d637d17553, []int{3}
}

func (m *SayGoodbyeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SayGoodbyeResponse.Unmarshal(m, b)
}
func (m *SayGoodbyeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SayGoodbyeResponse.Marshal(b, m, deterministic)
}
func (m *SayGoodbyeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SayGoodbyeResponse.Merge(m, src)
}
func (m *SayGoodbyeResponse) XXX_Size() int {
	return xxx_messageInfo_SayGoodbyeResponse.Size(m)
}
func (m *SayGoodbyeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SayGoodbyeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SayGoodbyeResponse proto.InternalMessageInfo

func (m *SayGoodbyeResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*SayHelloRequest)(nil), "hello.SayHelloRequest")
	proto.RegisterType((*SayHelloResponse)(nil), "hello.SayHelloResponse")
	proto.RegisterType((*SayGoodbyeRequest)(nil), "hello.SayGoodbyeRequest")
	proto.RegisterType((*SayGoodbyeResponse)(nil), "hello.SayGoodbyeResponse")
}

func init() { proto.RegisterFile("protos/hello.proto", fileDescriptor_525aa0d637d17553) }

var fileDescriptor_525aa0d637d17553 = []byte{
	// 259 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2a, 0x28, 0xca, 0x2f,
	0xc9, 0x2f, 0xd6, 0xcf, 0x48, 0xcd, 0xc9, 0xc9, 0xd7, 0x03, 0x73, 0x84, 0x58, 0xc1, 0x1c, 0x29,
	0x99, 0xf4, 0xfc, 0xfc, 0xf4, 0x9c, 0x54, 0xfd, 0xc4, 0x82, 0x4c, 0xfd, 0xc4, 0xbc, 0xbc, 0xfc,
	0x92, 0xc4, 0x92, 0xcc, 0xfc, 0xbc, 0x62, 0x88, 0x22, 0x25, 0x55, 0x2e, 0xfe, 0xe0, 0xc4, 0x4a,
	0x0f, 0x90, 0xca, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x21, 0x21, 0x2e, 0x96, 0xbc, 0xc4,
	0xdc, 0x54, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x30, 0x5b, 0x49, 0x87, 0x4b, 0x00, 0xa1,
	0xac, 0xb8, 0x20, 0x3f, 0xaf, 0x38, 0x55, 0x48, 0x82, 0x8b, 0x3d, 0x37, 0xb5, 0xb8, 0x38, 0x31,
	0x1d, 0xa6, 0x14, 0xc6, 0x55, 0x52, 0xe7, 0x12, 0x0c, 0x4e, 0xac, 0x74, 0xcf, 0xcf, 0x4f, 0x49,
	0xaa, 0x4c, 0xc5, 0x67, 0xac, 0x1e, 0x97, 0x10, 0xb2, 0x42, 0x42, 0x06, 0x1b, 0x5d, 0x63, 0xe4,
	0xe2, 0x01, 0x3b, 0x22, 0x38, 0xb5, 0xa8, 0x2c, 0x33, 0x39, 0x55, 0x28, 0x8e, 0x8b, 0x03, 0xe6,
	0x2e, 0x21, 0x31, 0x3d, 0x88, 0xef, 0xd1, 0xfc, 0x23, 0x25, 0x8e, 0x21, 0x0e, 0xb1, 0x47, 0x49,
	0xb1, 0xe9, 0xf2, 0x93, 0xc9, 0x4c, 0xd2, 0x4a, 0x62, 0xfa, 0x65, 0x86, 0xfa, 0xa9, 0x15, 0x89,
	0xb9, 0x05, 0x39, 0xa9, 0xc5, 0xfa, 0xc5, 0x89, 0x95, 0xf1, 0x60, 0x0d, 0x56, 0x8c, 0x5a, 0x42,
	0x69, 0x5c, 0x5c, 0x08, 0x07, 0x0a, 0x49, 0x20, 0x4c, 0x42, 0xf5, 0x9c, 0x94, 0x24, 0x16, 0x19,
	0xa8, 0x2d, 0xca, 0x60, 0x5b, 0x64, 0x95, 0x24, 0x30, 0x6c, 0x49, 0x87, 0xa8, 0xb4, 0x62, 0xd4,
	0x4a, 0x62, 0x03, 0xc7, 0x86, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xf8, 0x01, 0xe4, 0x55, 0xc8,
	0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HelloServiceClient is the client API for HelloService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HelloServiceClient interface {
	// SayHello接口
	SayHello(ctx context.Context, in *SayHelloRequest, opts ...grpc.CallOption) (*SayHelloResponse, error)
	// SayGoodbye接口
	SayGoodbye(ctx context.Context, in *SayGoodbyeRequest, opts ...grpc.CallOption) (*SayGoodbyeResponse, error)
}

type helloServiceClient struct {
	cc *grpc.ClientConn
}

func NewHelloServiceClient(cc *grpc.ClientConn) HelloServiceClient {
	return &helloServiceClient{cc}
}

func (c *helloServiceClient) SayHello(ctx context.Context, in *SayHelloRequest, opts ...grpc.CallOption) (*SayHelloResponse, error) {
	out := new(SayHelloResponse)
	err := c.cc.Invoke(ctx, "/hello.HelloService/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helloServiceClient) SayGoodbye(ctx context.Context, in *SayGoodbyeRequest, opts ...grpc.CallOption) (*SayGoodbyeResponse, error) {
	out := new(SayGoodbyeResponse)
	err := c.cc.Invoke(ctx, "/hello.HelloService/SayGoodbye", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HelloServiceServer is the server API for HelloService service.
type HelloServiceServer interface {
	// SayHello接口
	SayHello(context.Context, *SayHelloRequest) (*SayHelloResponse, error)
	// SayGoodbye接口
	SayGoodbye(context.Context, *SayGoodbyeRequest) (*SayGoodbyeResponse, error)
}

// UnimplementedHelloServiceServer can be embedded to have forward compatible implementations.
type UnimplementedHelloServiceServer struct {
}

func (*UnimplementedHelloServiceServer) SayHello(ctx context.Context, req *SayHelloRequest) (*SayHelloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (*UnimplementedHelloServiceServer) SayGoodbye(ctx context.Context, req *SayGoodbyeRequest) (*SayGoodbyeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayGoodbye not implemented")
}

func RegisterHelloServiceServer(s *grpc.Server, srv HelloServiceServer) {
	s.RegisterService(&_HelloService_serviceDesc, srv)
}

func _HelloService_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SayHelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloServiceServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hello.HelloService/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloServiceServer).SayHello(ctx, req.(*SayHelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HelloService_SayGoodbye_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SayGoodbyeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloServiceServer).SayGoodbye(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hello.HelloService/SayGoodbye",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloServiceServer).SayGoodbye(ctx, req.(*SayGoodbyeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _HelloService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hello.HelloService",
	HandlerType: (*HelloServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _HelloService_SayHello_Handler,
		},
		{
			MethodName: "SayGoodbye",
			Handler:    _HelloService_SayGoodbye_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/hello.proto",
}
