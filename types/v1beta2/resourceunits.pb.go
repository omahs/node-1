// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: akash/base/v1beta2/resourceunits.proto

package v1beta2

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

// ResourceUnits describes all available resources types for deployment/node etc
// if field is nil resource is not present in the given data-structure
type ResourceUnits struct {
	CPU       *CPU      `protobuf:"bytes,1,opt,name=cpu,proto3" json:"cpu,omitempty" yaml:"cpu,omitempty"`
	Memory    *Memory   `protobuf:"bytes,2,opt,name=memory,proto3" json:"memory,omitempty" yaml:"memory,omitempty"`
	Storage   Volumes   `protobuf:"bytes,3,rep,name=storage,proto3,castrepeated=Volumes" json:"storage,omitempty" yaml:"storage,omitempty"`
	Endpoints Endpoints `protobuf:"bytes,4,rep,name=endpoints,proto3,castrepeated=Endpoints" json:"endpoints" yaml:"endpoints"`
}

func (m *ResourceUnits) Reset()         { *m = ResourceUnits{} }
func (m *ResourceUnits) String() string { return proto.CompactTextString(m) }
func (*ResourceUnits) ProtoMessage()    {}
func (*ResourceUnits) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a653b54b68ae16d, []int{0}
}
func (m *ResourceUnits) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ResourceUnits) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ResourceUnits.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ResourceUnits) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResourceUnits.Merge(m, src)
}
func (m *ResourceUnits) XXX_Size() int {
	return m.Size()
}
func (m *ResourceUnits) XXX_DiscardUnknown() {
	xxx_messageInfo_ResourceUnits.DiscardUnknown(m)
}

var xxx_messageInfo_ResourceUnits proto.InternalMessageInfo

func (m *ResourceUnits) GetCPU() *CPU {
	if m != nil {
		return m.CPU
	}
	return nil
}

func (m *ResourceUnits) GetMemory() *Memory {
	if m != nil {
		return m.Memory
	}
	return nil
}

func (m *ResourceUnits) GetStorage() Volumes {
	if m != nil {
		return m.Storage
	}
	return nil
}

func (m *ResourceUnits) GetEndpoints() Endpoints {
	if m != nil {
		return m.Endpoints
	}
	return nil
}

func init() {
	proto.RegisterType((*ResourceUnits)(nil), "akash.base.v1beta2.ResourceUnits")
}

func init() {
	proto.RegisterFile("akash/base/v1beta2/resourceunits.proto", fileDescriptor_3a653b54b68ae16d)
}

var fileDescriptor_3a653b54b68ae16d = []byte{
	// 397 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xcf, 0x8a, 0xda, 0x50,
	0x14, 0x87, 0x93, 0x46, 0x14, 0x23, 0x82, 0x0d, 0x82, 0xc1, 0x96, 0x5c, 0x9b, 0x45, 0x11, 0xda,
	0x26, 0x34, 0x76, 0xe5, 0xaa, 0x44, 0x5c, 0x16, 0x24, 0xc5, 0x2e, 0xba, 0x19, 0x92, 0x78, 0x89,
	0x41, 0x93, 0x1b, 0x72, 0x6f, 0x46, 0xf2, 0x16, 0xf3, 0x08, 0xb3, 0x9e, 0x27, 0xc9, 0xd2, 0xe5,
	0xac, 0xee, 0x0c, 0x71, 0x33, 0xb8, 0xf4, 0x09, 0x86, 0xfc, 0x1b, 0x71, 0x14, 0x77, 0x97, 0x73,
	0xbe, 0xdf, 0xf9, 0x0e, 0x97, 0xc3, 0x7f, 0x35, 0x57, 0x26, 0x5e, 0xaa, 0x96, 0x89, 0xa1, 0x7a,
	0xfb, 0xd3, 0x82, 0xc4, 0xd4, 0xd4, 0x10, 0x62, 0x14, 0x85, 0x36, 0x8c, 0x7c, 0x97, 0x60, 0x25,
	0x08, 0x11, 0x41, 0x82, 0x90, 0x73, 0x4a, 0xc6, 0x29, 0x25, 0xd7, 0xef, 0x3a, 0xc8, 0x41, 0x79,
	0x5b, 0xcd, 0x5e, 0x05, 0xd9, 0xff, 0x72, 0x65, 0xe2, 0x15, 0x04, 0xfa, 0x8b, 0x00, 0xb9, 0x3e,
	0x29, 0x10, 0x39, 0xe1, 0xf8, 0xb6, 0x51, 0xa6, 0xe6, 0xd9, 0x1e, 0xc2, 0x0d, 0xcf, 0xd9, 0x41,
	0x24, 0xb2, 0x03, 0x76, 0xd8, 0xd2, 0x7a, 0xca, 0xf9, 0x3e, 0xca, 0x64, 0x36, 0xd7, 0x7f, 0x25,
	0x14, 0xb0, 0x29, 0x05, 0xdc, 0x64, 0x36, 0xdf, 0x53, 0xd0, 0xb6, 0x83, 0xe8, 0x3b, 0xf2, 0x5c,
	0x02, 0xbd, 0x80, 0xc4, 0x07, 0x0a, 0xba, 0xb1, 0xe9, 0xad, 0xc7, 0xf2, 0x49, 0x59, 0x36, 0xb2,
	0xc9, 0x82, 0xc3, 0xd7, 0x3d, 0xe8, 0xa1, 0x30, 0x16, 0x3f, 0xe4, 0x8e, 0xfe, 0x25, 0xc7, 0x9f,
	0x9c, 0xd0, 0x47, 0x99, 0x66, 0x4f, 0x41, 0xa7, 0x48, 0x9c, 0x28, 0x7a, 0x85, 0xe2, 0x7d, 0x47,
	0x36, 0xca, 0xf1, 0xc2, 0x86, 0x6f, 0x60, 0x82, 0x42, 0xd3, 0x81, 0x22, 0x37, 0xe0, 0x86, 0x2d,
	0xed, 0xd3, 0x25, 0xd3, 0xdf, 0x02, 0xd1, 0x7f, 0x27, 0x14, 0x30, 0x7b, 0x0a, 0x3e, 0x96, 0x99,
	0x13, 0x97, 0x58, 0xb8, 0xce, 0x5a, 0xf2, 0xc3, 0x13, 0x68, 0xfc, 0x43, 0xeb, 0xc8, 0x83, 0xd8,
	0xa8, 0x6c, 0x82, 0xcf, 0x37, 0xab, 0x6f, 0xc6, 0x62, 0x2d, 0x57, 0x7f, 0xbe, 0xa4, 0x9e, 0x96,
	0x90, 0xae, 0x95, 0xee, 0x63, 0xec, 0x40, 0x41, 0xa7, 0x70, 0xbe, 0x95, 0x32, 0x57, 0xb3, 0x8a,
	0x60, 0xe3, 0xc8, 0x8e, 0x6b, 0x2f, 0xf7, 0x80, 0xd5, 0xa7, 0x49, 0x2a, 0xb1, 0xdb, 0x54, 0x62,
	0x9f, 0x53, 0x89, 0xbd, 0xdb, 0x49, 0xcc, 0x76, 0x27, 0x31, 0x8f, 0x3b, 0x89, 0xf9, 0xff, 0xcd,
	0x71, 0xc9, 0x32, 0xb2, 0x14, 0x1b, 0x79, 0x6a, 0xbe, 0xc6, 0x0f, 0x1f, 0x92, 0x0d, 0x0a, 0x57,
	0xaa, 0x8f, 0x16, 0x50, 0x25, 0x71, 0x00, 0x71, 0x75, 0x20, 0x56, 0x3d, 0x3f, 0x8c, 0xd1, 0x6b,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x8c, 0xe5, 0x6d, 0x41, 0xb2, 0x02, 0x00, 0x00,
}

func (this *ResourceUnits) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ResourceUnits)
	if !ok {
		that2, ok := that.(ResourceUnits)
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
	if !this.CPU.Equal(that1.CPU) {
		return false
	}
	if !this.Memory.Equal(that1.Memory) {
		return false
	}
	if len(this.Storage) != len(that1.Storage) {
		return false
	}
	for i := range this.Storage {
		if !this.Storage[i].Equal(&that1.Storage[i]) {
			return false
		}
	}
	if len(this.Endpoints) != len(that1.Endpoints) {
		return false
	}
	for i := range this.Endpoints {
		if !this.Endpoints[i].Equal(&that1.Endpoints[i]) {
			return false
		}
	}
	return true
}
func (m *ResourceUnits) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ResourceUnits) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ResourceUnits) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Endpoints) > 0 {
		for iNdEx := len(m.Endpoints) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Endpoints[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintResourceunits(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Storage) > 0 {
		for iNdEx := len(m.Storage) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Storage[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintResourceunits(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.Memory != nil {
		{
			size, err := m.Memory.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintResourceunits(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.CPU != nil {
		{
			size, err := m.CPU.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintResourceunits(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintResourceunits(dAtA []byte, offset int, v uint64) int {
	offset -= sovResourceunits(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ResourceUnits) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CPU != nil {
		l = m.CPU.Size()
		n += 1 + l + sovResourceunits(uint64(l))
	}
	if m.Memory != nil {
		l = m.Memory.Size()
		n += 1 + l + sovResourceunits(uint64(l))
	}
	if len(m.Storage) > 0 {
		for _, e := range m.Storage {
			l = e.Size()
			n += 1 + l + sovResourceunits(uint64(l))
		}
	}
	if len(m.Endpoints) > 0 {
		for _, e := range m.Endpoints {
			l = e.Size()
			n += 1 + l + sovResourceunits(uint64(l))
		}
	}
	return n
}

func sovResourceunits(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozResourceunits(x uint64) (n int) {
	return sovResourceunits(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ResourceUnits) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowResourceunits
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
			return fmt.Errorf("proto: ResourceUnits: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ResourceUnits: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CPU", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResourceunits
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
				return ErrInvalidLengthResourceunits
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthResourceunits
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.CPU == nil {
				m.CPU = &CPU{}
			}
			if err := m.CPU.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Memory", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResourceunits
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
				return ErrInvalidLengthResourceunits
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthResourceunits
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Memory == nil {
				m.Memory = &Memory{}
			}
			if err := m.Memory.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Storage", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResourceunits
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
				return ErrInvalidLengthResourceunits
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthResourceunits
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Storage = append(m.Storage, Storage{})
			if err := m.Storage[len(m.Storage)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Endpoints", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowResourceunits
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
				return ErrInvalidLengthResourceunits
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthResourceunits
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Endpoints = append(m.Endpoints, Endpoint{})
			if err := m.Endpoints[len(m.Endpoints)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipResourceunits(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthResourceunits
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
func skipResourceunits(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowResourceunits
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
					return 0, ErrIntOverflowResourceunits
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
					return 0, ErrIntOverflowResourceunits
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
				return 0, ErrInvalidLengthResourceunits
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupResourceunits
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthResourceunits
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthResourceunits        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowResourceunits          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupResourceunits = fmt.Errorf("proto: unexpected end of group")
)
