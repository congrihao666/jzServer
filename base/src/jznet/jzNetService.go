package jznet

import (
       "fmt"
       "net"
       //"github.com/golang/protobuf/proto"
       //"common/MsgDefine"
       //"common/charinitmsg"
       )

type NetService struct {
	tcpService *jzTCPService
	name string
	terminate bool
	servertype uint32
}

func NewNetService(name string,servertype uint32) *NetService {
	netservice := &NetService{
			tcpService:NewTcpService("test"),
			name:name,
			terminate :false,
			servertype :servertype,
		}
	return netservice
}

func (netservice *NetService) NewTcpTask(conn  *net.TCPConn) {
	fmt.Println(conn.RemoteAddr())
	fmt.Println(conn.LocalAddr())
	/*
	go func (conn *net.TCPConn) {
		readdata := make([]byte,1000)

		for true {
			_,err := conn.Read(readdata)
			if err == nil {
				fmt.Println("========",readdata)
				msg := &MsgDefine.BaseMsg{}
				proto.Unmarshal(readdata[4:],msg)
				basemsg := msg.String()
				fmt.Println("basemsg",basemsg)
				fmt.Printf("%s===\n",basemsg)
			} else {
					fmt.Println("close================ error")
					conn.Close()
					break
				}
		}
		
	}(conn)
	*/
}

func (netservice *NetService) Accept() (conn *net.TCPConn,err error) {
	return netservice.tcpService.Accept()
}

func (netservice *NetService) Init(port string) bool {
	netservice.tcpService = NewTcpService(netservice.name)
	if !netservice.tcpService.Bind(port) {
		fmt.Println("bind false")
		return false
	}
	fmt.Println("NetService init succee")
	return true
}

func (netservice *NetService) ReloadConfig() {
}

func (netservice *NetService) Terminate() {
	fmt.Println("NetService Terminate=======")
	netservice.tcpService.Final()
	netservice.terminate = true
}

func (netservice *NetService) IsTerminate() bool {
	return netservice.terminate
}

func (netservice *NetService) Start() {

}

func (netservice *NetService) Final() {
}

func (netservice *NetService) GetType() uint32 {
	return netservice.servertype
}

func (netservice *NetService) GetIP() string {
	return netservice.tcpService.GetIP()
}