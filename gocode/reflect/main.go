package main

import (
	"fmt"
	"reflect"
)

// 反射

func reflectType(any interface{}) {
	//不知道别人调用这个函数的时候会传什么类型的变量
	//1. 方式1：类型断言
	//2. 方式2：反射
	obj := reflect.TypeOf(any)
	fmt.Println(obj, obj.Name(), obj.Kind()) //obj.Name 类型信息，obj.Kind 种类信息
}
func reflectValue(any interface{}) {
	//获取any的值
	v := reflect.ValueOf(any)
	fmt.Printf("%v,%T\n", v, v)
	k := v.Kind() // 获取v的种类
	//根据K的终端转换值类型
	switch k {
	case reflect.Int32:
		ret := int32(v.Int())
		fmt.Println(ret)
	case reflect.Float32:
		ret := float32(v.Float())
		fmt.Println(ret)
	}

}

// 修改反射值信息
func reflectSetValue(any interface{}) {
	v := reflect.ValueOf(any)
	switch v.Elem().Kind() { //需要修改指针类型参数，才有意义，使用Elem()函数来根据指针获取对应的值
	case reflect.Int32:
		v.Elem().SetInt(1998)
		fmt.Println(v)
	case reflect.Float32:
		v.Elem().SetFloat(3.14159)
		fmt.Println(v)
	}
}

type Cat struct {
}

type Dog struct {
}

func main() {
	/*reflectType("asad")
	reflectType(123.34)
	reflectType('1')
	reflectType(true)
	//结构体的类型
	var d Dog
	reflectType(d)
	var c Cat
	reflectType(c)*/

	var a int32 = 100
	var b float32 = 123.123

	reflectValue(a)
	reflectValue(b)

	var c int32 = 123
	reflectSetValue(&c)
}
