// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: secret/registration/v1beta1/msg.proto

package types

import (
	bytes "bytes"
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_scrtlabs_SecretNetwork_x_registration_remote_attestation "github.com/scrtlabs/SecretNetwork/x/registration/remote_attestation"
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

type RaAuthenticate struct {
	Sender      github_com_cosmos_cosmos_sdk_types.AccAddress                                   `protobuf:"bytes,1,opt,name=sender,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"sender,omitempty"`
	Certificate github_com_scrtlabs_SecretNetwork_x_registration_remote_attestation.Certificate `protobuf:"bytes,2,opt,name=certificate,proto3,casttype=github.com/scrtlabs/SecretNetwork/x/registration/remote_attestation.Certificate" json:"ra_cert"`
}

func (m *RaAuthenticate) Reset()         { *m = RaAuthenticate{} }
func (m *RaAuthenticate) String() string { return proto.CompactTextString(m) }
func (*RaAuthenticate) ProtoMessage()    {}
func (*RaAuthenticate) Descriptor() ([]byte, []int) {
	return fileDescriptor_91e653c4cfa6dfea, []int{0}
}
func (m *RaAuthenticate) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RaAuthenticate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RaAuthenticate.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RaAuthenticate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RaAuthenticate.Merge(m, src)
}
func (m *RaAuthenticate) XXX_Size() int {
	return m.Size()
}
func (m *RaAuthenticate) XXX_DiscardUnknown() {
	xxx_messageInfo_RaAuthenticate.DiscardUnknown(m)
}

var xxx_messageInfo_RaAuthenticate proto.InternalMessageInfo

type MasterKey struct {
	Bytes []byte `protobuf:"bytes,1,opt,name=bytes,proto3" json:"bytes,omitempty"`
}

func (m *MasterKey) Reset()         { *m = MasterKey{} }
func (m *MasterKey) String() string { return proto.CompactTextString(m) }
func (*MasterKey) ProtoMessage()    {}
func (*MasterKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_91e653c4cfa6dfea, []int{1}
}
func (m *MasterKey) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MasterKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MasterKey.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MasterKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MasterKey.Merge(m, src)
}
func (m *MasterKey) XXX_Size() int {
	return m.Size()
}
func (m *MasterKey) XXX_DiscardUnknown() {
	xxx_messageInfo_MasterKey.DiscardUnknown(m)
}

var xxx_messageInfo_MasterKey proto.InternalMessageInfo

type Key struct {
	Key []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key"`
}

func (m *Key) Reset()         { *m = Key{} }
func (m *Key) String() string { return proto.CompactTextString(m) }
func (*Key) ProtoMessage()    {}
func (*Key) Descriptor() ([]byte, []int) {
	return fileDescriptor_91e653c4cfa6dfea, []int{2}
}
func (m *Key) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Key) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Key.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Key) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Key.Merge(m, src)
}
func (m *Key) XXX_Size() int {
	return m.Size()
}
func (m *Key) XXX_DiscardUnknown() {
	xxx_messageInfo_Key.DiscardUnknown(m)
}

var xxx_messageInfo_Key proto.InternalMessageInfo

func init() {
	proto.RegisterType((*RaAuthenticate)(nil), "secret.registration.v1beta1.RaAuthenticate")
	proto.RegisterType((*MasterKey)(nil), "secret.registration.v1beta1.MasterKey")
	proto.RegisterType((*Key)(nil), "secret.registration.v1beta1.Key")
}

func init() {
	proto.RegisterFile("secret/registration/v1beta1/msg.proto", fileDescriptor_91e653c4cfa6dfea)
}

var fileDescriptor_91e653c4cfa6dfea = []byte{
	// 344 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0xbf, 0x4e, 0xeb, 0x30,
	0x14, 0x87, 0xe3, 0x5b, 0xdd, 0x56, 0xd7, 0xf7, 0xea, 0x0e, 0x51, 0x87, 0x02, 0x92, 0x53, 0x2a,
	0x21, 0xb1, 0x34, 0x56, 0xc5, 0x03, 0xa0, 0x96, 0x09, 0x21, 0x40, 0x0a, 0x1b, 0x03, 0x95, 0xe3,
	0x1c, 0xd2, 0xa8, 0x4d, 0x5c, 0xd9, 0xa7, 0x40, 0x36, 0x1e, 0x81, 0xc7, 0xe0, 0x51, 0x3a, 0x76,
	0x64, 0x8a, 0x20, 0xdd, 0xfa, 0x08, 0x9d, 0x50, 0xfe, 0x48, 0x94, 0x91, 0xc5, 0xf6, 0xd1, 0xf9,
	0xf9, 0x3b, 0xfa, 0x6c, 0x7a, 0x64, 0x40, 0x6a, 0x40, 0xae, 0x21, 0x8c, 0x0c, 0x6a, 0x81, 0x91,
	0x4a, 0xf8, 0xc3, 0xc0, 0x07, 0x14, 0x03, 0x1e, 0x9b, 0xd0, 0x9d, 0x6b, 0x85, 0xca, 0x3e, 0xa8,
	0x62, 0xee, 0x6e, 0xcc, 0xad, 0x63, 0xfb, 0xed, 0x50, 0x85, 0xaa, 0xcc, 0xf1, 0xe2, 0x54, 0x5d,
	0xe9, 0x65, 0x84, 0xfe, 0xf7, 0xc4, 0x70, 0x81, 0x13, 0x48, 0x30, 0x92, 0x02, 0xc1, 0x3e, 0xa7,
	0x4d, 0x03, 0x49, 0x00, 0xba, 0x43, 0xba, 0xe4, 0xf8, 0xdf, 0x68, 0xb0, 0xcd, 0x9c, 0x7e, 0x18,
	0xe1, 0x64, 0xe1, 0xbb, 0x52, 0xc5, 0x5c, 0x2a, 0x13, 0x2b, 0x53, 0x6f, 0x7d, 0x13, 0x4c, 0x39,
	0xa6, 0x73, 0x30, 0xee, 0x50, 0xca, 0x61, 0x10, 0x68, 0x30, 0xc6, 0xab, 0x01, 0xf6, 0x33, 0xa1,
	0x7f, 0x25, 0x68, 0x8c, 0xee, 0x4b, 0x74, 0xe7, 0x57, 0x09, 0xbc, 0xdb, 0x64, 0x4e, 0x4b, 0x8b,
	0x71, 0xd1, 0xd9, 0x66, 0xce, 0xf5, 0x0e, 0xdb, 0x48, 0x8d, 0x33, 0xe1, 0x1b, 0x7e, 0x53, 0x9a,
	0x5c, 0x01, 0x3e, 0x2a, 0x3d, 0xe5, 0x4f, 0xdf, 0xcd, 0x35, 0xc4, 0x0a, 0x61, 0x2c, 0x10, 0xc1,
	0x60, 0x65, 0x79, 0xf6, 0x35, 0xc5, 0xdb, 0x1d, 0xd9, 0x3b, 0xa4, 0x7f, 0x2e, 0x85, 0x41, 0xd0,
	0x17, 0x90, 0xda, 0x6d, 0xfa, 0xdb, 0x4f, 0x11, 0x4c, 0x65, 0xe6, 0x55, 0x45, 0xaf, 0x4b, 0x1b,
	0x45, 0x73, 0x8f, 0x36, 0xa6, 0x90, 0xd6, 0xd2, 0xad, 0x4d, 0xe6, 0x14, 0xa5, 0x57, 0x2c, 0x23,
	0xb1, 0xfc, 0x60, 0xd6, 0x6b, 0xce, 0xc8, 0x32, 0x67, 0x64, 0x95, 0x33, 0xf2, 0x9e, 0x33, 0xf2,
	0xb2, 0x66, 0xd6, 0x6a, 0xcd, 0xac, 0xb7, 0x35, 0xb3, 0x6e, 0x4f, 0x7f, 0x2c, 0x11, 0x25, 0x08,
	0x3a, 0x11, 0xb3, 0xea, 0xf5, 0xfc, 0x66, 0xf9, 0x1f, 0x27, 0x9f, 0x01, 0x00, 0x00, 0xff, 0xff,
	0xc1, 0x35, 0x10, 0xbc, 0xeb, 0x01, 0x00, 0x00,
}

func (this *RaAuthenticate) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RaAuthenticate)
	if !ok {
		that2, ok := that.(RaAuthenticate)
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
	if !bytes.Equal(this.Sender, that1.Sender) {
		return false
	}
	if !bytes.Equal(this.Certificate, that1.Certificate) {
		return false
	}
	return true
}
func (this *MasterKey) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MasterKey)
	if !ok {
		that2, ok := that.(MasterKey)
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
	if !bytes.Equal(this.Bytes, that1.Bytes) {
		return false
	}
	return true
}
func (this *Key) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Key)
	if !ok {
		that2, ok := that.(Key)
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
	if !bytes.Equal(this.Key, that1.Key) {
		return false
	}
	return true
}
func (m *RaAuthenticate) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RaAuthenticate) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RaAuthenticate) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Certificate) > 0 {
		i -= len(m.Certificate)
		copy(dAtA[i:], m.Certificate)
		i = encodeVarintMsg(dAtA, i, uint64(len(m.Certificate)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintMsg(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MasterKey) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MasterKey) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MasterKey) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Bytes) > 0 {
		i -= len(m.Bytes)
		copy(dAtA[i:], m.Bytes)
		i = encodeVarintMsg(dAtA, i, uint64(len(m.Bytes)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Key) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Key) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Key) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVarintMsg(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintMsg(dAtA []byte, offset int, v uint64) int {
	offset -= sovMsg(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *RaAuthenticate) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovMsg(uint64(l))
	}
	l = len(m.Certificate)
	if l > 0 {
		n += 1 + l + sovMsg(uint64(l))
	}
	return n
}

func (m *MasterKey) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Bytes)
	if l > 0 {
		n += 1 + l + sovMsg(uint64(l))
	}
	return n
}

func (m *Key) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovMsg(uint64(l))
	}
	return n
}

func sovMsg(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMsg(x uint64) (n int) {
	return sovMsg(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RaAuthenticate) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsg
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
			return fmt.Errorf("proto: RaAuthenticate: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RaAuthenticate: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
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
				return ErrInvalidLengthMsg
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = append(m.Sender[:0], dAtA[iNdEx:postIndex]...)
			if m.Sender == nil {
				m.Sender = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Certificate", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
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
				return ErrInvalidLengthMsg
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Certificate = append(m.Certificate[:0], dAtA[iNdEx:postIndex]...)
			if m.Certificate == nil {
				m.Certificate = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMsg
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
func (m *MasterKey) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsg
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
			return fmt.Errorf("proto: MasterKey: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MasterKey: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bytes", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
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
				return ErrInvalidLengthMsg
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Bytes = append(m.Bytes[:0], dAtA[iNdEx:postIndex]...)
			if m.Bytes == nil {
				m.Bytes = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMsg
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
func (m *Key) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMsg
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
			return fmt.Errorf("proto: Key: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Key: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMsg
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
				return ErrInvalidLengthMsg
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMsg
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = append(m.Key[:0], dAtA[iNdEx:postIndex]...)
			if m.Key == nil {
				m.Key = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMsg(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMsg
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
func skipMsg(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMsg
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
					return 0, ErrIntOverflowMsg
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
					return 0, ErrIntOverflowMsg
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
				return 0, ErrInvalidLengthMsg
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMsg
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMsg
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMsg        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMsg          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMsg = fmt.Errorf("proto: unexpected end of group")
)
