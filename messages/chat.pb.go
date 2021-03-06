// Code generated by protoc-gen-go. DO NOT EDIT.
// source: chat.proto

package messages

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

type POWMessage struct {
	Nonce                int32    `protobuf:"varint,1,opt,name=nonce,proto3" json:"nonce,omitempty"`
	Pubkey               string   `protobuf:"bytes,2,opt,name=pubkey,proto3" json:"pubkey,omitempty"`
	Addr                 string   `protobuf:"bytes,3,opt,name=addr,proto3" json:"addr,omitempty"`
	Blocknum             int32    `protobuf:"varint,4,opt,name=blocknum,proto3" json:"blocknum,omitempty"`
	Difficulty           int32    `protobuf:"varint,5,opt,name=difficulty,proto3" json:"difficulty,omitempty"`
	Result               string   `protobuf:"bytes,6,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *POWMessage) Reset()         { *m = POWMessage{} }
func (m *POWMessage) String() string { return proto.CompactTextString(m) }
func (*POWMessage) ProtoMessage()    {}
func (*POWMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c585a45e2093e54, []int{0}
}

func (m *POWMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_POWMessage.Unmarshal(m, b)
}
func (m *POWMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_POWMessage.Marshal(b, m, deterministic)
}
func (m *POWMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_POWMessage.Merge(m, src)
}
func (m *POWMessage) XXX_Size() int {
	return xxx_messageInfo_POWMessage.Size(m)
}
func (m *POWMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_POWMessage.DiscardUnknown(m)
}

var xxx_messageInfo_POWMessage proto.InternalMessageInfo

func (m *POWMessage) GetNonce() int32 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *POWMessage) GetPubkey() string {
	if m != nil {
		return m.Pubkey
	}
	return ""
}

func (m *POWMessage) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func (m *POWMessage) GetBlocknum() int32 {
	if m != nil {
		return m.Blocknum
	}
	return 0
}

func (m *POWMessage) GetDifficulty() int32 {
	if m != nil {
		return m.Difficulty
	}
	return 0
}

func (m *POWMessage) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

func init() {
	proto.RegisterType((*POWMessage)(nil), "messages.POWMessage")
}

func init() { proto.RegisterFile("chat.proto", fileDescriptor_8c585a45e2093e54) }

var fileDescriptor_8c585a45e2093e54 = []byte{
	// 164 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0xce, 0x48, 0x2c,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xc8, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0x2d,
	0x56, 0x5a, 0xc4, 0xc8, 0xc5, 0x15, 0xe0, 0x1f, 0xee, 0x0b, 0xe1, 0x0b, 0x89, 0x70, 0xb1, 0xe6,
	0xe5, 0xe7, 0x25, 0xa7, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0xb0, 0x06, 0x41, 0x38, 0x42, 0x62, 0x5c,
	0x6c, 0x05, 0xa5, 0x49, 0xd9, 0xa9, 0x95, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x50, 0x9e,
	0x90, 0x10, 0x17, 0x4b, 0x62, 0x4a, 0x4a, 0x91, 0x04, 0x33, 0x58, 0x14, 0xcc, 0x16, 0x92, 0xe2,
	0xe2, 0x48, 0xca, 0xc9, 0x4f, 0xce, 0xce, 0x2b, 0xcd, 0x95, 0x60, 0x01, 0x1b, 0x02, 0xe7, 0x0b,
	0xc9, 0x71, 0x71, 0xa5, 0x64, 0xa6, 0xa5, 0x65, 0x26, 0x97, 0xe6, 0x94, 0x54, 0x4a, 0xb0, 0x82,
	0x65, 0x91, 0x44, 0x40, 0xf6, 0x14, 0xa5, 0x16, 0x97, 0xe6, 0x94, 0x48, 0xb0, 0x41, 0xec, 0x81,
	0xf0, 0x92, 0xd8, 0xc0, 0xae, 0x36, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xbb, 0xe6, 0x0b, 0xd8,
	0xc3, 0x00, 0x00, 0x00,
}
