package utils

import (
	"errors"
	"github.com/sirupsen/logrus"
	"net"
	"sync"
	"time"
)

/**
 *	TCP 当前链接列表
 *  维护当前TCP的 FID:CONN 的映射关系
 *  只负责TCP维护，不负责域名维护
 */

var TcpConnectMap *tcpConnectMap

type tcpConnectMap struct {
	connectInfo map[uint16]*connectInfo
	lock sync.Mutex //map线程锁
}

//tcp相关信息
type connectInfo struct {
	Conn net.Conn
	FlushTime int64 //心跳包更新时间(时间搓)
}


//初始化tcp链接维护表
func InitTcpConnectMap() {
	TcpConnectMap = &tcpConnectMap{
		connectInfo:make(map[uint16]*connectInfo),
	}
}


//设置一个新的连接
func (this *tcpConnectMap)SetNewConnectMap(fid uint16 , conn net.Conn){
	this.lock.Lock()
	defer this.lock.Unlock()
	this.connectInfo[fid] = &connectInfo{
		Conn:      conn,
		FlushTime: time.Now().Unix(),  //当前心跳时间搓
	}
}

//获取指定连接的对象
func (this *tcpConnectMap)GetConnectMap(fid uint16) (*connectInfo,error){
	connInfo,ok :=this.connectInfo[fid]
	if ok{
		return connInfo,nil
	}
	return nil,errors.New("tcp connect object not found")
}


//关闭TCP链接资源
func (this *tcpConnectMap)CloseConnect(fid uint16) error{
	connectSource ,err :=this.GetConnectMap(fid)
	if err!=nil{
		return err
	}
	_ = connectSource.Conn.Close()
	delete(this.connectInfo,fid)
	logrus.Info("TCP CONNECT IS CLOSE, fid:",fid)
	return nil
}


//刷新心跳包时间
func (this *tcpConnectMap)FlushHeaderTime(fid uint16) error {
	connectSource ,err :=this.GetConnectMap(fid)
	if err!=nil{
		return err
	}
	//刷新
	connectSource.FlushTime = time.Now().Unix()
	return nil
}

//获取所有连接
func (this *tcpConnectMap)GetAll(){
	for k,v := range this.connectInfo{
		logrus.Info("tcp fid:",k,",  tcp source:",v.Conn)
	}
}