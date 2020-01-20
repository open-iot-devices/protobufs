// Code generated by protoc-gen-go. DO NOT EDIT.
// source: openiot/common.proto

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

type Header struct {
	// Unique (at least at given local network) device ID. It must always
	// be un-encrypted to let device/controller find decryption key.
	// Should also match Message.device_id
	DeviceId uint64 `protobuf:"varint,1,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	// Encryption parameters
	//
	// Types that are valid to be assigned to Encryption:
	//	*Header_Plain
	//	*Header_AesIv
	Encryption           isHeader_Encryption `protobuf_oneof:"encryption"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Header) Reset()         { *m = Header{} }
func (m *Header) String() string { return proto.CompactTextString(m) }
func (*Header) ProtoMessage()    {}
func (*Header) Descriptor() ([]byte, []int) {
	return fileDescriptor_5aeeb5b759bd5053, []int{0}
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

type isHeader_Encryption interface {
	isHeader_Encryption()
}

type Header_Plain struct {
	Plain bool `protobuf:"varint,5,opt,name=plain,proto3,oneof"`
}

type Header_AesIv struct {
	AesIv []byte `protobuf:"bytes,6,opt,name=aes_iv,json=aesIv,proto3,oneof"`
}

func (*Header_Plain) isHeader_Encryption() {}

func (*Header_AesIv) isHeader_Encryption() {}

func (m *Header) GetEncryption() isHeader_Encryption {
	if m != nil {
		return m.Encryption
	}
	return nil
}

func (m *Header) GetPlain() bool {
	if x, ok := m.GetEncryption().(*Header_Plain); ok {
		return x.Plain
	}
	return false
}

func (m *Header) GetAesIv() []byte {
	if x, ok := m.GetEncryption().(*Header_AesIv); ok {
		return x.AesIv
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Header) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Header_Plain)(nil),
		(*Header_AesIv)(nil),
	}
}

type Message struct {
	// The same device id as in unencrypted Header.
	// Serves mostly to validate successful decryption.
	DeviceId uint64 `protobuf:"varint,1,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	// Auto-incremented sequence number to drop duplicates
	// Even if device / controller will be sending
	// updates every second - counter will last in about 136 years
	Sequence uint32 `protobuf:"varint,2,opt,name=sequence,proto3" json:"sequence,omitempty"`
	// Actual message
	//
	// Types that are valid to be assigned to Message:
	//	*Message_SystemJoinRequest
	//	*Message_SystemJoinResponse
	//	*Message_SystemLeaveRequest
	//	*Message_SystemLeaveResponse
	//	*Message_SystemDeviceInfoRequest
	//	*Message_SystemDeviceInfoResponse
	Message              isMessage_Message `protobuf_oneof:"message"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_5aeeb5b759bd5053, []int{1}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetDeviceId() uint64 {
	if m != nil {
		return m.DeviceId
	}
	return 0
}

func (m *Message) GetSequence() uint32 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

type isMessage_Message interface {
	isMessage_Message()
}

type Message_SystemJoinRequest struct {
	SystemJoinRequest *SystemJoinRequest `protobuf:"bytes,10,opt,name=system_join_request,json=systemJoinRequest,proto3,oneof"`
}

type Message_SystemJoinResponse struct {
	SystemJoinResponse *SystemJoinResponse `protobuf:"bytes,11,opt,name=system_join_response,json=systemJoinResponse,proto3,oneof"`
}

type Message_SystemLeaveRequest struct {
	SystemLeaveRequest *SystemLeaveRequest `protobuf:"bytes,12,opt,name=system_leave_request,json=systemLeaveRequest,proto3,oneof"`
}

type Message_SystemLeaveResponse struct {
	SystemLeaveResponse *SystemLeaveResponse `protobuf:"bytes,13,opt,name=system_leave_response,json=systemLeaveResponse,proto3,oneof"`
}

type Message_SystemDeviceInfoRequest struct {
	SystemDeviceInfoRequest *SystemDeviceInfoRequest `protobuf:"bytes,14,opt,name=system_device_info_request,json=systemDeviceInfoRequest,proto3,oneof"`
}

type Message_SystemDeviceInfoResponse struct {
	SystemDeviceInfoResponse *SystemDeviceInfoResponse `protobuf:"bytes,15,opt,name=system_device_info_response,json=systemDeviceInfoResponse,proto3,oneof"`
}

func (*Message_SystemJoinRequest) isMessage_Message() {}

func (*Message_SystemJoinResponse) isMessage_Message() {}

func (*Message_SystemLeaveRequest) isMessage_Message() {}

func (*Message_SystemLeaveResponse) isMessage_Message() {}

func (*Message_SystemDeviceInfoRequest) isMessage_Message() {}

func (*Message_SystemDeviceInfoResponse) isMessage_Message() {}

func (m *Message) GetMessage() isMessage_Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *Message) GetSystemJoinRequest() *SystemJoinRequest {
	if x, ok := m.GetMessage().(*Message_SystemJoinRequest); ok {
		return x.SystemJoinRequest
	}
	return nil
}

func (m *Message) GetSystemJoinResponse() *SystemJoinResponse {
	if x, ok := m.GetMessage().(*Message_SystemJoinResponse); ok {
		return x.SystemJoinResponse
	}
	return nil
}

func (m *Message) GetSystemLeaveRequest() *SystemLeaveRequest {
	if x, ok := m.GetMessage().(*Message_SystemLeaveRequest); ok {
		return x.SystemLeaveRequest
	}
	return nil
}

func (m *Message) GetSystemLeaveResponse() *SystemLeaveResponse {
	if x, ok := m.GetMessage().(*Message_SystemLeaveResponse); ok {
		return x.SystemLeaveResponse
	}
	return nil
}

func (m *Message) GetSystemDeviceInfoRequest() *SystemDeviceInfoRequest {
	if x, ok := m.GetMessage().(*Message_SystemDeviceInfoRequest); ok {
		return x.SystemDeviceInfoRequest
	}
	return nil
}

func (m *Message) GetSystemDeviceInfoResponse() *SystemDeviceInfoResponse {
	if x, ok := m.GetMessage().(*Message_SystemDeviceInfoResponse); ok {
		return x.SystemDeviceInfoResponse
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Message) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Message_SystemJoinRequest)(nil),
		(*Message_SystemJoinResponse)(nil),
		(*Message_SystemLeaveRequest)(nil),
		(*Message_SystemLeaveResponse)(nil),
		(*Message_SystemDeviceInfoRequest)(nil),
		(*Message_SystemDeviceInfoResponse)(nil),
	}
}

func init() {
	proto.RegisterType((*Header)(nil), "openiot.Header")
	proto.RegisterType((*Message)(nil), "openiot.Message")
}

func init() { proto.RegisterFile("openiot/common.proto", fileDescriptor_5aeeb5b759bd5053) }

var fileDescriptor_5aeeb5b759bd5053 = []byte{
	// 346 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xc1, 0x4e, 0xea, 0x40,
	0x14, 0x86, 0xe9, 0x0d, 0x14, 0x38, 0xc0, 0x35, 0x0e, 0x28, 0x4d, 0x71, 0x51, 0x59, 0x75, 0x85,
	0x89, 0xbe, 0x01, 0x71, 0x51, 0x0c, 0xc6, 0x64, 0x7c, 0x80, 0xa6, 0xb4, 0x07, 0x33, 0x86, 0xce,
	0xd4, 0x9e, 0xda, 0x84, 0x17, 0xf5, 0x79, 0x0c, 0x9d, 0x5a, 0x10, 0x2a, 0xcb, 0xf9, 0xff, 0x93,
	0xef, 0x7c, 0x93, 0x19, 0x18, 0xa9, 0x04, 0xa5, 0x50, 0xd9, 0x5d, 0xa8, 0xe2, 0x58, 0xc9, 0x59,
	0x92, 0xaa, 0x4c, 0xb1, 0x76, 0x99, 0xda, 0x55, 0x4d, 0x5b, 0xca, 0x30, 0xd6, 0xf5, 0x34, 0x02,
	0xd3, 0xc3, 0x20, 0xc2, 0x94, 0x4d, 0xa0, 0x1b, 0x61, 0x2e, 0x42, 0xf4, 0x45, 0x64, 0x19, 0x8e,
	0xe1, 0x36, 0x79, 0x47, 0x07, 0x8b, 0x88, 0x5d, 0x43, 0x2b, 0xd9, 0x04, 0x42, 0x5a, 0x2d, 0xc7,
	0x70, 0x3b, 0x5e, 0x83, 0xeb, 0x23, 0x1b, 0x83, 0x19, 0x20, 0xf9, 0x22, 0xb7, 0x4c, 0xc7, 0x70,
	0xfb, 0xbb, 0x22, 0x40, 0x5a, 0xe4, 0xf3, 0x3e, 0x00, 0xca, 0x30, 0xdd, 0x26, 0x99, 0x50, 0x72,
	0xfa, 0xd5, 0x84, 0xf6, 0x33, 0x12, 0x05, 0x6f, 0x78, 0x7e, 0x8f, 0x0d, 0x1d, 0xc2, 0x8f, 0x4f,
	0x94, 0x21, 0x5a, 0xff, 0x1c, 0xc3, 0x1d, 0xf0, 0xea, 0xcc, 0x96, 0x30, 0xd4, 0xea, 0xfe, 0xbb,
	0x12, 0xd2, 0x4f, 0x77, 0x39, 0x65, 0x16, 0x38, 0x86, 0xdb, 0xbb, 0xb7, 0x67, 0xe5, 0xf5, 0x66,
	0xaf, 0xc5, 0xcc, 0x93, 0x12, 0x92, 0xeb, 0x09, 0xaf, 0xc1, 0x2f, 0xe9, 0x38, 0x64, 0x2f, 0x30,
	0xfa, 0x4d, 0xa3, 0x44, 0x49, 0x42, 0xab, 0x57, 0xe0, 0x26, 0xb5, 0x38, 0x3d, 0xe2, 0x35, 0x38,
	0xa3, 0x93, 0xf4, 0x00, 0xb8, 0xc1, 0x20, 0xc7, 0xca, 0xaf, 0x5f, 0x0b, 0x5c, 0xee, 0x66, 0xf6,
	0x82, 0x25, 0xf0, 0x30, 0x65, 0x1c, 0xae, 0x8e, 0x80, 0xa5, 0xe2, 0xa0, 0x20, 0xde, 0xd4, 0x13,
	0x2b, 0xc7, 0x21, 0x9d, 0xc6, 0xcc, 0x07, 0xbb, 0x64, 0xfe, 0xbc, 0x81, 0x5c, 0xab, 0x4a, 0xf5,
	0x7f, 0x01, 0x76, 0x8e, 0xc0, 0x8f, 0xfa, 0x71, 0xe4, 0x5a, 0xed, 0x7d, 0xc7, 0x54, 0x5f, 0xb1,
	0x15, 0x4c, 0x6a, 0x17, 0x94, 0xea, 0x17, 0xc5, 0x86, 0xdb, 0x33, 0x1b, 0x2a, 0x7f, 0x8b, 0xfe,
	0xe8, 0xe6, 0x5d, 0x68, 0xc7, 0xfa, 0x33, 0xad, 0xcc, 0xe2, 0x17, 0x3f, 0x7c, 0x07, 0x00, 0x00,
	0xff, 0xff, 0xa6, 0xea, 0x13, 0x93, 0xfc, 0x02, 0x00, 0x00,
}
