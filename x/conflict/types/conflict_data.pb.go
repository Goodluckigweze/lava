// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: conflict/conflict_data.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/lavanet/lava/x/pairing/types"
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

type ResponseConflict struct {
	ConflictRelayData0 *ConflictRelayData `protobuf:"bytes,1,opt,name=conflictRelayData0,proto3" json:"conflictRelayData0,omitempty"`
	ConflictRelayData1 *ConflictRelayData `protobuf:"bytes,2,opt,name=conflictRelayData1,proto3" json:"conflictRelayData1,omitempty"`
}

func (m *ResponseConflict) Reset()         { *m = ResponseConflict{} }
func (m *ResponseConflict) String() string { return proto.CompactTextString(m) }
func (*ResponseConflict) ProtoMessage()    {}
func (*ResponseConflict) Descriptor() ([]byte, []int) {
	return fileDescriptor_d7f63a98ab02ebfa, []int{0}
}
func (m *ResponseConflict) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ResponseConflict) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ResponseConflict.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ResponseConflict) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseConflict.Merge(m, src)
}
func (m *ResponseConflict) XXX_Size() int {
	return m.Size()
}
func (m *ResponseConflict) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseConflict.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseConflict proto.InternalMessageInfo

func (m *ResponseConflict) GetConflictRelayData0() *ConflictRelayData {
	if m != nil {
		return m.ConflictRelayData0
	}
	return nil
}

func (m *ResponseConflict) GetConflictRelayData1() *ConflictRelayData {
	if m != nil {
		return m.ConflictRelayData1
	}
	return nil
}

type ConflictRelayData struct {
	Request *types.RelayRequest `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
	Reply   *types.RelayReply   `protobuf:"bytes,2,opt,name=reply,proto3" json:"reply,omitempty"`
}

func (m *ConflictRelayData) Reset()         { *m = ConflictRelayData{} }
func (m *ConflictRelayData) String() string { return proto.CompactTextString(m) }
func (*ConflictRelayData) ProtoMessage()    {}
func (*ConflictRelayData) Descriptor() ([]byte, []int) {
	return fileDescriptor_d7f63a98ab02ebfa, []int{1}
}
func (m *ConflictRelayData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ConflictRelayData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ConflictRelayData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ConflictRelayData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConflictRelayData.Merge(m, src)
}
func (m *ConflictRelayData) XXX_Size() int {
	return m.Size()
}
func (m *ConflictRelayData) XXX_DiscardUnknown() {
	xxx_messageInfo_ConflictRelayData.DiscardUnknown(m)
}

var xxx_messageInfo_ConflictRelayData proto.InternalMessageInfo

func (m *ConflictRelayData) GetRequest() *types.RelayRequest {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *ConflictRelayData) GetReply() *types.RelayReply {
	if m != nil {
		return m.Reply
	}
	return nil
}

type FinalizationConflict struct {
	RelayReply0 *types.RelayReply `protobuf:"bytes,1,opt,name=relayReply0,proto3" json:"relayReply0,omitempty"`
	RelayReply1 *types.RelayReply `protobuf:"bytes,2,opt,name=relayReply1,proto3" json:"relayReply1,omitempty"`
}

func (m *FinalizationConflict) Reset()         { *m = FinalizationConflict{} }
func (m *FinalizationConflict) String() string { return proto.CompactTextString(m) }
func (*FinalizationConflict) ProtoMessage()    {}
func (*FinalizationConflict) Descriptor() ([]byte, []int) {
	return fileDescriptor_d7f63a98ab02ebfa, []int{2}
}
func (m *FinalizationConflict) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FinalizationConflict) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FinalizationConflict.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FinalizationConflict) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FinalizationConflict.Merge(m, src)
}
func (m *FinalizationConflict) XXX_Size() int {
	return m.Size()
}
func (m *FinalizationConflict) XXX_DiscardUnknown() {
	xxx_messageInfo_FinalizationConflict.DiscardUnknown(m)
}

var xxx_messageInfo_FinalizationConflict proto.InternalMessageInfo

func (m *FinalizationConflict) GetRelayReply0() *types.RelayReply {
	if m != nil {
		return m.RelayReply0
	}
	return nil
}

func (m *FinalizationConflict) GetRelayReply1() *types.RelayReply {
	if m != nil {
		return m.RelayReply1
	}
	return nil
}

func init() {
	proto.RegisterType((*ResponseConflict)(nil), "lavanet.lava.conflict.ResponseConflict")
	proto.RegisterType((*ConflictRelayData)(nil), "lavanet.lava.conflict.ConflictRelayData")
	proto.RegisterType((*FinalizationConflict)(nil), "lavanet.lava.conflict.FinalizationConflict")
}

func init() { proto.RegisterFile("conflict/conflict_data.proto", fileDescriptor_d7f63a98ab02ebfa) }

var fileDescriptor_d7f63a98ab02ebfa = []byte{
	// 306 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x49, 0xce, 0xcf, 0x4b,
	0xcb, 0xc9, 0x4c, 0x2e, 0xd1, 0x87, 0x31, 0xe2, 0x53, 0x12, 0x4b, 0x12, 0xf5, 0x0a, 0x8a, 0xf2,
	0x4b, 0xf2, 0x85, 0x44, 0x73, 0x12, 0xcb, 0x12, 0xf3, 0x52, 0x4b, 0xf4, 0x40, 0xb4, 0x1e, 0x4c,
	0x85, 0x94, 0x48, 0x7a, 0x7e, 0x7a, 0x3e, 0x58, 0x85, 0x3e, 0x88, 0x05, 0x51, 0x2c, 0x25, 0x5c,
	0x90, 0x98, 0x59, 0x94, 0x99, 0x97, 0xae, 0x5f, 0x94, 0x9a, 0x93, 0x58, 0x09, 0x11, 0x54, 0x3a,
	0xc6, 0xc8, 0x25, 0x10, 0x94, 0x5a, 0x5c, 0x90, 0x9f, 0x57, 0x9c, 0xea, 0x0c, 0xd5, 0x2f, 0x14,
	0xc1, 0x25, 0x04, 0x33, 0x2b, 0x08, 0xa4, 0xd6, 0x25, 0xb1, 0x24, 0xd1, 0x40, 0x82, 0x51, 0x81,
	0x51, 0x83, 0xdb, 0x48, 0x43, 0x0f, 0xab, 0x9d, 0x7a, 0xce, 0xe8, 0x1a, 0x82, 0xb0, 0x98, 0x81,
	0xd5, 0x64, 0x43, 0x09, 0x26, 0x8a, 0x4d, 0x36, 0x54, 0xea, 0x64, 0xe4, 0x12, 0xc4, 0x50, 0x29,
	0x64, 0xc3, 0xc5, 0x5e, 0x94, 0x5a, 0x58, 0x9a, 0x5a, 0x5c, 0x02, 0x75, 0xbe, 0x12, 0xaa, 0x25,
	0xd0, 0x20, 0xd1, 0x03, 0xeb, 0x08, 0x82, 0xa8, 0x0c, 0x82, 0x69, 0x11, 0x32, 0xe3, 0x62, 0x2d,
	0x4a, 0x2d, 0xc8, 0xa9, 0x84, 0x3a, 0x50, 0x01, 0xaf, 0xde, 0x82, 0x9c, 0xca, 0x20, 0x88, 0x72,
	0xa5, 0x79, 0x8c, 0x5c, 0x22, 0x6e, 0x99, 0x79, 0x89, 0x39, 0x99, 0x55, 0x89, 0x25, 0x99, 0xf9,
	0x79, 0xf0, 0x80, 0x75, 0xe2, 0xe2, 0x2e, 0x82, 0xab, 0x86, 0x85, 0x28, 0x61, 0x63, 0x91, 0x35,
	0xa1, 0x9a, 0x61, 0x48, 0xb4, 0xd3, 0x90, 0x35, 0x39, 0x39, 0x9d, 0x78, 0x24, 0xc7, 0x78, 0xe1,
	0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x13, 0x1e, 0xcb, 0x31, 0x5c, 0x78, 0x2c, 0xc7, 0x70,
	0xe3, 0xb1, 0x1c, 0x43, 0x94, 0x46, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e, 0xae,
	0x3e, 0xd4, 0x48, 0x30, 0xad, 0x5f, 0x01, 0x4f, 0x80, 0xfa, 0x25, 0x95, 0x05, 0xa9, 0xc5, 0x49,
	0x6c, 0xe0, 0x04, 0x64, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x9b, 0x21, 0x1e, 0x57, 0xa2, 0x02,
	0x00, 0x00,
}

func (m *ResponseConflict) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ResponseConflict) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ResponseConflict) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ConflictRelayData1 != nil {
		{
			size, err := m.ConflictRelayData1.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintConflictData(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.ConflictRelayData0 != nil {
		{
			size, err := m.ConflictRelayData0.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintConflictData(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ConflictRelayData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ConflictRelayData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ConflictRelayData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Reply != nil {
		{
			size, err := m.Reply.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintConflictData(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Request != nil {
		{
			size, err := m.Request.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintConflictData(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *FinalizationConflict) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FinalizationConflict) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FinalizationConflict) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.RelayReply1 != nil {
		{
			size, err := m.RelayReply1.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintConflictData(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.RelayReply0 != nil {
		{
			size, err := m.RelayReply0.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintConflictData(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintConflictData(dAtA []byte, offset int, v uint64) int {
	offset -= sovConflictData(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ResponseConflict) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ConflictRelayData0 != nil {
		l = m.ConflictRelayData0.Size()
		n += 1 + l + sovConflictData(uint64(l))
	}
	if m.ConflictRelayData1 != nil {
		l = m.ConflictRelayData1.Size()
		n += 1 + l + sovConflictData(uint64(l))
	}
	return n
}

func (m *ConflictRelayData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Request != nil {
		l = m.Request.Size()
		n += 1 + l + sovConflictData(uint64(l))
	}
	if m.Reply != nil {
		l = m.Reply.Size()
		n += 1 + l + sovConflictData(uint64(l))
	}
	return n
}

func (m *FinalizationConflict) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.RelayReply0 != nil {
		l = m.RelayReply0.Size()
		n += 1 + l + sovConflictData(uint64(l))
	}
	if m.RelayReply1 != nil {
		l = m.RelayReply1.Size()
		n += 1 + l + sovConflictData(uint64(l))
	}
	return n
}

func sovConflictData(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozConflictData(x uint64) (n int) {
	return sovConflictData(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ResponseConflict) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConflictData
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
			return fmt.Errorf("proto: ResponseConflict: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ResponseConflict: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConflictRelayData0", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConflictData
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
				return ErrInvalidLengthConflictData
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthConflictData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ConflictRelayData0 == nil {
				m.ConflictRelayData0 = &ConflictRelayData{}
			}
			if err := m.ConflictRelayData0.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConflictRelayData1", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConflictData
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
				return ErrInvalidLengthConflictData
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthConflictData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ConflictRelayData1 == nil {
				m.ConflictRelayData1 = &ConflictRelayData{}
			}
			if err := m.ConflictRelayData1.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConflictData(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthConflictData
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
func (m *ConflictRelayData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConflictData
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
			return fmt.Errorf("proto: ConflictRelayData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ConflictRelayData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Request", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConflictData
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
				return ErrInvalidLengthConflictData
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthConflictData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Request == nil {
				m.Request = &types.RelayRequest{}
			}
			if err := m.Request.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reply", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConflictData
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
				return ErrInvalidLengthConflictData
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthConflictData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Reply == nil {
				m.Reply = &types.RelayReply{}
			}
			if err := m.Reply.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConflictData(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthConflictData
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
func (m *FinalizationConflict) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConflictData
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
			return fmt.Errorf("proto: FinalizationConflict: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FinalizationConflict: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RelayReply0", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConflictData
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
				return ErrInvalidLengthConflictData
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthConflictData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.RelayReply0 == nil {
				m.RelayReply0 = &types.RelayReply{}
			}
			if err := m.RelayReply0.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RelayReply1", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConflictData
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
				return ErrInvalidLengthConflictData
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthConflictData
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.RelayReply1 == nil {
				m.RelayReply1 = &types.RelayReply{}
			}
			if err := m.RelayReply1.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConflictData(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthConflictData
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
func skipConflictData(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowConflictData
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
					return 0, ErrIntOverflowConflictData
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
					return 0, ErrIntOverflowConflictData
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
				return 0, ErrInvalidLengthConflictData
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupConflictData
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthConflictData
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthConflictData        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowConflictData          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupConflictData = fmt.Errorf("proto: unexpected end of group")
)
