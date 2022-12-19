// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/authz/v1beta1/tx.proto

package authz

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	types "github.com/cosmos/cosmos-sdk/codec/types"
	_ "github.com/cosmos/cosmos-sdk/types/msgservice"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
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

// MsgGrant is a request type for Grant method. It declares authorization to the grantee
// on behalf of the granter with the provided expiration time.
type MsgGrant struct {
	Granter string `protobuf:"bytes,1,opt,name=granter,proto3" json:"granter,omitempty"`
	Grantee string `protobuf:"bytes,2,opt,name=grantee,proto3" json:"grantee,omitempty"`
	Grant   Grant  `protobuf:"bytes,3,opt,name=grant,proto3" json:"grant"`
}

func (m *MsgGrant) Reset()         { *m = MsgGrant{} }
func (m *MsgGrant) String() string { return proto.CompactTextString(m) }
func (*MsgGrant) ProtoMessage()    {}
func (*MsgGrant) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ceddab7d8589ad1, []int{0}
}
func (m *MsgGrant) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgGrant) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgGrant.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgGrant) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgGrant.Merge(m, src)
}
func (m *MsgGrant) XXX_Size() int {
	return m.Size()
}
func (m *MsgGrant) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgGrant.DiscardUnknown(m)
}

var xxx_messageInfo_MsgGrant proto.InternalMessageInfo

// MsgExecResponse defines the Msg/MsgExecResponse response type.
type MsgExecResponse struct {
	Results [][]byte `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
}

func (m *MsgExecResponse) Reset()         { *m = MsgExecResponse{} }
func (m *MsgExecResponse) String() string { return proto.CompactTextString(m) }
func (*MsgExecResponse) ProtoMessage()    {}
func (*MsgExecResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ceddab7d8589ad1, []int{1}
}
func (m *MsgExecResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgExecResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgExecResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgExecResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgExecResponse.Merge(m, src)
}
func (m *MsgExecResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgExecResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgExecResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgExecResponse proto.InternalMessageInfo

// MsgExec attempts to execute the provided messages using
// authorizations granted to the grantee. Each message should have only
// one signer corresponding to the granter of the authorization.
type MsgExec struct {
	Grantee string `protobuf:"bytes,1,opt,name=grantee,proto3" json:"grantee,omitempty"`
	// Authorization Msg requests to execute. Each msg must implement Authorization interface
	// The x/authz will try to find a grant matching (msg.signers[0], grantee, MsgTypeURL(msg))
	// triple and validate it.
	Msgs []*types.Any `protobuf:"bytes,2,rep,name=msgs,proto3" json:"msgs,omitempty"`
}

func (m *MsgExec) Reset()         { *m = MsgExec{} }
func (m *MsgExec) String() string { return proto.CompactTextString(m) }
func (*MsgExec) ProtoMessage()    {}
func (*MsgExec) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ceddab7d8589ad1, []int{2}
}
func (m *MsgExec) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgExec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgExec.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgExec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgExec.Merge(m, src)
}
func (m *MsgExec) XXX_Size() int {
	return m.Size()
}
func (m *MsgExec) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgExec.DiscardUnknown(m)
}

var xxx_messageInfo_MsgExec proto.InternalMessageInfo

// MsgGrantResponse defines the Msg/MsgGrant response type.
type MsgGrantResponse struct {
}

func (m *MsgGrantResponse) Reset()         { *m = MsgGrantResponse{} }
func (m *MsgGrantResponse) String() string { return proto.CompactTextString(m) }
func (*MsgGrantResponse) ProtoMessage()    {}
func (*MsgGrantResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ceddab7d8589ad1, []int{3}
}
func (m *MsgGrantResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgGrantResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgGrantResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgGrantResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgGrantResponse.Merge(m, src)
}
func (m *MsgGrantResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgGrantResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgGrantResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgGrantResponse proto.InternalMessageInfo

// MsgRevoke revokes any authorization with the provided sdk.Msg type on the
// granter's account with that has been granted to the grantee.
type MsgRevoke struct {
	Granter    string `protobuf:"bytes,1,opt,name=granter,proto3" json:"granter,omitempty"`
	Grantee    string `protobuf:"bytes,2,opt,name=grantee,proto3" json:"grantee,omitempty"`
	MsgTypeUrl string `protobuf:"bytes,3,opt,name=msg_type_url,json=msgTypeUrl,proto3" json:"msg_type_url,omitempty"`
}

func (m *MsgRevoke) Reset()         { *m = MsgRevoke{} }
func (m *MsgRevoke) String() string { return proto.CompactTextString(m) }
func (*MsgRevoke) ProtoMessage()    {}
func (*MsgRevoke) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ceddab7d8589ad1, []int{4}
}
func (m *MsgRevoke) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRevoke) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRevoke.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRevoke) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRevoke.Merge(m, src)
}
func (m *MsgRevoke) XXX_Size() int {
	return m.Size()
}
func (m *MsgRevoke) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRevoke.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRevoke proto.InternalMessageInfo

// MsgRevokeResponse defines the Msg/MsgRevokeResponse response type.
type MsgRevokeResponse struct {
}

func (m *MsgRevokeResponse) Reset()         { *m = MsgRevokeResponse{} }
func (m *MsgRevokeResponse) String() string { return proto.CompactTextString(m) }
func (*MsgRevokeResponse) ProtoMessage()    {}
func (*MsgRevokeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3ceddab7d8589ad1, []int{5}
}
func (m *MsgRevokeResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRevokeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRevokeResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRevokeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRevokeResponse.Merge(m, src)
}
func (m *MsgRevokeResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgRevokeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRevokeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRevokeResponse proto.InternalMessageInfo

func init() {
	//proto.RegisterType((*MsgGrant)(nil), "cosmos.authz.v1beta1.MsgGrant")
	//proto.RegisterType((*MsgExecResponse)(nil), "cosmos.authz.v1beta1.MsgExecResponse")
	//proto.RegisterType((*MsgExec)(nil), "cosmos.authz.v1beta1.MsgExec")
	//proto.RegisterType((*MsgGrantResponse)(nil), "cosmos.authz.v1beta1.MsgGrantResponse")
	//proto.RegisterType((*MsgRevoke)(nil), "cosmos.authz.v1beta1.MsgRevoke")
	//proto.RegisterType((*MsgRevokeResponse)(nil), "cosmos.authz.v1beta1.MsgRevokeResponse")
}

func init() { }//proto.RegisterFile("cosmos/authz/v1beta1/tx.proto", fileDescriptor_3ceddab7d8589ad1) }

var fileDescriptor_3ceddab7d8589ad1 = []byte{
	// 567 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x54, 0x31, 0x6f, 0xd3, 0x4e,
	0x14, 0xcf, 0x35, 0x4d, 0xf3, 0xcf, 0xb5, 0xd2, 0x9f, 0xba, 0x91, 0x70, 0x0d, 0x75, 0x2d, 0xd3,
	0x42, 0x14, 0x88, 0xad, 0x84, 0x2d, 0x62, 0x49, 0x24, 0x84, 0x84, 0x88, 0x90, 0x0c, 0x2c, 0x2c,
	0x91, 0x93, 0x1c, 0x17, 0x2b, 0xb1, 0x2f, 0xf2, 0x9d, 0xa3, 0xa4, 0x13, 0x62, 0x64, 0xe2, 0x63,
	0x30, 0xa1, 0x0c, 0x1d, 0xf9, 0x00, 0x11, 0x53, 0xc5, 0x80, 0x98, 0x10, 0x24, 0x43, 0x3e, 0x06,
	0xc8, 0x77, 0xbe, 0x90, 0xa2, 0xb4, 0x74, 0x62, 0xb1, 0xef, 0xbd, 0xdf, 0xef, 0xbd, 0x7b, 0xbf,
	0xf7, 0x9e, 0x0e, 0x1e, 0xb4, 0x09, 0xf5, 0x09, 0xb5, 0xdd, 0x88, 0x75, 0x4f, 0xec, 0x61, 0xb9,
	0x85, 0x98, 0x5b, 0xb6, 0xd9, 0xc8, 0x1a, 0x84, 0x84, 0x11, 0x25, 0x2f, 0x60, 0x8b, 0xc3, 0x56,
	0x02, 0x6b, 0xfb, 0xc2, 0xdb, 0xe4, 0x1c, 0x3b, 0xa1, 0x70, 0x43, 0xcb, 0x63, 0x82, 0x89, 0xf0,
	0xc7, 0xa7, 0xc4, 0xbb, 0x8f, 0x09, 0xc1, 0x7d, 0x64, 0x73, 0xab, 0x15, 0xbd, 0xb2, 0xdd, 0x60,
	0x9c, 0x40, 0xc6, 0xda, 0x02, 0xc4, 0x7d, 0x82, 0x71, 0x3d, 0x61, 0xf8, 0x14, 0xdb, 0xc3, 0x72,
	0xfc, 0x4b, 0x80, 0x5d, 0xd7, 0xf7, 0x02, 0x62, 0xf3, 0xaf, 0x70, 0x99, 0x5f, 0x00, 0xfc, 0xaf,
	0x41, 0xf1, 0xa3, 0xd0, 0x0d, 0x98, 0x52, 0x81, 0x59, 0x1c, 0x1f, 0x50, 0xa8, 0x02, 0x03, 0x14,
	0x72, 0x75, 0xf5, 0xf3, 0x69, 0x49, 0x2a, 0xaa, 0x75, 0x3a, 0x21, 0xa2, 0xf4, 0x19, 0x0b, 0xbd,
	0x00, 0x3b, 0x92, 0xf8, 0x3b, 0x06, 0xa9, 0x1b, 0x57, 0x8b, 0x41, 0xca, 0x03, 0x98, 0xe1, 0x47,
	0x35, 0x6d, 0x80, 0xc2, 0x76, 0xe5, 0x86, 0xb5, 0xae, 0x69, 0x16, 0xaf, 0xa9, 0x9e, 0x9b, 0x7e,
	0x3b, 0x4c, 0xbd, 0x5f, 0x4c, 0x8a, 0xc0, 0x11, 0x41, 0xd5, 0xa3, 0x37, 0x8b, 0x49, 0x51, 0xde,
	0xff, 0x76, 0x31, 0x29, 0xee, 0x89, 0xf0, 0x12, 0xed, 0xf4, 0x6c, 0xa9, 0xc5, 0xbc, 0x0b, 0xff,
	0x6f, 0x50, 0xfc, 0x70, 0x84, 0xda, 0x0e, 0xa2, 0x03, 0x12, 0x50, 0xa4, 0xa8, 0x30, 0x1b, 0x22,
	0x1a, 0xf5, 0x19, 0x55, 0x81, 0x91, 0x2e, 0xec, 0x38, 0xd2, 0x34, 0x3f, 0x00, 0x98, 0x4d, 0xd8,
	0xab, 0x82, 0xc0, 0x55, 0x05, 0x3d, 0x86, 0x9b, 0x3e, 0xc5, 0x54, 0xdd, 0x30, 0xd2, 0x85, 0xed,
	0x4a, 0xde, 0x12, 0xd3, 0xb3, 0xe4, 0xf4, 0xac, 0x5a, 0x30, 0xae, 0x1b, 0x9f, 0x4e, 0x4b, 0x37,
	0x69, 0xa7, 0x67, 0x35, 0x28, 0xbe, 0x67, 0x08, 0xa9, 0xb5, 0x88, 0x75, 0x49, 0xe8, 0x9d, 0xb8,
	0xcc, 0x23, 0x81, 0xc3, 0x73, 0x54, 0x6f, 0xad, 0xc8, 0x43, 0xb1, 0x3c, 0xe5, 0xbc, 0xbc, 0xb8,
	0x48, 0x53, 0x81, 0xd7, 0xa4, 0x52, 0x29, 0xcf, 0xfc, 0x08, 0x60, 0xae, 0x41, 0xb1, 0x83, 0x86,
	0xa4, 0x87, 0xfe, 0xd9, 0x2c, 0x0d, 0xb8, 0xe3, 0x53, 0xdc, 0x64, 0xe3, 0x01, 0x6a, 0x46, 0x61,
	0x9f, 0x8f, 0x34, 0xe7, 0x40, 0x9f, 0xe2, 0xe7, 0xe3, 0x01, 0x7a, 0x11, 0xf6, 0xab, 0xc7, 0x7f,
	0xce, 0x2b, 0x7f, 0x5e, 0x90, 0x28, 0xd8, 0xdc, 0x83, 0xbb, 0x4b, 0x43, 0x6a, 0xaa, 0xfc, 0x04,
	0x30, 0xdd, 0xa0, 0x58, 0x79, 0x0a, 0x33, 0x62, 0x45, 0xf5, 0xf5, 0xbb, 0x22, 0x9b, 0xa1, 0xdd,
	0xbe, 0x1c, 0x5f, 0xee, 0xc2, 0x13, 0xb8, 0xc9, 0xa7, 0x7d, 0x70, 0x21, 0x3f, 0x86, 0xb5, 0xe3,
	0x4b, 0xe1, 0x65, 0x36, 0x07, 0x6e, 0x25, 0x6d, 0x3f, 0xbc, 0x30, 0x40, 0x10, 0xb4, 0x3b, 0x7f,
	0x21, 0xc8, 0x9c, 0x5a, 0xe6, 0x75, 0xbc, 0xf4, 0xf5, 0xfa, 0xf4, 0x87, 0x9e, 0x9a, 0xce, 0x74,
	0x70, 0x36, 0xd3, 0xc1, 0xf7, 0x99, 0x0e, 0xde, 0xcd, 0xf5, 0xd4, 0xd9, 0x5c, 0x4f, 0x7d, 0x9d,
	0xeb, 0xa9, 0x97, 0x47, 0xd8, 0x63, 0xdd, 0xa8, 0x65, 0xb5, 0x89, 0x9f, 0x3c, 0x2b, 0xf6, 0x4a,
	0x73, 0x47, 0xe2, 0x59, 0x68, 0x6d, 0xf1, 0x45, 0xbc, 0xff, 0x2b, 0x00, 0x00, 0xff, 0xff, 0x71,
	0x28, 0xfe, 0xc6, 0xbc, 0x04, 0x00, 0x00,
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
	// Grant grants the provided authorization to the grantee on the granter's
	// account with the provided expiration time. If there is already a grant
	// for the given (granter, grantee, Authorization) triple, then the grant
	// will be overwritten.
	Grant(ctx context.Context, in *MsgGrant, opts ...grpc.CallOption) (*MsgGrantResponse, error)
	// Exec attempts to execute the provided messages using
	// authorizations granted to the grantee. Each message should have only
	// one signer corresponding to the granter of the authorization.
	Exec(ctx context.Context, in *MsgExec, opts ...grpc.CallOption) (*MsgExecResponse, error)
	// Revoke revokes any authorization corresponding to the provided method name on the
	// granter's account that has been granted to the grantee.
	Revoke(ctx context.Context, in *MsgRevoke, opts ...grpc.CallOption) (*MsgRevokeResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) Grant(ctx context.Context, in *MsgGrant, opts ...grpc.CallOption) (*MsgGrantResponse, error) {
	out := new(MsgGrantResponse)
	err := c.cc.Invoke(ctx, "/cosmos.authz.v1beta1.Msg/Grant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) Exec(ctx context.Context, in *MsgExec, opts ...grpc.CallOption) (*MsgExecResponse, error) {
	out := new(MsgExecResponse)
	err := c.cc.Invoke(ctx, "/cosmos.authz.v1beta1.Msg/Exec", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) Revoke(ctx context.Context, in *MsgRevoke, opts ...grpc.CallOption) (*MsgRevokeResponse, error) {
	out := new(MsgRevokeResponse)
	err := c.cc.Invoke(ctx, "/cosmos.authz.v1beta1.Msg/Revoke", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// Grant grants the provided authorization to the grantee on the granter's
	// account with the provided expiration time. If there is already a grant
	// for the given (granter, grantee, Authorization) triple, then the grant
	// will be overwritten.
	Grant(context.Context, *MsgGrant) (*MsgGrantResponse, error)
	// Exec attempts to execute the provided messages using
	// authorizations granted to the grantee. Each message should have only
	// one signer corresponding to the granter of the authorization.
	Exec(context.Context, *MsgExec) (*MsgExecResponse, error)
	// Revoke revokes any authorization corresponding to the provided method name on the
	// granter's account that has been granted to the grantee.
	Revoke(context.Context, *MsgRevoke) (*MsgRevokeResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) Grant(ctx context.Context, req *MsgGrant) (*MsgGrantResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Grant not implemented")
}
func (*UnimplementedMsgServer) Exec(ctx context.Context, req *MsgExec) (*MsgExecResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Exec not implemented")
}
func (*UnimplementedMsgServer) Revoke(ctx context.Context, req *MsgRevoke) (*MsgRevokeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Revoke not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_Grant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgGrant)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Grant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.authz.v1beta1.Msg/Grant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Grant(ctx, req.(*MsgGrant))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_Exec_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgExec)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Exec(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.authz.v1beta1.Msg/Exec",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Exec(ctx, req.(*MsgExec))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_Revoke_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRevoke)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Revoke(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.authz.v1beta1.Msg/Revoke",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Revoke(ctx, req.(*MsgRevoke))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cosmos.authz.v1beta1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Grant",
			Handler:    _Msg_Grant_Handler,
		},
		{
			MethodName: "Exec",
			Handler:    _Msg_Exec_Handler,
		},
		{
			MethodName: "Revoke",
			Handler:    _Msg_Revoke_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cosmos/authz/v1beta1/tx.proto",
}

func (m *MsgGrant) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgGrant) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgGrant) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Grant.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.Grantee) > 0 {
		i -= len(m.Grantee)
		copy(dAtA[i:], m.Grantee)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Grantee)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Granter) > 0 {
		i -= len(m.Granter)
		copy(dAtA[i:], m.Granter)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Granter)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgExecResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgExecResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgExecResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Results) > 0 {
		for iNdEx := len(m.Results) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Results[iNdEx])
			copy(dAtA[i:], m.Results[iNdEx])
			i = encodeVarintTx(dAtA, i, uint64(len(m.Results[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *MsgExec) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgExec) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgExec) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Msgs) > 0 {
		for iNdEx := len(m.Msgs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Msgs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.Grantee) > 0 {
		i -= len(m.Grantee)
		copy(dAtA[i:], m.Grantee)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Grantee)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgGrantResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgGrantResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgGrantResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgRevoke) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRevoke) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRevoke) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.MsgTypeUrl) > 0 {
		i -= len(m.MsgTypeUrl)
		copy(dAtA[i:], m.MsgTypeUrl)
		i = encodeVarintTx(dAtA, i, uint64(len(m.MsgTypeUrl)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Grantee) > 0 {
		i -= len(m.Grantee)
		copy(dAtA[i:], m.Grantee)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Grantee)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Granter) > 0 {
		i -= len(m.Granter)
		copy(dAtA[i:], m.Granter)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Granter)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgRevokeResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRevokeResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRevokeResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
func (m *MsgGrant) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Granter)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Grantee)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.Grant.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func (m *MsgExecResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Results) > 0 {
		for _, b := range m.Results {
			l = len(b)
			n += 1 + l + sovTx(uint64(l))
		}
	}
	return n
}

func (m *MsgExec) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Grantee)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if len(m.Msgs) > 0 {
		for _, e := range m.Msgs {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	return n
}

func (m *MsgGrantResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgRevoke) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Granter)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Grantee)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.MsgTypeUrl)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgRevokeResponse) Size() (n int) {
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
func (m *MsgGrant) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgGrant: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgGrant: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Granter", wireType)
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
			m.Granter = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Grantee", wireType)
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
			m.Grantee = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Grant", wireType)
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
			if err := m.Grant.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *MsgExecResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgExecResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgExecResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Results", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Results = append(m.Results, make([]byte, postIndex-iNdEx))
			copy(m.Results[len(m.Results)-1], dAtA[iNdEx:postIndex])
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
func (m *MsgExec) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgExec: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgExec: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Grantee", wireType)
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
			m.Grantee = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Msgs", wireType)
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
			m.Msgs = append(m.Msgs, &types.Any{})
			if err := m.Msgs[len(m.Msgs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *MsgGrantResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgGrantResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgGrantResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *MsgRevoke) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgRevoke: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRevoke: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Granter", wireType)
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
			m.Granter = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Grantee", wireType)
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
			m.Grantee = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MsgTypeUrl", wireType)
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
			m.MsgTypeUrl = string(dAtA[iNdEx:postIndex])
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
func (m *MsgRevokeResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgRevokeResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRevokeResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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