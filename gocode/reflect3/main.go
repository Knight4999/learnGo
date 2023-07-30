package main

import (
	"fmt"
	"reflect"
)

// 结构体反射--方法

type Student struct {
	Name string `json:"user_name" ini:"s_user"`
	Age  int    `json:"user_age" ini:"s_age"`
}

// Sleep 定义两个方法
func (s Student) Sleep() string {
	str := "好好学习，天天睡觉"
	fmt.Println(str)
	return str
}

func (s Student) Study() string {
	str := "好好学习，天天向上"
	fmt.Println(str)
	return str
}

func printMethod(x interface{}) {
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)

	fmt.Println(t.NumField())
	for i := 0; i < v.NumMethod(); i++ { //NumMethod() 获取结构体方法的数量
		methodType := v.Method(i).Type() //利用索引下标方式
		fmt.Printf("method name:%v\n", t.Method(i).Name)
		fmt.Printf("method type:%v\n", methodType)
		//通过反射调用方法，并传递参数，参数必须是[]reflect.Value类型
		var args = []reflect.Value{}
		v.Method(i).Call(args)
	}
}
func main() {
	stu := Student{"Jerry", 40}
	printMethod(stu)
}
