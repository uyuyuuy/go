package data_migration

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	myconfig "github/data_migration/config"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"time"
)

var oldDb *gorm.DB
var newDb *gorm.DB
var redisClien *redis.Client

func init() {

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

	newDbConfig := config.NewDobiDatabase
	newDb, err = gorm.Open(newDbConfig.DriverName, newDbConfig.DataSourceName)
	if err != nil {
		panic(err)
	}
	newDb.LogMode(true)


	redisConfig := config.RedisDatabase
	var redis_options *redis.Options = &redis.Options{
		Addr:	redisConfig.Addr,
		Password:	redisConfig.Password,
		DB:15}
	redisClien = redis.NewClient(redis_options)

}

func UserData() {

	defer oldDb.Close()
	defer newDb.Close()
	defer redisClien.Close()

	//var address_model oldModel.Address
	//oldDb.Select("address,uid,publicKey").Where(oldModel.Address{
	//	Address:"NVcHJEzm4pPnzndWM47nFPkHmociER54xs",
	//}).Limit(1).Find(&address_model)

	var user_lock sync.Mutex
	var user_lock_group sync.WaitGroup


	for i := 0; i < 100; i ++ {
		user_lock_group.Add(1)
		go DoUserData(&user_lock, &user_lock_group, i)
		time.Sleep(time.Duration(1))
	}

	user_lock_group.Wait()

	os.Exit(110)




}


func DoUserData(user_lock *sync.Mutex, user_lock_group *sync.WaitGroup, i int) bool {
	user_lock.Lock()
	defer user_lock.Unlock()
	defer user_lock_group.Done()

	log.Print("NO", i)

	type UserScan struct {
		UID int64	`gorm:"column:uid"`
		Email string	`gorm:"column:email"`
		Password	string	`gorm:"column:pwd"`
		TradePassword	string	`gorm:"column:pwdtrade"`
		Prand	string	`gorm:"column:prand"`
		Created	int64	`gorm:"column:created"`
		CreateIP string	`gorm:"column:createip"`
		Source	string	`gorm:"column:source"`
		RegisterType	string		`gorm:"column:registertype"`
		FromUID	int64		`gorm:"column:from_uid"`
		Area	string	`gorm:"column:area"`
		GoogleKey	string	`gorm:"column:google_key"`
		Realname	string	`gorm:"column:realname"`
		CardType	string	`gorm:"column:cardtype"`
		IDCard	string	`gorm:"column:idcard"`
		Status	string	`gorm:"column:status"`
		FrontFace	string	`gorm:"column:frontFace"`
		BackFace	string	`gorm:"column:backFace"`
		Handkeep	string	`gorm:"column:handkeep"`
		AutonymCreated	int64	`gorm:"column:autonym_created"`
		AutonymUpdated	int64	`gorm:"column:autonym_updated"`
		AdminID	int64	`gorm:"column:admin"`
		Content	string	`gorm:"column:content"`
		Country	string	`gorm:"column:country"`
	}

	//var LastUID int64

	//字符串类型  string

	var LastUID int
	var uid string
	uid, err := redisClien.Get("LastUID").Result()
	if err != nil {
		panic(err)
	}
	LastUID, err = strconv.Atoi(uid)

	var userScan UserScan
	var selectStr string
	selectStr = "user.uid,user.email,user.name,user.pwd,user.pwdtrade,user.prand," +
		"user.created,user.createip,user.source,user.registertype,user.from_uid,user.area,user.google_key," +
		"autonym.name realname,autonym.cardtype,autonym.idcard,autonym.status,autonym.frontFace,autonym.backFace,autonym.handkeep," +
		"autonym.created autonym_created,autonym.updated autonym_updated,autonym.admin,autonym.content,autonym.country"
	oldDb.Table("user").Select(selectStr).Joins("inner join autonym on autonym.uid = user.uid").Where("user.uid > ?", LastUID).Order("user.uid asc").Limit(1).Scan(&userScan)

	if userScan.UID == 0 {
		panic("No data")
		return false
	}
	fmt.Println(userScan)

	LastUID++
	err = redisClien.Set("LastUID", LastUID, 0).Err()
	if err != nil {
		panic(err)
	}

	return true


	//os.Exit(0)
	//
	//
	//
	//var userModel newModel.UserMain
	//userModel.AccountName = "test02"
	//userModel.Password = "sssdds"
	//userModel.AreaCode = "86"
	//userModel.AvatarUrl = "sdsds"
	//userModel.Email = "842276678@qq.com"
	//userModel.Mobile = "13316588988"
	//
	//insertResult := newDb.Create(&userModel)
	//if insertResult.RowsAffected > 0 {
	//	log.Fatal("插入成功")
	//} else {
	//	log.Fatal(fmt.Println(insertResult.GetErrors()))
	//}

}
