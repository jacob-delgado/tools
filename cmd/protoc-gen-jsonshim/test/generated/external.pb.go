// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: external.proto

package generated

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

type ExternalSimple struct {
	FieldC               uint32                         `protobuf:"varint,1,opt,name=fieldC,proto3" json:"fieldC,omitempty"`
	FieldD               *ExternalSimple_ExternalNested `protobuf:"bytes,2,opt,name=fieldD,proto3" json:"fieldD,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *ExternalSimple) Reset()         { *m = ExternalSimple{} }
func (m *ExternalSimple) String() string { return proto.CompactTextString(m) }
func (*ExternalSimple) ProtoMessage()    {}
func (*ExternalSimple) Descriptor() ([]byte, []int) {
	return fileDescriptor_2b7268f56e161ef5, []int{0}
}
func (m *ExternalSimple) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ExternalSimple) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ExternalSimple.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ExternalSimple) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExternalSimple.Merge(m, src)
}
func (m *ExternalSimple) XXX_Size() int {
	return m.Size()
}
func (m *ExternalSimple) XXX_DiscardUnknown() {
	xxx_messageInfo_ExternalSimple.DiscardUnknown(m)
}

var xxx_messageInfo_ExternalSimple proto.InternalMessageInfo

func (m *ExternalSimple) GetFieldC() uint32 {
	if m != nil {
		return m.FieldC
	}
	return 0
}

func (m *ExternalSimple) GetFieldD() *ExternalSimple_ExternalNested {
	if m != nil {
		return m.FieldD
	}
	return nil
}

type ExternalSimple_ExternalNested struct {
	FieldA               map[string]string `protobuf:"bytes,1,rep,name=fieldA,proto3" json:"fieldA,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ExternalSimple_ExternalNested) Reset()         { *m = ExternalSimple_ExternalNested{} }
func (m *ExternalSimple_ExternalNested) String() string { return proto.CompactTextString(m) }
func (*ExternalSimple_ExternalNested) ProtoMessage()    {}
func (*ExternalSimple_ExternalNested) Descriptor() ([]byte, []int) {
	return fileDescriptor_2b7268f56e161ef5, []int{0, 0}
}
func (m *ExternalSimple_ExternalNested) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ExternalSimple_ExternalNested) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ExternalSimple_ExternalNested.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ExternalSimple_ExternalNested) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExternalSimple_ExternalNested.Merge(m, src)
}
func (m *ExternalSimple_ExternalNested) XXX_Size() int {
	return m.Size()
}
func (m *ExternalSimple_ExternalNested) XXX_DiscardUnknown() {
	xxx_messageInfo_ExternalSimple_ExternalNested.DiscardUnknown(m)
}

var xxx_messageInfo_ExternalSimple_ExternalNested proto.InternalMessageInfo

func (m *ExternalSimple_ExternalNested) GetFieldA() map[string]string {
	if m != nil {
		return m.FieldA
	}
	return nil
}

func init() {
	proto.RegisterType((*ExternalSimple)(nil), "istio.tools.test.ExternalSimple")
	proto.RegisterType((*ExternalSimple_ExternalNested)(nil), "istio.tools.test.ExternalSimple.ExternalNested")
	proto.RegisterMapType((map[string]string)(nil), "istio.tools.test.ExternalSimple.ExternalNested.FieldAEntry")
}

func init() { proto.RegisterFile("external.proto", fileDescriptor_2b7268f56e161ef5) }

var fileDescriptor_2b7268f56e161ef5 = []byte{
	// 225 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4b, 0xad, 0x28, 0x49,
	0x2d, 0xca, 0x4b, 0xcc, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0xc8, 0x2c, 0x2e, 0xc9,
	0xcc, 0xd7, 0x2b, 0xc9, 0xcf, 0xcf, 0x29, 0xd6, 0x2b, 0x49, 0x2d, 0x2e, 0x51, 0x9a, 0xc2, 0xc4,
	0xc5, 0xe7, 0x0a, 0x55, 0x14, 0x9c, 0x99, 0x5b, 0x90, 0x93, 0x2a, 0x24, 0xc6, 0xc5, 0x96, 0x96,
	0x99, 0x9a, 0x93, 0xe2, 0x2c, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x1b, 0x04, 0xe5, 0x09, 0xb9, 0x43,
	0xc5, 0x5d, 0x24, 0x98, 0x14, 0x18, 0x35, 0xb8, 0x8d, 0xf4, 0xf5, 0xd0, 0x4d, 0xd3, 0x43, 0x35,
	0x09, 0xce, 0xf5, 0x4b, 0x2d, 0x2e, 0x49, 0x4d, 0x81, 0x1a, 0xe4, 0x22, 0xb5, 0x80, 0x11, 0x61,
	0x27, 0x44, 0x4a, 0x28, 0x18, 0x6a, 0xb6, 0xa3, 0x04, 0xa3, 0x02, 0xb3, 0x06, 0xb7, 0x91, 0x35,
	0x89, 0x66, 0xeb, 0xb9, 0x81, 0x75, 0xbb, 0xe6, 0x95, 0x14, 0x55, 0x42, 0xed, 0x71, 0x94, 0xb2,
	0xe4, 0xe2, 0x46, 0x12, 0x16, 0x12, 0xe0, 0x62, 0xce, 0x4e, 0xad, 0x04, 0x7b, 0x8a, 0x33, 0x08,
	0xc4, 0x14, 0x12, 0xe1, 0x62, 0x2d, 0x4b, 0xcc, 0x29, 0x4d, 0x05, 0x7b, 0x88, 0x33, 0x08, 0xc2,
	0xb1, 0x62, 0xb2, 0x60, 0x74, 0x12, 0x3f, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07,
	0x8f, 0xe4, 0x18, 0xa3, 0x38, 0xd3, 0x53, 0xf3, 0x52, 0x8b, 0x12, 0x4b, 0x52, 0x53, 0x92, 0xd8,
	0xc0, 0x01, 0x69, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x5b, 0x0d, 0x9c, 0xbe, 0x5a, 0x01, 0x00,
	0x00,
}

func (m *ExternalSimple) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ExternalSimple) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ExternalSimple) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.FieldD != nil {
		{
			size, err := m.FieldD.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintExternal(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.FieldC != 0 {
		i = encodeVarintExternal(dAtA, i, uint64(m.FieldC))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ExternalSimple_ExternalNested) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ExternalSimple_ExternalNested) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ExternalSimple_ExternalNested) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.FieldA) > 0 {
		for k := range m.FieldA {
			v := m.FieldA[k]
			baseI := i
			i -= len(v)
			copy(dAtA[i:], v)
			i = encodeVarintExternal(dAtA, i, uint64(len(v)))
			i--
			dAtA[i] = 0x12
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintExternal(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintExternal(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintExternal(dAtA []byte, offset int, v uint64) int {
	offset -= sovExternal(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ExternalSimple) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.FieldC != 0 {
		n += 1 + sovExternal(uint64(m.FieldC))
	}
	if m.FieldD != nil {
		l = m.FieldD.Size()
		n += 1 + l + sovExternal(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ExternalSimple_ExternalNested) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.FieldA) > 0 {
		for k, v := range m.FieldA {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovExternal(uint64(len(k))) + 1 + len(v) + sovExternal(uint64(len(v)))
			n += mapEntrySize + 1 + sovExternal(uint64(mapEntrySize))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovExternal(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozExternal(x uint64) (n int) {
	return sovExternal(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ExternalSimple) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowExternal
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
			return fmt.Errorf("proto: ExternalSimple: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ExternalSimple: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FieldC", wireType)
			}
			m.FieldC = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExternal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FieldC |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FieldD", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExternal
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
				return ErrInvalidLengthExternal
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthExternal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.FieldD == nil {
				m.FieldD = &ExternalSimple_ExternalNested{}
			}
			if err := m.FieldD.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipExternal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthExternal
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthExternal
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ExternalSimple_ExternalNested) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowExternal
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
			return fmt.Errorf("proto: ExternalNested: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ExternalNested: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FieldA", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowExternal
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
				return ErrInvalidLengthExternal
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthExternal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.FieldA == nil {
				m.FieldA = make(map[string]string)
			}
			var mapkey string
			var mapvalue string
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowExternal
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowExternal
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthExternal
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthExternal
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var stringLenmapvalue uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowExternal
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapvalue |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapvalue := int(stringLenmapvalue)
					if intStringLenmapvalue < 0 {
						return ErrInvalidLengthExternal
					}
					postStringIndexmapvalue := iNdEx + intStringLenmapvalue
					if postStringIndexmapvalue < 0 {
						return ErrInvalidLengthExternal
					}
					if postStringIndexmapvalue > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = string(dAtA[iNdEx:postStringIndexmapvalue])
					iNdEx = postStringIndexmapvalue
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipExternal(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthExternal
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.FieldA[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipExternal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthExternal
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthExternal
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipExternal(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowExternal
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
					return 0, ErrIntOverflowExternal
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowExternal
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
				return 0, ErrInvalidLengthExternal
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthExternal
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowExternal
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipExternal(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthExternal
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthExternal = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowExternal   = fmt.Errorf("proto: integer overflow")
)
