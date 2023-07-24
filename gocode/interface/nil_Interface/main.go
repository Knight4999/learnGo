package main

import (
	"fmt"
)

//空接口

func main() {
	// 方法中没有定义任何方法，该接口就是空接口
	// 任意类型都实现了空接口，因此空接口可以存储任何类型的值
	var x interface{}
	x = "hello"
	fmt.Println(x)
	x = true
	fmt.Println(x)
	x = 100
	fmt.Println(x)

	//类型断言
	/*_, ok := x.(int)*/
	switch v := x.(type) {
	case string:
		fmt.Printf("是字符串类型：%v\n", v)
	case int:
		fmt.Printf("是整数类型：%v\n", v)
	case bool:
		fmt.Printf("是布尔类型：%v\n", v)
	default:
		fmt.Printf("%v不清楚是什么类型\n", v)
	}

	/*var m = make(map[string]interface{}, 16)
	m["name"] = "李白"
	m["age"] = 18
	m["hobby"] = []string{"篮球", "足球", "dota2"}
	fmt.Println(m)*/

}
