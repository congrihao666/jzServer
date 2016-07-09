// Code generated by protoc-gen-go.
// source: ZoneServerMsg.proto
// DO NOT EDIT!

/*
Package ZoneServerMsg is a generated protocol buffer package.

It is generated from these files:
	ZoneServerMsg.proto

It has these top-level messages:
	ZoneServerMsg
	ForWardZoneStartupMsg
	BackWardStartUpMsg
	ServerInfo
	BackWardToOtherRegServerInfo
	BackWardToServerRegOtherServers
*/
package ZoneServerMsg

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

type ZoneServerMsg struct {
	// Types that are valid to be assigned to Msg:
	//	*ZoneServerMsg_ForWardZoneStartupMsg
	Msg              isZoneServerMsg_Msg `protobuf_oneof:"Msg"`
	XXX_unrecognized []byte              `json:"-"`
}

func (m *ZoneServerMsg) Reset()                    { *m = ZoneServerMsg{} }
func (m *ZoneServerMsg) String() string            { return proto.CompactTextString(m) }
func (*ZoneServerMsg) ProtoMessage()               {}
func (*ZoneServerMsg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type isZoneServerMsg_Msg interface {
	isZoneServerMsg_Msg()
}

type ZoneServerMsg_ForWardZoneStartupMsg struct {
	ForWardZoneStartupMsg *ForWardZoneStartupMsg `protobuf:"bytes,1,opt,name=ForWardZoneStartupMsg,oneof"`
}

func (*ZoneServerMsg_ForWardZoneStartupMsg) isZoneServerMsg_Msg() {}

func (m *ZoneServerMsg) GetMsg() isZoneServerMsg_Msg {
	if m != nil {
		return m.Msg
	}
	return nil
}

func (m *ZoneServerMsg) GetForWardZoneStartupMsg() *ForWardZoneStartupMsg {
	if x, ok := m.GetMsg().(*ZoneServerMsg_ForWardZoneStartupMsg); ok {
		return x.ForWardZoneStartupMsg
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ZoneServerMsg) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ZoneServerMsg_OneofMarshaler, _ZoneServerMsg_OneofUnmarshaler, _ZoneServerMsg_OneofSizer, []interface{}{
		(*ZoneServerMsg_ForWardZoneStartupMsg)(nil),
	}
}

func _ZoneServerMsg_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ZoneServerMsg)
	// Msg
	switch x := m.Msg.(type) {
	case *ZoneServerMsg_ForWardZoneStartupMsg:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ForWardZoneStartupMsg); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("ZoneServerMsg.Msg has unexpected type %T", x)
	}
	return nil
}

func _ZoneServerMsg_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ZoneServerMsg)
	switch tag {
	case 1: // Msg.ForWardZoneStartupMsg
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ForWardZoneStartupMsg)
		err := b.DecodeMessage(msg)
		m.Msg = &ZoneServerMsg_ForWardZoneStartupMsg{msg}
		return true, err
	default:
		return false, nil
	}
}

func _ZoneServerMsg_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ZoneServerMsg)
	// Msg
	switch x := m.Msg.(type) {
	case *ZoneServerMsg_ForWardZoneStartupMsg:
		s := proto.Size(x.ForWardZoneStartupMsg)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// 向ZoneServer请求启动
type ForWardZoneStartupMsg struct {
	Servertype       *uint32 `protobuf:"varint,1,req,name=servertype" json:"servertype,omitempty"`
	PstrIP           *string `protobuf:"bytes,2,req,name=pstrIP" json:"pstrIP,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ForWardZoneStartupMsg) Reset()                    { *m = ForWardZoneStartupMsg{} }
func (m *ForWardZoneStartupMsg) String() string            { return proto.CompactTextString(m) }
func (*ForWardZoneStartupMsg) ProtoMessage()               {}
func (*ForWardZoneStartupMsg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ForWardZoneStartupMsg) GetServertype() uint32 {
	if m != nil && m.Servertype != nil {
		return *m.Servertype
	}
	return 0
}

func (m *ForWardZoneStartupMsg) GetPstrIP() string {
	if m != nil && m.PstrIP != nil {
		return *m.PstrIP
	}
	return ""
}

// 启动返回
type BackWardStartUpMsg struct {
	ServerID         *uint32 `protobuf:"varint,1,req,name=serverID" json:"serverID,omitempty"`
	VerifyRet        *uint32 `protobuf:"varint,2,req,name=verifyRet" json:"verifyRet,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *BackWardStartUpMsg) Reset()                    { *m = BackWardStartUpMsg{} }
func (m *BackWardStartUpMsg) String() string            { return proto.CompactTextString(m) }
func (*BackWardStartUpMsg) ProtoMessage()               {}
func (*BackWardStartUpMsg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *BackWardStartUpMsg) GetServerID() uint32 {
	if m != nil && m.ServerID != nil {
		return *m.ServerID
	}
	return 0
}

func (m *BackWardStartUpMsg) GetVerifyRet() uint32 {
	if m != nil && m.VerifyRet != nil {
		return *m.VerifyRet
	}
	return 0
}

type ServerInfo struct {
	ServerID         *uint32 `protobuf:"varint,1,req,name=serverID" json:"serverID,omitempty"`
	ServerType       *uint32 `protobuf:"varint,2,req,name=serverType" json:"serverType,omitempty"`
	ExtIP            *string `protobuf:"bytes,3,req,name=extIP" json:"extIP,omitempty"`
	ServerName       *string `protobuf:"bytes,4,req,name=serverName" json:"serverName,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ServerInfo) Reset()                    { *m = ServerInfo{} }
func (m *ServerInfo) String() string            { return proto.CompactTextString(m) }
func (*ServerInfo) ProtoMessage()               {}
func (*ServerInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ServerInfo) GetServerID() uint32 {
	if m != nil && m.ServerID != nil {
		return *m.ServerID
	}
	return 0
}

func (m *ServerInfo) GetServerType() uint32 {
	if m != nil && m.ServerType != nil {
		return *m.ServerType
	}
	return 0
}

func (m *ServerInfo) GetExtIP() string {
	if m != nil && m.ExtIP != nil {
		return *m.ExtIP
	}
	return ""
}

func (m *ServerInfo) GetServerName() string {
	if m != nil && m.ServerName != nil {
		return *m.ServerName
	}
	return ""
}

// 通知其他Server 到当前Server信息
type BackWardToOtherRegServerInfo struct {
	ServerInfo       *ServerInfo `protobuf:"bytes,1,req,name=serverInfo" json:"serverInfo,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *BackWardToOtherRegServerInfo) Reset()                    { *m = BackWardToOtherRegServerInfo{} }
func (m *BackWardToOtherRegServerInfo) String() string            { return proto.CompactTextString(m) }
func (*BackWardToOtherRegServerInfo) ProtoMessage()               {}
func (*BackWardToOtherRegServerInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *BackWardToOtherRegServerInfo) GetServerInfo() *ServerInfo {
	if m != nil {
		return m.ServerInfo
	}
	return nil
}

// 通知当前Server 其他Server信息
type BackWardToServerRegOtherServers struct {
	ServerList       []*ServerInfo `protobuf:"bytes,1,rep,name=serverList" json:"serverList,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *BackWardToServerRegOtherServers) Reset()                    { *m = BackWardToServerRegOtherServers{} }
func (m *BackWardToServerRegOtherServers) String() string            { return proto.CompactTextString(m) }
func (*BackWardToServerRegOtherServers) ProtoMessage()               {}
func (*BackWardToServerRegOtherServers) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *BackWardToServerRegOtherServers) GetServerList() []*ServerInfo {
	if m != nil {
		return m.ServerList
	}
	return nil
}

func init() {
	proto.RegisterType((*ZoneServerMsg)(nil), "ZoneServerMsg.ZoneServerMsg")
	proto.RegisterType((*ForWardZoneStartupMsg)(nil), "ZoneServerMsg.ForWardZoneStartupMsg")
	proto.RegisterType((*BackWardStartUpMsg)(nil), "ZoneServerMsg.BackWardStartUpMsg")
	proto.RegisterType((*ServerInfo)(nil), "ZoneServerMsg.ServerInfo")
	proto.RegisterType((*BackWardToOtherRegServerInfo)(nil), "ZoneServerMsg.BackWardToOtherRegServerInfo")
	proto.RegisterType((*BackWardToServerRegOtherServers)(nil), "ZoneServerMsg.BackWardToServerRegOtherServers")
}

var fileDescriptor0 = []byte{
	// 268 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x90, 0xd1, 0x4b, 0x02, 0x41,
	0x10, 0xc6, 0xd3, 0xcb, 0xc8, 0x91, 0x8b, 0xda, 0x08, 0x2e, 0x08, 0x8a, 0xa3, 0x07, 0x5f, 0xf2,
	0xc1, 0xb7, 0xe8, 0x4d, 0x2a, 0x3a, 0xc8, 0x12, 0x53, 0x82, 0xa0, 0x87, 0xa3, 0x46, 0x93, 0xc8,
	0x3d, 0x66, 0xa7, 0xc8, 0xff, 0xbe, 0xdd, 0xd9, 0x6c, 0xbb, 0xb8, 0x7c, 0x9c, 0x6f, 0xe7, 0xf7,
	0xfb, 0x86, 0x85, 0xdd, 0x07, 0x3d, 0xc7, 0x3b, 0xa4, 0x0f, 0xa4, 0xbe, 0x99, 0x76, 0x0a, 0xd2,
	0xac, 0x55, 0x5c, 0x0a, 0xd3, 0x47, 0x28, 0x07, 0xea, 0x02, 0xf6, 0x2e, 0x35, 0xdd, 0xe7, 0xf4,
	0x2c, 0x39, 0xe7, 0xc4, 0xef, 0x85, 0x7d, 0x48, 0x6a, 0x47, 0xb5, 0x76, 0xab, 0x7b, 0xdc, 0x29,
	0x4b, 0x2b, 0x77, 0xaf, 0xd6, 0x7a, 0x0d, 0x88, 0x9c, 0xfe, 0xec, 0x1f, 0x9b, 0x52, 0x00, 0x46,
	0x24, 0xbc, 0x28, 0xd0, 0xba, 0xeb, 0xed, 0x58, 0x6d, 0xc1, 0x46, 0x61, 0x98, 0xb2, 0x41, 0x52,
	0xb7, 0x73, 0x33, 0x3d, 0x05, 0xd5, 0xcb, 0x9f, 0x5e, 0x1d, 0x2d, 0xe4, 0x58, 0xc8, 0x6d, 0xd8,
	0xf4, 0x64, 0x76, 0xfe, 0xcd, 0xed, 0x40, 0xd3, 0x8e, 0xb3, 0xc9, 0x62, 0x88, 0x2c, 0x68, 0x9c,
	0x8e, 0x01, 0xfc, 0x8d, 0xd9, 0x7c, 0xa2, 0x2b, 0x90, 0x9f, 0xfa, 0x91, 0xab, 0x17, 0x46, 0xc5,
	0xd0, 0xc0, 0x4f, 0xb6, 0xed, 0x91, 0x6b, 0x0f, 0x2b, 0x37, 0xf9, 0x1b, 0x26, 0xeb, 0x72, 0x51,
	0x1f, 0x0e, 0x96, 0x17, 0x8d, 0xf4, 0x2d, 0xbf, 0x20, 0x0d, 0x71, 0xfa, 0xab, 0xe8, 0x64, 0xc9,
	0xb8, 0x49, 0xaa, 0x5a, 0xdd, 0xfd, 0x3f, 0x3f, 0x16, 0xd6, 0xd3, 0x01, 0x1c, 0x06, 0x9d, 0xcf,
	0xad, 0x4f, 0xbc, 0x7e, 0x32, 0xc1, 0x78, 0x3d, 0x33, 0x6c, 0x8d, 0xd1, 0x4a, 0xe3, 0x57, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xa0, 0x20, 0xbc, 0xf4, 0xf3, 0x01, 0x00, 0x00,
}
