package dbs

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

var MysqlDb *sql.DB
var MysqlDbErr error

const (
	USER_NAME = "root"
	PASS_WORD = "root"
	HOST      = "localhost"
	PORT      = "3306"
	DATABASE  = "mysql"
	CHARSET   = "utf8"
)

func AddDevice(deviceName string) {
	Init()
	stmt, _ := MysqlDb.Prepare("insert devices set devicename=?")
	_, _ = stmt.Exec(deviceName)
	_ = MysqlDb.Close()
}

func DeleteDevice(deviceName string) {
	Init()
	stmt, err := MysqlDb.Prepare("delete from devices where devicename=?")
	if err != nil {
		log.Println(err)
	}
	_, _ = stmt.Exec(deviceName)

	_ = MysqlDb.Close()
}

func FindAllDevices() []string {
	Init()
	rows, err := MysqlDb.Query("select devicename from devices")

	if err != nil {
		log.Println(err)
	}

	fmt.Println(rows)

	var result []string
	var devicename string

	for rows.Next() {
		_ = rows.Scan(&devicename)

		result = append(result, devicename)

		//fmt.Printf("username: %s\n", dname)
	}

	_ = MysqlDb.Close()
	return result
}

func Init() {
	dbDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", USER_NAME, PASS_WORD, HOST, PORT, DATABASE, CHARSET)

	// 打开连接失败
	MysqlDb, MysqlDbErr = sql.Open("mysql", dbDSN)
	//defer MysqlDb.Close();
	if MysqlDbErr != nil {
		log.Println("dbDSN: " + dbDSN)
		panic("数据源配置不正确: " + MysqlDbErr.Error())
	}

	// 最大连接数
	MysqlDb.SetMaxOpenConns(100)
	// 闲置连接数
	MysqlDb.SetMaxIdleConns(20)
	// 最大连接周期
	MysqlDb.SetConnMaxLifetime(100 * time.Second)

	if MysqlDbErr = MysqlDb.Ping(); nil != MysqlDbErr {
		panic("数据库链接失败: " + MysqlDbErr.Error())
	}
}
