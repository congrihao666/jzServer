package jznet
import (
       "fmt"
       "os"
       "os/signal"
       "syscall"
       "math/rand"
       "time"
       "net"
      )


const (
	ZoneServerType uint32 = 1 + iota
	RecordServerType
)

type IServiceInterface interface {
     Init(port string) bool
     IsTerminate() bool
     Final()
     Terminate() 
     ReloadConfig()
     Accept() (*net.TCPConn,error)
     NewTcpTask(*net.TCPConn)
     ParseMsg(data []byte)
     Start()
}

func StartService(service_interface IServiceInterface,port string) bool {
	if !Init(service_interface,port) {
		fmt.Println("service_interface failure")
		return false
	} 
	
	service_interface.Start()
	
	for  !service_interface.IsTerminate() {
		fmt.Println("main IsTerminate")
		if (!ServiceCallback(service_interface)) {
			fmt.Println("service callback return")
			break
		}
	}
	
	service_interface.Final()
	return true
}

func ServiceCallback(service_interface IServiceInterface) bool{
	conn,err := service_interface.Accept()
	if err != nil {
		return false
	}
	service_interface.NewTcpTask(conn)
	return true
}

func Init(service_interface IServiceInterface,port string) bool {
	sigs := make(chan os.Signal)
	signal.Notify(sigs,syscall.SIGINT,syscall.SIGQUIT,syscall.SIGABRT,syscall.SIGTERM)
	
	sigs_hub := make(chan os.Signal)
	signal.Notify(sigs_hub,syscall.SIGHUP)
	
	go func() {
		for {
			select {
				case <- sigs:{
			     		fmt.Println("Terminate")
					service_interface.Terminate()
					break
			     	}
				case <- sigs_hub: {
					fmt.Println("reloadConfig")
					service_interface.ReloadConfig()
				}
			
			}	
	    	
		}
	}()

	rand.New(rand.NewSource(time.Now().Unix()))
	return service_interface.Init(port)
}

