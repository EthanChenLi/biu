package utils

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"io"
	"net"
)

//tcp 沾包处理

const TCP_DATA_LENGTH = 12


// Encode 将消息编码
func Encode(message *TcpMessage) ([]byte, error) {
	//创建一个存放bytes字节的缓冲
	dataBuff := bytes.NewBuffer([]byte{})
	//写dataLen
	if err := binary.Write(dataBuff, binary.LittleEndian, message.MessageLength); err != nil {
		return nil, err
	}
	//写msgID
	if err := binary.Write(dataBuff, binary.LittleEndian, message.MessageId); err != nil {
		return nil, err
	}
	//写body数据
	if err := binary.Write(dataBuff, binary.LittleEndian, message.Body); err != nil {
		return nil, err
	}
	return dataBuff.Bytes(),nil
}

// Decode 解码消息
func Decode(conn net.Conn) (*TcpMessage, error) {
	//截图消息头的长度
	headData := make([]byte, TCP_DATA_LENGTH)
	if _,err := io.ReadFull(conn,headData);err != nil{
		return nil,err
	}
	dataBuff := bytes.NewReader(headData)
	tcpMessage := &TcpMessage{}

	//读dataLen
	if err := binary.Read(dataBuff, binary.LittleEndian, &tcpMessage.MessageLength); err != nil {
		return nil, err
	}

	//读msgID
	if err := binary.Read(dataBuff, binary.LittleEndian, &tcpMessage.MessageId); err != nil {
		return nil, err
	}

	//截取body的长度
	body := make([]byte, tcpMessage.MessageLength)
	_,err := io.ReadFull(conn,body)
	if err != nil{
		return nil,err
	}
	logrus.Info("收到的原始消息:",string(body))
	//解析tcpMessage的body内容
	tcpBody := &TcpBody{}
	if err := json.Unmarshal(body,&tcpBody); err!=nil{
		return nil,err
	}
	tcpMessage.BodyStruct = *tcpBody
	return tcpMessage,nil
}