// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api.proto

package api

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

type OperationType int32

const (
	OperationType_INSERT OperationType = 0
	OperationType_DELETE OperationType = 1
)

var OperationType_name = map[int32]string{
	0: "INSERT",
	1: "DELETE",
}

var OperationType_value = map[string]int32{
	"INSERT": 0,
	"DELETE": 1,
}

func (x OperationType) String() string {
	return proto.EnumName(OperationType_name, int32(x))
}

func (OperationType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{0}
}

type EchoRequest struct {
	Body                 string   `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EchoRequest) Reset()         { *m = EchoRequest{} }
func (m *EchoRequest) String() string { return proto.CompactTextString(m) }
func (*EchoRequest) ProtoMessage()    {}
func (*EchoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{0}
}

func (m *EchoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EchoRequest.Unmarshal(m, b)
}
func (m *EchoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EchoRequest.Marshal(b, m, deterministic)
}
func (m *EchoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EchoRequest.Merge(m, src)
}
func (m *EchoRequest) XXX_Size() int {
	return xxx_messageInfo_EchoRequest.Size(m)
}
func (m *EchoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EchoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EchoRequest proto.InternalMessageInfo

func (m *EchoRequest) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

type EchoReply struct {
	Body                 string   `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EchoReply) Reset()         { *m = EchoReply{} }
func (m *EchoReply) String() string { return proto.CompactTextString(m) }
func (*EchoReply) ProtoMessage()    {}
func (*EchoReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{1}
}

func (m *EchoReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EchoReply.Unmarshal(m, b)
}
func (m *EchoReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EchoReply.Marshal(b, m, deterministic)
}
func (m *EchoReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EchoReply.Merge(m, src)
}
func (m *EchoReply) XXX_Size() int {
	return xxx_messageInfo_EchoReply.Size(m)
}
func (m *EchoReply) XXX_DiscardUnknown() {
	xxx_messageInfo_EchoReply.DiscardUnknown(m)
}

var xxx_messageInfo_EchoReply proto.InternalMessageInfo

func (m *EchoReply) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

// Wchar is W-character c
type Wchar struct {
	// Id is the id of c
	Id *Wid `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Alpha is alphabetical value of the effective character of c
	CodePoint int32 `protobuf:"varint,2,opt,name=code_point,json=codePoint,proto3" json:"code_point,omitempty"`
	// Visible is {true | false}, if the character c is visible
	Visible bool `protobuf:"varint,3,opt,name=visible,proto3" json:"visible,omitempty"`
	// PreviousId is the id of the previous W-character of c
	PreviousId *Wid `protobuf:"bytes,4,opt,name=previousId,proto3" json:"previousId,omitempty"`
	// NextId is the id of the next W-character of c
	NextId               *Wid     `protobuf:"bytes,5,opt,name=nextId,proto3" json:"nextId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Wchar) Reset()         { *m = Wchar{} }
func (m *Wchar) String() string { return proto.CompactTextString(m) }
func (*Wchar) ProtoMessage()    {}
func (*Wchar) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{2}
}

func (m *Wchar) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Wchar.Unmarshal(m, b)
}
func (m *Wchar) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Wchar.Marshal(b, m, deterministic)
}
func (m *Wchar) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Wchar.Merge(m, src)
}
func (m *Wchar) XXX_Size() int {
	return xxx_messageInfo_Wchar.Size(m)
}
func (m *Wchar) XXX_DiscardUnknown() {
	xxx_messageInfo_Wchar.DiscardUnknown(m)
}

var xxx_messageInfo_Wchar proto.InternalMessageInfo

func (m *Wchar) GetId() *Wid {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *Wchar) GetCodePoint() int32 {
	if m != nil {
		return m.CodePoint
	}
	return 0
}

func (m *Wchar) GetVisible() bool {
	if m != nil {
		return m.Visible
	}
	return false
}

func (m *Wchar) GetPreviousId() *Wid {
	if m != nil {
		return m.PreviousId
	}
	return nil
}

func (m *Wchar) GetNextId() *Wid {
	if m != nil {
		return m.NextId
	}
	return nil
}

// Wid is the id of W-character
type Wid struct {
	// The identifier of a site (a peer)
	Ns string `protobuf:"bytes,1,opt,name=ns,proto3" json:"ns,omitempty"`
	// The local clock of the W-character is generated on a site
	Ng                   int64    `protobuf:"varint,2,opt,name=ng,proto3" json:"ng,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Wid) Reset()         { *m = Wid{} }
func (m *Wid) String() string { return proto.CompactTextString(m) }
func (*Wid) ProtoMessage()    {}
func (*Wid) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{3}
}

func (m *Wid) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Wid.Unmarshal(m, b)
}
func (m *Wid) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Wid.Marshal(b, m, deterministic)
}
func (m *Wid) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Wid.Merge(m, src)
}
func (m *Wid) XXX_Size() int {
	return xxx_messageInfo_Wid.Size(m)
}
func (m *Wid) XXX_DiscardUnknown() {
	xxx_messageInfo_Wid.DiscardUnknown(m)
}

var xxx_messageInfo_Wid proto.InternalMessageInfo

func (m *Wid) GetNs() string {
	if m != nil {
		return m.Ns
	}
	return ""
}

func (m *Wid) GetNg() int64 {
	if m != nil {
		return m.Ng
	}
	return 0
}

type Operation struct {
	Type                 OperationType `protobuf:"varint,1,opt,name=type,proto3,enum=api.OperationType" json:"type,omitempty"`
	C                    *Wchar        `protobuf:"bytes,2,opt,name=c,proto3" json:"c,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Operation) Reset()         { *m = Operation{} }
func (m *Operation) String() string { return proto.CompactTextString(m) }
func (*Operation) ProtoMessage()    {}
func (*Operation) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{4}
}

func (m *Operation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Operation.Unmarshal(m, b)
}
func (m *Operation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Operation.Marshal(b, m, deterministic)
}
func (m *Operation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Operation.Merge(m, src)
}
func (m *Operation) XXX_Size() int {
	return xxx_messageInfo_Operation.Size(m)
}
func (m *Operation) XXX_DiscardUnknown() {
	xxx_messageInfo_Operation.DiscardUnknown(m)
}

var xxx_messageInfo_Operation proto.InternalMessageInfo

func (m *Operation) GetType() OperationType {
	if m != nil {
		return m.Type
	}
	return OperationType_INSERT
}

func (m *Operation) GetC() *Wchar {
	if m != nil {
		return m.C
	}
	return nil
}

func init() {
	proto.RegisterEnum("api.OperationType", OperationType_name, OperationType_value)
	proto.RegisterType((*EchoRequest)(nil), "api.EchoRequest")
	proto.RegisterType((*EchoReply)(nil), "api.EchoReply")
	proto.RegisterType((*Wchar)(nil), "api.Wchar")
	proto.RegisterType((*Wid)(nil), "api.Wid")
	proto.RegisterType((*Operation)(nil), "api.Operation")
}

func init() { proto.RegisterFile("api.proto", fileDescriptor_00212fb1f9d3bf1c) }

var fileDescriptor_00212fb1f9d3bf1c = []byte{
	// 315 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x51, 0x5d, 0x4b, 0xc3, 0x30,
	0x14, 0x5d, 0xda, 0x6e, 0xae, 0x77, 0x38, 0xc6, 0x7d, 0x0a, 0x82, 0x58, 0x0b, 0x6a, 0xd9, 0xc3,
	0x1e, 0xea, 0x5f, 0xb0, 0x0f, 0x03, 0xbf, 0x88, 0x83, 0x3e, 0x4a, 0xd7, 0x84, 0x2d, 0x30, 0x92,
	0xd8, 0x66, 0xc3, 0xfe, 0x21, 0x7f, 0xa7, 0x24, 0x9b, 0x32, 0x87, 0x6f, 0xf7, 0x9e, 0x73, 0x72,
	0x6e, 0xee, 0x3d, 0x10, 0x57, 0x46, 0xce, 0x4c, 0xa3, 0xad, 0xc6, 0xb0, 0x32, 0x32, 0xbd, 0x86,
	0x51, 0x51, 0xaf, 0x35, 0x13, 0x1f, 0x5b, 0xd1, 0x5a, 0x44, 0x88, 0x96, 0x9a, 0x77, 0x94, 0x24,
	0x24, 0x8b, 0x99, 0xaf, 0xd3, 0x2b, 0x88, 0xf7, 0x12, 0xb3, 0xe9, 0xfe, 0x15, 0x7c, 0x11, 0xe8,
	0x97, 0xf5, 0xba, 0x6a, 0x90, 0x42, 0x20, 0xb9, 0xe7, 0x46, 0xf9, 0x70, 0xe6, 0x46, 0x95, 0x92,
	0xb3, 0x40, 0x72, 0xbc, 0x04, 0xa8, 0x35, 0x17, 0xef, 0x46, 0x4b, 0x65, 0x69, 0x90, 0x90, 0xac,
	0xcf, 0x62, 0x87, 0xbc, 0x3a, 0x00, 0x29, 0x9c, 0xed, 0x64, 0x2b, 0x97, 0x1b, 0x41, 0xc3, 0x84,
	0x64, 0x43, 0xf6, 0xd3, 0x62, 0x06, 0x60, 0x1a, 0xb1, 0x93, 0x7a, 0xdb, 0xce, 0x39, 0x8d, 0x4e,
	0xac, 0x8f, 0x38, 0x4c, 0x60, 0xa0, 0xc4, 0xa7, 0x9d, 0x73, 0xda, 0x3f, 0x51, 0x1d, 0xf0, 0xf4,
	0x06, 0xc2, 0x52, 0x72, 0x1c, 0x43, 0xa0, 0xda, 0xc3, 0x06, 0x81, 0x6a, 0x7d, 0xbf, 0xf2, 0x7f,
	0x0a, 0x59, 0xa0, 0x56, 0xe9, 0x13, 0xc4, 0x2f, 0x46, 0x34, 0x95, 0x95, 0x5a, 0xe1, 0x2d, 0x44,
	0xb6, 0x33, 0xc2, 0xcb, 0xc7, 0x39, 0x7a, 0xcf, 0x5f, 0x76, 0xd1, 0x19, 0xc1, 0x3c, 0x8f, 0x14,
	0x48, 0xed, 0x3d, 0x46, 0x39, 0xec, 0x07, 0xbb, 0x8b, 0x30, 0x52, 0x4f, 0xef, 0xe0, 0xfc, 0xcf,
	0x03, 0x04, 0x18, 0xcc, 0x9f, 0xdf, 0x0a, 0xb6, 0x98, 0xf4, 0x5c, 0xfd, 0x50, 0x3c, 0x16, 0x8b,
	0x62, 0x42, 0xf2, 0x1c, 0xa2, 0x52, 0x6b, 0x8b, 0x53, 0x88, 0xdc, 0xc1, 0x71, 0xe2, 0x7d, 0x8e,
	0xe2, 0xb9, 0x18, 0x1f, 0x21, 0x66, 0xd3, 0xa5, 0xbd, 0xe5, 0xc0, 0x67, 0x79, 0xff, 0x1d, 0x00,
	0x00, 0xff, 0xff, 0x11, 0x65, 0xcd, 0xee, 0xd8, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// WootClient is the client API for Woot service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type WootClient interface {
	Echo(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (*EchoReply, error)
}

type wootClient struct {
	cc grpc.ClientConnInterface
}

func NewWootClient(cc grpc.ClientConnInterface) WootClient {
	return &wootClient{cc}
}

func (c *wootClient) Echo(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (*EchoReply, error) {
	out := new(EchoReply)
	err := c.cc.Invoke(ctx, "/api.Woot/Echo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WootServer is the server API for Woot service.
type WootServer interface {
	Echo(context.Context, *EchoRequest) (*EchoReply, error)
}

// UnimplementedWootServer can be embedded to have forward compatible implementations.
type UnimplementedWootServer struct {
}

func (*UnimplementedWootServer) Echo(ctx context.Context, req *EchoRequest) (*EchoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Echo not implemented")
}

func RegisterWootServer(s *grpc.Server, srv WootServer) {
	s.RegisterService(&_Woot_serviceDesc, srv)
}

func _Woot_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EchoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WootServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Woot/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WootServer).Echo(ctx, req.(*EchoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Woot_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Woot",
	HandlerType: (*WootServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Echo",
			Handler:    _Woot_Echo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}
