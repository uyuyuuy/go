package main

import (
	log "github.com/sirupsen/logrus"
	"os"
)



func main() {


	//go内置原生日志库 log，就三个方法 log.Fatal() print、panic、fatal
	//log.Fatal()
	//log.Print()

	log.SetFormatter(&log.JSONFormatter{})

	file,_ := os.OpenFile("1.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	log.SetOutput(file)

	log.SetLevel(log.WarnLevel)

	log.WithFields(log.Fields{
		"test": "test",
	}).Warn("Wrong")

	//var log = logrus.New()
	//log.Out = os.Stdout
	//log.Formatter = &logrus.JSONFormatter{}
	//log.WithFields(logrus.Fields{
	//	"test": "test",
	//}).Warn("something is wrong!")



}
