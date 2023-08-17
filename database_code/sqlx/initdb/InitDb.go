package initdb

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"os"
)

var DB *sqlx.DB

func InitDB() (err error) {
	dsn := "root:Wzh123123.@tcp(127.0.0.1:3306)/dbtest"
	DB, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Println("connect database failed：", err)
		os.Exit(1)
	}
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(10)
	return
}
