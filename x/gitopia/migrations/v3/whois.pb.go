// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: gitopia/whois.proto

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

type OwnerType int32

const (
	OwnerType_USER OwnerType = 0
	OwnerType_DAO  OwnerType = 1
)

var OwnerType_name = map[int32]string{
	0: "USER",
	1: "DAO",
}

var OwnerType_value = map[string]int32{
	"USER": 0,
	"DAO":  1,
}

func (x OwnerType) String() string {
	return proto.EnumName(OwnerType_name, int32(x))
}

func (OwnerType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_eb936eb6838a72a6, []int{0}
}

type Whois struct {
	Creator   string    `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Id        uint64    `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Name      string    `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Address   string    `protobuf:"bytes,4,opt,name=address,proto3" json:"address,omitempty"`
	OwnerType OwnerType `protobuf:"varint,5,opt,name=ownerType,proto3,enum=gitopia.gitopia.gitopia.v3.OwnerType" json:"ownerType,omitempty"`
}

func (m *Whois) Reset()         { *m = Whois{} }
func (m *Whois) String() string { return proto.CompactTextString(m) }
func (*Whois) ProtoMessage()    {}
func (*Whois) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb936eb6838a72a6, []int{0}
}
func (m *Whois) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Whois) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Whois.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Whois) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Whois.Merge(m, src)
}
func (m *Whois) XXX_Size() int {
	return m.Size()
}
func (m *Whois) XXX_DiscardUnknown() {
	xxx_messageInfo_Whois.DiscardUnknown(m)
}

var xxx_messageInfo_Whois proto.InternalMessageInfo

func (m *Whois) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *Whois) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Whois) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Whois) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Whois) GetOwnerType() OwnerType {
	if m != nil {
		return m.OwnerType
	}
	return OwnerType_USER
}

func init() {
	proto.RegisterEnum("gitopia.gitopia.gitopia.v3.OwnerType", OwnerType_name, OwnerType_value)
	proto.RegisterType((*Whois)(nil), "gitopia.gitopia.gitopia.v3.Whois")
}

func init() { proto.RegisterFile("gitopia/whois.proto", fileDescriptor_eb936eb6838a72a6) }

var fileDescriptor_eb936eb6838a72a6 = []byte{
	// 267 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4e, 0xcf, 0x2c, 0xc9,
	0x2f, 0xc8, 0x4c, 0xd4, 0x2f, 0xcf, 0xc8, 0xcf, 0x2c, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x92, 0x82, 0x0a, 0xea, 0xa1, 0xd3, 0x65, 0xc6, 0x52, 0x22, 0xe9, 0xf9, 0xe9, 0xf9, 0x60, 0x65,
	0xfa, 0x20, 0x16, 0x44, 0x87, 0xd2, 0x12, 0x46, 0x2e, 0xd6, 0x70, 0x90, 0x09, 0x42, 0x12, 0x5c,
	0xec, 0xc9, 0x45, 0xa9, 0x89, 0x25, 0xf9, 0x45, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x30,
	0xae, 0x10, 0x1f, 0x17, 0x53, 0x66, 0x8a, 0x04, 0x93, 0x02, 0xa3, 0x06, 0x4b, 0x10, 0x53, 0x66,
	0x8a, 0x90, 0x10, 0x17, 0x4b, 0x5e, 0x62, 0x6e, 0xaa, 0x04, 0x33, 0x58, 0x19, 0x98, 0x0d, 0xd2,
	0x9d, 0x98, 0x92, 0x52, 0x94, 0x5a, 0x5c, 0x2c, 0xc1, 0x02, 0xd1, 0x0d, 0xe5, 0x0a, 0x39, 0x73,
	0x71, 0xe6, 0x97, 0xe7, 0xa5, 0x16, 0x85, 0x54, 0x16, 0xa4, 0x4a, 0xb0, 0x2a, 0x30, 0x6a, 0xf0,
	0x19, 0xa9, 0xea, 0xe1, 0x76, 0xa7, 0x9e, 0x3f, 0x4c, 0x71, 0x10, 0x42, 0x9f, 0x96, 0x1c, 0x17,
	0x27, 0x5c, 0x5c, 0x88, 0x83, 0x8b, 0x25, 0x34, 0xd8, 0x35, 0x48, 0x80, 0x41, 0x88, 0x9d, 0x8b,
	0xd9, 0xc5, 0xd1, 0x5f, 0x80, 0xd1, 0x29, 0xe8, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18,
	0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e, 0xe1, 0xc2, 0x63, 0x39, 0x86, 0x1b, 0x8f, 0xe5,
	0x18, 0xa2, 0x2c, 0xd2, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0x61, 0x41,
	0x06, 0xa3, 0x2b, 0xe0, 0xac, 0xdc, 0xcc, 0xf4, 0xa2, 0xc4, 0x92, 0xcc, 0xfc, 0xbc, 0x62, 0xfd,
	0x32, 0x63, 0xeb, 0x92, 0xca, 0x82, 0xd4, 0xe2, 0x24, 0x36, 0x70, 0x08, 0x19, 0x03, 0x02, 0x00,
	0x00, 0xff, 0xff, 0x85, 0x2e, 0x19, 0x75, 0x6a, 0x01, 0x00, 0x00,
}

func (m *Whois) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Whois) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Whois) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.OwnerType != 0 {
		i = encodeVarintWhois(dAtA, i, uint64(m.OwnerType))
		i--
		dAtA[i] = 0x28
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintWhois(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintWhois(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Id != 0 {
		i = encodeVarintWhois(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintWhois(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintWhois(dAtA []byte, offset int, v uint64) int {
	offset -= sovWhois(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Whois) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovWhois(uint64(l))
	}
	if m.Id != 0 {
		n += 1 + sovWhois(uint64(m.Id))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovWhois(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovWhois(uint64(l))
	}
	if m.OwnerType != 0 {
		n += 1 + sovWhois(uint64(m.OwnerType))
	}
	return n
}

func sovWhois(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozWhois(x uint64) (n int) {
	return sovWhois(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Whois) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWhois
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
			return fmt.Errorf("proto: Whois: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Whois: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWhois
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
				return ErrInvalidLengthWhois
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWhois
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWhois
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWhois
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
				return ErrInvalidLengthWhois
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWhois
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWhois
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
				return ErrInvalidLengthWhois
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWhois
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OwnerType", wireType)
			}
			m.OwnerType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWhois
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.OwnerType |= OwnerType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipWhois(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthWhois
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
func skipWhois(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowWhois
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
					return 0, ErrIntOverflowWhois
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
					return 0, ErrIntOverflowWhois
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
				return 0, ErrInvalidLengthWhois
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupWhois
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthWhois
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthWhois        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowWhois          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupWhois = fmt.Errorf("proto: unexpected end of group")
)
