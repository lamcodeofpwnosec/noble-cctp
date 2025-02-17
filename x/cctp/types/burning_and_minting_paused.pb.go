// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: circle/cctp/v1/burning_and_minting_paused.proto

package types

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
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

// *
// Message format for BurningAndMintingPaused
// @param paused true if paused, false if not paused
type BurningAndMintingPaused struct {
	Paused bool `protobuf:"varint,1,opt,name=paused,proto3" json:"paused,omitempty"`
}

func (m *BurningAndMintingPaused) Reset()         { *m = BurningAndMintingPaused{} }
func (m *BurningAndMintingPaused) String() string { return proto.CompactTextString(m) }
func (*BurningAndMintingPaused) ProtoMessage()    {}
func (*BurningAndMintingPaused) Descriptor() ([]byte, []int) {
	return fileDescriptor_7f12bb101f745034, []int{0}
}
func (m *BurningAndMintingPaused) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BurningAndMintingPaused) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BurningAndMintingPaused.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BurningAndMintingPaused) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BurningAndMintingPaused.Merge(m, src)
}
func (m *BurningAndMintingPaused) XXX_Size() int {
	return m.Size()
}
func (m *BurningAndMintingPaused) XXX_DiscardUnknown() {
	xxx_messageInfo_BurningAndMintingPaused.DiscardUnknown(m)
}

var xxx_messageInfo_BurningAndMintingPaused proto.InternalMessageInfo

func (m *BurningAndMintingPaused) GetPaused() bool {
	if m != nil {
		return m.Paused
	}
	return false
}

func init() {
	proto.RegisterType((*BurningAndMintingPaused)(nil), "circle.cctp.v1.BurningAndMintingPaused")
}

func init() {
	proto.RegisterFile("circle/cctp/v1/burning_and_minting_paused.proto", fileDescriptor_7f12bb101f745034)
}

var fileDescriptor_7f12bb101f745034 = []byte{
	// 179 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x4f, 0xce, 0x2c, 0x4a,
	0xce, 0x49, 0xd5, 0x4f, 0x4e, 0x2e, 0x29, 0xd0, 0x2f, 0x33, 0xd4, 0x4f, 0x2a, 0x2d, 0xca, 0xcb,
	0xcc, 0x4b, 0x8f, 0x4f, 0xcc, 0x4b, 0x89, 0xcf, 0xcd, 0xcc, 0x2b, 0x01, 0xb1, 0x0b, 0x12, 0x4b,
	0x8b, 0x53, 0x53, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0xf8, 0x20, 0x1a, 0xf4, 0x40, 0x1a,
	0xf4, 0xca, 0x0c, 0x95, 0x0c, 0xb9, 0xc4, 0x9d, 0x20, 0x7a, 0x1c, 0xf3, 0x52, 0x7c, 0x21, 0x3a,
	0x02, 0xc0, 0x1a, 0x84, 0xc4, 0xb8, 0xd8, 0x20, 0x5a, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x82,
	0xa0, 0x3c, 0x27, 0xb7, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e,
	0x71, 0xc2, 0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18, 0x6e, 0x3c, 0x96, 0x63, 0x88, 0xd2, 0x49,
	0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0x85, 0x3a, 0x2c, 0x2d, 0x33, 0x4f, 0x3f,
	0x2f, 0x3f, 0x29, 0x27, 0x55, 0x17, 0xec, 0xc2, 0x0a, 0x88, 0x43, 0x4b, 0x2a, 0x0b, 0x52, 0x8b,
	0x93, 0xd8, 0xc0, 0x2e, 0x32, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x72, 0x62, 0x88, 0xaf, 0xc4,
	0x00, 0x00, 0x00,
}

func (m *BurningAndMintingPaused) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BurningAndMintingPaused) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BurningAndMintingPaused) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Paused {
		i--
		if m.Paused {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintBurningAndMintingPaused(dAtA []byte, offset int, v uint64) int {
	offset -= sovBurningAndMintingPaused(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *BurningAndMintingPaused) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Paused {
		n += 2
	}
	return n
}

func sovBurningAndMintingPaused(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozBurningAndMintingPaused(x uint64) (n int) {
	return sovBurningAndMintingPaused(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BurningAndMintingPaused) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBurningAndMintingPaused
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
			return fmt.Errorf("proto: BurningAndMintingPaused: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BurningAndMintingPaused: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Paused", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBurningAndMintingPaused
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Paused = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipBurningAndMintingPaused(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBurningAndMintingPaused
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
func skipBurningAndMintingPaused(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBurningAndMintingPaused
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
					return 0, ErrIntOverflowBurningAndMintingPaused
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
					return 0, ErrIntOverflowBurningAndMintingPaused
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
				return 0, ErrInvalidLengthBurningAndMintingPaused
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupBurningAndMintingPaused
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthBurningAndMintingPaused
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthBurningAndMintingPaused        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBurningAndMintingPaused          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupBurningAndMintingPaused = fmt.Errorf("proto: unexpected end of group")
)
