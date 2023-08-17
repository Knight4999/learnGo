package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"os"
)

// SQLX ,相当于database/sql包的加强版，有更多扩展功能。
// 展示增删改查

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

// 查询单行数据
func queryRow() {
	sqlstr := "SELECT id, name, age FROM user WHERE id = ?"
	var u user
	err := db.Get(&u, sqlstr, 1) // -> db.queryrow.Scan
	if err != nil {
		fmt.Println("query failed,", err)
		return
	}
	fmt.Printf("id:%d name:%s age:%d\n", u.Id, u.Name, u.Age)
}

// 多行查询
func queryMutiRowDemo() {
	sqlstr := "SELECT id, name, age FROM user WHERE id > ?"
	var users []user
	err := db.Select(&users, sqlstr, 0)
	if err != nil {
		fmt.Println("query failed,", err)
		return
	}
	fmt.Printf("users:%#v\n", users)
}

// 增删改

func insertDemo() {
	sqlstr := "INSERT INTO user(name,age) VALUES (?,?)"
	rows, err := db.Exec(sqlstr, "王兵", 51)
	if err != nil {
		fmt.Println("insert failed:", err)
		return
	}
	theID, err := rows.LastInsertId()
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		return
	}
	fmt.Printf("insert success, the id is %d.\n", theID)
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Println("初始化数据库失败：", err)
	}
	queryRow()
	queryMutiRowDemo()
	insertDemo()
}
