package data_migration

import (
	"errors"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sirupsen/logrus"
	"os"
	"reflect"
	"strconv"
	"sync"
	"time"
	oldModel "github/data_migration/models/old"
)

var Market map[string]int

func init() {

	Market["BTC"] = 1
	Market["ETH"] = 2
	Market["USDT"] = 3
	Market["DOB"] = 4
}

//交易区


func Coin() (err error) {
	defer func() {
		if p := recover();p != nil{
			//异常日志
			Logrus.Warn(reflect.ValueOf(p).String())
			err = errors.New(reflect.ValueOf(p).String())
		}
	}()
	defer oldDb.Close()
	defer newDbTrade.Close()
	defer redisClien.Close()

	err = nil

	//日志设置
	Logrus = logrus.New()
	Logrus.SetLevel(logrus.InfoLevel)
	Logrus.SetFormatter(&logrus.JSONFormatter{})

	logfile, err := os.OpenFile("./src/github/data_migration/log/coin.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0777)
	if err != nil {
		panic("logfile is error")
	}
	Logrus.SetOutput(logfile)




	var user_lock sync.Mutex
	var user_lock_group sync.WaitGroup

	for true {
		user_lock_group.Add(1)
		go DoCoin(&user_lock, &user_lock_group)
		time.Sleep(1000 * time.Millisecond)	//100毫秒，一秒十次
	}

	//阻塞，等待
	user_lock_group.Wait()

	return	err

}

/**
用户基本信息和实名认证信息同步
*/
func DoCoin(user_lock *sync.Mutex, user_lock_group *sync.WaitGroup) {
	user_lock.Lock()
	defer user_lock.Unlock()
	defer user_lock_group.Done()

	var LastCoinID uint64	//无符号数字
	var coinID string
	coinID, err := redisClien.Get("Coin_LastID").Result()
	if err != nil {
		panic(err)
	}
	LastCoinID, err = strconv.ParseUint(coinID, 10,64)
	defer func() {
		err = redisClien.Set("Coin_LastID", LastCoinID, 0).Err()
		if err != nil {
			Logrus.Warn("Redis 键Coin_LastID更新失败：", LastCoinID)
			panic(err)
		}

	}()

	//var coinModel oldModel.
	//oldDb.Where("id > ?", LastCoinID).Limit(1).Find(&addressModel)
	//if (oldModel.Address{}) != addressModel {
	//
	//
	//	newDbTrade.Create()
	//}


	return

}
