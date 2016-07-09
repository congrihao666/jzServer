package jznet

import (
	"bytes"
	"encoding/binary"
)

type MsgPack struct {
	rawbuffer *bytes.Buffer
	totallen uint32
}

const (
      	MessageHeaderLen = 4  //数据包头长度
	Mask = 0x40000000
	)

var (
    byteOrder = binary.LittleEndian
)

func NewMsgPack() *MsgPack {
        return &MsgPack {
		rawbuffer : bytes.NewBuffer([]byte("")),
		totallen : 0,
	}
}

func (pack *MsgPack) AddMsg(data []byte) {
	curlen := (uint32)(len(data))
	binary.Write(pack.rawbuffer,byteOrder,curlen)
	pack.rawbuffer.Write(data)
	pack.totallen += curlen
}

func (pack* MsgPack) GetMsg() []byte {
	alldata := make([]byte,pack.totallen + MessageHeaderLen)
	bytebuffer := bytes.NewBuffer([]byte{})
	mask := pack.totallen | Mask
	binary.Write(bytebuffer,byteOrder,(uint32)(mask))
	copy(alldata,bytebuffer.Bytes())
	copy(alldata[4:],pack.rawbuffer.Bytes())
	return alldata
}

func (pack* MsgPack) GetDirMsg(data []byte) []byte {
	datalen := len(data)
	alldata := make([]byte,MessageHeaderLen + datalen)
	bytewrite := bytes.NewBuffer([]byte{})
	binary.Write(bytewrite,byteOrder,(uint32)(datalen))
	copy(alldata,bytewrite.Bytes())
	copy(alldata[MessageHeaderLen:],data)
	return alldata
}

func (pack* MsgPack) GetLen() uint32 {
	return pack.totallen
}

func (pack* MsgPack)Reset() {
	pack.rawbuffer.Reset()
	pack.totallen = 0
}