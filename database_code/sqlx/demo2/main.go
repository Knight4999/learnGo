package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"os"
)

// NamedExec 和 NamedQuery

type user struct {
	Id   int
	Name string
	Age  int
}

var db *sqlx.DB

func initDB() (err error) {
	dsn := "root:Wzh123123.@tcp(127.0.0.1:3306)/dbtest"
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Println("connect database failed:", err)
		os.Exit(1)
	}
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)
	return
}

// DB.NamedExec方法用来绑定SQL语句与结构体或map中的同名字段。
func insertUserDemo() (err error) {
	sqlstr := "INSERT INTO user(name,age) VALUES (:name,:age)"
	_, err = db.Exec(sqlstr, map[string]interface{}{
		"name": "李火旺",
		"age":  18,
	})
	return
}

// DB.NamedQuery方法用来绑定SQL语句与结构体或map中的同名字段。支持查询
func namedQuery() {
	sqlStr := "SELECT * FROM user WHERE name=:name"
	// 使用map做命名查询
	rows, err := db.NamedQuery(sqlStr, map[string]interface{}{"name": "七米"})
	if err != nil {
		fmt.Printf("db.NamedQuery failed, err:%v\n", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var u user
		err := rows.StructScan(&u)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			continue
		}
		fmt.Printf("user:%#v\n", u)
	}

	u := user{
		Name: "七米",
	}
	// 使用结构体命名查询，根据结构体字段的 db tag进行映射
	rows, err = db.NamedQuery(sqlStr, u)
	if err != nil {
		fmt.Printf("db.NamedQuery failed, err:%v\n", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var u user
		err := rows.StructScan(&u)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			continue
		}
		fmt.Printf("user:%#v\n", u)
	}
}
func main() {

}
