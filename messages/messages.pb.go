// Code generated by protoc-gen-go. DO NOT EDIT.
// source: messages.proto

/*
Package rpic_messages is a generated protocol buffer package.

It is generated from these files:
	messages.proto

It has these top-level messages:
	CtlMessage
	ImgMessage
	ClientInfoMessage
	ClientCommand
*/
package rpic_messages

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CtlMessage_Direction int32

const (
	CtlMessage_X CtlMessage_Direction = 0
	CtlMessage_Y CtlMessage_Direction = 1
)

var CtlMessage_Direction_name = map[int32]string{
	0: "X",
	1: "Y",
}
var CtlMessage_Direction_value = map[string]int32{
	"X": 0,
	"Y": 1,
}

func (x CtlMessage_Direction) String() string {
	return proto.EnumName(CtlMessage_Direction_name, int32(x))
}
func (CtlMessage_Direction) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

type CtlMessage_CtlType int32

const (
	CtlMessage_STOP    CtlMessage_CtlType = 0
	CtlMessage_VEHICLE CtlMessage_CtlType = 1
	CtlMessage_CAMERA  CtlMessage_CtlType = 2
)

var CtlMessage_CtlType_name = map[int32]string{
	0: "STOP",
	1: "VEHICLE",
	2: "CAMERA",
}
var CtlMessage_CtlType_value = map[string]int32{
	"STOP":    0,
	"VEHICLE": 1,
	"CAMERA":  2,
}

func (x CtlMessage_CtlType) String() string {
	return proto.EnumName(CtlMessage_CtlType_name, int32(x))
}
func (CtlMessage_CtlType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 1} }

type CtlMessage struct {
	Value     int32                `protobuf:"varint,1,opt,name=value" json:"value,omitempty"`
	Direction CtlMessage_Direction `protobuf:"varint,2,opt,name=direction,enum=rpic_messages.CtlMessage_Direction" json:"direction,omitempty"`
	CtlType   CtlMessage_CtlType   `protobuf:"varint,3,opt,name=ctl_type,json=ctlType,enum=rpic_messages.CtlMessage_CtlType" json:"ctl_type,omitempty"`
}

func (m *CtlMessage) Reset()                    { *m = CtlMessage{} }
func (m *CtlMessage) String() string            { return proto.CompactTextString(m) }
func (*CtlMessage) ProtoMessage()               {}
func (*CtlMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CtlMessage) GetValue() int32 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *CtlMessage) GetDirection() CtlMessage_Direction {
	if m != nil {
		return m.Direction
	}
	return CtlMessage_X
}

func (m *CtlMessage) GetCtlType() CtlMessage_CtlType {
	if m != nil {
		return m.CtlType
	}
	return CtlMessage_STOP
}

type ImgMessage struct {
	Time  *google_protobuf.Timestamp `protobuf:"bytes,1,opt,name=time" json:"time,omitempty"`
	Value string                     `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (m *ImgMessage) Reset()                    { *m = ImgMessage{} }
func (m *ImgMessage) String() string            { return proto.CompactTextString(m) }
func (*ImgMessage) ProtoMessage()               {}
func (*ImgMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ImgMessage) GetTime() *google_protobuf.Timestamp {
	if m != nil {
		return m.Time
	}
	return nil
}

func (m *ImgMessage) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type ClientInfoMessage struct {
	ClientInfo string `protobuf:"bytes,1,opt,name=clientInfo" json:"clientInfo,omitempty"`
}

func (m *ClientInfoMessage) Reset()                    { *m = ClientInfoMessage{} }
func (m *ClientInfoMessage) String() string            { return proto.CompactTextString(m) }
func (*ClientInfoMessage) ProtoMessage()               {}
func (*ClientInfoMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ClientInfoMessage) GetClientInfo() string {
	if m != nil {
		return m.ClientInfo
	}
	return ""
}

type ClientCommand struct {
	Command string `protobuf:"bytes,1,opt,name=command" json:"command,omitempty"`
}

func (m *ClientCommand) Reset()                    { *m = ClientCommand{} }
func (m *ClientCommand) String() string            { return proto.CompactTextString(m) }
func (*ClientCommand) ProtoMessage()               {}
func (*ClientCommand) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ClientCommand) GetCommand() string {
	if m != nil {
		return m.Command
	}
	return ""
}

func init() {
	proto.RegisterType((*CtlMessage)(nil), "rpic_messages.CtlMessage")
	proto.RegisterType((*ImgMessage)(nil), "rpic_messages.ImgMessage")
	proto.RegisterType((*ClientInfoMessage)(nil), "rpic_messages.ClientInfoMessage")
	proto.RegisterType((*ClientCommand)(nil), "rpic_messages.ClientCommand")
	proto.RegisterEnum("rpic_messages.CtlMessage_Direction", CtlMessage_Direction_name, CtlMessage_Direction_value)
	proto.RegisterEnum("rpic_messages.CtlMessage_CtlType", CtlMessage_CtlType_name, CtlMessage_CtlType_value)
}

func init() { proto.RegisterFile("messages.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 308 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x90, 0xc1, 0x4f, 0xc2, 0x30,
	0x14, 0xc6, 0x29, 0x02, 0x63, 0x8f, 0x80, 0xb3, 0xf1, 0x30, 0x39, 0x18, 0x9c, 0x17, 0x4c, 0x4c,
	0x0f, 0x70, 0xf5, 0x42, 0x2a, 0x89, 0x24, 0x12, 0x4d, 0x25, 0x46, 0x4f, 0x64, 0x94, 0xb2, 0x2c,
	0x69, 0xd7, 0x85, 0x15, 0x13, 0xfe, 0x67, 0xff, 0x08, 0xc3, 0x4a, 0x99, 0x1e, 0x3c, 0xf5, 0xbd,
	0xaf, 0xdf, 0xfb, 0xda, 0xf7, 0x83, 0x9e, 0x12, 0x45, 0x11, 0x27, 0xa2, 0x20, 0xf9, 0x56, 0x1b,
	0x8d, 0xbb, 0xdb, 0x3c, 0xe5, 0x4b, 0x27, 0xf6, 0xcf, 0x4d, 0xaa, 0x44, 0x61, 0x62, 0x95, 0xdb,
	0xfb, 0xe8, 0x1b, 0x01, 0x50, 0x23, 0xe7, 0xd6, 0x80, 0x2f, 0xa1, 0xf9, 0x15, 0xcb, 0x9d, 0x08,
	0xd1, 0x00, 0x0d, 0x9b, 0xcc, 0x36, 0x78, 0x02, 0xfe, 0x3a, 0xdd, 0x0a, 0x6e, 0x52, 0x9d, 0x85,
	0xf5, 0x01, 0x1a, 0xf6, 0x46, 0xb7, 0xe4, 0x4f, 0x30, 0xa9, 0x32, 0xc8, 0xa3, 0xb3, 0xb2, 0x6a,
	0x0a, 0x3f, 0x40, 0x9b, 0x1b, 0xb9, 0x34, 0xfb, 0x5c, 0x84, 0x67, 0x65, 0xc2, 0xcd, 0xff, 0x09,
	0xd4, 0xc8, 0xc5, 0x3e, 0x17, 0xcc, 0xe3, 0xb6, 0x88, 0xae, 0xc0, 0x3f, 0xa5, 0xe2, 0x26, 0xa0,
	0x8f, 0xa0, 0x76, 0x38, 0x3e, 0x03, 0x14, 0xdd, 0x83, 0x77, 0xb4, 0xe3, 0x36, 0x34, 0xde, 0x16,
	0x2f, 0xaf, 0x41, 0x0d, 0x77, 0xc0, 0x7b, 0x9f, 0x3e, 0xcd, 0xe8, 0xf3, 0x34, 0x40, 0x18, 0xa0,
	0x45, 0x27, 0xf3, 0x29, 0x9b, 0x04, 0xf5, 0x88, 0x01, 0xcc, 0x54, 0xe2, 0xb6, 0x25, 0xd0, 0x38,
	0xf0, 0x28, 0x97, 0xed, 0x8c, 0xfa, 0x24, 0xd1, 0x3a, 0x91, 0xc2, 0x92, 0x59, 0xed, 0x36, 0x64,
	0xe1, 0x60, 0xb1, 0xd2, 0x57, 0xd1, 0x39, 0x30, 0xf0, 0x8f, 0x74, 0xa2, 0x31, 0x5c, 0x50, 0x99,
	0x8a, 0xcc, 0xcc, 0xb2, 0x8d, 0x76, 0xd1, 0xd7, 0x00, 0xfc, 0x24, 0x96, 0x0f, 0xf8, 0xec, 0x97,
	0x12, 0xdd, 0x41, 0xd7, 0x0e, 0x51, 0xad, 0x54, 0x9c, 0xad, 0x71, 0x08, 0x1e, 0xb7, 0xe5, 0xd1,
	0xed, 0xda, 0x55, 0xab, 0xfc, 0xcf, 0xf8, 0x27, 0x00, 0x00, 0xff, 0xff, 0x1a, 0x60, 0x0d, 0x0a,
	0xdb, 0x01, 0x00, 0x00,
}
