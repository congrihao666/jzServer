package jznet

import (
	"fmt"
	"net"
	//"time"
)
const (
      	T_MSEC = 2000		//轮询超时
	MAX_WAITQUQUE = 2000	//最大等待队列
      )

type jzTCPService struct {
	port string
	name string
	listener *net.TCPListener
}

func NewTcpService(name string) *jzTCPService {
	tcp := &jzTCPService{
			port:"",
			name:name,
		}
	return tcp
}

func (jzTcpService *jzTCPService) Bind (port string) bool {
     	fmt.Println("bind====================")
	addr,_ := net.ResolveTCPAddr("tcp",port) 	
	listener,err := net.ListenTCP("tcp",addr)
	if err != nil {
		fmt.Println("error ===",err)
		return false
	}
	jzTcpService.listener = listener
	jzTcpService.port = port
	return true
}

func (jzTcpService *jzTCPService) Accept() (*net.TCPConn,error) {
	fmt.Println("Accept=====",jzTcpService)
	return jzTcpService.listener.AcceptTCP()
}

func (jzTcpService *jzTCPService) Final() {
	jzTcpService.listener.Close()
}

func (tcpService *jzTCPService) GetIP() string {
	return tcpService.port
}