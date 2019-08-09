package data_migration

import (
	"errors"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sirupsen/logrus"
	newModel "github/data_migration/models/new"
	oldModel "github/data_migration/models/old"
	"log"
	"os"
	"reflect"
	"strconv"
	"sync"
)


func init() {
	log.Print("openapi init")

}


func OpenApi() (err error) {
	defer func() {
		if p := recover();p != nil{
			//异常日志
			Logrus.Warn(reflect.ValueOf(p).String())
			err = errors.New(reflect.ValueOf(p).String())
		}
	}()
	defer oldDb.Close()
	defer newDbCore.Close()
	defer redisClien.Close()

	err = nil

	//日志设置
	Logrus = logrus.New()
	Logrus.SetLevel(logrus.InfoLevel)
	Logrus.SetFormatter(&logrus.JSONFormatter{})

	logfile, err := os.OpenFile("./src/github/data_migration/log/openapi.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0777)
	if err != nil {
		panic("logfile is error")
	}
	Logrus.SetOutput(logfile)

	var user_lock sync.Mutex
	var user_lock_group sync.WaitGroup

	for i := 0; i < 1300; i ++ {
		user_lock_group.Add(1)
		go doOpenApi(&user_lock, &user_lock_group)
		//time.Sleep(1000 * time.Millisecond)	//100毫秒，一秒十次
	}

	//阻塞，等待
	user_lock_group.Wait()

	return	err

}

/**
用户基本信息和实名认证信息同步
*/
func doOpenApi(user_lock *sync.Mutex, user_lock_group *sync.WaitGroup) {
	user_lock.Lock()
	defer user_lock.Unlock()
	defer user_lock_group.Done()

	var LastID int64	//无符号数字
	var openid string
	openid, err := redisClien.Get("Openapi_LastID").Result()
	if err != nil {
		panic(err)
	}
	LastID, err = strconv.ParseInt(openid, 10,64)
	defer func() {
		err = redisClien.Set("Openapi_LastID", LastID, 0).Err()
		if err != nil {
			Logrus.Warn("Redis 键Openapi_LastID更新失败：", LastID)
			panic(err)
		}

	}()

	var openapi oldModel.Openapi
	oldDb.Table("openapi").Where("id > ? ", LastID).Order("id asc").Limit(1).Scan(&openapi)

	if openapi.ID > 0 {
		LastID = openapi.ID

		var newOpenapi newModel.Openapi
		newOpenapi.UID = openapi.UID
		newOpenapi.SecretKey = openapi.SecretKey
		newOpenapi.AccessKey = openapi.AccessKey
		newOpenapi.Ip = openapi.Ip

		if createErr := newDbCore.Create(&newOpenapi).Error; createErr != nil {
			Logrus.Warn(openapi.ID, "同步失败")
		} else {
			Logrus.Info(openapi.ID, "同步成功")
		}

	} else {
		Logrus.Warn("同步完毕")
	}

	return

}
