package jznet

import (
       "time"
       "net"
       )


const  (
      state_notuse = iota		//连接关闭状态
      state_verify   			//连接通过验证
      state_sync			//等待来自其他客户端的验证同步
      state_okay 			//连接通过验证，进入主循环
      state_recycle			//连接退出状态，回收
      )

const (
      	terminate_no	= iota		//没有结束任务
	terminate_active		//客户端主动断开连接
	terminate_passive		//服务器主动断开
      )

type ITcpTask interface {
	VerifyConn() int
	WaitSync() int
	RecycleConn() int
	AddToContainer() 
	RemoveFromContainer() 
	UniqueAdd() bool
	UniqueRemove() bool
	SetUnique()
	IsTerminateWait() bool
	TerminateWait()
	IsTerminate() bool
	Terminate(terminate uint32) 
	TerminateError() bool 
	GetState() uint32 
	SetState(state uint32)
	GetNextState()
	ResetState()
	GetStateString(state uint32) string
}

type jzTcpTask struct {
	socket *Socket			//jzSocket 对象 
	state uint32			//tcpTask状态
	taskType uint32			//taskType 类型
	terminate uint32		//
	terminate_wait bool		
	lifeTime time.Time
	uniqueVerified bool
 	checkSignal bool
	_tick_timer time.Timer
	tick bool			
}

func NewjzTcpTask(conn *net.TCPConn) *jzTcpTask {
	return nil
}

func (task* jzTcpTask) SetUnique() {
	task.uniqueVerified = true
}

func (task* jzTcpTask) IsTerminateWait() bool {
	return task.terminate_wait
}

func (task* jzTcpTask) TerminateWait() {
	task.terminate_wait = true
}

func (task* jzTcpTask) IsTerminate()bool {
	return task.terminate != terminate_no
}

func (task* jzTcpTask) Terminate(method uint32) {
	task.terminate = method
}
	
func (task* jzTcpTask) GetState() uint32 {
	return task.state
}

func (task* jzTcpTask) SetState(state uint32) {
	task.state = state
}

func (task* jzTcpTask) AddToContainer() {}
func (task* jzTcpTask) GetNextState() {
	old_state := task.GetState()
	switch old_state {
	       case state_notuse: {task.SetState(state_verify);	break;}
	       case state_verify: {task.SetState(state_sync);break;}
	       case state_sync:   {task.AddToContainer();task.socket.SetBuffer(true);task.SetState(state_recycle);break}
	       case state_recycle: {task.SetState(state_notuse);break;}
	}
}

func (task* jzTcpTask) ResetState() {
	task.SetState(state_recycle)
}

