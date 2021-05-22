package httpServer

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"ngrok/server/lib/message"
	"ngrok/server/lib/utils"
	"strings"
	"time"
)


/**
 http server
 */
func Bootstrap(){
	addr := fmt.Sprintf(":%s",utils.GetSection(utils.WebSection).Key("HTTP_PORT").String())
	logrus.Info("HTTP SERVER START SUCCESS,ADDR :",addr)
	http.HandleFunc("/",httpHandle)
	if err := http.ListenAndServe(addr, nil);err!=nil{
		logrus.Fatal("http start fail, err:",err)
	}
}

//http请求处理
func httpHandle(response http.ResponseWriter , request *http.Request){
		ctx,cancel := context.WithTimeout(context.Background(), 60 * time.Second)
		defer cancel()
		body, _ := ioutil.ReadAll(request.Body)
		//写入数据
		message.HttpMessageQueue<-message.HttpMessage{
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
			case <-ctx.Done():
				//请求超时
				logrus.Info("HTTP请求超时")
				//响应404请求
				response.WriteHeader(http.StatusRequestTimeout)
				_,_ =response.Write([]byte("Target Server Request Time Out"))
				break
		}
}


func getUrlKey(referer string) string {
	newString:= strings.Replace(referer,"http://","",-1)
	newString=strings.Replace(newString,"https://","",-1)
	return strings.Split(newString,".")[0]
}