// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: gitopia/member.proto

package types

import (
	fmt "fmt"
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

type MemberRole int32

const (
	MemberRole_MEMBER MemberRole = 0
	MemberRole_OWNER  MemberRole = 1
)

var MemberRole_name = map[int32]string{
	0: "MEMBER",
	1: "OWNER",
}

var MemberRole_value = map[string]int32{
	"MEMBER": 0,
	"OWNER":  1,
}

func (x MemberRole) String() string {
	return proto.EnumName(MemberRole_name, int32(x))
}

func (MemberRole) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_eb50d9e34f0ece82, []int{0}
}

type Member struct {
	Id         uint64     `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Address    string     `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	DaoAddress string     `protobuf:"bytes,3,opt,name=daoAddress,proto3" json:"daoAddress,omitempty"`
	Role       MemberRole `protobuf:"varint,4,opt,name=role,proto3,enum=gitopia.gitopia.gitopia.MemberRole" json:"role,omitempty"`
}

func (m *Member) Reset()         { *m = Member{} }
func (m *Member) String() string { return proto.CompactTextString(m) }
func (*Member) ProtoMessage()    {}
func (*Member) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb50d9e34f0ece82, []int{0}
}
func (m *Member) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Member) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Member.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Member) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Member.Merge(m, src)
}
func (m *Member) XXX_Size() int {
	return m.Size()
}
func (m *Member) XXX_DiscardUnknown() {
	xxx_messageInfo_Member.DiscardUnknown(m)
}

var xxx_messageInfo_Member proto.InternalMessageInfo

func (m *Member) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Member) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Member) GetDaoAddress() string {
	if m != nil {
		return m.DaoAddress
	}
	return ""
}

func (m *Member) GetRole() MemberRole {
	if m != nil {
		return m.Role
	}
	return MemberRole_MEMBER
}

func init() {
	proto.RegisterEnum("gitopia.gitopia.gitopia.MemberRole", MemberRole_name, MemberRole_value)
	proto.RegisterType((*Member)(nil), "gitopia.gitopia.gitopia.Member")
}

func init() { proto.RegisterFile("gitopia/member.proto", fileDescriptor_eb50d9e34f0ece82) }

var fileDescriptor_eb50d9e34f0ece82 = []byte{
	// 227 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x49, 0xcf, 0x2c, 0xc9,
	0x2f, 0xc8, 0x4c, 0xd4, 0xcf, 0x4d, 0xcd, 0x4d, 0x4a, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0x12, 0x87, 0x8a, 0xea, 0xa1, 0xd1, 0x4a, 0xdd, 0x8c, 0x5c, 0x6c, 0xbe, 0x60, 0x95, 0x42,
	0x7c, 0x5c, 0x4c, 0x99, 0x29, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x2c, 0x41, 0x4c, 0x99, 0x29, 0x42,
	0x12, 0x5c, 0xec, 0x89, 0x29, 0x29, 0x45, 0xa9, 0xc5, 0xc5, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0x9c,
	0x41, 0x30, 0xae, 0x90, 0x1c, 0x17, 0x57, 0x4a, 0x62, 0xbe, 0x23, 0x54, 0x92, 0x19, 0x2c, 0x89,
	0x24, 0x22, 0x64, 0xce, 0xc5, 0x52, 0x94, 0x9f, 0x93, 0x2a, 0xc1, 0xa2, 0xc0, 0xa8, 0xc1, 0x67,
	0xa4, 0xac, 0x87, 0xc3, 0x72, 0x3d, 0x88, 0xc5, 0x41, 0xf9, 0x39, 0xa9, 0x41, 0x60, 0x0d, 0x5a,
	0xca, 0x5c, 0x5c, 0x08, 0x31, 0x21, 0x2e, 0x2e, 0x36, 0x5f, 0x57, 0x5f, 0x27, 0xd7, 0x20, 0x01,
	0x06, 0x21, 0x4e, 0x2e, 0x56, 0xff, 0x70, 0x3f, 0xd7, 0x20, 0x01, 0x46, 0x27, 0x97, 0x13, 0x8f,
	0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86, 0x0b,
	0x8f, 0xe5, 0x18, 0x6e, 0x3c, 0x96, 0x63, 0x88, 0xd2, 0x4a, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2,
	0x4b, 0xce, 0xcf, 0xd5, 0x87, 0x05, 0x03, 0x8c, 0xae, 0x80, 0xb3, 0x4a, 0x2a, 0x0b, 0x52, 0x8b,
	0x93, 0xd8, 0xc0, 0x01, 0x63, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xe7, 0x6d, 0x78, 0x5d, 0x30,
	0x01, 0x00, 0x00,
}

func (m *Member) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Member) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Member) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Role != 0 {
		i = encodeVarintMember(dAtA, i, uint64(m.Role))
		i--
		dAtA[i] = 0x20
	}
	if len(m.DaoAddress) > 0 {
		i -= len(m.DaoAddress)
		copy(dAtA[i:], m.DaoAddress)
		i = encodeVarintMember(dAtA, i, uint64(len(m.DaoAddress)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintMember(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintMember(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintMember(dAtA []byte, offset int, v uint64) int {
	offset -= sovMember(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Member) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovMember(uint64(m.Id))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovMember(uint64(l))
	}
	l = len(m.DaoAddress)
	if l > 0 {
		n += 1 + l + sovMember(uint64(l))
	}
	if m.Role != 0 {
		n += 1 + sovMember(uint64(m.Role))
	}
	return n
}

func sovMember(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMember(x uint64) (n int) {
	return sovMember(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Member) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMember
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
			return fmt.Errorf("proto: Member: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Member: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMember
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
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMember
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
				return ErrInvalidLengthMember
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMember
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DaoAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMember
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
				return ErrInvalidLengthMember
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMember
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DaoAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Role", wireType)
			}
			m.Role = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMember
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Role |= MemberRole(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMember(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMember
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
func skipMember(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMember
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
					return 0, ErrIntOverflowMember
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
					return 0, ErrIntOverflowMember
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
				return 0, ErrInvalidLengthMember
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMember
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMember
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMember        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMember          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMember = fmt.Errorf("proto: unexpected end of group")
)
