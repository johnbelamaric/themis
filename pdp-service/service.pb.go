// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

/*
Package service is a generated protocol buffer package.

It is generated from these files:
	service.proto

It has these top-level messages:
	Attribute
	Request
	Response
*/
package service

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

type Response_Effect int32

const (
	Response_DENY            Response_Effect = 0
	Response_PERMIT          Response_Effect = 1
	Response_NOTAPPLICABLE   Response_Effect = 2
	Response_INDETERMINATE   Response_Effect = 3
	Response_INDETERMINATED  Response_Effect = 4
	Response_INDETERMINATEP  Response_Effect = 5
	Response_INDETERMINATEDP Response_Effect = 6
)

var Response_Effect_name = map[int32]string{
	0: "DENY",
	1: "PERMIT",
	2: "NOTAPPLICABLE",
	3: "INDETERMINATE",
	4: "INDETERMINATED",
	5: "INDETERMINATEP",
	6: "INDETERMINATEDP",
}
var Response_Effect_value = map[string]int32{
	"DENY":            0,
	"PERMIT":          1,
	"NOTAPPLICABLE":   2,
	"INDETERMINATE":   3,
	"INDETERMINATED":  4,
	"INDETERMINATEP":  5,
	"INDETERMINATEDP": 6,
}

func (x Response_Effect) String() string {
	return proto.EnumName(Response_Effect_name, int32(x))
}
func (Response_Effect) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2, 0} }

type Attribute struct {
	Id    string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Type  string `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
	Value string `protobuf:"bytes,3,opt,name=value" json:"value,omitempty"`
}

func (m *Attribute) Reset()                    { *m = Attribute{} }
func (m *Attribute) String() string            { return proto.CompactTextString(m) }
func (*Attribute) ProtoMessage()               {}
func (*Attribute) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Attribute) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Attribute) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Attribute) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type Request struct {
	Attributes []*Attribute `protobuf:"bytes,1,rep,name=attributes" json:"attributes,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Request) GetAttributes() []*Attribute {
	if m != nil {
		return m.Attributes
	}
	return nil
}

type Response struct {
	Effect     Response_Effect `protobuf:"varint,1,opt,name=effect,enum=service.Response_Effect" json:"effect,omitempty"`
	Reason     string          `protobuf:"bytes,2,opt,name=reason" json:"reason,omitempty"`
	Obligation []*Attribute    `protobuf:"bytes,3,rep,name=obligation" json:"obligation,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Response) GetEffect() Response_Effect {
	if m != nil {
		return m.Effect
	}
	return Response_DENY
}

func (m *Response) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

func (m *Response) GetObligation() []*Attribute {
	if m != nil {
		return m.Obligation
	}
	return nil
}

func init() {
	proto.RegisterType((*Attribute)(nil), "service.Attribute")
	proto.RegisterType((*Request)(nil), "service.Request")
	proto.RegisterType((*Response)(nil), "service.Response")
	proto.RegisterEnum("service.Response_Effect", Response_Effect_name, Response_Effect_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for PDP service

type PDPClient interface {
	Validate(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	NewValidationStream(ctx context.Context, opts ...grpc.CallOption) (PDP_NewValidationStreamClient, error)
}

type pDPClient struct {
	cc *grpc.ClientConn
}

func NewPDPClient(cc *grpc.ClientConn) PDPClient {
	return &pDPClient{cc}
}

func (c *pDPClient) Validate(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/service.PDP/Validate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pDPClient) NewValidationStream(ctx context.Context, opts ...grpc.CallOption) (PDP_NewValidationStreamClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_PDP_serviceDesc.Streams[0], c.cc, "/service.PDP/NewValidationStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &pDPNewValidationStreamClient{stream}
	return x, nil
}

type PDP_NewValidationStreamClient interface {
	Send(*Request) error
	Recv() (*Response, error)
	grpc.ClientStream
}

type pDPNewValidationStreamClient struct {
	grpc.ClientStream
}

func (x *pDPNewValidationStreamClient) Send(m *Request) error {
	return x.ClientStream.SendMsg(m)
}

func (x *pDPNewValidationStreamClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for PDP service

type PDPServer interface {
	Validate(context.Context, *Request) (*Response, error)
	NewValidationStream(PDP_NewValidationStreamServer) error
}

func RegisterPDPServer(s *grpc.Server, srv PDPServer) {
	s.RegisterService(&_PDP_serviceDesc, srv)
}

func _PDP_Validate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PDPServer).Validate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.PDP/Validate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PDPServer).Validate(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _PDP_NewValidationStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PDPServer).NewValidationStream(&pDPNewValidationStreamServer{stream})
}

type PDP_NewValidationStreamServer interface {
	Send(*Response) error
	Recv() (*Request, error)
	grpc.ServerStream
}

type pDPNewValidationStreamServer struct {
	grpc.ServerStream
}

func (x *pDPNewValidationStreamServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func (x *pDPNewValidationStreamServer) Recv() (*Request, error) {
	m := new(Request)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _PDP_serviceDesc = grpc.ServiceDesc{
	ServiceName: "service.PDP",
	HandlerType: (*PDPServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Validate",
			Handler:    _PDP_Validate_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "NewValidationStream",
			Handler:       _PDP_NewValidationStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "service.proto",
}

func init() { proto.RegisterFile("service.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 337 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xc1, 0x6b, 0xea, 0x40,
	0x10, 0xc6, 0x4d, 0xa2, 0x51, 0xe7, 0xa1, 0x6f, 0x1d, 0x1f, 0x8f, 0xd0, 0x93, 0xe4, 0xe4, 0x49,
	0x6c, 0x7a, 0x2e, 0x34, 0x6d, 0x72, 0x10, 0x6c, 0x1a, 0xd2, 0x50, 0xe8, 0x71, 0xd5, 0xb1, 0x2c,
	0xd8, 0xac, 0x4d, 0x56, 0x4b, 0x7b, 0xeb, 0xbd, 0x7f, 0x74, 0x71, 0x8d, 0x56, 0x91, 0x42, 0x6f,
	0x99, 0x2f, 0xbf, 0x6f, 0xe6, 0xe3, 0x63, 0xa1, 0x55, 0x50, 0xbe, 0x16, 0x53, 0x1a, 0x2c, 0x73,
	0xa9, 0x24, 0xd6, 0xcb, 0xd1, 0x0d, 0xa1, 0xe9, 0x2b, 0x95, 0x8b, 0xc9, 0x4a, 0x11, 0xb6, 0xc1,
	0x14, 0x33, 0xc7, 0xe8, 0x19, 0xfd, 0x66, 0x62, 0x8a, 0x19, 0x22, 0x54, 0xd5, 0xdb, 0x92, 0x1c,
	0x53, 0x2b, 0xfa, 0x1b, 0xff, 0x41, 0x6d, 0xcd, 0x17, 0x2b, 0x72, 0x2c, 0x2d, 0x6e, 0x07, 0xf7,
	0x12, 0xea, 0x09, 0xbd, 0xac, 0xa8, 0x50, 0xe8, 0x01, 0xf0, 0xdd, 0xc6, 0xc2, 0x31, 0x7a, 0x56,
	0xff, 0x8f, 0x87, 0x83, 0xdd, 0xf9, 0xfd, 0xb1, 0xe4, 0x80, 0x72, 0x3f, 0x4d, 0x68, 0x24, 0x54,
	0x2c, 0x65, 0x56, 0x10, 0x0e, 0xc1, 0xa6, 0xf9, 0x9c, 0xa6, 0x4a, 0x27, 0x69, 0x7b, 0xce, 0xde,
	0xbc, 0x43, 0x06, 0xa1, 0xfe, 0x9f, 0x94, 0x1c, 0xfe, 0x07, 0x3b, 0x27, 0x5e, 0xc8, 0xac, 0x4c,
	0x5a, 0x4e, 0x9b, 0x28, 0x72, 0xb2, 0x10, 0x4f, 0x5c, 0x09, 0x99, 0x39, 0xd6, 0xcf, 0x51, 0xbe,
	0x29, 0xf7, 0xc3, 0x00, 0x7b, 0xbb, 0x1e, 0x1b, 0x50, 0x0d, 0xc2, 0xe8, 0x91, 0x55, 0x10, 0xc0,
	0x8e, 0xc3, 0xe4, 0x76, 0x94, 0x32, 0x03, 0x3b, 0xd0, 0x8a, 0xee, 0x52, 0x3f, 0x8e, 0xc7, 0xa3,
	0x1b, 0xff, 0x7a, 0x1c, 0x32, 0x73, 0x23, 0x8d, 0xa2, 0x20, 0x4c, 0x37, 0x48, 0xe4, 0xa7, 0x21,
	0xb3, 0x10, 0xa1, 0x7d, 0x24, 0x05, 0xac, 0x7a, 0xa2, 0xc5, 0xac, 0x86, 0x5d, 0xf8, 0x7b, 0xcc,
	0xc5, 0xcc, 0xf6, 0xde, 0xc1, 0x8a, 0x83, 0x18, 0xcf, 0xa1, 0xf1, 0xc0, 0x17, 0x62, 0xc6, 0x15,
	0x21, 0x3b, 0x28, 0x41, 0xf7, 0x7c, 0xd6, 0x39, 0xa9, 0xc5, 0xad, 0xe0, 0x15, 0x74, 0x23, 0x7a,
	0x2d, 0x5d, 0x42, 0x66, 0xf7, 0x2a, 0x27, 0xfe, 0xfc, 0x4b, 0x77, 0xdf, 0x18, 0x1a, 0x13, 0x5b,
	0x3f, 0x90, 0x8b, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x71, 0xa4, 0xc0, 0xf5, 0x31, 0x02, 0x00,
	0x00,
}
