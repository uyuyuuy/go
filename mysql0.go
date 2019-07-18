package github

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
	"reflect"
)

const (
	LevelDefault int = iota
	LevelReadUncommitted
	LevelReadCommitted
	LevelWriteCommitted
	LevelRepeatableRead
	LevelSnapshot
)


const DATABASE_TYPE  = "mysql"
const DNS  =  "dobi_tra:123###00@tcp(172.16.8.222:3306)/morecoin"


func main() {

	fmt.Println(LevelDefault)
	fmt.Println(LevelReadUncommitted)
	fmt.Println(LevelReadCommitted)
	fmt.Println(LevelWriteCommitted)
	fmt.Println(LevelRepeatableRead)
	fmt.Println(LevelSnapshot)


	var new_mysql mysql_tool
	open, err := new_mysql.connection(DATABASE_TYPE, DNS)
	fmt.Println(sql.Drivers())

	if open == false {
		log.Fatal(err)
		fmt.Println("error")
		os.Exit(11)
	}

	rows, err := new_mysql.query("select uid,address from address where address != '' limit 1")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(rows.Columns())

	for rows.Next() {
		var uid int
		var address string

		err = rows.Scan(&uid, &address)
		fmt.Println(uid)
		fmt.Println(address)
	}
	fmt.Println(getTypes(rows))
	fmt.Println(rows)

	fmt.Println("over")
	os.Exit(0)
}


type mysql_tool struct {
	db *sql.DB

}

//数据库连接
func (m *mysql_tool) connection(d_type string, dns string) (open bool, err error){
	open = true
	con, err := sql.Open(d_type, dns)
	if err != nil {
		log.Fatal(err)
		open = false
	}
	m.db = con
	return
}

//数据库查询
func (m *mysql_tool) query(sql string) (rows *sql.Rows, err error){
	rows, err = m.db.Query(sql)
	if err != nil {
		log.Fatal(err)
	}
	return
}

//数据库更新（增删改）
func (m *mysql_tool) exec(sql string) (res bool, err error) {
	stmt, err := m.db.Prepare(sql)
	stmt.Exec()
	return
}

//数据库断开
func (m *mysql_tool) close() error{
	err := m.db.Close()
	return err
}





func getTypes(v interface{}) string {
	return fmt.Sprintf("%T", v)
	return reflect.TypeOf(v).String()
}




