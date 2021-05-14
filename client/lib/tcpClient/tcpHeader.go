package tcpClient

import (
	"net"
	"ngrok/client/lib/utils"
	"time"
)

const HEADER_TIME_INTERVAL = 10 //时间间隔
const HEADER_CONTENT = "~"

func CreateTcpHeader(conn net.Conn){
	ticker := time.NewTicker( HEADER_TIME_INTERVAL * time.Second )
	for _ = range ticker.C {
		content ,_ :=utils.Encode(
			utils.BuildTcpMessage(
				utils.TCP_MESSAGE_TYPE_HEADER, HEADER_CONTENT,nil,
			),
		)
		_, _ =conn.Write(content)
	}
}