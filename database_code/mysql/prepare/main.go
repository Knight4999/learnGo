package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

// Mysql 预处理。
/*
1.把SQL语句分成两部分，命令部分与数据部分。
2.先把命令部分发送给MySQL服务端，MySQL服务端进行SQL预处理。
3.然后把数据部分发送给MySQL服务端，MySQL服务端对SQL语句进行占位符替换。
4.MySQL服务端执行完整的SQL语句并将结果返回给客户端。
*/
type user struct {
	id   int
	name string
	age  int
}

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", "root:Wzh123123.@tcp(127.0.0.1:3306)/dbtest")
	if err != nil {
		fmt.Printf("connect database failed,err:%v \n", err)
		os.Exit(1)
	}
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
}

// 预编译查询
func prepareQueryDemo() {
	sqlstr := "SELECT * FROM user WHERE id > ?"
	stmt, err := db.Prepare(sqlstr)
	if err != nil {
		fmt.Printf("Prepare sql failed,err: %v\n", err)
		return
	}
	defer stmt.Close()
	rows, err := stmt.Query(2)
	if err != nil {
		fmt.Printf("query failed,err:%v\n", err)
		return
	}
	defer rows.Close()
	// 循环读取结果集中的数据
	for rows.Next() {
		var u user
		err := rows.Scan(&u.id, &u.name, &u.age)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		fmt.Printf("id:%d name:%s age:%d\n", u.id, u.name, u.age)
	}
}

// 预编译增删改
func perpareInsertDemo() {
	sqlstr := "INSERT INTO user(name,age) VALUES (?,?)"
	stmt, err := db.Prepare(sqlstr)
	if err != nil {
		fmt.Println("prepare failed：", err)
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec("湖北", 23)
	if err != nil {
		fmt.Println("Insert failed：", err)
		return
	}
	fmt.Println("Insert Success!")
}
func main() {
	prepareQueryDemo()
	perpareInsertDemo()
}
