package dao

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var DB *sql.DB

func InitDB() error {
	fmt.Println("Mysql init.")
	var err error
	DB, err = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/eating_snake") //eating_snake库
	if err != nil {
		log.Panicln("err:", err.Error())
		return err
	}
	err = DB.Ping()
	if err != nil {
		log.Panicln("err:", err.Error())
		return err
	}
	DB.SetMaxOpenConns(10) //最大链接数
	DB.SetMaxIdleConns(10) //最大空闲链接数
	return nil
}
