package main

import (
	"github.com/sirupsen/logrus"
	"os"
)



func main() {


	//go内置原生日志库 log，就三个方法 log.Fatal() print、panic、fatal
	//log.Fatal()
	//log.Print()

	//创建实例
	//logsss := logrus.Logger{}		//这个是logrus的logger类型
	logLogrus := logrus.New() //官方文档，就是返回logger结构指针

	//通过实例的方法设置实例的属性
	logLogrus.SetFormatter(&logrus.JSONFormatter{})
	file,_ := os.OpenFile("1.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	logLogrus.SetOutput(file)
	logLogrus.SetLevel(logrus.WarnLevel)
	logLogrus.WithFields(logrus.Fields{
		"test": "test",
	}).Warn("Wrong")


	//直接设置实例化的属性
	logLogrus2 := logrus.New()
	logLogrus2.Out = os.Stdout
	logLogrus2.Level = logrus.WarnLevel
	logLogrus2.Formatter = &logrus.JSONFormatter{}
	logLogrus2.WithFields(logrus.Fields{
		"test": "test",
	}).Warn("something is wrong!")



}
