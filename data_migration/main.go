package data_migration

import (
	"github.com/BurntSushi/toml"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sirupsen/logrus"
	myconfig "github/data_migration/config"
	"log"
	"path/filepath"
)

var oldDb *gorm.DB
var newDbCore *gorm.DB
var newDbTrade *gorm.DB
var redisClien *redis.Client
var Logrus *logrus.Logger


func init() {
	log.Print("main init")

	var config myconfig.Config

	//获取绝对路径
	filePath, err := filepath.Abs("./src/github/data_migration/config/config.toml")
	if err != nil {
		panic(err)
	}
	//fmt.Println(filePath)

	//映射数据库连接配置
	if _, err := toml.DecodeFile(filePath, &config); err != nil {
		panic(err)
	}

	oldDbConfig := config.OldDobiDatabase
	oldDb, err = gorm.Open(oldDbConfig.DriverName, oldDbConfig.DataSourceName)
	if err != nil {
		panic(err)
	}
	oldDb.LogMode(true)

	newDbCoreConfig := config.NewDobiDatabase.Core
	newDbCore, err = gorm.Open(newDbCoreConfig.DriverName, newDbCoreConfig.DataSourceName)
	if err != nil {
		panic(err)
	}
	newDbCore.LogMode(true)

	newDbTradeConfig := config.NewDobiDatabase.Trade
	newDbTrade, err = gorm.Open(newDbTradeConfig.DriverName, newDbTradeConfig.DataSourceName)
	if err != nil {
		panic(err)
	}
	newDbTrade.LogMode(true)

	redisConfig := config.RedisDatabase
	var redis_options = &redis.Options{
		Addr:	redisConfig.Addr,
		Password:	redisConfig.Password,
		DB:	redisConfig.DB,
	}
	redisClien = redis.NewClient(redis_options)

}





