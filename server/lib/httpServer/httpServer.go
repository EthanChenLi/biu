package httpServer

import (
	"encoding/base64"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"ngrok/server/lib/message"
	"strings"
	"time"
)

const ADDR =  ":12138"

/**
 http server
 */
func Bootstrap(){
	logrus.Info("HTTP SERVER START SUCCESS,ADDR :",ADDR)
	http.HandleFunc("/",httpHandle)
	if err := http.ListenAndServe(ADDR, nil);err!=nil{
		logrus.Fatal("http start fail, err:",err)
	}
}

//http请求处理
func httpHandle(response http.ResponseWriter , request *http.Request){
		body, _ := ioutil.ReadAll(request.Body)
		//写入数据
		logrus.Println("收到HTTP消息，写入管道")
		messageId := time.Now().UnixNano()
		message.HttpMessageQueue<-message.HttpMessage{
			MessageId: messageId,
			MessageType:  message.MESSAGE_TYPE_HTTP,
			TargetKey: getUrlKey(request.Host),
			HttpRequest:  message.HttpRequest{
				Body:       body,
				Method:     request.Method,
				Header:     request.Header,
				Host:       request.Host,
				RequestUri: request.RequestURI,
			},
		}
		select{
			case respConent := <-message.HttpResponseMessageQueue:
				//base64 解码
				content ,err := base64.StdEncoding.DecodeString(respConent.Content)
				if err != nil{
					logrus.Warning("base64 encode error:",err)
				}
				//回复内容
				for key,item := range respConent.Headers {
					response.Header().Set(key,item[0])
				}
				_,err =response.Write(content)
				if err !=nil{
					logrus.Warning("http resp fail:",err)
				}
			case errMsg:= <-message.ErrorQueue:
				logrus.Warning(errMsg)
				//响应404请求
				response.WriteHeader(http.StatusNotFound)
				_,_ =response.Write([]byte("Target Page does not exist or client is closed"))
		}
}


func getUrlKey(referer string) string {
	newString:= strings.Replace(referer,"http://","",-1)
	newString=strings.Replace(newString,"https://","",-1)
	return strings.Split(newString,".")[0]
}