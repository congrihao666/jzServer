package jzcode

import (
       "github.com/golang/protobuf/proto"
)

func DecodePB(data []byte,message proto.Message) {
	proto.Unmarshal(data,message)
}

func EncodePB(message proto.Message) ([]byte,error) {
	return proto.Marshal(message)
}
