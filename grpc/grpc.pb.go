// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grpc.proto

package grpc

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

// The request message containing the user's name.
type CustomerContactRequest struct {
	Data                 string   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CustomerContactRequest) Reset()         { *m = CustomerContactRequest{} }
func (m *CustomerContactRequest) String() string { return proto.CompactTextString(m) }
func (*CustomerContactRequest) ProtoMessage()    {}
func (*CustomerContactRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bedfbfc9b54e5600, []int{0}
}

func (m *CustomerContactRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CustomerContactRequest.Unmarshal(m, b)
}
func (m *CustomerContactRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CustomerContactRequest.Marshal(b, m, deterministic)
}
func (m *CustomerContactRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustomerContactRequest.Merge(m, src)
}
func (m *CustomerContactRequest) XXX_Size() int {
	return xxx_messageInfo_CustomerContactRequest.Size(m)
}
func (m *CustomerContactRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CustomerContactRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CustomerContactRequest proto.InternalMessageInfo

func (m *CustomerContactRequest) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

// The response message containing the greetings
type CustomerContactReply struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CustomerContactReply) Reset()         { *m = CustomerContactReply{} }
func (m *CustomerContactReply) String() string { return proto.CompactTextString(m) }
func (*CustomerContactReply) ProtoMessage()    {}
func (*CustomerContactReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_bedfbfc9b54e5600, []int{1}
}

func (m *CustomerContactReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CustomerContactReply.Unmarshal(m, b)
}
func (m *CustomerContactReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CustomerContactReply.Marshal(b, m, deterministic)
}
func (m *CustomerContactReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustomerContactReply.Merge(m, src)
}
func (m *CustomerContactReply) XXX_Size() int {
	return xxx_messageInfo_CustomerContactReply.Size(m)
}
func (m *CustomerContactReply) XXX_DiscardUnknown() {
	xxx_messageInfo_CustomerContactReply.DiscardUnknown(m)
}

var xxx_messageInfo_CustomerContactReply proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CustomerContactRequest)(nil), "grpc.CustomerContactRequest")
	proto.RegisterType((*CustomerContactReply)(nil), "grpc.CustomerContactReply")
}

func init() { proto.RegisterFile("grpc.proto", fileDescriptor_bedfbfc9b54e5600) }

var fileDescriptor_bedfbfc9b54e5600 = []byte{
	// 142 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x2f, 0x2a, 0x48,
	0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0xb1, 0x95, 0x74, 0xb8, 0xc4, 0x9c, 0x4b,
	0x8b, 0x4b, 0xf2, 0x73, 0x53, 0x8b, 0x9c, 0xf3, 0xf3, 0x4a, 0x12, 0x93, 0x4b, 0x82, 0x52, 0x0b,
	0x4b, 0x53, 0x8b, 0x4b, 0x84, 0x84, 0xb8, 0x58, 0x52, 0x12, 0x4b, 0x12, 0x25, 0x18, 0x15, 0x18,
	0x35, 0x38, 0x83, 0xc0, 0x6c, 0x25, 0x31, 0x2e, 0x11, 0x0c, 0xd5, 0x05, 0x39, 0x95, 0x46, 0x45,
	0x18, 0xa6, 0x04, 0xa7, 0x16, 0x95, 0x65, 0x26, 0xa7, 0x0a, 0x45, 0x70, 0x49, 0x39, 0xa6, 0xa4,
	0xf8, 0x17, 0x85, 0x16, 0xa4, 0x24, 0x96, 0xa4, 0xa2, 0x29, 0x12, 0x92, 0xd1, 0x03, 0x3b, 0x08,
	0xbb, 0x0b, 0xa4, 0xa4, 0x70, 0xc8, 0x16, 0xe4, 0x54, 0x2a, 0x31, 0x24, 0xb1, 0x81, 0xbd, 0x61,
	0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x32, 0x0c, 0x1e, 0xb3, 0xd4, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CustomerContactServiceClient is the client API for CustomerContactService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CustomerContactServiceClient interface {
	AddOrUpdateCustomerContact(ctx context.Context, in *CustomerContactRequest, opts ...grpc.CallOption) (*CustomerContactReply, error)
}

type customerContactServiceClient struct {
	cc *grpc.ClientConn
}

func NewCustomerContactServiceClient(cc *grpc.ClientConn) CustomerContactServiceClient {
	return &customerContactServiceClient{cc}
}

func (c *customerContactServiceClient) AddOrUpdateCustomerContact(ctx context.Context, in *CustomerContactRequest, opts ...grpc.CallOption) (*CustomerContactReply, error) {
	out := new(CustomerContactReply)
	err := c.cc.Invoke(ctx, "/grpc.CustomerContactService/AddOrUpdateCustomerContact", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CustomerContactServiceServer is the server API for CustomerContactService service.
type CustomerContactServiceServer interface {
	AddOrUpdateCustomerContact(context.Context, *CustomerContactRequest) (*CustomerContactReply, error)
}

func RegisterCustomerContactServiceServer(s *grpc.Server, srv CustomerContactServiceServer) {
	s.RegisterService(&_CustomerContactService_serviceDesc, srv)
}

func _CustomerContactService_AddOrUpdateCustomerContact_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CustomerContactRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerContactServiceServer).AddOrUpdateCustomerContact(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.CustomerContactService/AddOrUpdateCustomerContact",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerContactServiceServer).AddOrUpdateCustomerContact(ctx, req.(*CustomerContactRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CustomerContactService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.CustomerContactService",
	HandlerType: (*CustomerContactServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddOrUpdateCustomerContact",
			Handler:    _CustomerContactService_AddOrUpdateCustomerContact_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc.proto",
}
