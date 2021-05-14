package utils

import "sync"

const MAX = 65535 //计数器最大数

//tcpServer 连接计数器
type TcpCount struct {
	count uint16  //0~65535
	lock sync.Mutex
}

//tcp计数器
func NewTcpCounter()*TcpCount {
	return &TcpCount{count: 0}
}

func (this *TcpCount)IncrTcpCount() uint16{
	this.lock.Lock()
	defer this.lock.Unlock()
	if this.count>= MAX {
		this.count =0
	}
	this.count+=1
	return this.count
}