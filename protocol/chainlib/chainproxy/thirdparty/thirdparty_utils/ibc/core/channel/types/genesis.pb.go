// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ibc/core/channel/v1/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

// GenesisState defines the ibc channel submodule's genesis state.
type GenesisState struct {
	Channels         []IdentifiedChannel `protobuf:"bytes,1,rep,name=channels,proto3,casttype=IdentifiedChannel" json:"channels"`
	Acknowledgements []PacketState       `protobuf:"bytes,2,rep,name=acknowledgements,proto3" json:"acknowledgements"`
	Commitments      []PacketState       `protobuf:"bytes,3,rep,name=commitments,proto3" json:"commitments"`
	Receipts         []PacketState       `protobuf:"bytes,4,rep,name=receipts,proto3" json:"receipts"`
	SendSequences    []PacketSequence    `protobuf:"bytes,5,rep,name=send_sequences,json=sendSequences,proto3" json:"send_sequences" yaml:"send_sequences"`
	RecvSequences    []PacketSequence    `protobuf:"bytes,6,rep,name=recv_sequences,json=recvSequences,proto3" json:"recv_sequences" yaml:"recv_sequences"`
	AckSequences     []PacketSequence    `protobuf:"bytes,7,rep,name=ack_sequences,json=ackSequences,proto3" json:"ack_sequences" yaml:"ack_sequences"`
	// the sequence for the next generated channel identifier
	NextChannelSequence uint64 `protobuf:"varint,8,opt,name=next_channel_sequence,json=nextChannelSequence,proto3" json:"next_channel_sequence,omitempty" yaml:"next_channel_sequence"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb06ec201f452595, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetChannels() []IdentifiedChannel {
	if m != nil {
		return m.Channels
	}
	return nil
}

func (m *GenesisState) GetAcknowledgements() []PacketState {
	if m != nil {
		return m.Acknowledgements
	}
	return nil
}

func (m *GenesisState) GetCommitments() []PacketState {
	if m != nil {
		return m.Commitments
	}
	return nil
}

func (m *GenesisState) GetReceipts() []PacketState {
	if m != nil {
		return m.Receipts
	}
	return nil
}

func (m *GenesisState) GetSendSequences() []PacketSequence {
	if m != nil {
		return m.SendSequences
	}
	return nil
}

func (m *GenesisState) GetRecvSequences() []PacketSequence {
	if m != nil {
		return m.RecvSequences
	}
	return nil
}

func (m *GenesisState) GetAckSequences() []PacketSequence {
	if m != nil {
		return m.AckSequences
	}
	return nil
}

func (m *GenesisState) GetNextChannelSequence() uint64 {
	if m != nil {
		return m.NextChannelSequence
	}
	return 0
}

// PacketSequence defines the genesis type necessary to retrieve and store
// next send and receive sequences.
type PacketSequence struct {
	PortId    string `protobuf:"bytes,1,opt,name=port_id,json=portId,proto3" json:"port_id,omitempty" yaml:"port_id"`
	ChannelId string `protobuf:"bytes,2,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty" yaml:"channel_id"`
	Sequence  uint64 `protobuf:"varint,3,opt,name=sequence,proto3" json:"sequence,omitempty"`
}

func (m *PacketSequence) Reset()         { *m = PacketSequence{} }
func (m *PacketSequence) String() string { return proto.CompactTextString(m) }
func (*PacketSequence) ProtoMessage()    {}
func (*PacketSequence) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb06ec201f452595, []int{1}
}
func (m *PacketSequence) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PacketSequence) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PacketSequence.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PacketSequence) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PacketSequence.Merge(m, src)
}
func (m *PacketSequence) XXX_Size() int {
	return m.Size()
}
func (m *PacketSequence) XXX_DiscardUnknown() {
	xxx_messageInfo_PacketSequence.DiscardUnknown(m)
}

var xxx_messageInfo_PacketSequence proto.InternalMessageInfo

func (m *PacketSequence) GetPortId() string {
	if m != nil {
		return m.PortId
	}
	return ""
}

func (m *PacketSequence) GetChannelId() string {
	if m != nil {
		return m.ChannelId
	}
	return ""
}

func (m *PacketSequence) GetSequence() uint64 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

// disabled because of duplicate proto type error
func init() {
	// proto.RegisterType((*GenesisState)(nil), "ibc.core.channel.v1.GenesisState")
	// proto.RegisterType((*PacketSequence)(nil), "ibc.core.channel.v1.PacketSequence")
}

func init() { proto.RegisterFile("ibc/core/channel/v1/genesis.proto", fileDescriptor_cb06ec201f452595) }

var fileDescriptor_cb06ec201f452595 = []byte{
	// 507 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0x87, 0xe3, 0x36, 0x4d, 0xd3, 0x6d, 0x13, 0xd1, 0x6d, 0x23, 0x99, 0xa8, 0xd8, 0xc6, 0x48,
	0x28, 0x12, 0xaa, 0x4d, 0xa1, 0x42, 0x82, 0xa3, 0x39, 0x40, 0x6e, 0x68, 0xcb, 0x09, 0x09, 0x45,
	0xf6, 0x7a, 0xea, 0xae, 0x12, 0x7b, 0x83, 0x77, 0x13, 0xe8, 0x53, 0xc0, 0x63, 0xf5, 0xd8, 0x23,
	0x27, 0x0b, 0x25, 0x6f, 0x90, 0x23, 0x27, 0xe4, 0xbf, 0x49, 0xd4, 0x08, 0x51, 0x6e, 0xde, 0x99,
	0xdf, 0x7c, 0xdf, 0x1c, 0x3c, 0xe8, 0x31, 0xf3, 0xa8, 0x4d, 0x79, 0x0c, 0x36, 0xbd, 0x72, 0xa3,
	0x08, 0x46, 0xf6, 0xf4, 0xcc, 0x0e, 0x20, 0x02, 0xc1, 0x84, 0x35, 0x8e, 0xb9, 0xe4, 0xf8, 0x88,
	0x79, 0xd4, 0x4a, 0x23, 0x56, 0x11, 0xb1, 0xa6, 0x67, 0xdd, 0xe3, 0x80, 0x07, 0x3c, 0xeb, 0xdb,
	0xe9, 0x57, 0x1e, 0xed, 0x6e, 0xa4, 0x95, 0x53, 0x59, 0xc4, 0x9c, 0xef, 0xa0, 0x83, 0x77, 0x39,
	0xff, 0x42, 0xba, 0x12, 0xf0, 0x67, 0xd4, 0x2c, 0x12, 0x42, 0x55, 0x8c, 0xed, 0xde, 0xfe, 0x8b,
	0xa7, 0xd6, 0x06, 0xa3, 0xd5, 0xf7, 0x21, 0x92, 0xec, 0x92, 0x81, 0xff, 0x36, 0x2f, 0x3a, 0x0f,
	0x6f, 0x12, 0xbd, 0xf6, 0x3b, 0xd1, 0x0f, 0xef, 0xb4, 0x48, 0x85, 0xc4, 0x04, 0x3d, 0x70, 0xe9,
	0x30, 0xe2, 0x5f, 0x47, 0xe0, 0x07, 0x10, 0x42, 0x24, 0x85, 0xba, 0x95, 0x69, 0x8c, 0x8d, 0x9a,
	0x0f, 0x2e, 0x1d, 0x82, 0xcc, 0x56, 0x73, 0xea, 0xa9, 0x80, 0xdc, 0x99, 0xc7, 0xef, 0xd1, 0x3e,
	0xe5, 0x61, 0xc8, 0x64, 0x8e, 0xdb, 0xbe, 0x17, 0x6e, 0x75, 0x14, 0x3b, 0xa8, 0x19, 0x03, 0x05,
	0x36, 0x96, 0x42, 0xad, 0xdf, 0x0b, 0x53, 0xcd, 0x61, 0x86, 0xda, 0x02, 0x22, 0x7f, 0x20, 0xe0,
	0xcb, 0x04, 0x22, 0x0a, 0x42, 0xdd, 0xc9, 0x48, 0x4f, 0xfe, 0x46, 0x2a, 0xb2, 0xce, 0xa3, 0x14,
	0xb6, 0x48, 0xf4, 0xce, 0xb5, 0x1b, 0x8e, 0xde, 0x98, 0xeb, 0x20, 0x93, 0xb4, 0xd2, 0x42, 0x19,
	0xce, 0x54, 0x31, 0xd0, 0xe9, 0x8a, 0xaa, 0xf1, 0xdf, 0xaa, 0x75, 0x90, 0x49, 0x5a, 0x69, 0x61,
	0xa9, 0xba, 0x44, 0x2d, 0x97, 0x0e, 0x57, 0x4c, 0xbb, 0xff, 0x6e, 0x3a, 0x29, 0x4c, 0xc7, 0xb9,
	0x69, 0x8d, 0x63, 0x92, 0x03, 0x97, 0x0e, 0x97, 0x9e, 0x8f, 0xa8, 0x13, 0xc1, 0x37, 0x39, 0x28,
	0x68, 0x55, 0x50, 0x6d, 0x1a, 0x4a, 0xaf, 0xee, 0x18, 0x8b, 0x44, 0x3f, 0xc9, 0x31, 0x1b, 0x63,
	0x26, 0x39, 0x4a, 0xeb, 0xc5, 0x7f, 0x57, 0x62, 0xcd, 0xef, 0x0a, 0x6a, 0xaf, 0x2f, 0x85, 0x9f,
	0xa1, 0xdd, 0x31, 0x8f, 0xe5, 0x80, 0xf9, 0xaa, 0x62, 0x28, 0xbd, 0x3d, 0x07, 0x2f, 0x12, 0xbd,
	0x9d, 0xa3, 0x8b, 0x86, 0x49, 0x1a, 0xe9, 0x57, 0xdf, 0xc7, 0xe7, 0x08, 0x95, 0x26, 0xe6, 0xab,
	0x5b, 0x59, 0xbe, 0xb3, 0x48, 0xf4, 0xc3, 0x3c, 0xbf, 0xec, 0x99, 0x64, 0xaf, 0x78, 0xf4, 0x7d,
	0xdc, 0x45, 0xcd, 0x6a, 0xfd, 0xed, 0x74, 0x7d, 0x52, 0xbd, 0x9d, 0x8b, 0x9b, 0x99, 0xa6, 0xdc,
	0xce, 0x34, 0xe5, 0xd7, 0x4c, 0x53, 0x7e, 0xcc, 0xb5, 0xda, 0xed, 0x5c, 0xab, 0xfd, 0x9c, 0x6b,
	0xb5, 0x4f, 0xaf, 0x03, 0x26, 0xaf, 0x26, 0x9e, 0x45, 0x79, 0x68, 0x53, 0x2e, 0x42, 0x2e, 0x6c,
	0xe6, 0xd1, 0xd3, 0x80, 0xdb, 0xd3, 0x57, 0x76, 0xc8, 0xfd, 0xc9, 0x08, 0x44, 0x7e, 0xd4, 0xcf,
	0xcf, 0x4f, 0xcb, 0xbb, 0x96, 0xd7, 0x63, 0x10, 0x5e, 0x23, 0xbb, 0xe9, 0x97, 0x7f, 0x02, 0x00,
	0x00, 0xff, 0xff, 0x65, 0x82, 0x18, 0x7d, 0x46, 0x04, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.NextChannelSequence != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.NextChannelSequence))
		i--
		dAtA[i] = 0x40
	}
	if len(m.AckSequences) > 0 {
		for iNdEx := len(m.AckSequences) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.AckSequences[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x3a
		}
	}
	if len(m.RecvSequences) > 0 {
		for iNdEx := len(m.RecvSequences) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.RecvSequences[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.SendSequences) > 0 {
		for iNdEx := len(m.SendSequences) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.SendSequences[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.Receipts) > 0 {
		for iNdEx := len(m.Receipts) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Receipts[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Commitments) > 0 {
		for iNdEx := len(m.Commitments) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Commitments[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Acknowledgements) > 0 {
		for iNdEx := len(m.Acknowledgements) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Acknowledgements[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Channels) > 0 {
		for iNdEx := len(m.Channels) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Channels[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *PacketSequence) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PacketSequence) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PacketSequence) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Sequence != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.Sequence))
		i--
		dAtA[i] = 0x18
	}
	if len(m.ChannelId) > 0 {
		i -= len(m.ChannelId)
		copy(dAtA[i:], m.ChannelId)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.ChannelId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.PortId) > 0 {
		i -= len(m.PortId)
		copy(dAtA[i:], m.PortId)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.PortId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Channels) > 0 {
		for _, e := range m.Channels {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Acknowledgements) > 0 {
		for _, e := range m.Acknowledgements {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Commitments) > 0 {
		for _, e := range m.Commitments {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Receipts) > 0 {
		for _, e := range m.Receipts {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.SendSequences) > 0 {
		for _, e := range m.SendSequences {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.RecvSequences) > 0 {
		for _, e := range m.RecvSequences {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.AckSequences) > 0 {
		for _, e := range m.AckSequences {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.NextChannelSequence != 0 {
		n += 1 + sovGenesis(uint64(m.NextChannelSequence))
	}
	return n
}

func (m *PacketSequence) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.PortId)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.ChannelId)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.Sequence != 0 {
		n += 1 + sovGenesis(uint64(m.Sequence))
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Channels", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Channels = append(m.Channels, IdentifiedChannel{})
			if err := m.Channels[len(m.Channels)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Acknowledgements", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Acknowledgements = append(m.Acknowledgements, PacketState{})
			if err := m.Acknowledgements[len(m.Acknowledgements)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Commitments", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Commitments = append(m.Commitments, PacketState{})
			if err := m.Commitments[len(m.Commitments)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Receipts", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Receipts = append(m.Receipts, PacketState{})
			if err := m.Receipts[len(m.Receipts)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SendSequences", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SendSequences = append(m.SendSequences, PacketSequence{})
			if err := m.SendSequences[len(m.SendSequences)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RecvSequences", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RecvSequences = append(m.RecvSequences, PacketSequence{})
			if err := m.RecvSequences[len(m.RecvSequences)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AckSequences", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AckSequences = append(m.AckSequences, PacketSequence{})
			if err := m.AckSequences[len(m.AckSequences)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NextChannelSequence", wireType)
			}
			m.NextChannelSequence = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NextChannelSequence |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *PacketSequence) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: PacketSequence: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PacketSequence: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PortId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PortId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChannelId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChannelId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sequence", wireType)
			}
			m.Sequence = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Sequence |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
