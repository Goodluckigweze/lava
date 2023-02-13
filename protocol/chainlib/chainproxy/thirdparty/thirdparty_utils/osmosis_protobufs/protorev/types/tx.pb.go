// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: osmosis/protorev/v1beta1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// MsgSetHotRoutes defines the Msg/SetHotRoutes request type.
type MsgSetHotRoutes struct {
	// admin is the account that is authorized to set the hot routes.
	Admin string `protobuf:"bytes,1,opt,name=admin,proto3" json:"admin,omitempty"`
	// hot_routes is the list of hot routes to set.
	HotRoutes []*TokenPairArbRoutes `protobuf:"bytes,2,rep,name=hot_routes,json=hotRoutes,proto3" json:"hot_routes,omitempty"`
}

func (m *MsgSetHotRoutes) Reset()         { *m = MsgSetHotRoutes{} }
func (m *MsgSetHotRoutes) String() string { return proto.CompactTextString(m) }
func (*MsgSetHotRoutes) ProtoMessage()    {}
func (*MsgSetHotRoutes) Descriptor() ([]byte, []int) {
	return fileDescriptor_2783dce032fc6954, []int{0}
}
func (m *MsgSetHotRoutes) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSetHotRoutes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSetHotRoutes.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSetHotRoutes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSetHotRoutes.Merge(m, src)
}
func (m *MsgSetHotRoutes) XXX_Size() int {
	return m.Size()
}
func (m *MsgSetHotRoutes) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSetHotRoutes.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSetHotRoutes proto.InternalMessageInfo

func (m *MsgSetHotRoutes) GetAdmin() string {
	if m != nil {
		return m.Admin
	}
	return ""
}

func (m *MsgSetHotRoutes) GetHotRoutes() []*TokenPairArbRoutes {
	if m != nil {
		return m.HotRoutes
	}
	return nil
}

// MsgSetHotRoutesResponse defines the Msg/SetHotRoutes response type.
type MsgSetHotRoutesResponse struct {
}

func (m *MsgSetHotRoutesResponse) Reset()         { *m = MsgSetHotRoutesResponse{} }
func (m *MsgSetHotRoutesResponse) String() string { return proto.CompactTextString(m) }
func (*MsgSetHotRoutesResponse) ProtoMessage()    {}
func (*MsgSetHotRoutesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2783dce032fc6954, []int{1}
}
func (m *MsgSetHotRoutesResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSetHotRoutesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSetHotRoutesResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSetHotRoutesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSetHotRoutesResponse.Merge(m, src)
}
func (m *MsgSetHotRoutesResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgSetHotRoutesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSetHotRoutesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSetHotRoutesResponse proto.InternalMessageInfo

// MsgSetDeveloperAccount defines the Msg/SetDeveloperAccount request type.
type MsgSetDeveloperAccount struct {
	// admin is the account that is authorized to set the developer account.
	Admin string `protobuf:"bytes,1,opt,name=admin,proto3" json:"admin,omitempty"`
	// developer_account is the account that will receive a portion of the profits
	// from the protorev module.
	DeveloperAccount string `protobuf:"bytes,2,opt,name=developer_account,json=developerAccount,proto3" json:"developer_account,omitempty"`
}

func (m *MsgSetDeveloperAccount) Reset()         { *m = MsgSetDeveloperAccount{} }
func (m *MsgSetDeveloperAccount) String() string { return proto.CompactTextString(m) }
func (*MsgSetDeveloperAccount) ProtoMessage()    {}
func (*MsgSetDeveloperAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_2783dce032fc6954, []int{2}
}
func (m *MsgSetDeveloperAccount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSetDeveloperAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSetDeveloperAccount.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSetDeveloperAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSetDeveloperAccount.Merge(m, src)
}
func (m *MsgSetDeveloperAccount) XXX_Size() int {
	return m.Size()
}
func (m *MsgSetDeveloperAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSetDeveloperAccount.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSetDeveloperAccount proto.InternalMessageInfo

func (m *MsgSetDeveloperAccount) GetAdmin() string {
	if m != nil {
		return m.Admin
	}
	return ""
}

func (m *MsgSetDeveloperAccount) GetDeveloperAccount() string {
	if m != nil {
		return m.DeveloperAccount
	}
	return ""
}

// MsgSetDeveloperAccountResponse defines the Msg/SetDeveloperAccount response
// type.
type MsgSetDeveloperAccountResponse struct {
}

func (m *MsgSetDeveloperAccountResponse) Reset()         { *m = MsgSetDeveloperAccountResponse{} }
func (m *MsgSetDeveloperAccountResponse) String() string { return proto.CompactTextString(m) }
func (*MsgSetDeveloperAccountResponse) ProtoMessage()    {}
func (*MsgSetDeveloperAccountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2783dce032fc6954, []int{3}
}
func (m *MsgSetDeveloperAccountResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSetDeveloperAccountResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSetDeveloperAccountResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSetDeveloperAccountResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSetDeveloperAccountResponse.Merge(m, src)
}
func (m *MsgSetDeveloperAccountResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgSetDeveloperAccountResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSetDeveloperAccountResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSetDeveloperAccountResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgSetHotRoutes)(nil), "osmosis.protorev.v1beta1.MsgSetHotRoutes")
	proto.RegisterType((*MsgSetHotRoutesResponse)(nil), "osmosis.protorev.v1beta1.MsgSetHotRoutesResponse")
	proto.RegisterType((*MsgSetDeveloperAccount)(nil), "osmosis.protorev.v1beta1.MsgSetDeveloperAccount")
	proto.RegisterType((*MsgSetDeveloperAccountResponse)(nil), "osmosis.protorev.v1beta1.MsgSetDeveloperAccountResponse")
}

func init() { proto.RegisterFile("osmosis/protorev/v1beta1/tx.proto", fileDescriptor_2783dce032fc6954) }

var fileDescriptor_2783dce032fc6954 = []byte{
	// 404 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0xcf, 0x4a, 0xe3, 0x50,
	0x14, 0xc6, 0x9b, 0x94, 0x19, 0xe8, 0x9d, 0x81, 0x99, 0xc9, 0x14, 0x8d, 0x41, 0x42, 0x0d, 0xa2,
	0x95, 0xb6, 0xb9, 0xa6, 0xed, 0xc2, 0x6d, 0xc5, 0x85, 0x20, 0x15, 0x89, 0xae, 0x74, 0x51, 0x6e,
	0xda, 0x4b, 0x1a, 0x6c, 0x73, 0x42, 0xee, 0x6d, 0xa8, 0x5b, 0x9f, 0x40, 0x70, 0xe7, 0x83, 0xf8,
	0x0c, 0x2e, 0x0b, 0x6e, 0x04, 0x37, 0xd2, 0xfa, 0x20, 0x62, 0xd2, 0xb4, 0x1a, 0x1b, 0xd4, 0xdd,
	0xfd, 0xf3, 0x3b, 0xdf, 0xf7, 0x9d, 0x73, 0xd0, 0x1a, 0xb0, 0x3e, 0x30, 0x87, 0x61, 0xcf, 0x07,
	0x0e, 0x3e, 0x0d, 0x70, 0x60, 0x58, 0x94, 0x13, 0x03, 0xf3, 0xa1, 0x1e, 0xbe, 0x49, 0xf2, 0x14,
	0xd1, 0x63, 0x44, 0x9f, 0x22, 0x4a, 0xde, 0x06, 0x1b, 0xc2, 0x57, 0xfc, 0x7a, 0x8a, 0x00, 0x65,
	0xd5, 0x06, 0xb0, 0x7b, 0x14, 0x13, 0xcf, 0xc1, 0xc4, 0x75, 0x81, 0x13, 0xee, 0x80, 0x3b, 0x2d,
	0x57, 0x36, 0x53, 0x0d, 0x67, 0xf2, 0xe1, 0x41, 0xe3, 0xe8, 0x4f, 0x93, 0xd9, 0xc7, 0x94, 0xef,
	0x03, 0x37, 0x61, 0xc0, 0x29, 0x93, 0xf2, 0xe8, 0x07, 0xe9, 0xf4, 0x1d, 0x57, 0x16, 0x0a, 0x42,
	0x31, 0x67, 0x46, 0x17, 0xe9, 0x00, 0xa1, 0x2e, 0xf0, 0x96, 0x1f, 0x32, 0xb2, 0x58, 0xc8, 0x16,
	0x7f, 0x55, 0xcb, 0x7a, 0x5a, 0x68, 0xfd, 0x04, 0xce, 0xa9, 0x7b, 0x44, 0x1c, 0xbf, 0xe1, 0x5b,
	0x91, 0xae, 0x99, 0xeb, 0xc6, 0x16, 0xda, 0x0a, 0x5a, 0x4e, 0xb8, 0x9a, 0x94, 0x79, 0xe0, 0x32,
	0xaa, 0x9d, 0xa1, 0xa5, 0xe8, 0x6b, 0x8f, 0x06, 0xb4, 0x07, 0x1e, 0xf5, 0x1b, 0xed, 0x36, 0x0c,
	0x5c, 0x9e, 0x92, 0xab, 0x84, 0xfe, 0x75, 0x62, 0xb2, 0x45, 0x22, 0x54, 0x16, 0x43, 0xe2, 0x6f,
	0x27, 0x21, 0xa1, 0x15, 0x90, 0xba, 0x58, 0x3c, 0xb6, 0xaf, 0x3e, 0x8a, 0x28, 0xdb, 0x64, 0xb6,
	0x74, 0x23, 0xa0, 0xdf, 0xef, 0xa6, 0xb2, 0x95, 0xde, 0x6b, 0xa2, 0x15, 0xc5, 0xf8, 0x32, 0x3a,
	0xeb, 0xba, 0x7c, 0x79, 0xff, 0x7c, 0x2d, 0x6e, 0x68, 0xeb, 0x38, 0x5e, 0x5c, 0x60, 0xd4, 0xe6,
	0xcb, 0x63, 0x94, 0xb7, 0xe6, 0xd3, 0x97, 0x6e, 0x05, 0xf4, 0x7f, 0xd1, 0x84, 0xb6, 0x3f, 0x33,
	0x4e, 0x56, 0x28, 0x3b, 0xdf, 0xad, 0x98, 0x25, 0xae, 0x85, 0x89, 0x2b, 0x5a, 0x29, 0x3d, 0xf1,
	0x87, 0xbd, 0xec, 0x1e, 0xde, 0x8d, 0x55, 0x61, 0x34, 0x56, 0x85, 0xa7, 0xb1, 0x2a, 0x5c, 0x4d,
	0xd4, 0xcc, 0x68, 0xa2, 0x66, 0x1e, 0x26, 0x6a, 0xe6, 0xb4, 0x6e, 0x3b, 0xbc, 0x3b, 0xb0, 0xf4,
	0x36, 0xf4, 0x63, 0xc1, 0x4a, 0x8f, 0x58, 0xec, 0x8d, 0x7a, 0x1d, 0x0f, 0xe7, 0xfa, 0xfc, 0xc2,
	0xa3, 0xcc, 0xfa, 0x19, 0xde, 0x6b, 0x2f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x30, 0x42, 0x52, 0xeb,
	0x60, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// SetHotRoutes sets the hot routes that will be explored when creating
	// cyclic arbitrage routes. Can only be called by the admin account.
	SetHotRoutes(ctx context.Context, in *MsgSetHotRoutes, opts ...grpc.CallOption) (*MsgSetHotRoutesResponse, error)
	// SetDeveloperAccount sets the account that can withdraw a portion of the
	// profits from the protorev module. This will be Skip's address.
	SetDeveloperAccount(ctx context.Context, in *MsgSetDeveloperAccount, opts ...grpc.CallOption) (*MsgSetDeveloperAccountResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) SetHotRoutes(ctx context.Context, in *MsgSetHotRoutes, opts ...grpc.CallOption) (*MsgSetHotRoutesResponse, error) {
	out := new(MsgSetHotRoutesResponse)
	err := c.cc.Invoke(ctx, "/osmosis.protorev.v1beta1.Msg/SetHotRoutes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) SetDeveloperAccount(ctx context.Context, in *MsgSetDeveloperAccount, opts ...grpc.CallOption) (*MsgSetDeveloperAccountResponse, error) {
	out := new(MsgSetDeveloperAccountResponse)
	err := c.cc.Invoke(ctx, "/osmosis.protorev.v1beta1.Msg/SetDeveloperAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// SetHotRoutes sets the hot routes that will be explored when creating
	// cyclic arbitrage routes. Can only be called by the admin account.
	SetHotRoutes(context.Context, *MsgSetHotRoutes) (*MsgSetHotRoutesResponse, error)
	// SetDeveloperAccount sets the account that can withdraw a portion of the
	// profits from the protorev module. This will be Skip's address.
	SetDeveloperAccount(context.Context, *MsgSetDeveloperAccount) (*MsgSetDeveloperAccountResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) SetHotRoutes(ctx context.Context, req *MsgSetHotRoutes) (*MsgSetHotRoutesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetHotRoutes not implemented")
}
func (*UnimplementedMsgServer) SetDeveloperAccount(ctx context.Context, req *MsgSetDeveloperAccount) (*MsgSetDeveloperAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetDeveloperAccount not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_SetHotRoutes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSetHotRoutes)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SetHotRoutes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/osmosis.protorev.v1beta1.Msg/SetHotRoutes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SetHotRoutes(ctx, req.(*MsgSetHotRoutes))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_SetDeveloperAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSetDeveloperAccount)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SetDeveloperAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/osmosis.protorev.v1beta1.Msg/SetDeveloperAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SetDeveloperAccount(ctx, req.(*MsgSetDeveloperAccount))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "osmosis.protorev.v1beta1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetHotRoutes",
			Handler:    _Msg_SetHotRoutes_Handler,
		},
		{
			MethodName: "SetDeveloperAccount",
			Handler:    _Msg_SetDeveloperAccount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "osmosis/protorev/v1beta1/tx.proto",
}

func (m *MsgSetHotRoutes) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSetHotRoutes) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSetHotRoutes) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.HotRoutes) > 0 {
		for iNdEx := len(m.HotRoutes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.HotRoutes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Admin) > 0 {
		i -= len(m.Admin)
		copy(dAtA[i:], m.Admin)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Admin)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgSetHotRoutesResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSetHotRoutesResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSetHotRoutesResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgSetDeveloperAccount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSetDeveloperAccount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSetDeveloperAccount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.DeveloperAccount) > 0 {
		i -= len(m.DeveloperAccount)
		copy(dAtA[i:], m.DeveloperAccount)
		i = encodeVarintTx(dAtA, i, uint64(len(m.DeveloperAccount)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Admin) > 0 {
		i -= len(m.Admin)
		copy(dAtA[i:], m.Admin)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Admin)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgSetDeveloperAccountResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSetDeveloperAccountResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSetDeveloperAccountResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgSetHotRoutes) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Admin)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if len(m.HotRoutes) > 0 {
		for _, e := range m.HotRoutes {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	return n
}

func (m *MsgSetHotRoutesResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgSetDeveloperAccount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Admin)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.DeveloperAccount)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgSetDeveloperAccountResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgSetHotRoutes) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgSetHotRoutes: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSetHotRoutes: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Admin", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Admin = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HotRoutes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.HotRoutes = append(m.HotRoutes, &TokenPairArbRoutes{})
			if err := m.HotRoutes[len(m.HotRoutes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgSetHotRoutesResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgSetHotRoutesResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSetHotRoutesResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgSetDeveloperAccount) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgSetDeveloperAccount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSetDeveloperAccount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Admin", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Admin = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeveloperAccount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DeveloperAccount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgSetDeveloperAccountResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgSetDeveloperAccountResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSetDeveloperAccountResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
