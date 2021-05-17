package utils

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/ini.v1"
)

const TcpSection = "TCP"
const WebSection = "WEB"

var Config *ini.File


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
