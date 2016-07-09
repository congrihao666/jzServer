package jznet

import (
	"net"
)


type INetClientInterface interface {
	ParseMsg([]byte)
	Terminate() 
}

type NetClient struct {
	socket *Socket
	callback INetClientInterface
}

func NewNetClient(callback INetClientInterface) *NetClient {
	return &NetClient {
	       	callback:callback,
	       }
}

func (netclient *NetClient) Connect(serverAddr string) bool {
	addr, err := net.ResolveTCPAddr("tcp",serverAddr)
	if err != nil {
		return false
	}
	
	conn,err2 := net.DialTCP("tcp",nil,addr)
	if err2 != nil {
		return false
	}
	netclient.socket = NewSocket(conn)

	return true
}


func (netclient *NetClient) Start() {
	netclient.socket.Start()
	go netclient.ParseMsg()
}

func (netclient *NetClient) SendCmd(data []byte) {
	netclient.socket.SendCmd(data)
}


func (netclient *NetClient) ParseMsg() {
	readdata := make([]byte,1024)
	for true {
		_,err := netclient.socket.ReadData(readdata)
		if err != nil {
			netclient.Terminate()
			break
		}
		
		netclient.callback.ParseMsg(readdata)
	}
}

func (netclient *NetClient) Terminate() {
	netclient.callback.Terminate()
}