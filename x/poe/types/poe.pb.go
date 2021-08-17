// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: confio/poe/v1beta1/poe.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
	_ "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/cosmos-sdk/x/staking/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/regen-network/cosmos-proto"
	_ "github.com/tendermint/tendermint/proto/tendermint/types"
	_ "google.golang.org/protobuf/types/known/durationpb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
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

// PoEContractType type of PoE contract
type PoEContractType int32

const (
	PoEContractTypeUndefined  PoEContractType = 0
	PoEContractTypeStaking    PoEContractType = 1
	PoEContractTypeValset     PoEContractType = 2
	PoEContractTypeEngagement PoEContractType = 3
	PoEContractTypeMixer      PoEContractType = 4
)

var PoEContractType_name = map[int32]string{
	0: "UNDEFINED",
	1: "STAKING",
	2: "VALSET",
	3: "ENGAGEMENT",
	4: "MIXER",
}

var PoEContractType_value = map[string]int32{
	"UNDEFINED":  0,
	"STAKING":    1,
	"VALSET":     2,
	"ENGAGEMENT": 3,
	"MIXER":      4,
}

func (x PoEContractType) String() string {
	return proto.EnumName(PoEContractType_name, int32(x))
}

func (PoEContractType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_df6d9ea68813554a, []int{0}
}

// Params defines the parameters for the PoE module.
type Params struct {
	// historical_entries is the number of historical entries to persist.
	HistoricalEntries uint32 `protobuf:"varint,4,opt,name=historical_entries,json=historicalEntries,proto3" json:"historical_entries,omitempty" yaml:"historical_entries"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_df6d9ea68813554a, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetHistoricalEntries() uint32 {
	if m != nil {
		return m.HistoricalEntries
	}
	return 0
}

func init() {
	proto.RegisterEnum("confio.poe.v1beta1.PoEContractType", PoEContractType_name, PoEContractType_value)
	proto.RegisterType((*Params)(nil), "confio.poe.v1beta1.Params")
}

func init() { proto.RegisterFile("confio/poe/v1beta1/poe.proto", fileDescriptor_df6d9ea68813554a) }

var fileDescriptor_df6d9ea68813554a = []byte{
	// 465 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x92, 0x41, 0x8b, 0xd3, 0x40,
	0x14, 0xc7, 0x13, 0xad, 0x55, 0x07, 0xc4, 0x18, 0x56, 0x69, 0x43, 0x37, 0x0d, 0x75, 0x45, 0x51,
	0x6c, 0x58, 0xbc, 0xed, 0x41, 0xe8, 0xba, 0x63, 0x29, 0x6e, 0xc3, 0xb2, 0xed, 0x2e, 0xe2, 0x65,
	0x9d, 0x24, 0xd3, 0xd9, 0xc1, 0x66, 0x26, 0xcc, 0x4c, 0x65, 0xfb, 0x0d, 0x24, 0x27, 0x8f, 0x5e,
	0x02, 0x0b, 0x7e, 0x19, 0x8f, 0x7b, 0xf4, 0x24, 0xd2, 0x5e, 0xf6, 0xec, 0x27, 0x90, 0x66, 0x52,
	0x0b, 0xf1, 0x12, 0xde, 0xff, 0xfd, 0x7e, 0xbc, 0x24, 0xbc, 0x07, 0x5a, 0x11, 0x67, 0x13, 0xca,
	0xfd, 0x94, 0x63, 0xff, 0xf3, 0x6e, 0x88, 0x15, 0xda, 0x5d, 0xd5, 0xdd, 0x54, 0x70, 0xc5, 0x6d,
	0x5b, 0xd3, 0xee, 0xaa, 0x53, 0x52, 0x67, 0x8b, 0x70, 0xc2, 0x0b, 0xec, 0xaf, 0x2a, 0x6d, 0x3a,
	0x4d, 0xc2, 0x39, 0x99, 0x62, 0xbf, 0x48, 0xe1, 0x6c, 0xe2, 0x23, 0x36, 0x2f, 0x91, 0x5b, 0x45,
	0xf1, 0x4c, 0x20, 0x45, 0x39, 0x2b, 0x79, 0xbb, 0xca, 0x15, 0x4d, 0xb0, 0x54, 0x28, 0x49, 0xd7,
	0xb3, 0x23, 0x2e, 0x13, 0x2e, 0xcf, 0xf4, 0x4b, 0x75, 0x58, 0xcf, 0xd6, 0xc9, 0x0f, 0x91, 0xdc,
	0x7c, 0x7f, 0xc4, 0xe9, 0x7a, 0xf6, 0x4e, 0xc9, 0xa5, 0x42, 0x9f, 0x28, 0x23, 0xff, 0x94, 0x32,
	0x97, 0x56, 0x4b, 0x61, 0x16, 0x63, 0x91, 0x50, 0xa6, 0x7c, 0x35, 0x4f, 0xb1, 0xd4, 0x4f, 0x4d,
	0x3b, 0x1f, 0x41, 0xfd, 0x08, 0x09, 0x94, 0x48, 0xfb, 0x10, 0xd8, 0xe7, 0x54, 0x2a, 0x2e, 0x68,
	0x84, 0xa6, 0x67, 0x98, 0x29, 0x41, 0xb1, 0x6c, 0xd4, 0x3c, 0xf3, 0xd9, 0xbd, 0xfd, 0xed, 0x3f,
	0xbf, 0xda, 0xcd, 0x39, 0x4a, 0xa6, 0x7b, 0x9d, 0xff, 0x9d, 0xce, 0xf1, 0x83, 0x4d, 0x13, 0xea,
	0xde, 0xde, 0x9d, 0x6f, 0x97, 0x6d, 0xe3, 0xfa, 0xb2, 0x6d, 0x3e, 0xbf, 0x36, 0xc1, 0xfd, 0x23,
	0x0e, 0xdf, 0x70, 0xa6, 0x04, 0x8a, 0xd4, 0x78, 0x9e, 0x62, 0xfb, 0x05, 0xb8, 0x7b, 0x12, 0x1c,
	0xc0, 0xb7, 0x83, 0x00, 0x1e, 0x58, 0x86, 0xd3, 0xca, 0x72, 0xaf, 0x51, 0x71, 0x4e, 0x58, 0x8c,
	0x27, 0x94, 0xe1, 0xd8, 0x7e, 0x0a, 0x6e, 0x8f, 0xc6, 0xbd, 0x77, 0x83, 0xa0, 0x6f, 0x99, 0x8e,
	0x93, 0xe5, 0xde, 0xa3, 0x8a, 0x3a, 0xd2, 0xff, 0x6b, 0x3f, 0x01, 0xf5, 0xd3, 0xde, 0xe1, 0x08,
	0x8e, 0xad, 0x1b, 0x4e, 0x33, 0xcb, 0xbd, 0x87, 0x15, 0xef, 0x14, 0x4d, 0x25, 0x56, 0xf6, 0x4b,
	0x00, 0x60, 0xd0, 0xef, 0xf5, 0xe1, 0x10, 0x06, 0x63, 0xeb, 0xa6, 0xb3, 0x9d, 0xe5, 0x5e, 0xb3,
	0xa2, 0x42, 0x46, 0x10, 0xc1, 0x09, 0x66, 0xca, 0x7e, 0x0c, 0x6e, 0x0d, 0x07, 0xef, 0xe1, 0xb1,
	0x55, 0x73, 0x1a, 0x59, 0xee, 0x6d, 0x55, 0xcc, 0x21, 0xbd, 0xc0, 0xc2, 0xa9, 0x7d, 0xf9, 0xee,
	0x1a, 0xfb, 0xaf, 0x7f, 0x2c, 0x5c, 0xf3, 0x6a, 0xe1, 0x9a, 0xbf, 0x17, 0xae, 0xf9, 0x75, 0xe9,
	0x1a, 0x57, 0x4b, 0xd7, 0xf8, 0xb9, 0x74, 0x8d, 0x0f, 0x3b, 0x84, 0xaa, 0xf3, 0x59, 0xd8, 0x8d,
	0x78, 0xe2, 0x97, 0x47, 0xa9, 0x88, 0x40, 0x31, 0xf6, 0x2f, 0x8a, 0xeb, 0x2c, 0x56, 0x12, 0xd6,
	0x8b, 0x9d, 0xbc, 0xfa, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x33, 0xbe, 0xb7, 0x0e, 0xb8, 0x02, 0x00,
	0x00,
}

func (this *Params) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Params)
	if !ok {
		that2, ok := that.(Params)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.HistoricalEntries != that1.HistoricalEntries {
		return false
	}
	return true
}
func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.HistoricalEntries != 0 {
		i = encodeVarintPoe(dAtA, i, uint64(m.HistoricalEntries))
		i--
		dAtA[i] = 0x20
	}
	return len(dAtA) - i, nil
}

func encodeVarintPoe(dAtA []byte, offset int, v uint64) int {
	offset -= sovPoe(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.HistoricalEntries != 0 {
		n += 1 + sovPoe(uint64(m.HistoricalEntries))
	}
	return n
}

func sovPoe(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPoe(x uint64) (n int) {
	return sovPoe(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPoe
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field HistoricalEntries", wireType)
			}
			m.HistoricalEntries = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPoe
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.HistoricalEntries |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipPoe(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPoe
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
func skipPoe(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPoe
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
					return 0, ErrIntOverflowPoe
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
					return 0, ErrIntOverflowPoe
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
				return 0, ErrInvalidLengthPoe
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPoe
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPoe
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPoe        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPoe          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPoe = fmt.Errorf("proto: unexpected end of group")
)