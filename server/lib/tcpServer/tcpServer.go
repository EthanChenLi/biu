package tcpServer

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"log"
	"net"
	"ngrok/server/lib/message"
	"ngrok/server/lib/utils"
)

var counter	 *utils.TcpCount
/**
 初始化tcp的维护表
 */
func init(){
	//初始化tcp计数器
	counter = utils.NewTcpCounter()
}

/**
 *	tcpServer server 启动
 */
func Bootstrap()  {
	addr := fmt.Sprintf(":%s",utils.GetSection(utils.WebSection).Key("TCP_PORT").String())
	//启动tcp监听
	listen ,err := net.Listen("tcp",addr)
	if err !=nil{
		log.Panicln("tcpServer listen err: ",err)
		return
	}
	logrus.Info("TCP SERVER START SUCCESS, ADDR:",addr)

	for{
		conn,err := listen.Accept() //建立连接
		if err != nil{
			log.Println("tcpServer accept err:",err)
			continue
		}
		go process(conn) //处理单条连接请求
	}
}

//处理连接请求
func process(conn net.Conn){
	defer conn.Close()
	fid := counter.IncrTcpCount() //增加fid
	//更新链接表
	utils.TcpConnectMap.SetNewConnectMap(fid,conn)
	logrus.Info("有一个新的链接,fid:",fid)
	//循环从buff读取数据
	for{
		//拆包处理
		tcpMessage, err := utils.Decode(conn)
		if err != nil {
			//断开处理
			_ =  utils.TcpConnectMap.CloseConnect(fid)
			logrus.Warn("decode msg failed, err:", err)
			return
		}
		//判断消息类型
		switch tcpMessage.MessageId {
			case  utils.TCP_MESSAGE_TYPE_INIT:
				//更新域名映射表
				utils.TcpDomainMap.WriteTcpDomainMap(tcpMessage.BodyStruct.Content,fid)
		case utils.TCP_MESSAGE_TYPE_HEADER:
			//心跳包
			_ = utils.TcpConnectMap.FlushHeaderTime(fid)
		default:
			//普通消息回复
			message.HttpResponseMessageQueue<-message.HttpResponse{
				Headers: 	tcpMessage.BodyStruct.Headers,
				MessageId:  tcpMessage.MessageId,
				Content:    tcpMessage.BodyStruct.Content,
			}
		}
	}
}
