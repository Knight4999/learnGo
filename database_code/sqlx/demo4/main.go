package main

import (
	"database/sql/driver"
	"fmt"
	"github.com/jmoiron/sqlx"
	"learnGo/database_code/sqlx/initdb"
	"os"
	"strings"
)

// sqlx.In  批量插入数据

// User 创建与数据库字段对应的结构体
type User struct {
	ID   int    `db:"id"`
	Name string `db:"name"` //使用tag标签与数据库列一致
	Age  int    `db:"age"`
}

// Value 要想使用sqlx.IN,就要让结构体先实现Value接口
func (u User) Value() (driver.Value, error) {
	return []interface{}{u.Name, u.Age}, nil
}

// BatchInsertUsers 自行构造批量插入的语句
func BatchInsertUsers(users []*User) error {
	//存放（？，？）的slice
	valueStrings := make([]string, 0, len(users))
	//存放values的slice
	valueArgs := make([]interface{}, 0, len(users))
	//变量users，准备相关数据
	for _, u := range users {
		// 此处占位符要与插入值的个数对应
		valueStrings = append(valueStrings, "(?,?)")
		valueArgs = append(valueArgs, u.Name)
		valueArgs = append(valueArgs, u.Age)
	}
	stmt := fmt.Sprintf("INSERT INTO user(name,age) VALUES %s", strings.Join(valueStrings, ","))
	_, err := initdb.DB.Exec(stmt, valueArgs...)
	return err
}

// BatchInsertUsers2 使用sqlx.In帮我们拼接语句和参数, 注意传入的参数是[]interface{}
func BatchInsertUsers2(users []interface{}) error {
	// 如果arg实现了 driver.Valuer, sqlx.In 会通过调用 Value()来展开它
	query, args, _ := sqlx.In("INSERT INTO user (name, age) VALUES (?), (?), (?)", users...)
	fmt.Println("query sql string: ", query) // 查看生成的querystring
	fmt.Println("args: ", args)              // 查看生成的args
	// Exec执行查询而不返回任何行。参数用于查询中的任何占位符参数。
	result, err := initdb.DB.Exec(query, args...)
	var rows_affected int64
	rows_affected, err = result.RowsAffected() // 返回受更新、插入或删除影响的行数。并非每个数据库或数据库驱动程序都支持此功能。
	if err != nil {
		fmt.Printf("返回受更新、插入或删除影响的行数 failed, err: %v\n", err)
		return err
	}
	fmt.Println("受更新、插入或删除影响的行数: ", rows_affected)
	return nil
}

// BatchInsertUsers3 使用NamedExec来批量插入数据
func BatchInsertUsers3(users []*User) error {
	_, err := initdb.DB.NamedExec("INSERT INTO user (name,age) VALUES (:name,:age)", users)
	return err
}

// QueryByIDs 使用sqlx.In来查询 ID 在指定集合中的数据
func QueryByIDs(ids []int) (users []User, err error) {
	// In 展开args中的切片值，返回修改后的查询字符串和一个可以由数据库执行的新的arg列表。
	// “查询”应该使用“?”“bindVar。返回值使用' ?“bindVar
	query, args, _ := sqlx.In("SELECT name,age FROM user WHERE id IN(?)", ids)
	// Rebind 将查询从 QUESTION 转换为DB驱动程序的 bindvar 类型。
	query = initdb.DB.Rebind(query)
	err = initdb.DB.Select(&users, query, args...)
	if err != nil {
		return nil, err
	}
	return users, err
}

// QueryAndOrderByIDs 根据 ID 在指定集合中和指定顺序查询
func QueryAndOrderByIDs(ids []int) (users []User, err error) {
	// 创建一个字符串切片，大小为ids的长度
	strIDs := make([]string, 0, len(ids))
	// 将ids转换为字符串类型
	for _, id := range ids {
		// Sprintf根据格式说明符进行格式化，并返回结果字符串。
		strIDs = append(strIDs, fmt.Sprintf("%d", id))
	}
	// In展开args中的切片值，返回修改后的查询字符串和一个可以由数据库执行的新的arg列表。“查询”应该使用“?”“bindVar。返回值使用' ?“bindVar。
	query, args, err := sqlx.In("SELECT name, age FROM user WHERE id IN (?) ORDER BY FIND_IN_SET(id, ?)",
		ids, strings.Join(strIDs, ","))
	if err != nil {
		return
	}
	// Rebind 将查询从QUESTION转换为DB驱动程序的bindvar类型。
	query = initdb.DB.Rebind(query)
	// 执行查询 Select 使用此数据库。任何占位符参数都将被提供的参数替换。
	err = initdb.DB.Select(&users, query, args...)
	return
}

func main() {
	err := initdb.InitDB()
	if err != nil {
		fmt.Printf("init database failed,err:%v\n", err)
		os.Exit(1)
	}
	fmt.Println("init DB succeeded")
	// 批量插入
	/*u1 := User{Name: "李白", Age: 16}
	u2 := User{Name: "杜甫", Age: 42}
	u3 := User{Name: "王维", Age: 29}
	users := []interface{}{u1, u2, u3}

	err = BatchInsertUsers2(users)
	if err != nil {
		fmt.Printf("插入数据失败,err:%v\n", err)
	}*/

	/*u1 := User{Name: "刘备", Age: 54}
	u2 := User{Name: "张飞", Age: 32}
	u3 := User{Name: "关羽", Age: 48}
	users := []*User{&u1, &u2, &u3}

	err = BatchInsertUsers3(users)
	if err != nil {
		fmt.Println("插入数据失败！", err)
	}*/
	ids := []int{1, 2, 3, 4, 5, 6}
	users, err := QueryByIDs(ids)
	if err != nil {
		fmt.Println("查询失败：", err)
		return
	}
	for _, user := range users {
		fmt.Printf("user: %#v\n", user)
	}
	fmt.Println("************************************************")
	// FIND_IN_SET
	users, err = QueryAndOrderByIDs([]int{1, 15, 6, 2})
	if err != nil {
		fmt.Printf("query error: %v\n", err)
		return
	}
	fmt.Printf("query successful result users %v\n", users)
	for _, user := range users {
		fmt.Printf("user: %#v\n", user)
	}
}
