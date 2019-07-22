package data_migration

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/jinzhu/gorm"
	myconfig "github/data_migration/config"
	"path/filepath"

	old_model "github/data_migration/models/old"
)

var oldDb *gorm.DB
var newDb *gorm.DB

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
	oldDb, _ = gorm.Open(oldDbConfig.DriverName, oldDbConfig.DataSourceName)
	defer oldDb.Close()
	oldDb.LogMode(true)

	newDbConfig := config.NewDobiDatabase
	newDb, _ = gorm.Open(newDbConfig.DriverName, newDbConfig.DataSourceName)
	defer newDb.Close()
	newDb.LogMode(true)


}

func User_data() {
	var address_model old_model.Address
	oldDb.Select("address,uid,publicKey").Where(old_model.Address{
		Address:"NVcHJEzm4pPnzndWM47nFPkHmociER54xs",
	}).Limit(1).Find(&address_model)

	fmt.Println(address_model.Address)
	fmt.Println(address_model.Uid)
	fmt.Println(address_model.PublicKey)

	type Trade struct {
		F01 int64  `gorm:"column:F01"`
		F05 string	 `gorm:"column:F05"`
	}
	var trade []Trade
	newDb.Raw("select * from t5010 limit 10").Scan(&trade)
	fmt.Println(trade)

}
