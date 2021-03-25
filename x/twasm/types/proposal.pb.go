// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: confio/twasm/v1beta1/proposal.proto

package types

import (
	fmt "fmt"
	_ "github.com/CosmWasm/wasmd/x/wasm/types"
	_ "github.com/cosmos/cosmos-sdk/types"
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

// PromoteToPrivilegedContractProposal gov proposal content type to add
// "privileges" to a contract
type PromoteToPrivilegedContractProposal struct {
	// Title is a short summary
	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty" yaml:"title"`
	// Description is a human readable text
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty" yaml:"description"`
	// Contract is the address of the smart contract
	Contract string `protobuf:"bytes,3,opt,name=contract,proto3" json:"contract,omitempty" yaml:"contract"`
}

func (m *PromoteToPrivilegedContractProposal) Reset()      { *m = PromoteToPrivilegedContractProposal{} }
func (*PromoteToPrivilegedContractProposal) ProtoMessage() {}
func (*PromoteToPrivilegedContractProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_77ea8b6359ab7726, []int{0}
}
func (m *PromoteToPrivilegedContractProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PromoteToPrivilegedContractProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PromoteToPrivilegedContractProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PromoteToPrivilegedContractProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PromoteToPrivilegedContractProposal.Merge(m, src)
}
func (m *PromoteToPrivilegedContractProposal) XXX_Size() int {
	return m.Size()
}
func (m *PromoteToPrivilegedContractProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_PromoteToPrivilegedContractProposal.DiscardUnknown(m)
}

var xxx_messageInfo_PromoteToPrivilegedContractProposal proto.InternalMessageInfo

// PromoteToPrivilegedContractProposal gov proposal content type to remove
// "privileges" from a contract
type DemotePrivilegedContractProposal struct {
	// Title is a short summary
	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty" yaml:"title"`
	// Description is a human readable text
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty" yaml:"description"`
	// Contract is the address of the smart contract
	Contract string `protobuf:"bytes,3,opt,name=contract,proto3" json:"contract,omitempty" yaml:"contract"`
}

func (m *DemotePrivilegedContractProposal) Reset()      { *m = DemotePrivilegedContractProposal{} }
func (*DemotePrivilegedContractProposal) ProtoMessage() {}
func (*DemotePrivilegedContractProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_77ea8b6359ab7726, []int{1}
}
func (m *DemotePrivilegedContractProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DemotePrivilegedContractProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DemotePrivilegedContractProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DemotePrivilegedContractProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DemotePrivilegedContractProposal.Merge(m, src)
}
func (m *DemotePrivilegedContractProposal) XXX_Size() int {
	return m.Size()
}
func (m *DemotePrivilegedContractProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_DemotePrivilegedContractProposal.DiscardUnknown(m)
}

var xxx_messageInfo_DemotePrivilegedContractProposal proto.InternalMessageInfo

func init() {
	proto.RegisterType((*PromoteToPrivilegedContractProposal)(nil), "confio.twasm.v1beta1.PromoteToPrivilegedContractProposal")
	proto.RegisterType((*DemotePrivilegedContractProposal)(nil), "confio.twasm.v1beta1.DemotePrivilegedContractProposal")
}

func init() {
	proto.RegisterFile("confio/twasm/v1beta1/proposal.proto", fileDescriptor_77ea8b6359ab7726)
}

var fileDescriptor_77ea8b6359ab7726 = []byte{
	// 325 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x92, 0x41, 0x4a, 0x33, 0x31,
	0x14, 0xc7, 0x27, 0xdf, 0x87, 0xa2, 0x51, 0x50, 0xc6, 0x22, 0xa5, 0x8b, 0xb4, 0xa6, 0x50, 0x5c,
	0x4d, 0x28, 0x6e, 0xc4, 0x65, 0x75, 0xe9, 0xa2, 0x14, 0x57, 0xee, 0x32, 0x69, 0x1c, 0x03, 0x33,
	0xf3, 0x86, 0x24, 0x56, 0x7b, 0x0b, 0x8f, 0xe1, 0x05, 0x04, 0x8f, 0xd0, 0x65, 0x97, 0x5d, 0x15,
	0x3b, 0xbd, 0x41, 0x4f, 0x20, 0x4d, 0xa6, 0xa5, 0x1e, 0xc1, 0xdd, 0x30, 0xbf, 0xdf, 0xcb, 0xff,
	0x0f, 0xef, 0xe1, 0xb6, 0x80, 0xfc, 0x49, 0x01, 0xb3, 0xaf, 0xdc, 0x64, 0x6c, 0xd4, 0x8d, 0xa5,
	0xe5, 0x5d, 0x56, 0x68, 0x28, 0xc0, 0xf0, 0x34, 0x2a, 0x34, 0x58, 0x08, 0x6b, 0x5e, 0x8a, 0x9c,
	0x14, 0x55, 0x52, 0xa3, 0x96, 0x40, 0x02, 0x4e, 0x60, 0xeb, 0x2f, 0xef, 0x36, 0x88, 0x00, 0x93,
	0x81, 0x61, 0x31, 0x37, 0x72, 0xfb, 0x9e, 0x00, 0x95, 0x57, 0xfc, 0x62, 0xcd, 0x5d, 0xd8, 0xaf,
	0x44, 0x3b, 0x2e, 0xa4, 0xf1, 0x0a, 0xfd, 0x42, 0xb8, 0xdd, 0xd7, 0x90, 0x81, 0x95, 0x0f, 0xd0,
	0xd7, 0x6a, 0xa4, 0x52, 0x99, 0xc8, 0xe1, 0x2d, 0xe4, 0x56, 0x73, 0x61, 0xfb, 0x55, 0xb9, 0xb0,
	0x83, 0xf7, 0xac, 0xb2, 0xa9, 0xac, 0xa3, 0x16, 0xba, 0x3c, 0xec, 0x9d, 0xae, 0xe6, 0xcd, 0xe3,
	0x31, 0xcf, 0xd2, 0x1b, 0xea, 0x7e, 0xd3, 0x81, 0xc7, 0xe1, 0x35, 0x3e, 0x1a, 0x4a, 0x23, 0xb4,
	0x2a, 0xac, 0x82, 0xbc, 0xfe, 0xcf, 0xd9, 0xe7, 0xab, 0x79, 0x33, 0xf4, 0xf6, 0x0e, 0xa4, 0x83,
	0x5d, 0x35, 0x64, 0xf8, 0x40, 0x54, 0xa9, 0xf5, 0xff, 0x6e, 0xec, 0x6c, 0x35, 0x6f, 0x9e, 0xf8,
	0xb1, 0x0d, 0xa1, 0x83, 0xad, 0x44, 0x3f, 0x11, 0x6e, 0xdd, 0xc9, 0x75, 0xf3, 0x3f, 0xd5, 0xbb,
	0x77, 0x3f, 0x59, 0x90, 0x60, 0xb6, 0x20, 0xc1, 0x47, 0x49, 0xd0, 0xa4, 0x24, 0x68, 0x5a, 0x12,
	0xf4, 0x5d, 0x12, 0xf4, 0xbe, 0x24, 0xc1, 0x74, 0x49, 0x82, 0xd9, 0x92, 0x04, 0x8f, 0x9d, 0x44,
	0xd9, 0xe7, 0x97, 0x38, 0x12, 0x90, 0xb1, 0xcd, 0xdd, 0x24, 0x9a, 0x0f, 0x25, 0x7b, 0xab, 0x0e,
	0xc8, 0xad, 0x31, 0xde, 0x77, 0x7b, 0xbc, 0xfa, 0x09, 0x00, 0x00, 0xff, 0xff, 0xf3, 0x49, 0x11,
	0xca, 0x5d, 0x02, 0x00, 0x00,
}

func (this *PromoteToPrivilegedContractProposal) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*PromoteToPrivilegedContractProposal)
	if !ok {
		that2, ok := that.(PromoteToPrivilegedContractProposal)
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
	if this.Title != that1.Title {
		return false
	}
	if this.Description != that1.Description {
		return false
	}
	if this.Contract != that1.Contract {
		return false
	}
	return true
}
func (this *DemotePrivilegedContractProposal) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*DemotePrivilegedContractProposal)
	if !ok {
		that2, ok := that.(DemotePrivilegedContractProposal)
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
	if this.Title != that1.Title {
		return false
	}
	if this.Description != that1.Description {
		return false
	}
	if this.Contract != that1.Contract {
		return false
	}
	return true
}
func (m *PromoteToPrivilegedContractProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PromoteToPrivilegedContractProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PromoteToPrivilegedContractProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Contract) > 0 {
		i -= len(m.Contract)
		copy(dAtA[i:], m.Contract)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Contract)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *DemotePrivilegedContractProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DemotePrivilegedContractProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DemotePrivilegedContractProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Contract) > 0 {
		i -= len(m.Contract)
		copy(dAtA[i:], m.Contract)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Contract)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintProposal(dAtA []byte, offset int, v uint64) int {
	offset -= sovProposal(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PromoteToPrivilegedContractProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	l = len(m.Contract)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	return n
}

func (m *DemotePrivilegedContractProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	l = len(m.Contract)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	return n
}

func sovProposal(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozProposal(x uint64) (n int) {
	return sovProposal(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PromoteToPrivilegedContractProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProposal
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
			return fmt.Errorf("proto: PromoteToPrivilegedContractProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PromoteToPrivilegedContractProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Contract", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Contract = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProposal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProposal
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
func (m *DemotePrivilegedContractProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProposal
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
			return fmt.Errorf("proto: DemotePrivilegedContractProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DemotePrivilegedContractProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Contract", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Contract = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProposal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProposal
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
func skipProposal(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProposal
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
					return 0, ErrIntOverflowProposal
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
					return 0, ErrIntOverflowProposal
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
				return 0, ErrInvalidLengthProposal
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupProposal
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthProposal
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthProposal        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProposal          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupProposal = fmt.Errorf("proto: unexpected end of group")
)
