// Code generated by protoc-gen-go.
// source: SceneServerMsg.proto
// DO NOT EDIT!

/*
Package SceneServerMsg is a generated protocol buffer package.

It is generated from these files:
	SceneServerMsg.proto

It has these top-level messages:
	SceneServerMsg
	ForwardSceneMsg
*/
package SceneServerMsg

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type SceneServerMsg struct {
	// Types that are valid to be assigned to Msg:
	//	*SceneServerMsg_ForwardSceneMsg
	Msg              isSceneServerMsg_Msg `protobuf_oneof:"Msg"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *SceneServerMsg) Reset()                    { *m = SceneServerMsg{} }
func (m *SceneServerMsg) String() string            { return proto.CompactTextString(m) }
func (*SceneServerMsg) ProtoMessage()               {}
func (*SceneServerMsg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type isSceneServerMsg_Msg interface {
	isSceneServerMsg_Msg()
}

type SceneServerMsg_ForwardSceneMsg struct {
	ForwardSceneMsg *ForwardSceneMsg `protobuf:"bytes,1,opt,name=ForwardSceneMsg,oneof"`
}

func (*SceneServerMsg_ForwardSceneMsg) isSceneServerMsg_Msg() {}

func (m *SceneServerMsg) GetMsg() isSceneServerMsg_Msg {
	if m != nil {
		return m.Msg
	}
	return nil
}

func (m *SceneServerMsg) GetForwardSceneMsg() *ForwardSceneMsg {
	if x, ok := m.GetMsg().(*SceneServerMsg_ForwardSceneMsg); ok {
		return x.ForwardSceneMsg
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*SceneServerMsg) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _SceneServerMsg_OneofMarshaler, _SceneServerMsg_OneofUnmarshaler, _SceneServerMsg_OneofSizer, []interface{}{
		(*SceneServerMsg_ForwardSceneMsg)(nil),
	}
}

func _SceneServerMsg_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*SceneServerMsg)
	// Msg
	switch x := m.Msg.(type) {
	case *SceneServerMsg_ForwardSceneMsg:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ForwardSceneMsg); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("SceneServerMsg.Msg has unexpected type %T", x)
	}
	return nil
}

func _SceneServerMsg_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*SceneServerMsg)
	switch tag {
	case 1: // Msg.ForwardSceneMsg
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ForwardSceneMsg)
		err := b.DecodeMessage(msg)
		m.Msg = &SceneServerMsg_ForwardSceneMsg{msg}
		return true, err
	default:
		return false, nil
	}
}

func _SceneServerMsg_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*SceneServerMsg)
	// Msg
	switch x := m.Msg.(type) {
	case *SceneServerMsg_ForwardSceneMsg:
		s := proto.Size(x.ForwardSceneMsg)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type ForwardSceneMsg struct {
	ServerID         *uint32 `protobuf:"varint,1,req,name=serverID" json:"serverID,omitempty"`
	ServerName       *uint32 `protobuf:"varint,2,req,name=serverName" json:"serverName,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ForwardSceneMsg) Reset()                    { *m = ForwardSceneMsg{} }
func (m *ForwardSceneMsg) String() string            { return proto.CompactTextString(m) }
func (*ForwardSceneMsg) ProtoMessage()               {}
func (*ForwardSceneMsg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ForwardSceneMsg) GetServerID() uint32 {
	if m != nil && m.ServerID != nil {
		return *m.ServerID
	}
	return 0
}

func (m *ForwardSceneMsg) GetServerName() uint32 {
	if m != nil && m.ServerName != nil {
		return *m.ServerName
	}
	return 0
}

func init() {
	proto.RegisterType((*SceneServerMsg)(nil), "SceneServerMsg.SceneServerMsg")
	proto.RegisterType((*ForwardSceneMsg)(nil), "SceneServerMsg.ForwardSceneMsg")
}

var fileDescriptor0 = []byte{
	// 122 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x09, 0x4e, 0x4e, 0xcd,
	0x4b, 0x0d, 0x4e, 0x2d, 0x2a, 0x4b, 0x2d, 0xf2, 0x2d, 0x4e, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0xe2, 0x43, 0x15, 0x55, 0x0a, 0xe6, 0x42, 0x13, 0x11, 0xb2, 0xe2, 0xe2, 0x77, 0xcb, 0x2f,
	0x2a, 0x4f, 0x2c, 0x4a, 0x01, 0x4b, 0x00, 0x85, 0x24, 0x18, 0x15, 0x18, 0x35, 0xb8, 0x8d, 0xe4,
	0xf5, 0xd0, 0x4c, 0x44, 0x53, 0xe6, 0xc1, 0xe0, 0xc4, 0xca, 0xc5, 0x0c, 0x32, 0xd4, 0x1c, 0xc3,
	0x08, 0x21, 0x01, 0x2e, 0x8e, 0x62, 0xb0, 0x46, 0x4f, 0x17, 0xa0, 0x71, 0x4c, 0x1a, 0xbc, 0x42,
	0x42, 0x5c, 0x5c, 0x10, 0x11, 0xbf, 0xc4, 0xdc, 0x54, 0x09, 0x26, 0x90, 0x18, 0x20, 0x00, 0x00,
	0xff, 0xff, 0x82, 0x5d, 0x39, 0x19, 0xb4, 0x00, 0x00, 0x00,
}