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


/**
 tcp 客户端
 */
func Bootstrap(){
	tcpAdd := fmt.Sprintf("%s:%s",
		utils.GetSection(utils.BaseSection).Key("SERVER_IP").String(),
		utils.GetSection(utils.WebSection).Key("SERVER_PORT").String(),
		)
	conn,err := net.Dial("tcp", tcpAdd)
	if err!=nil{
		logrus.Panic("tcp dial fail, err:",err)
		return
	}
	//初始化链接消息
	content ,_ :=utils.Encode(utils.BuildTcpMessage(
		utils.TCP_MESSAGE_TYPE_INIT,utils.GetSection(utils.BaseSection).Key("KEY").String(),nil,
		))
	_, err =conn.Write(content)
	if err!=nil{
		logrus.Warning("tcp init fail,err",err)
		return
	}

	logrus.Info("服务器连接成功")
	go CreateTcpHeader(conn) //创建心跳任务
	//读取服务端消息

	for {
		//消息解码
		//拆包处理
		tcpMessage, err := utils.Decode(conn)
		if err != nil {
			_ = conn.Close()
			utils.ErrorQueue<-err //异常退出
			break
		}
		go messageHandle(tcpMessage,conn)
	}
}

func messageHandle(msg *utils.TcpMessage,conn net.Conn){
   data := &message.HttpMessage{}
   json.Unmarshal([]byte(msg.BodyStruct.Content),&data)
   logrus.Info("tcp发来的消息：",data)
   webAddr := utils.GetSection(utils.WebSection).Key("WEB_ADDR").String()
   targetUrl:=fmt.Sprintf("http://%s%s",webAddr,data.HttpRequest.RequestUri)
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