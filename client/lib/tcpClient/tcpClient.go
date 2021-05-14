package tcpClient

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net"
	"net/http"
	"ngrok/client/lib/message"
	"ngrok/client/lib/utils"
)




var TargetIp string
var TargetPort string
var TargetHost string

/**
 tcp 客户端
 */
func Bootstrap(tcpAdd string){
	conn,err := net.Dial("tcp", tcpAdd)
	if err!=nil{
		logrus.Panic("tcp dial fail, err:",err)
		return
	}
	//初始化链接消息
	content ,_ :=utils.Encode(utils.BuildTcpMessage(
		utils.TCP_MESSAGE_TYPE_INIT,TargetHost,nil,
		))
	_, err =conn.Write(content)
	if err!=nil{
		logrus.Warning("tcp init fail,err",err)
		return
	}

	logrus.Info("tcp连接成功")
	go CreateTcpHeader(conn) //创建心跳任务
	//读取服务端消息

	for {
		//消息解码
		//拆包处理
		tcpMessage, err := utils.Decode(conn)
		if err != nil {
			logrus.Warning("TCP READY ERROR:",err)
			conn.Close()
			break
		}
		go messageHandle(tcpMessage,conn)
	}
}

func messageHandle(msg *utils.TcpMessage,conn net.Conn){


   data := &message.HttpMessage{}
   json.Unmarshal([]byte(msg.BodyStruct.Content),&data)
   logrus.Info("tcp发来的消息：",data)

   targetUrl:=fmt.Sprintf("http://%s:%s",TargetIp,TargetPort)+data.HttpRequest.RequestUri
   logrus.Info("目标请求HTTP地址：",targetUrl)

   req,err:= http.NewRequest(data.HttpRequest.Method,targetUrl,bytes.NewBuffer(data.HttpRequest.Body))

   if err !=nil {
   	 logrus.Warning("htt request err:",err)
	   return
   }
   //设置header头
	for k,item := range data.HttpRequest.Header{
		req.Header.Set(k,item[0])
	}

	client := &http.Client{}
	resp, err := client.Do(req)

	defer resp.Body.Close()
	if err != nil {
		logrus.Warning("htt request err:",err)
		return
	}
	httpContent, err := ioutil.ReadAll(resp.Body)
	//base64
	httpContentByBase64 := base64.StdEncoding.EncodeToString(httpContent)
	if err !=nil {
		logrus.Warning("htt ioutil err:",err)
		return
	}

	//回复给server端
	 content ,_ :=utils.Encode(
	 	utils.BuildTcpMessage(data.MessageId,httpContentByBase64,resp.Header),
	 	)
	_, _ = conn.Write(content)
}