package message

import (
	"encoding/json"
	"errors"
	"github.com/sirupsen/logrus"
	"net/http"
	"ngrok/server/lib/utils"
	"time"
)

const MESSAGE_TYPE_HTTP = "HTTP"
const MESSAGE_TYPE_TCP = "TCP"

var HttpMessageQueue chan HttpMessage  //HTTP消息管道
var HttpResponseMessageQueue chan HttpResponse  //HTTP消息管道
var ErrorQueue chan error


type HttpResponse struct {
	MessageId int64 //消息ID
	Content string //返回http消息
	Headers map[string][]string //http header
}


type HttpMessage struct {
	MessageType string //消息类型  TCP HTTP
	TargetKey string  //目标地址的KEY
	HttpRequest HttpRequest
}

type HttpRequest struct{
	Body []byte
	Method string
	Header http.Header
	Host string
	RequestUri string
}

func NewMessageQueue()  {
	HttpMessageQueue = make(chan HttpMessage,1)
	HttpResponseMessageQueue = make(chan HttpResponse)
	ErrorQueue  = make(chan error)
}


func SendHttpRequestMessage(){
	for{
		select{
		//读到http消息
		case requestDate :=<-HttpMessageQueue:
			go pushHttpContent(requestDate)
		}
	}
}

//推送http内容到tcp
func pushHttpContent(request HttpMessage){
	fid,ok := utils.TcpDomainMap.TcpDomainMap[request.TargetKey]
	if ok{
		//校验一下fid是否已经离线
		_,err := utils.TcpConnectMap.GetConnectMap(fid)
		if err != nil {
			ErrorQueue<- errors.New("客户端已经离线")
			return
		}
		conn,err := utils.TcpConnectMap.GetConnectMap(fid)
		if err!=nil{
			logrus.Info("GET TCP CONN ERR :",err)
			return
		}
		requestByte,err := json.Marshal(request)
		if err!= nil{
			logrus.Info("STRUCT TO JSON FAIL :",err)
			return
		}
		//使用编码
		messageId := time.Now().UnixNano()
		content ,_ :=utils.Encode(utils.BuildTcpMessage(
			messageId,string(requestByte),nil,
		))
		_, _ = conn.Conn.Write(content) //发送给TCP客户端

	}else{
		ErrorQueue<- errors.New("域名列表找不到该目标地址")
		return
	}
}