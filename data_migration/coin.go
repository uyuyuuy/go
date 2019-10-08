package data_migration

import (
	"errors"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"reflect"
	"strings"
	"time"

	newModel "github/data_migration/models/new"
	oldModel "github/data_migration/models/old"
)

var Market = make(map[string]int)

func init() {
	log.Print("coin init")

}

func doMarket() (err error) {

	InMarket := make(map[string]string)
	InMarket["BTC"] = "Bitcoin"
	InMarket["ETH"] = "Ethereum"
	InMarket["USDT"] = "USDT"
	InMarket["DOB"] = "DOB"

	valueStrings := make([]string, 0, len(InMarket))
	valueArgs := make([]interface{}, 0, 2 * len(InMarket))

	CoinID := make(map[string]int)
	rows, err := newDbCore.Model(&newModel.Coin{}).Select("F01,F02").Where("F02 IN (?)", []string{"BTC","ETH","USDT","DOB"}).Rows()
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var F02 string
		var F01 int
		_ = rows.Scan(&F01, &F02)
		CoinID[F02] = F01
	}

	Logrus.Info("交易区：",CoinID)

	for Name, AssetName := range InMarket {
		valueStrings = append(valueStrings, "(?,?,?)")
		valueArgs = append(valueArgs, CoinID[Name])
		valueArgs = append(valueArgs, Name)
		valueArgs = append(valueArgs, AssetName)
	}

	sqlStr := fmt.Sprintf("INSERT INTO t6014(F02,F03,F04) VALUES %s", strings.Join(valueStrings, ","))
	if err := newDbTrade.Exec(sqlStr,valueArgs...).Error; err != nil{
		Logrus.Warn(err, "交易区同步失败")
		panic(err)
	} else {
		Logrus.Info("交易区同步成功")
	}

	return
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

	//1.先同步币种信息
	var coin []oldModel.Coin
	selectStr := "name,asset_name,minout,maxout,out_limit,rate_out,out_status,in_status,logo,number_float,coin_transfer,`describe`"
	oldDb.Select(selectStr).Find(&coin)

	for _,c := range coin {
		var newCoin newModel.Coin
		newCoin.CreateTime = "2019-08-22 00:00:00"
		newCoin.Name = strings.ToUpper(c.Name)
		newCoin.AssetName = c.AssetName
		newCoin.MinOut = c.MinOut
		newCoin.MaxOut = c.MaxOut
		newCoin.OutLimit = c.OutLimit
		newCoin.OutRate = c.OutRate
		if c.OutStatus == 0 {
			newCoin.OutStatus = "S"
		} else {
			newCoin.OutStatus = "F"
		}
		if c.InStatus == 0 {
			newCoin.InStatus = "S"
		} else {
			newCoin.InStatus = "F"
		}
		newCoin.Logo = c.Logo
		newCoin.OutFloatNumberLimit = c.NumberFloat
		if c.CoinTransfer == 0 {
			newCoin.MoveStatus = "S"
		} else {
			newCoin.MoveStatus = "F"
		}
		newCoin.Description = c.Describe

		if e := newDbCore.Create(&newCoin).Error; e != nil {
			Logrus.Warn(c.Name, "：失败", e)
		}

	}

	time.Sleep(5 * time.Second)
	//2.再插入交易区
	doMarket()

	time.Sleep(5 * time.Second)
	//3.最后同步交易对
	var coinpair []oldModel.CoinPair
	selectStr = "coin_pair.coin_from,coin_pair.coin_to,coin_pair.rate,coin_pair.rate_buy,coin_pair.min_trade,coin_pair.max_trade,coin_pair.status,coin_pair.price_float,coin_pair.number_float,coin_pair.order_by"
	oldDb.Select(selectStr).Find(&coinpair)

	for _,c := range coinpair {

		var newCoinPari newModel.CoinPair
		var F01 int
		row := newDbTrade.Model(&newModel.Market{}).Select("F01").Where("F03 = ?", strings.ToUpper(c.CoinTo)).Row()

		row.Scan(&F01)

		rows, _ := newDbCore.Model(&newModel.Coin{}).Select("F01,F02").Where("F02 IN (?)", []string{strings.ToUpper(c.CoinFrom), strings.ToUpper(c.CoinTo)}).Rows()
		for rows.Next() {
			var coinName string
			var coinID int
			_ = rows.Scan(&coinID, &coinName)
			switch coinName {
			case strings.ToUpper(c.CoinFrom):
				newCoinPari.SaleCoinID = coinID
			case strings.ToUpper(c.CoinTo):
				newCoinPari.BuyCoinID = coinID
			default:
				Logrus.Warn("CoinFrom,CoinTo对不上")
			}
		}

		newCoinPari.MarketID = F01
		newCoinPari.SaleRate = c.RateSale
		newCoinPari.BuyRate = c.RateBuy
		newCoinPari.BuyMin = c.MinTrade
		newCoinPari.SaleMin = c.MinTrade
		newCoinPari.BuyMax = c.MaxTrade
		newCoinPari.SaleMax = c.MaxTrade
		if c.Status == 0 {
			newCoinPari.IsDisplay = "S"
			newCoinPari.IsOpen = "S"
		} else {
			newCoinPari.IsDisplay = "F"
			newCoinPari.IsOpen = "F"
		}
		newCoinPari.PricePrecision = c.PairPriceFloat
		newCoinPari.NumberPrecision = c.PairNumberFloat
		newCoinPari.Order = c.OrderBy
		newCoinPari.CreateTime = "2019-08-22 00:00:00"

		if e := newDbTrade.Create(&newCoinPari).Error; e != nil {
			Logrus.Warn("交易对失败", e)
		}

	}
	return 
}
