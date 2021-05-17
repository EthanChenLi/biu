package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/wzshiming/ctc"
	"ngrok/client/lib/tcpClient"
	"ngrok/client/lib/utils"
	"strings"
)

func init()  {
	utils.NewConfig()
	utils.NewQueue()
}


func main() {
	fmt.Println(ctc.ForegroundRed,
`_____   _____   _       _____   _____   __   _        _____   __   _   _   _____   _____   _   _  
/  ___| /  _  \ | |     |  _  \ | ____| |  \ | |      /  ___/ |  \ | | | | |_   _| /  ___| | | | | 
| |     | | | | | |     | | | | | |__   |   \| |      | |___  |   \| | | |   | |   | |     | |_| | 
| |  _  | | | | | |     | | | | |  __|  | |\   |      \___  \ | |\   | | |   | |   | |     |  _  | 
| |_| | | |_| | | |___  | |_| | | |___  | | \  |       ___| | | | \  | | |   | |   | |___  | | | | 
\_____/ \_____/ |_____| |_____/ |_____| |_|  \_|      /_____/ |_|  \_| |_|   |_|   \_____| |_| |_| `)
	logrus.Info("服务启动中...")

	serverType := utils.GetSection(utils.BaseSection).Key("TYPE").String()
	serverTypeUpper := strings.ToUpper(serverType)
	//web服务
	if serverTypeUpper == utils.SERVER_TYPE_HTTP || serverTypeUpper == utils.SERVER_TYPE_ALL {
		go tcpClient.Bootstrap()
	}
	//tcp代理
	if serverTypeUpper == utils.SERVER_TYPE_TCP || serverTypeUpper == utils.SERVER_TYPE_ALL {
		//TODO -- 下个版本迭代
	}
	select {
		case err := <-utils.ErrorQueue:
			logrus.Fatal("服务异常退出,error:",err)
	}
}
