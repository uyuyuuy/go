package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	//uuid "github.com/satori/go.uuid"
	"os"

	//uuid "github.com/satori/go.uuid"
	"math/rand"
	"time"
)

var DDDDD *gorm.DB

const driverName = "mysql"
const dataSourceName = "dobi_tra:123###00@tcp(172.16.8.222:3306)/morecoin?charset=utf8&parseTime=True&loc=Local"


//定义表字段
type FillTable struct {
	UID int64	`gorm:"column:uid"`
	Wallet string	`gorm:"column:wallet"`
	Tid string		`gorm:"column:txid"`
	Number float64	`gorm:"column:number"`
	PlatformFee	float64	`gorm:"column:platform_fee"`
	NumberReal	float64	`gorm:"column:number_real"`
	OptType string	`gorm:"column:opt_type"`
	Status string	`gorm:"column:status"`
	Created int64	`gorm:"column:created"`
	OrderSn string	`gorm:"column:order_sn"`
	IsOut int	`gorm:"column:is_out"`
}

//定义表名
func (FillTable) TableName() string {
	return "exchange_new"
}




func init()  {
	var err error
	DDDDD, err := gorm.Open(driverName, dataSourceName)
	if err != nil {
		panic(err)
	}
	DDDDD.LogMode(true)

	rand.Seed(time.Now().UnixNano())
}



func main() {

	defer DDDDD.Close()
	//defer func() {
	//	if err := recover(); err != nil {
	//		fmt.Println(err)
	//	}
	//	fmt.Println("ssssssss")
	//}()

	//定义插入记录数量
	//var count = 10


	statusSlice := []string{"等待","确认中","成功","已取消","冻结中"}
	optTypeSlice := []string{"in","out"}
	var FfillTable FillTable
	//fillTable = FillTable{}
	//for i := 0; i < count; i++  {

		//orderSn, _ := uuid.NewV4()

	FfillTable.Wallet = RandStr(2,32)
	FfillTable.Tid = RandStr(2,32)
	FfillTable.Status = statusSlice[rand.Intn(4)]
		//fillTable.OrderSn = orderSn.String()
	FfillTable.Number = 1.025666
	FfillTable.OptType = optTypeSlice[rand.Intn(1)]
	FfillTable.NumberReal = 1.0000
	FfillTable.Created = 1565073155
	FfillTable.IsOut = 1

		fmt.Println(FfillTable)
	DDDDD.Create(&FfillTable)
		//if err := DB.Create(&fillTable).Error; err != nil {
		//	panic(err)
		//}
		//fillTable = FillTable{}
		time.Sleep(1 * time.Second)
	//}
	os.Exit(110)

}



func RandStr(t int, len int) string {
	str := ""
	var randData map[string]interface{}
	randData = make( map[string]interface{})
	number := "0123456789"
	letterA := "abcdefghijklmnopqrstuvwxyz"
	letterB := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	switch t {
	case 0:
		randData["randStr"] = number
		randData["randStrLeng"] = 10
	case 1:
		randData["randStr"] = letterA + letterB
		randData["randStrLeng"] = 52
	case 2:
		randData["randStr"] = number + letterA + letterB
		randData["randStrLeng"] = 62
	default:
		randData["randStr"] = number
		randData["randStrLeng"] = 10
	}


	for i := 0; i < len; i++ {
		n := rand.Intn(randData["randStrLeng"].(int) - 1)
		str += randData["randStr"].(string)[n:n+1]
	}

	return	str
}


