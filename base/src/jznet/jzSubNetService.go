package jznet

import (
       "ZoneServerMsg"
       "jzcode"
)

type SubNetService struct {
	*NetService
	zoneClient *ZoneClient
}

func NewSubNetService(name string, servertype uint32,serverinstance IServiceInterface) *SubNetService {
	subNetService := &SubNetService{
		NetService :NewNetService(name,servertype),
	}
	
	subNetService.zoneClient = NewZoneClient(serverinstance)
	return subNetService
}

func (subNetService *SubNetService) Init(zoneServerAddr string,localServerAddr string) bool {
	if !subNetService.NetService.Init(localServerAddr) {
		return false
	}
	if !subNetService.zoneClient.Connect(zoneServerAddr) {
		return false
	}
	 
	servertype := subNetService.GetType()
	pstrIP := subNetService.GetIP()	
	forwardmsg := &ZoneServerMsg.ForWardZoneStartupMsg{
		Servertype :&servertype,
		PstrIP:&pstrIP,
	}

	msg := &ZoneServerMsg.ZoneServerMsg{Msg:&ZoneServerMsg.ZoneServerMsg_ForWardZoneStartupMsg{forwardmsg,},}
	pbdata,err := jzcode.EncodePB(msg)
	if err != nil {
		return false
	}
	subNetService.zoneClient.SendCmd(pbdata)
	
	return true
}

func (subNetService *SubNetService) ParseMsg(data []byte) bool {
	return true
}

func (subNetService *SubNetService) Start() {
	subNetService.NetService.Start()
	subNetService.zoneClient.Start()
}