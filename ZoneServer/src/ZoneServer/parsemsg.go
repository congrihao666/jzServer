package main

import (
	"jzcode"
	"ZoneServerMsg"
	"fmt"
)

const (
	Check_OK = iota
	Check_Failure
)

func ParseZoneMsg(data []byte,zoneServer* ZoneServer) int {
	msg := &ZoneServerMsg.ZoneServerMsg{}
	jzcode.DecodePB(data[4:],msg)
	fmt.Println("msg=======",msg)
	fmt.Println("data====",data)
	switch msg.Msg.(type) {
		case *ZoneServerMsg.ZoneServerMsg_ForWardZoneStartupMsg:
				
		case nil:
		        return Check_Failure
	}
	return Check_Failure
}
