package data_migration

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github/data_migration/mycore"
	"strings"

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
	log.Print("userinfo init")
}


func UserData() (err error) {

	defer oldDb.Close()
	defer newDbCore.Close()
	defer redisClien.Close()
	defer func() {
		if p := recover();p != nil{
			//异常日志
			Logrus.Warn(reflect.ValueOf(p).String())
			err = errors.New(reflect.ValueOf(p).String())
		}
	}()
	err = nil


	//日志设置
	Logrus = logrus.New()
	Logrus.SetLevel(logrus.InfoLevel)
	Logrus.SetFormatter(&logrus.JSONFormatter{})

	logfile, err := os.OpenFile("./src/github/data_migration/log/userinfo.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0777)
	if err != nil {
		panic("logfile is error")
	}
	Logrus.SetOutput(logfile)


	var user_lock sync.Mutex
	var user_lock_group sync.WaitGroup

	for i := 0; i < 10; i ++ {
		user_lock_group.Add(1)
		go doUserData(&user_lock, &user_lock_group)
		time.Sleep(1000 * time.Millisecond)	//100毫秒，一秒十次
	}

	//阻塞，等待
	user_lock_group.Wait()

	return	err

}

/**
用户基本信息和实名认证信息同步
 */
func doUserData(user_lock *sync.Mutex, user_lock_group *sync.WaitGroup) {
	user_lock.Lock()
	defer user_lock.Unlock()
	defer user_lock_group.Done()

	var LastUID uint64	//无符号数字
	var uid string
	uid, err := redisClien.Get("LastUID").Result()
	if err != nil {
		panic(err)
	}
	LastUID, err = strconv.ParseUint(uid, 10,64)


	var userScan oldModel.UserScan
	var selectStr string
	selectStr = "user.uid,user.email,user.mo,user.name,user.pwd,user.pwdtrade,user.prand," +
		"user.created,user.createip,user.source,user.registertype,user.from_uid,user.area,user.google_key," +
		"autonym.id autonym_id,autonym.name realname,autonym.cardtype,autonym.idcard,autonym.status,autonym.frontFace,autonym.backFace,autonym.handkeep," +
		"autonym.created autonym_created,autonym.updated autonym_updated,autonym.admin,autonym.content,autonym.country"
	oldDb.Table("user").Select(selectStr).Joins("left join autonym on autonym.uid = user.uid").Where("user.uid > ? ", LastUID).Order("user.uid asc").Limit(1).Scan(&userScan)

	if userScan.UID == 0 {
		panic("No data")
	} else {
		Logrus.Info("开始同步：", userScan.UID)
	}

	defer func() {
		err = redisClien.Set("LastUID", userScan.UID, 0).Err()
		if err != nil {
			Logrus.Warn("Redis 键LastUID更新失败：", userScan.UID)
			panic(err)
		}
	}()

	//插入新数据库，uid不能变
	var userMain newModel.UserMain
	var userInfo newModel.UserInfo
	var googelAuthenticator newModel.GoogleAuthenticator
	var userAutonym	newModel.Autonym

	//user_main表
	userMain.UID = userScan.UID
	switch userScan.RegisterType {
		case 0:
			if userScan.Email != "" {
				userMain.AccountName.String = userScan.Email
				userMain.AccountName.Valid = true
			} else if userScan.Mobile != "" {
				userMain.AccountName.String = userScan.Mobile
				userMain.AccountName.Valid = true
			} else {
				Logrus.Warn(strconv.FormatUint(userScan.UID,10) + "无手机号码和邮箱")
			}
		case 1:
			userMain.AccountName.String = userScan.Email
			userMain.AccountName.Valid = true
		case 2:
			userMain.AccountName.String = userScan.Mobile
			userMain.AccountName.Valid = true
		default:
			Logrus.Warn(strconv.FormatUint(userScan.UID,10) + " RegisterType is wrong ")
	}

	userMain.Password = userScan.Password
	userMain.AreaCode = userScan.Area

	if userScan.Mobile != "" {
		userMain.Mobile.String = userScan.Mobile
		userMain.Mobile.Valid = true
	}
	if userScan.Email != "" {
		userMain.Email.String = userScan.Email
		userMain.Email.Valid = true
	}

	//googel_authenticator 表
	googelAuthenticator.UserID = userScan.UID

	//autonym表
	userAutonym.UserID = userScan.UID
	userAutonym.Front = userScan.FrontFace
	userAutonym.HandFront = userScan.Handkeep
	userAutonym.Reverse = userScan.BackFace
	userAutonym.Remark = userScan.Content

	//userinfo
	userInfo.UserID = userScan.UID
	userInfo.ID = userScan.IDCard
	userInfo.Name = ""
	//userInfo.CountryID = userScan.Country

	tm := time.Unix(userScan.Created, 0)
	userInfo.RegisterTime = tm.Format("2019-07-29 20:00:00")

	//证件类型
	if userScan.CardType != 0 {
		switch userScan.CardType {
		case 1:
			userInfo.CredentialsType = "JZ"
		case 2:
			userInfo.CredentialsType = "HZ"
		}
	}

	//实名认证状态
	if userScan.AutonymID != 0 {
		switch userScan.Status {
		case 0:
			userInfo.VipStatus = "Checking"
		case 1:
			userInfo.VipStatus = "Checking"
		case 2:
			userInfo.VipStatus = "Yes"
		case 3:
			userInfo.VipStatus = "No"
		}
	}

	//执行数据同步
	tx := newDbCore.Begin()
	dberr := make([]error, 6)

	dberr[0] = tx.Create(&userMain).Error
	if userScan.GoogleKey != "" {
		dberr[1] = tx.Model(&googelAuthenticator).Update(newModel.GoogleAuthenticator{SecretKey:userScan.GoogleKey}).Error
	}
	dberr[2] = tx.Create(&userAutonym).Error
	dberr[3] = tx.Create(&userInfo).Error


	//用户地址
	var userAddress []oldModel.Address
	oldDb.Where(oldModel.Address{UID:userScan.UID}).Find(&userAddress)
	Logrus.Warn(userAddress)

	if len(userAddress) > 0 {
		addressCount := len(userAddress)
		valueStrings := make([]string, addressCount)
		valueArgs := make([]interface{}, 5 * addressCount)

		for i := 0; i < addressCount; i++ {
			valueStrings = append(valueStrings, "(?,?,?,?,?)")
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

		dberr[4] = tx.Exec(sqlStr,valueArgs...).Error

	} else {
		Logrus.Warn(userScan.UID, "无地址数据")
	}

	//用户资产
	var userAsset oldModel.Asset
	oldDb.Where(oldModel.Asset{UID:userScan.UID}).Find(&userAsset)

	assetMap := mycore.Struct2Map(userAsset)
	fmt.Println(assetMap)

	var newCoin []newModel.Coin
	newDbCore.Find(&newCoin)
	//fmt.Println(newCoin)
	//fmt.Println(len(newCoin))

	assetCount := len(newCoin)
	valueStrings := make([]string, 0, assetCount)
	valueArgs := make([]interface{}, 0, 4 * assetCount)

	for _, coin := range newCoin {
		coinName := coin.Name
		coinID := coin.ID
		valueStrings = append(valueStrings, "(?,?,?,?)")
		valueArgs = append(valueArgs, userScan.UID)
		valueArgs = append(valueArgs, coinID)	//币种ID
		if v, ok := assetMap[strings.ToUpper(coinName)+"OVER"]; ok {
			valueArgs = append(valueArgs, v)
		} else {
			valueArgs = append(valueArgs, 0)
		}
		if v, ok := assetMap[strings.ToUpper(coinName)+"LOCK"]; ok {
			valueArgs = append(valueArgs, v)
		} else {
			valueArgs = append(valueArgs, 0)
		}
	}
	sqlStr := fmt.Sprintf("INSERT INTO t6025(F02,F03,F04,F05) VALUES %s", strings.Join(valueStrings, ","))
	Logrus.Info("valueStrings：", valueStrings)
	Logrus.Info("strings.Join(valueStrings,,)：", strings.Join(valueStrings, ","))
	Logrus.Info("valueArgs：", valueArgs)
	Logrus.Info("sqlStr：", sqlStr)
	dberr[5] = tx.Exec(sqlStr,valueArgs...).Error

	//os.Exit(0)

	isHaveError := false
	for _, e := range dberr {
		if e != nil {
			isHaveError = true
			Logrus.Warn(e)
			break
		}
	}


	if isHaveError {
		tx.Rollback()
		Logrus.Warn("同步失败：", userScan.UID)
	} else {
		tx.Commit()
		Logrus.Warn("同步成功：", userScan.UID)
	}
	return

}
