package main

import (
       	"jznet"
	"net"
	"fmt"
       )

type ZoneServer struct {
	*jznet.NetService
	
}

var zoneserver *ZoneServer
func NewZoneServer(name string) *ZoneServer {
	zoneserver = &ZoneServer {
		NetService:jznet.NewNetService(name,jznet.ZoneServerType),
	} 
	
	return zoneserver
}

func GetService() *ZoneServer {
	return zoneserver
}

func (zoneService *ZoneServer) NewTcpTask(conn *net.TCPConn) {
	remoteAddr := conn.RemoteAddr()
	go func(conn *net.TCPConn) {
		readdata := make ([]byte, 1000)
		for true {
			_,err := conn.Read(readdata)
			if err != nil {
				conn.Close()
				break
			} else {
				
				ret := ParseZoneMsg(readdata[4:],zoneService)
				if ret == Check_OK {
				}
			}
		}
	}(conn)
	
	fmt.Println("remoteAddr=======",remoteAddr)	
}

func (zoneService* ZoneServer) Init(port string) bool {
	if !zoneService.NetService.Init(port) {
		return false
	}  
	
	return true
}

func (zoneService* ZoneServer) Terminate() {
	zoneService.NetService.Terminate()
}

func (zoneservice *ZoneServer) ParseMsg(data []byte) {
}
