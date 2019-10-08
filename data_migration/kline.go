package data_migration

import (
	"encoding/json"
	"errors"
	"github.com/go-redis/redis"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sirupsen/logrus"
	newModel "github/data_migration/models/new"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
	"sync"
)


func init() {
	log.Print("news init")

}


func Kline() (err error) {
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

	logfile, err := os.OpenFile("./src/github/data_migration/log/kline.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0777)
	if err != nil {
		panic("logfile is error")
	}
	Logrus.SetOutput(logfile)

	var user_lock sync.Mutex
	var user_lock_group sync.WaitGroup

	var newCoin []newModel.Coin
	newDbCore.Limit(5).Find(&newCoin)
	Logrus.Info("同步K线的币种", newCoin)

	for _, coin := range newCoin {
		user_lock_group.Add(1)
		go doKline(&user_lock, &user_lock_group, coin.ID)
	}

	//阻塞，等待
	user_lock_group.Wait()

	return	err

}

/**
用户基本信息和实名认证信息同步
*/
func doKline(user_lock *sync.Mutex, user_lock_group *sync.WaitGroup, coin_id int64) {
	user_lock.Lock()
	defer user_lock.Unlock()
	defer user_lock_group.Done()

	//日志设置
	Logrus = logrus.New()
	Logrus.SetLevel(logrus.InfoLevel)
	Logrus.SetFormatter(&logrus.JSONFormatter{})

	logfile, err := os.OpenFile("./src/github/data_migration/log/kline-"+strconv.FormatInt(coin_id, 10)+".log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0777)
	if err != nil {
		panic("logfile is error")
	}
	Logrus.SetOutput(logfile)

	type coinPair struct{
		ID int64
		BuyCoin	string
		SaleCoin	string
	}

	//var coinPairList []coinPair
	rows, err := newDbCore.Raw("SELECT pair.F01 pair_id,buy.F02 buy_coin,sale.F02 sale_coin FROM dobi_trade.t6015 pair LEFT JOIN dobi_core.t6013 buy ON pair.F03 = buy.F01 LEFT JOIN dobi_core.t6013 sale ON pair.F04 = sale.F01 where pair.F03 = ?", coin_id).Rows()
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var pair_id int64
		var buy_coin string
		var sale_coin string
		rows.Scan(&pair_id, &buy_coin, &sale_coin)

		pair_name := strings.ToLower(buy_coin) +"_"+ strings.ToLower(sale_coin)
		Logrus.Info("同步K线的交易对", pair_name)

		coinPairSlice := [][]string{}
		coinPairSlice = append(coinPairSlice, []string{pair_name+"tradeline_1m", "60"})
		coinPairSlice = append(coinPairSlice, []string{pair_name+"tradeline_3m", "180"})
		coinPairSlice = append(coinPairSlice, []string{pair_name+"tradeline_5m", "300"})
		coinPairSlice = append(coinPairSlice, []string{pair_name+"tradeline_15m", "900"})
		coinPairSlice = append(coinPairSlice, []string{pair_name+"tradeline_30m", "1800"})

		for _, coin_pair := range coinPairSlice {

			tableName := "t6019_"+strconv.FormatInt(pair_id,10)+"_"+coin_pair[1]

			dropSql := "DROP TABLE IF EXISTS " + tableName + ";"
			if err := newDbTrade.Exec(dropSql).Error; err != nil {
				panic(err)
			}
			createSql := "CREATE TABLE "+ tableName + "(`F01` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '编号'," +
				"`F02` decimal(20,8) DEFAULT NULL COMMENT '开盘价',"+
				"`F03` decimal(20,8) DEFAULT NULL COMMENT '收盘价',"+
				"`F04` decimal(20,8) DEFAULT NULL COMMENT '最高价',"+
				"`F05` decimal(20,8) DEFAULT NULL COMMENT '最低价',"+
				"`F06` decimal(20,8) DEFAULT NULL COMMENT '收盘成交额',"+
				"`F07` bigint(20) DEFAULT NULL COMMENT 'K线时间',"+
				"`F08` bigint(20) DEFAULT NULL COMMENT '创建时间',"+
				"`F09` bigint(20) DEFAULT NULL COMMENT '成交量',"+
				"`F10` text COMMENT 'Json数据',"+
				"`F11` int(11) DEFAULT NULL COMMENT '时间类型'," +
				"`F12` bigint(20) DEFAULT NULL COMMENT '市场编号'," +
				"PRIMARY KEY (`F01`) USING BTREE, KEY `F02` (`F12`,`F11`) USING BTREE, KEY `F07` (`F07`) USING BTREE" +
				") ENGINE=MyISAM AUTO_INCREMENT=0 DEFAULT CHARSET=utf8mb4 ROW_FORMAT=DYNAMIC COMMENT='k线交易_" + strconv.FormatInt(pair_id,10) + "_" + coin_pair[1]+ "秒';"
			if err := newDbTrade.Exec(createSql).Error; err != nil {
				panic(err)
			}

			klineData, err := redisClien5.Get(coin_pair[0]).Result()
			//fmt.Println(coin_pair)
			if err == redis.Nil {
				Logrus.Info(coin_pair[0] + ":Error")
				continue
			}

			mapKlineData := make(map[string]interface{})

			json.Unmarshal([]byte(klineData), &mapKlineData)

			if mapKlineData["datas"] != nil {
				klineList := mapKlineData["datas"].(map[string]interface{})["data"].([]interface{})

				for _, k :=	range klineList {
					var klineStruct newModel.Kline

					klineStruct.KlineTime = int64(k.([]interface{})[0].(float64))/1000
					klineStruct.OpenPrice = k.([]interface{})[1].(float64)
					klineStruct.HightPrice = k.([]interface{})[2].(float64)
					klineStruct.LowPrice = k.([]interface{})[3].(float64)
					klineStruct.ClosePrice = k.([]interface{})[4].(float64)
					klineStruct.TotalNumber = k.([]interface{})[5].(float64)
					klineStruct.MarketID = pair_id
					klineStruct.TimeType,_ = strconv.ParseInt(coin_pair[1], 10, 64)

					newDbTrade.Table(tableName).Create(klineStruct)
				}
				Logrus.Info(mapKlineData)

			}

		}

		//t6019_market_n
		coinPairSliceN := [][]string{}
		coinPairSliceN = append(coinPairSliceN, []string{pair_name+"tradeline_1h", "3600"})	//1小时
		coinPairSliceN = append(coinPairSliceN, []string{pair_name+"tradeline_2h", "7200"})	//2小时
		coinPairSliceN = append(coinPairSliceN, []string{pair_name+"tradeline_4h", "14400"})	//4小时
		for _, coin_pair := range coinPairSliceN {

			tableName := "t6019_market_n"

			klineData, err := redisClien5.Get(coin_pair[0]).Result()
			//fmt.Println(coin_pair)
			if err == redis.Nil {
				Logrus.Info(coin_pair[0] + ":Error")
				continue
			}

			mapKlineData := make(map[string]interface{})

			json.Unmarshal([]byte(klineData), &mapKlineData)

			if mapKlineData["datas"] != nil {
				klineList := mapKlineData["datas"].(map[string]interface{})["data"].([]interface{})

				for _, k :=	range klineList {
					var klineStruct newModel.Kline

					klineStruct.KlineTime = int64(k.([]interface{})[0].(float64))/1000
					klineStruct.OpenPrice = k.([]interface{})[1].(float64)
					klineStruct.HightPrice = k.([]interface{})[2].(float64)
					klineStruct.LowPrice = k.([]interface{})[3].(float64)
					klineStruct.ClosePrice = k.([]interface{})[4].(float64)
					klineStruct.TotalNumber = k.([]interface{})[5].(float64)
					klineStruct.MarketID = pair_id
					klineStruct.TimeType,_ = strconv.ParseInt(coin_pair[1], 10, 64)

					newDbTrade.Table(tableName).Create(klineStruct)
				}
				Logrus.Info(mapKlineData)

			}
		}

		//t6019_market_m
		coinPairSliceM := [][]string{}
		coinPairSliceM = append(coinPairSliceM, []string{pair_name+"tradeline_6h", "21600"})	//6小时
		coinPairSliceM = append(coinPairSliceM, []string{pair_name+"tradeline_12h", "43200"})	//12小时
		coinPairSliceM = append(coinPairSliceM, []string{pair_name+"tradeline_1d", "86400"})	//1天
		for _, coin_pair := range coinPairSliceM {

			tableName := "t6019_market_m"
			klineData, err := redisClien5.Get(coin_pair[0]).Result()
			//fmt.Println(coin_pair)
			if err == redis.Nil {
				Logrus.Info(coin_pair[0] + ":Error")
				continue
			}

			mapKlineData := make(map[string]interface{})

			json.Unmarshal([]byte(klineData), &mapKlineData)

			if mapKlineData["datas"] != nil {
				klineList := mapKlineData["datas"].(map[string]interface{})["data"].([]interface{})

				for _, k :=	range klineList {
					var klineStruct newModel.Kline

					klineStruct.KlineTime = int64(k.([]interface{})[0].(float64))/1000
					klineStruct.OpenPrice = k.([]interface{})[1].(float64)
					klineStruct.HightPrice = k.([]interface{})[2].(float64)
					klineStruct.LowPrice = k.([]interface{})[3].(float64)
					klineStruct.ClosePrice = k.([]interface{})[4].(float64)
					klineStruct.TotalNumber = k.([]interface{})[5].(float64)
					klineStruct.MarketID = pair_id
					klineStruct.TimeType,_ = strconv.ParseInt(coin_pair[1], 10, 64)

					newDbTrade.Table(tableName).Create(klineStruct)
				}
				Logrus.Info(mapKlineData)

			}
		}


	}

	return

}
