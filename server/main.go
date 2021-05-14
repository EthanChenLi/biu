package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/wzshiming/ctc"
	"ngrok/server/lib/httpServer"
	"ngrok/server/lib/message"
	"ngrok/server/lib/tcpServer"
	"ngrok/server/lib/utils"
)


func init(){
	fmt.Println(ctc.ForegroundRed,`_____   _____   _       _____   _____   __   _        _____   __   _   _   _____   _____   _   _  
/  ___| /  _  \ | |     |  _  \ | ____| |  \ | |      /  ___/ |  \ | | | | |_   _| /  ___| | | | | 
| |     | | | | | |     | | | | | |__   |   \| |      | |___  |   \| | | |   | |   | |     | |_| | 
| |  _  | | | | | |     | | | | |  __|  | |\   |      \___  \ | |\   | | |   | |   | |     |  _  | 
| |_| | | |_| | | |___  | |_| | | |___  | | \  |       ___| | | | \  | | |   | |   | |___  | | | | 
\_____/ \_____/ |_____| |_____/ |_____| |_|  \_|      /_____/ |_|  \_| |_|   |_|   \_____| |_| |_| `)
	fmt.Println()
	logrus.Info("初始化全局变量")
	message.NewMessageQueue() //初始化消息管道
	//初始化TCP域名映射表
	utils.InitTcpDomainMap()
	//初始化TCP链接维护表
	utils.InitTcpConnectMap()

}

func main() {
	go tcpServer.Bootstrap() //启动tcp服务
	go httpServer.Bootstrap() //启动HTTP服务
	go message.SendHttpRequestMessage() //发送消息
	utils.TcpHeaderTimer() //心跳包
}
