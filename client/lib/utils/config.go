package utils

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/ini.v1"
)

var Config *ini.File

const TcpSection = "TCP"
const WebSection = "WEB"
const BaseSection = "BASE"

const SERVER_TYPE_TCP = "TCP"
const SERVER_TYPE_HTTP = "HTTP"
const SERVER_TYPE_ALL = "ALL"



//初始化配置文件
func NewConfig()  {
	var err error
	Config, err = ini.Load("conf.ini")
	if err != nil {
		logrus.Fatal("读取配置失败")
	}
}

func GetSection(section string) *ini.Section{
	return Config.Section(section)
}

