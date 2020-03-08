// Code generated by protoc-gen-go. DO NOT EDIT.
// source: openiot/system.proto

package openiot

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type EncryptionType int32

const (
	EncryptionType_PLAIN   EncryptionType = 0
	EncryptionType_AES_ECB EncryptionType = 1
)

var EncryptionType_name = map[int32]string{
	0: "PLAIN",
	1: "AES_ECB",
}

var EncryptionType_value = map[string]int32{
	"PLAIN":   0,
	"AES_ECB": 1,
}

func (x EncryptionType) String() string {
	return proto.EnumName(EncryptionType_name, int32(x))
}

func (EncryptionType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_477c448daeb2d28d, []int{0}
}

// Always unencrypted message header
// to be able to find decryption key by device id.
type Header struct {
	// Unique (at least within network) device ID.
	DeviceId uint64 `protobuf:"varint,1,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	// CRC of message data before encryption (if desired).
	Crc uint32 `protobuf:"varint,2,opt,name=crc,proto3" json:"crc,omitempty"`
	// Indicates that message contains key exchange request / response
	KeyExchange          bool     `protobuf:"varint,3,opt,name=key_exchange,json=keyExchange,proto3" json:"key_exchange,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Header) Reset()         { *m = Header{} }
func (m *Header) String() string { return proto.CompactTextString(m) }
func (*Header) ProtoMessage()    {}
func (*Header) Descriptor() ([]byte, []int) {
	return fileDescriptor_477c448daeb2d28d, []int{0}
}

func (m *Header) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Header.Unmarshal(m, b)
}
func (m *Header) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Header.Marshal(b, m, deterministic)
}
func (m *Header) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Header.Merge(m, src)
}
func (m *Header) XXX_Size() int {
	return xxx_messageInfo_Header.Size(m)
}
func (m *Header) XXX_DiscardUnknown() {
	xxx_messageInfo_Header.DiscardUnknown(m)
}

var xxx_messageInfo_Header proto.InternalMessageInfo

func (m *Header) GetDeviceId() uint64 {
	if m != nil {
		return m.DeviceId
	}
	return 0
}

func (m *Header) GetCrc() uint32 {
	if m != nil {
		return m.Crc
	}
	return 0
}

func (m *Header) GetKeyExchange() bool {
	if m != nil {
		return m.KeyExchange
	}
	return false
}

// A message followed by Header and describes common
// parameters for all device / controller messages.
type MessageInfo struct {
	// Auto-incremented sequence number to drop duplicates
	// uint32 should last about 136 years even if device / controller
	// will be sending messages every second
	Sequence             uint32   `protobuf:"varint,1,opt,name=sequence,proto3" json:"sequence,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MessageInfo) Reset()         { *m = MessageInfo{} }
func (m *MessageInfo) String() string { return proto.CompactTextString(m) }
func (*MessageInfo) ProtoMessage()    {}
func (*MessageInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_477c448daeb2d28d, []int{1}
}

func (m *MessageInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageInfo.Unmarshal(m, b)
}
func (m *MessageInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageInfo.Marshal(b, m, deterministic)
}
func (m *MessageInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageInfo.Merge(m, src)
}
func (m *MessageInfo) XXX_Size() int {
	return xxx_messageInfo_MessageInfo.Size(m)
}
func (m *MessageInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageInfo.DiscardUnknown(m)
}

var xxx_messageInfo_MessageInfo proto.InternalMessageInfo

func (m *MessageInfo) GetSequence() uint32 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

// Diffie-Hellman key exchange: device to controller.
type KeyExchangeRequest struct {
	// Diffie-Hellman modulus
	DhP uint64 `protobuf:"varint,1,opt,name=dh_p,json=dhP,proto3" json:"dh_p,omitempty"`
	// Diffie-Hellman base
	DhG uint64 `protobuf:"varint,2,opt,name=dh_g,json=dhG,proto3" json:"dh_g,omitempty"`
	// "Public" key of device
	DhA []uint32 `protobuf:"varint,3,rep,packed,name=dh_a,json=dhA,proto3" json:"dh_a,omitempty"`
	// Desired encryption type
	EncryptionType       EncryptionType `protobuf:"varint,4,opt,name=encryption_type,json=encryptionType,proto3,enum=openiot.EncryptionType" json:"encryption_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *KeyExchangeRequest) Reset()         { *m = KeyExchangeRequest{} }
func (m *KeyExchangeRequest) String() string { return proto.CompactTextString(m) }
func (*KeyExchangeRequest) ProtoMessage()    {}
func (*KeyExchangeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_477c448daeb2d28d, []int{2}
}

func (m *KeyExchangeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeyExchangeRequest.Unmarshal(m, b)
}
func (m *KeyExchangeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeyExchangeRequest.Marshal(b, m, deterministic)
}
func (m *KeyExchangeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyExchangeRequest.Merge(m, src)
}
func (m *KeyExchangeRequest) XXX_Size() int {
	return xxx_messageInfo_KeyExchangeRequest.Size(m)
}
func (m *KeyExchangeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyExchangeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_KeyExchangeRequest proto.InternalMessageInfo

func (m *KeyExchangeRequest) GetDhP() uint64 {
	if m != nil {
		return m.DhP
	}
	return 0
}

func (m *KeyExchangeRequest) GetDhG() uint64 {
	if m != nil {
		return m.DhG
	}
	return 0
}

func (m *KeyExchangeRequest) GetDhA() []uint32 {
	if m != nil {
		return m.DhA
	}
	return nil
}

func (m *KeyExchangeRequest) GetEncryptionType() EncryptionType {
	if m != nil {
		return m.EncryptionType
	}
	return EncryptionType_PLAIN
}

// Diffie-Hellman key exchange: controller to device.
type KeyExchangeResponse struct {
	DhB                  []uint32 `protobuf:"varint,1,rep,packed,name=dh_b,json=dhB,proto3" json:"dh_b,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KeyExchangeResponse) Reset()         { *m = KeyExchangeResponse{} }
func (m *KeyExchangeResponse) String() string { return proto.CompactTextString(m) }
func (*KeyExchangeResponse) ProtoMessage()    {}
func (*KeyExchangeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_477c448daeb2d28d, []int{3}
}

func (m *KeyExchangeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeyExchangeResponse.Unmarshal(m, b)
}
func (m *KeyExchangeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeyExchangeResponse.Marshal(b, m, deterministic)
}
func (m *KeyExchangeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyExchangeResponse.Merge(m, src)
}
func (m *KeyExchangeResponse) XXX_Size() int {
	return xxx_messageInfo_KeyExchangeResponse.Size(m)
}
func (m *KeyExchangeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyExchangeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_KeyExchangeResponse proto.InternalMessageInfo

func (m *KeyExchangeResponse) GetDhB() []uint32 {
	if m != nil {
		return m.DhB
	}
	return nil
}

// Request to join IoT network from device.
// If device uses encryption it must initiate key exchange
// prior to join.
type JoinRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Manufacturer         string   `protobuf:"bytes,2,opt,name=manufacturer,proto3" json:"manufacturer,omitempty"`
	ProductUrl           string   `protobuf:"bytes,3,opt,name=product_url,json=productUrl,proto3" json:"product_url,omitempty"`
	ProtobufUrl          string   `protobuf:"bytes,4,opt,name=protobuf_url,json=protobufUrl,proto3" json:"protobuf_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JoinRequest) Reset()         { *m = JoinRequest{} }
func (m *JoinRequest) String() string { return proto.CompactTextString(m) }
func (*JoinRequest) ProtoMessage()    {}
func (*JoinRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_477c448daeb2d28d, []int{4}
}

func (m *JoinRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JoinRequest.Unmarshal(m, b)
}
func (m *JoinRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JoinRequest.Marshal(b, m, deterministic)
}
func (m *JoinRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JoinRequest.Merge(m, src)
}
func (m *JoinRequest) XXX_Size() int {
	return xxx_messageInfo_JoinRequest.Size(m)
}
func (m *JoinRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_JoinRequest.DiscardUnknown(m)
}

var xxx_messageInfo_JoinRequest proto.InternalMessageInfo

func (m *JoinRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *JoinRequest) GetManufacturer() string {
	if m != nil {
		return m.Manufacturer
	}
	return ""
}

func (m *JoinRequest) GetProductUrl() string {
	if m != nil {
		return m.ProductUrl
	}
	return ""
}

func (m *JoinRequest) GetProtobufUrl() string {
	if m != nil {
		return m.ProtobufUrl
	}
	return ""
}

// Request to join response from controller to device.
type JoinResponse struct {
	// Server's name
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Server's timestamp
	Timestamp            int64    `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JoinResponse) Reset()         { *m = JoinResponse{} }
func (m *JoinResponse) String() string { return proto.CompactTextString(m) }
func (*JoinResponse) ProtoMessage()    {}
func (*JoinResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_477c448daeb2d28d, []int{5}
}

func (m *JoinResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JoinResponse.Unmarshal(m, b)
}
func (m *JoinResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JoinResponse.Marshal(b, m, deterministic)
}
func (m *JoinResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JoinResponse.Merge(m, src)
}
func (m *JoinResponse) XXX_Size() int {
	return xxx_messageInfo_JoinResponse.Size(m)
}
func (m *JoinResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_JoinResponse.DiscardUnknown(m)
}

var xxx_messageInfo_JoinResponse proto.InternalMessageInfo

func (m *JoinResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *JoinResponse) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

// Deveice request to leave IoT network
type LeaveRequest struct {
	Reason               string   `protobuf:"bytes,1,opt,name=reason,proto3" json:"reason,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LeaveRequest) Reset()         { *m = LeaveRequest{} }
func (m *LeaveRequest) String() string { return proto.CompactTextString(m) }
func (*LeaveRequest) ProtoMessage()    {}
func (*LeaveRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_477c448daeb2d28d, []int{6}
}

func (m *LeaveRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LeaveRequest.Unmarshal(m, b)
}
func (m *LeaveRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LeaveRequest.Marshal(b, m, deterministic)
}
func (m *LeaveRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LeaveRequest.Merge(m, src)
}
func (m *LeaveRequest) XXX_Size() int {
	return xxx_messageInfo_LeaveRequest.Size(m)
}
func (m *LeaveRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LeaveRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LeaveRequest proto.InternalMessageInfo

func (m *LeaveRequest) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

// Controller response to leave network request
type LeaveResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LeaveResponse) Reset()         { *m = LeaveResponse{} }
func (m *LeaveResponse) String() string { return proto.CompactTextString(m) }
func (*LeaveResponse) ProtoMessage()    {}
func (*LeaveResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_477c448daeb2d28d, []int{7}
}

func (m *LeaveResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LeaveResponse.Unmarshal(m, b)
}
func (m *LeaveResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LeaveResponse.Marshal(b, m, deterministic)
}
func (m *LeaveResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LeaveResponse.Merge(m, src)
}
func (m *LeaveResponse) XXX_Size() int {
	return xxx_messageInfo_LeaveResponse.Size(m)
}
func (m *LeaveResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LeaveResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LeaveResponse proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("openiot.EncryptionType", EncryptionType_name, EncryptionType_value)
	proto.RegisterType((*Header)(nil), "openiot.Header")
	proto.RegisterType((*MessageInfo)(nil), "openiot.MessageInfo")
	proto.RegisterType((*KeyExchangeRequest)(nil), "openiot.KeyExchangeRequest")
	proto.RegisterType((*KeyExchangeResponse)(nil), "openiot.KeyExchangeResponse")
	proto.RegisterType((*JoinRequest)(nil), "openiot.JoinRequest")
	proto.RegisterType((*JoinResponse)(nil), "openiot.JoinResponse")
	proto.RegisterType((*LeaveRequest)(nil), "openiot.LeaveRequest")
	proto.RegisterType((*LeaveResponse)(nil), "openiot.LeaveResponse")
}

func init() { proto.RegisterFile("openiot/system.proto", fileDescriptor_477c448daeb2d28d) }

var fileDescriptor_477c448daeb2d28d = []byte{
	// 419 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0x51, 0x6f, 0xd3, 0x30,
	0x10, 0x80, 0x09, 0x09, 0x5d, 0x73, 0x69, 0xbb, 0x62, 0x10, 0x44, 0x80, 0x44, 0xf0, 0x03, 0x0a,
	0x3c, 0x14, 0x09, 0xfe, 0xc0, 0x3a, 0x54, 0x41, 0x61, 0xa0, 0xc9, 0xb0, 0x17, 0x5e, 0x22, 0x37,
	0xb9, 0xb6, 0xd1, 0x1a, 0xdb, 0xd8, 0xce, 0x44, 0x7e, 0x02, 0x4f, 0xfc, 0x65, 0x54, 0xd7, 0xed,
	0x54, 0x69, 0x6f, 0xbe, 0x2f, 0x97, 0xbb, 0xef, 0xee, 0xe0, 0xb1, 0x54, 0x28, 0x6a, 0x69, 0xdf,
	0x99, 0xce, 0x58, 0x6c, 0x26, 0x4a, 0x4b, 0x2b, 0xc9, 0x89, 0xa7, 0xf4, 0x17, 0xf4, 0x3e, 0x23,
	0xaf, 0x50, 0x93, 0xe7, 0x10, 0x57, 0x78, 0x53, 0x97, 0x58, 0xd4, 0x55, 0x1a, 0x64, 0x41, 0x1e,
	0xb1, 0xfe, 0x0e, 0xcc, 0x2b, 0x32, 0x86, 0xb0, 0xd4, 0x65, 0x7a, 0x3f, 0x0b, 0xf2, 0x21, 0xdb,
	0x3e, 0xc9, 0x2b, 0x18, 0x5c, 0x63, 0x57, 0xe0, 0x9f, 0x72, 0xcd, 0xc5, 0x0a, 0xd3, 0x30, 0x0b,
	0xf2, 0x3e, 0x4b, 0xae, 0xb1, 0x9b, 0x79, 0x44, 0xdf, 0x40, 0xf2, 0x0d, 0x8d, 0xe1, 0x2b, 0x9c,
	0x8b, 0xa5, 0x24, 0xcf, 0xa0, 0x6f, 0xf0, 0x77, 0x8b, 0xa2, 0x44, 0x57, 0x7f, 0xc8, 0x0e, 0x31,
	0xfd, 0x17, 0x00, 0xf9, 0x7a, 0xfb, 0x2b, 0xdb, 0x72, 0x63, 0xc9, 0x43, 0x88, 0xaa, 0x75, 0xa1,
	0xbc, 0x4e, 0x58, 0xad, 0x2f, 0x3d, 0x5a, 0x39, 0x15, 0x87, 0x3e, 0x79, 0xc4, 0xd3, 0x30, 0x0b,
	0xb7, 0x76, 0xd5, 0x7a, 0x4a, 0xce, 0xe0, 0x14, 0x45, 0xa9, 0x3b, 0x65, 0x6b, 0x29, 0x0a, 0xdb,
	0x29, 0x4c, 0xa3, 0x2c, 0xc8, 0x47, 0xef, 0x9f, 0x4e, 0xfc, 0xe4, 0x93, 0xd9, 0xe1, 0xfb, 0xcf,
	0x4e, 0x21, 0x1b, 0xe1, 0x51, 0x4c, 0x73, 0x78, 0x74, 0x24, 0x64, 0x94, 0x14, 0x06, 0x7d, 0xaf,
	0x45, 0x1a, 0xec, 0x7b, 0x9d, 0xd3, 0xbf, 0x01, 0x24, 0x5f, 0x64, 0x2d, 0xf6, 0xd2, 0x04, 0x22,
	0xc1, 0x9b, 0xdd, 0x8c, 0x31, 0x73, 0x6f, 0x42, 0x61, 0xd0, 0x70, 0xd1, 0x2e, 0x79, 0x69, 0x5b,
	0x8d, 0xda, 0xd9, 0xc7, 0xec, 0x88, 0x91, 0x97, 0x90, 0x28, 0x2d, 0xab, 0xb6, 0xb4, 0x45, 0xab,
	0x37, 0x6e, 0xa1, 0x31, 0x03, 0x8f, 0xae, 0xf4, 0x66, 0xbb, 0x72, 0x77, 0xbd, 0x45, 0xbb, 0x74,
	0x19, 0x91, 0xcb, 0x48, 0xf6, 0xec, 0x4a, 0x6f, 0xe8, 0x19, 0x0c, 0x76, 0x2a, 0x5e, 0xf7, 0x2e,
	0x97, 0x17, 0x10, 0xdb, 0xba, 0x41, 0x63, 0x79, 0xa3, 0x9c, 0x48, 0xc8, 0x6e, 0x01, 0x7d, 0x0d,
	0x83, 0x0b, 0xe4, 0x37, 0x87, 0x13, 0x3c, 0x81, 0x9e, 0x46, 0x6e, 0xa4, 0xf0, 0x35, 0x7c, 0x44,
	0x4f, 0x61, 0xe8, 0xf3, 0x76, 0xad, 0xde, 0xe6, 0x30, 0x3a, 0x5e, 0x29, 0x89, 0xe1, 0xc1, 0xe5,
	0xc5, 0x74, 0xfe, 0x7d, 0x7c, 0x8f, 0x24, 0x70, 0x32, 0x9d, 0xfd, 0x28, 0x66, 0x1f, 0xcf, 0xc7,
	0xc1, 0xa2, 0xe7, 0x8c, 0x3f, 0xfc, 0x0f, 0x00, 0x00, 0xff, 0xff, 0xed, 0xb4, 0xdf, 0xd4, 0x9b,
	0x02, 0x00, 0x00,
}
