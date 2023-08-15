package main

// go 语言操作Mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB //表示连接数据库的 实例对象

func initDB() (err error) {
	//init函数，初始化连接数据库操作。
	dsn := "root:Wzh123123.@tcp(127.0.0.1:3306)/dbtest"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("连接数据库失败!")
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}
	db.SetMaxOpenConns(10) //设置数据库最大连接数
	db.SetMaxIdleConns(10) //设置连接池中的最大闲置连接数。
	return nil
}

type user struct {
	id   int
	name string
	age  int
}

// 单行查询
func queryRowDemo() {
	sqlstr := "SELECT id,name,age FROM user WHERE id = ?"
	var u user
	// 非常重要：确保QueryRow之后调用Scan方法，否则持有的数据库链接不会被释放
	err := db.QueryRow(sqlstr, 1).Scan(&u.id, &u.name, &u.age)
	if err != nil {
		fmt.Printf("scan failed, err:%v\n", err)
		return
	}
	fmt.Printf("id:%d name:%s age:%d\n", u.id, u.name, u.age)
}

// 多行查询
func queryDemo() {
	sqlstr := "SELECT id,name,age FROM user WHERE id >?"
	rows, err := db.Query(sqlstr, 2)
	if err != nil {
		fmt.Println("解析错误", err)
	}
	// 非常重要：关闭rows释放持有的数据库链接
	defer rows.Close()
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

// 插入、更新、删除数据
func execInsertDemo() {
	sqlstr := "INSERT INTO user(name,age) VALUES (?,?),(?,?),(?,?)"
	result, err := db.Exec(sqlstr, "李白", 43, "杜甫", 23, "王维", 30)
	if err != nil {
		fmt.Println("插入数据失败:", err)
	}
	n, _ := result.LastInsertId()
	fmt.Printf("insert success, the id is %d.\n", n)
}
func execDeleteDemo() {
	sqlstr := "DELETE FROM user WHERE id BETWEEN ? AND ?"
	result, err := db.Exec(sqlstr, 12, 14)
	if err != nil {
		fmt.Println("删除数据失败:", err)
	}
	rows, err := result.RowsAffected()
	fmt.Printf("删除了%d 条数据\n", rows)
}

func execUpdataDemo() {
	sqlstr := "UPDATE user SET name=? WHERE id =?"
	result, err := db.Exec(sqlstr, "black", 6)
	if err != nil {
		fmt.Println("修改数据失败:", err)
	}
	rows, err := result.RowsAffected()
	fmt.Printf("修改了%d 条数据\n", rows)
}
func main() {
	err := initDB()
	if err != nil {
		fmt.Println("初始化数据库连接失败：", err)
		return
	}
	defer db.Close()
	queryRowDemo()
	queryDemo()
	//execInsertDemo()
	//execDeleteDemo()
	execUpdataDemo()
	queryDemo()

}
