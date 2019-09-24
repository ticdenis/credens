// Code generated by protoc-gen-go. DO NOT EDIT.
// source: libs/apps/grpc/proto/service.proto

package src

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type SayHelloRequest struct {
	To                   string   `protobuf:"bytes,1,opt,name=to,proto3" json:"to,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SayHelloRequest) Reset()         { *m = SayHelloRequest{} }
func (m *SayHelloRequest) String() string { return proto.CompactTextString(m) }
func (*SayHelloRequest) ProtoMessage()    {}
func (*SayHelloRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c8aa1b01565d2b85, []int{0}
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

func (m *SayHelloRequest) GetTo() string {
	if m != nil {
		return m.To
	}
	return ""
}

type SayHelloResponse struct {
	Result               string   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SayHelloResponse) Reset()         { *m = SayHelloResponse{} }
func (m *SayHelloResponse) String() string { return proto.CompactTextString(m) }
func (*SayHelloResponse) ProtoMessage()    {}
func (*SayHelloResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c8aa1b01565d2b85, []int{1}
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

func (m *SayHelloResponse) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

func init() {
	proto.RegisterType((*SayHelloRequest)(nil), "proto.SayHelloRequest")
	proto.RegisterType((*SayHelloResponse)(nil), "proto.SayHelloResponse")
}

func init() {
	proto.RegisterFile("libs/apps/grpc/proto/service.proto", fileDescriptor_c8aa1b01565d2b85)
}

var fileDescriptor_c8aa1b01565d2b85 = []byte{
	// 167 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x2e, 0x2e, 0x4a, 0xd6,
	0x2f, 0x2d, 0x4e, 0x2d, 0x8a, 0xcf, 0xcc, 0x2b, 0x49, 0x2d, 0x4a, 0x4b, 0x4c, 0x4e, 0xd5, 0x4f,
	0x2f, 0x2a, 0x48, 0xd6, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0xd7, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c,
	0x4e, 0xd5, 0x03, 0xf3, 0x84, 0x58, 0xc1, 0x94, 0x92, 0x22, 0x17, 0x7f, 0x70, 0x62, 0xa5, 0x47,
	0x6a, 0x4e, 0x4e, 0x7e, 0x50, 0x6a, 0x61, 0x69, 0x6a, 0x71, 0x89, 0x10, 0x1f, 0x17, 0x53, 0x49,
	0xbe, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x53, 0x49, 0xbe, 0x92, 0x16, 0x97, 0x00, 0x42,
	0x49, 0x71, 0x41, 0x7e, 0x5e, 0x71, 0xaa, 0x90, 0x18, 0x17, 0x5b, 0x51, 0x6a, 0x71, 0x69, 0x4e,
	0x09, 0x54, 0x1d, 0x94, 0x67, 0xe4, 0xcd, 0xc5, 0xe5, 0x18, 0xe0, 0x19, 0x0c, 0xb1, 0x49, 0xc8,
	0x96, 0x8b, 0x03, 0xa6, 0x53, 0x48, 0x0c, 0x62, 0xaf, 0x1e, 0x9a, 0x6d, 0x52, 0xe2, 0x18, 0xe2,
	0x10, 0x2b, 0x94, 0x18, 0x92, 0xd8, 0xc0, 0x32, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x2c,
	0x3c, 0xea, 0xac, 0xd8, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// APIServiceClient is the client API for APIService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type APIServiceClient interface {
	SayHello(ctx context.Context, in *SayHelloRequest, opts ...grpc.CallOption) (*SayHelloResponse, error)
}

type aPIServiceClient struct {
	cc *grpc.ClientConn
}

func NewAPIServiceClient(cc *grpc.ClientConn) APIServiceClient {
	return &aPIServiceClient{cc}
}

func (c *aPIServiceClient) SayHello(ctx context.Context, in *SayHelloRequest, opts ...grpc.CallOption) (*SayHelloResponse, error) {
	out := new(SayHelloResponse)
	err := c.cc.Invoke(ctx, "/proto.APIService/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// APIServiceServer is the server API for APIService service.
type APIServiceServer interface {
	SayHello(context.Context, *SayHelloRequest) (*SayHelloResponse, error)
}

// UnimplementedAPIServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAPIServiceServer struct {
}

func (*UnimplementedAPIServiceServer) SayHello(ctx context.Context, req *SayHelloRequest) (*SayHelloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}

func RegisterAPIServiceServer(s *grpc.Server, srv APIServiceServer) {
	s.RegisterService(&_APIService_serviceDesc, srv)
}

func _APIService_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SayHelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServiceServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.APIService/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServiceServer).SayHello(ctx, req.(*SayHelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _APIService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.APIService",
	HandlerType: (*APIServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _APIService_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "libs/apps/grpc/proto/service.proto",
}