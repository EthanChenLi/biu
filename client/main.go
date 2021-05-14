package main

import (
	"flag"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/wzshiming/ctc"
	"ngrok/client/lib/tcpClient"
)


const TcpAddr = "您的服务端IP:22222"

func main() {

	fmt.Println(ctc.ForegroundRed,`_____   _____   _       _____   _____   __   _        _____   __   _   _   _____   _____   _   _  
/  ___| /  _  \ | |     |  _  \ | ____| |  \ | |      /  ___/ |  \ | | | | |_   _| /  ___| | | | | 
| |     | | | | | |     | | | | | |__   |   \| |      | |___  |   \| | | |   | |   | |     | |_| | 
| |  _  | | | | | |     | | | | |  __|  | |\   |      \___  \ | |\   | | |   | |   | |     |  _  | 
| |_| | | |_| | | |___  | |_| | | |___  | | \  |       ___| | | | \  | | |   | |   | |___  | | | | 
\_____/ \_____/ |_____| |_____/ |_____| |_|  \_|      /_____/ |_|  \_| |_|   |_|   \_____| |_| |_| `)
	fmt.Println()
	flag.StringVar(&tcpClient.TargetHost,"host", "", "客户端访问的唯一KEY")
	flag.StringVar(&tcpClient.TargetIp,  "ip", "127.0.0.1", "HTTP客户端的IP地址")
	flag.StringVar(&tcpClient.TargetPort,"port", "80", "HTTP客户端的端口地址")
	flag.Parse()
	if tcpClient.TargetHost == ""{
		logrus.Panic("客户端访问KEY不能为空,请输入 : -host [YOUR KEY]")
	}
	//启动TCP链接
	tcpClient.Bootstrap(TcpAddr)
}
