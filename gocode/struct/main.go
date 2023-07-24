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
	structA()
}

type Person1 struct {
	Name    string
	Gender  string
	Age     int
	Address //地址结构体
	Email   //结构体字段冲突，需要指明具体结构体
}
type Email struct {
	Addr       string
	UpdateTime string
}

// 地址结构体
type Address struct {
	Province   string //省份
	City       string
	UpdateTime string
}

func structA() {
	p1 := Person1{
		"李火旺",
		"男",
		18,
		Address{
			"北京",
			"朝阳区",
			"2021-01-01 10:10:00",
		},
		Email{
			"1001284@sql.com",
			"2022-10-23 12:23:11",
		},
	}
	fmt.Printf("%#v\n", p1)
	fmt.Println(p1.Name, p1.Gender, p1.Age, p1.Province,
		p1.City, p1.Address.UpdateTime, p1.Addr, p1.Email.UpdateTime)
}
