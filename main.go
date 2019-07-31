package main

import (
	"fmt"
	//"github.com/BurntSushi/toml"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github/data_migration"
	"github/data_migration/config"

	//myconfig "github/data_migration/config"
	//"github/data_migration/models/old"

	//"github/data_migration/models/old"

	//models_a "github/data_migration/models/new" //别名在签名
	//"github/data_migration/models/old"
	"os"
	//"path/filepath"
	"reflect"
)

//var config myconfig.Config

func init() {
	//获取绝对路径
	//filePath, err := filepath.Abs("./data_migration/config/config.toml")
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(filePath)

	//映射数据库连接配置
	//if _, err := toml.DecodeFile(filePath, &config); err != nil {
	//	panic(err)
	//}

	//打印
	//fmt.Printf("%+v\n", config)		//可以用来打印结构体
	//fmt.Printf("%+v\n", config.Database)
	//fmt.Printf("%+v\n", config.Common)
	//
	//s := fmt.Sprintf("a %s", "string")
	//fmt.Println(s)	//a string
	//fmt.Fprintf(os.Stderr, "an %s\n", "error")

}



func main() {

	data_migration.UserData()

	//data_user()
	os.Exit(111)

	//连接数据库
	db, err := gorm.Open(config.MysqlConfig{}.DriverName, config.MysqlConfig.DataSourceName)
	defer db.Close()
	if err != nil {
		panic("failed to connect database")
	}

	//开启debug模式
	db.LogMode(true)
	//db.SetLogger(gorm.Logger{revel.TRACE})
	//log.New(os.Stdout, "\r\n", 0)

	//更新数据库
	if config.Common.InitDatabase == 1 {
		init_databases(db)
	}



	//fmt.Println(mproduct)
	var pp models_a.Product
	pp.ID = 11
	//models_a.Product{Code:"02"}
	//up := db.Debug().Model(pp).Update("code", "03").RowsAffected
	del := db.Debug().Delete(pp).RowsAffected
	//fmt.Println(up)
	fmt.Println(del)

	os.Exit(1)

	var mproduct []models_a.Product
	where := models_a.Product{Code:"01"}

	fields := []string{"code", "price"}
	db.Select(fields).Where(&where).Find(&mproduct)

	//fmt.Println(product)
	for i, p := range mproduct{
		fmt.Println(i,p)
		fmt.Println(p.Code)
		fmt.Println(p.Price)
		fmt.Println(p.ID)
	}



	os.Exit(0)

	//var inter interface{} = []interface{}
	//
	//values := reflect.ValueOf(inter)
	//
	//fmt.Println(values)
	//
	//row := db.Raw("select * from product where code = ?", "01").Row() // (*sql.Rows, error)
	//
	//row.Scan(&inter)
	//fmt.Println(inter)
	//os.Exit(0)




	//删除
	del_product := models_a.Product{ID:2}
	db.Delete(&del_product)
	//db.Unscoped().Where("code = ?", "05").Delete(models_a.Product{})   //加Unscoped() 就是永久删除

	// 使用User结构体创建名为`deleted_users`的表
	//db.Table("deleted_product").CreateTable(&models_a.Product{})


	//添加
	add_product := models_a.Product{Code: "01", Price:1000}
	db.Create(&add_product)

	//接收结果
	//var product []models_a.Product
	//where := models_a.Product{Code:"01"}
	//db.Where(&where).Find(&product)
	//fmt.Println(product) // find product with id 1





	//不使用model
	/*
	rows, err := db.Raw("select * from product where code = ?", "01").Rows() // (*sql.Rows, error)
	defer rows.Close()

	type tmp_row struct {
		id int
		code string
		price int
	}
	result := []tmp_row{}
	for rows.Next() {
		var id int
		var code string
		var price int
		rows.Scan(&id, &code, &price)

		var tmp_d tmp_row
		tmp_d.id = id
		tmp_d.code = code
		tmp_d.price = price

		result = append(result, tmp_d)

	}
	*/


	//row := db.Raw("select * from product where code = ?", "01").Row() // (*sql.Rows, error)


	//user_data()

}


func data_user (){


	os.Exit(0)






	os.Exit(110)


}



//检查数据库表结构，只会添加缺少的字段，不会删除/更改当前数据，如果表不存在则创建
func init_databases(db *gorm.DB) {
	// 关闭复数表名，如果设置为true，`User`表的表名就会是`user`，而不是`users`
	db.SingularTable(true)

	//db.AutoMigrate(&models_a.Product{})
	//db.AutoMigrate(&old.User{})
}


func getType(v interface{}) string {
	return fmt.Sprintf("%T", v)
	return reflect.TypeOf(v).String()
}




/*
单表数据查询可以用model，多表查询需要自定义返回数据结构类型
type Results struct {

}
var results Results
db.Table("users").Select("users.name, emails.email").Joins("left join emails on emails.user_id = users.id").Scan(&results)

执行原生SQL / Raw
db.Exec("DROP TABLE users;")
db.Raw("SELECT name, age FROM users WHERE name = ?", 3).Scan(&result)


Model / Table
db.Model(&User{}).Where("name = ?", "jinzhu").Count(&count)
db.Table("users").Where("name = ?", "jinzhu").Count(&count)


//获取改变的数量，用于Update,Delete
RowsAffected


//更新
UpdateColumn/UpdateColumn
Update/Updates


 */