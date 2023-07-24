package main

import (
	"fmt"
)

// 结构体匿名字段

type Person struct {
	string //匿名字段
	int8
	//string 匿名字段内内容，字段类型必须唯一
}

func main() {
	f1 := Person{
		"阿萨德",
		20,
	}
	fmt.Println(f1)
	fmt.Println(f1.string, f1.int8) //通过调用字段类型调用值

}

type Person1 struct {
	Name     string
	Gender   string
	Age      int
	Province string //省份
	City     string
}

func structA() {

}
