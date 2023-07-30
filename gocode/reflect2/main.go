package main

import (
	"fmt"
	"reflect"
)

//结构体反射

type Student struct {
	Name  string `json:"name" ini:"s_name"`
	Score int    `json:"score" ini:"s_score"`
}

func main() {
	stu1 := Student{
		"明天",
		90,
	}

	//通过反射去获取结构体中的所有字段信息
	t := reflect.TypeOf(stu1)
	fmt.Printf("name:%v kind:%v\n", t.Name(), t.Kind())
	//遍历结构体变量的所有字段
	for i := 0; i < t.NumField(); i++ {
		sf := t.Field(i)
		fmt.Printf("name:%v type:%v tag:%v \n", sf.Name, sf.Type, sf.Tag)
		fmt.Println(sf.Tag.Get("json"), sf.Tag.Get("ini"))
	}

	//根据名字取结构体中的字段
	sf, ok := t.FieldByName("Score")
	if ok {
		fmt.Printf("name:%v type:%v tag:%v \n", sf.Name, sf.Type, sf.Tag)
	}
}
