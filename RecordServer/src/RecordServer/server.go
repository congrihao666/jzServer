package main

import (
       "jznet"
       //"net"
       //"fmt"
       )

type RecordServer struct {
	*jznet.SubNetService
}

func NewRecordServer(name string) *RecordServer {
	recordServer := &RecordServer {}
	recordServer.SubNetService = jznet.NewSubNetService("recordServer",jznet.RecordServerType,recordServer)
	return recordServer
}

func (recordserver *RecordServer) ParseMsg(data []byte) {
	if recordserver.SubNetService.ParseMsg(data) {
		return
	}
}

func (recordServer* RecordServer) Init(localServerAddr string) bool {
	if !recordServer.SubNetService.Init("10.1.100.39:31892",localServerAddr) {
		return false
	}
	return true
}
