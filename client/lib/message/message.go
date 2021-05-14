package message

import "net/http"


const MESSAGE_TYPE_HTTP = "HTTP"
const MESSAGE_TYPE_TCP = "TCP"


type HttpMessage struct {
	MessageId int64 //消息ID
	MessageType string //消息类型  TCP HTTP
	TargetKey   string  //目标地址的KEY
	HttpRequest HttpRequest
}

type HttpRequest struct{
	Body []byte
	Method string
	Header http.Header
	Host string
	RequestUri string
}
