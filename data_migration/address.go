package data_migration

import (
	"errors"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sirupsen/logrus"
	//"github/data_migration"
	"log"
	"os"
	"reflect"
	"strings"

	oldModel "github/data_migration/models/old"
	"strconv"
	"sync"
	"time"
)

//该文件程序已移植到userinfo.go文件中


func init() {
	log.Print("address init")

}


func userAddress() (err error) {
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

	logfile, err := os.OpenFile("./src/github/data_migration/log/address.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0777)
	if err != nil {
		panic("logfile is error")
	}
	Logrus.SetOutput(logfile)

	var user_lock sync.Mutex
	var user_lock_group sync.WaitGroup

	for i := 0; i < 10; i ++ {
		user_lock_group.Add(1)
		go doUserAddress(&user_lock, &user_lock_group)
		time.Sleep(1000 * time.Millisecond)	//100毫秒，一秒十次
	}

	//阻塞，等待
	user_lock_group.Wait()

	return	err

}

/**
用户基本信息和实名认证信息同步
*/
func doUserAddress(user_lock *sync.Mutex, user_lock_group *sync.WaitGroup) {
	user_lock.Lock()
	defer user_lock.Unlock()
	defer user_lock_group.Done()

	var LastUID uint64	//无符号数字
	var uid string
	uid, err := redisClien.Get("Address_LastUID").Result()
	if err != nil {
		panic(err)
	}
	LastUID, err = strconv.ParseUint(uid, 10,64)
	defer func() {
		err = redisClien.Set("Address_LastUID", LastUID, 0).Err()
		if err != nil {
			Logrus.Warn("Redis 键Address_LastUID更新失败：", LastUID)
			panic(err)
		}

	}()

	type UserInfo struct {
		Uid uint64
	}
	var userInfo UserInfo
	oldDb.Table("user").Select("uid").Where("user.uid > ? ", LastUID).Order("uid asc").Limit(1).Scan(&userInfo)

	if userInfo.Uid > 0 {

		var userAddress []oldModel.Address
		oldDb.Where(oldModel.Address{UID:userInfo.Uid}).Find(&userAddress)
		Logrus.Warn(userAddress)

		if len(userAddress) > 0 {
			LastUID = userInfo.Uid
			Logrus.Info("开始同步：", userInfo.Uid)

			addressCount := len(userAddress)
			valueStrings := make([]string, 0, addressCount)
			valueArgs := make([]interface{}, 0, 5 * addressCount)

			for i := 0; i < addressCount; i++ {
				valueStrings[i] = "(?,?,?,?,?)"
				valueArgs = append(valueArgs, userAddress[i].UID)
				valueArgs = append(valueArgs, userAddress[i].Coin)	//币种ID
				//valueArgs = append(valueArgs, i)
				valueArgs = append(valueArgs, userAddress[i].Address)
				valueArgs = append(valueArgs, userAddress[i].Created)
				valueArgs = append(valueArgs, userAddress[i].Label)
			}

			sqlStr := fmt.Sprintf("INSERT INTO t6012_1(F02,F03,F04,F05,F08) VALUES %s", strings.Join(valueStrings, ","))
			Logrus.Info("valueStrings：", valueStrings)
			Logrus.Info("strings.Join(valueStrings,,)：", strings.Join(valueStrings, ","))
			Logrus.Info("valueArgs：", valueArgs)
			Logrus.Info("sqlStr：", sqlStr)
			if err := newDbCore.Exec(sqlStr,valueArgs...).Error; err != nil{
				Logrus.Warn("同步失败：", userInfo.Uid)
				panic(err)
			} else {
				Logrus.Info("同步成功：", userInfo.Uid)
			}

		} else {
			Logrus.Warn(userInfo.Uid, "无地址数据")
		}
	}

	return

}
