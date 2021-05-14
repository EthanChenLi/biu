package utils

import "encoding/json"

type TcpMessage struct {
	MessageLength uint32 `json:"message_length"` //消息长度 4字节
	MessageId int64 `json:"message_id"`  // 消息ID ， 4字节
	Body []byte `json:"body"` //消息内容 JSON
	BodyStruct TcpBody
}

type TcpBody struct {
	Content string `json:"content"`
	Headers map[string][]string `json:"headers"`
}

const TCP_MESSAGE_TYPE_INIT = 1 //初始化消息
const TCP_MESSAGE_TYPE_HEADER = 101 //TCP心跳包


func BuildTcpMessage(messageId int64,content string,headers map[string][]string) *TcpMessage {
	body := TcpBody{
		Content: content,
		Headers: headers,
	}
	bodyJson ,_ :=json.Marshal(body)
	return &TcpMessage{
		MessageLength: uint32(len(bodyJson)), //内容长度+header头长度
		MessageId: messageId,
		Body: bodyJson,
	}

}