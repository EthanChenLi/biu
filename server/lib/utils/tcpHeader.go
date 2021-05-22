package utils

import (
	"github.com/sirupsen/logrus"
	"time"
)

const HEADER_TIME_INTERVAL = 5 //时间间隔
const HEADER_TIMEOUT = 20 //超时时间（秒）


//tcp  心跳包
func TcpHeaderTimer(){
	ticker := time.NewTicker( HEADER_TIME_INTERVAL * time.Second )
	for _ = range ticker.C{
		for key,item := range TcpConnectMap.connectInfo{
			if item != nil && item.FlushTime + HEADER_TIMEOUT < time.Now().Unix() {
				//超时
				logrus.Info("tcp客户端心跳超时,fid: ",key)
				//关闭TCP资源
				_ = TcpConnectMap.CloseConnect(key)
			}
		}
	}
}

