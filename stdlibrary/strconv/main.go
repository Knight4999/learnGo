package main

import (
	"fmt"
	"strconv"
)

// Go语言中strconv包实现了基本数据类型和其字符串表示的相互转换

func main() {
	s1 := "100"
	i1, _ := strconv.Atoi(s1) //Atoi 函数，将字符串转化为基本数据类型
	fmt.Printf("%T\n", i1)
	fmt.Println(i1)

	i2 := 10086
	s2 := strconv.Itoa(i2) //Itoa 函数，将基本数据类型转化为字符串
	fmt.Printf("%T\n", s2)
	fmt.Println(s2)

	//Paser 系列函数 用于将字符串转化为给定的值
	b, err := strconv.ParseBool("t") //它接受1、0、t、f、T、F、true、false、True、False、TRUE、FALSE
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%T,%v\n", b, b)

	//Format 系列函数，将给定类型的值解析成字符串

	t := strconv.FormatBool(true)
	fmt.Printf("%T,%v\n", t, t)

	//append 系列函数
	buf := make([]byte, 10)
	buf = strconv.AppendBool(buf, true)
	fmt.Println(string(buf))
	//quote 系列函数
	s := strconv.Quote(`"Fran & Freddie's Diner	☺"`)
	fmt.Println(s)
}
