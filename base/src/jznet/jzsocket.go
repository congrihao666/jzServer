package jznet

import (
       	"fmt"
       	"time"
       	"sync"
	"net"
       )

type Socket struct {
	datachannel chan []byte
	msgpack *MsgPack
	bufferflag bool
	bufferchan chan bool
	lock *sync.Mutex
	conn* net.TCPConn
}

func NewSocket(conn *net.TCPConn) *Socket {
	return &Socket {
		datachannel : make(chan []byte,1),
		msgpack : NewMsgPack(),
		bufferflag : false,
		bufferchan : make(chan bool),
		lock	  : &sync.Mutex{},
		conn 	  : conn,
	}
}

func (socket *Socket) Start() {
     go socket.handleChannelData()
     go socket.syncData()
     //go socket.handleRecvData()
}


func (socket *Socket) SendCmd(data []byte) {
	socket.datachannel <- data
}

func (socket* Socket) SetBuffer(buffer bool) {
	var flag []byte
	if buffer  {
		flag = []byte("1")
	} else {
		flag = []byte("0")
	}
	
	socket.datachannel <- flag
}

func (socket* Socket) syncDirectData(data []byte) {
	alldata := socket.msgpack.GetDirMsg(data)
	socket.conn.Write(alldata)
}


func (socket *Socket) ReadData(readdata []byte) (int,error){
	return socket.conn.Read(readdata)
}

/*
func (socket* Socket) handleRecvData() {
	readBytes := make([]byte,1024) 
	for {
		_,err := socket.conn.Read(readBytes)
		if err != nil {
			socket.final()
			break
		}
	}
}
*/

func (socket* Socket) handleChannelData() {
	for true {
		data := <- socket.datachannel
		
		if len(data) == 1 {
			if string(data) == "0" {
				socket.bufferflag = false 
			} else if string(data) == "1" {
				socket.bufferflag = true
			}
		} else {
		
			if socket.bufferflag  {	
				socket.syncDirectData(data)

			} else {
				socket.lock.Lock()
			
				socket.msgpack.AddMsg(data)	
			
				socket.lock.Unlock()
			}
		}
	}
}

func (socket* Socket) syncData() {
     	for true {
		time.Sleep(10 * time.Millisecond)

		packlen := socket.msgpack.GetLen()
		if packlen > 0 {
			socket.lock.Lock()
			rawdata := socket.msgpack.GetMsg()
			socket.msgpack.Reset()
			socket.lock.Unlock()
	
			sendlen,err := socket.conn.Write(rawdata)
			//sendlen,err := socket.conn.Write([]byte("11111"))
			if err != nil {
				fmt.Println("send syncData error")
			}
			
			if sendlen !=  len(rawdata) {
				fmt.Println("sync data length not equal len")
			}
			
			fmt.Println("write len",sendlen,len(rawdata))
			fmt.Println("senddata ==",rawdata,byte(180))
		}
	}
}

func (socket *Socket) final() {

}