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
	"time"
)


func init() {
	log.Print("news init")

}


func News() (err error) {
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

	logfile, err := os.OpenFile("./src/github/data_migration/log/news.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0777)
	if err != nil {
		panic("logfile is error")
	}
	Logrus.SetOutput(logfile)

	var user_lock sync.Mutex
	var user_lock_group sync.WaitGroup

	for i := 0; i < 1300; i ++ {
		user_lock_group.Add(1)
		go doNews(&user_lock, &user_lock_group)
		//time.Sleep(1000 * time.Millisecond)	//100毫秒，一秒十次
	}

	//阻塞，等待
	user_lock_group.Wait()

	return	err

}

/**
用户基本信息和实名认证信息同步
*/
func doNews(user_lock *sync.Mutex, user_lock_group *sync.WaitGroup) {
	user_lock.Lock()
	defer user_lock.Unlock()
	defer user_lock_group.Done()

	var LastID int64	//无符号数字
	var newsid string
	newsid, err := redisClien.Get("News_LastID").Result()
	if err != nil {
		panic(err)
	}
	LastID, err = strconv.ParseInt(newsid, 10,64)
	defer func() {
		err = redisClien.Set("News_LastID", LastID, 0).Err()
		if err != nil {
			Logrus.Warn("Redis 键News_LastID更新失败：", LastID)
			panic(err)
		}

	}()

	var news oldModel.News
	oldDb.Table("openapi").Where("id > ? ", LastID).Order("id asc").Limit(1).Scan(&news)

	if news.ID > 0 {
		LastID = news.ID

		var newNews newModel.News
		newNews.Title = news.Title
		newNews.Content = news.Content

		categorySlice := make([]string, 10)
		categorySlice[2] = "Notice"
		categorySlice[4] = "Activity"
		categorySlice[5] = "Project"
		categorySlice[6] = "Information"
		newNews.Type = categorySlice[news.Category]

		tm := time.Unix(news.Created, 0)
		newNews.CreateTime = tm.Format("2019-08-09 15:00:00")

		if news.LanguageCode == "cn" {
			news.LanguageCode = "zh"
		}
		newNews.Lang = news.LanguageCode
		newNews.Pv = news.Click
		newNews.From = news.Source
		newNews.Sort = news.Sort

		if createErr := newDbCore.Create(&newNews).Error; createErr != nil {
			Logrus.Warn(news.ID, "同步失败")
		} else {
			Logrus.Info(news.ID, "同步成功")
		}

	} else {
		Logrus.Warn("同步完毕")
	}

	return

}
