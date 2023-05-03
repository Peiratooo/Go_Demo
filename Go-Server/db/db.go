package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

type User struct {
	uid      int
	name     string
	email    string
	password string
}

func InitDb() *xorm.Engine {
	var (
		userName  = "root"
		password  = "Aa2461655"
		ipAddress = "127.0.0.1"
		port      = 3306
		dbName    = "gostudy"
		charset   = "utf8mb4"
	)
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", userName, password, ipAddress, port, dbName, charset)
	engine, err := xorm.NewEngine("mysql", dataSourceName)
	if err != nil {
		fmt.Print(err)
	}
	return engine
}

var Engine = InitDb()
