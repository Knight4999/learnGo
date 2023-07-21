package main

import "fmt"

//switch case
func main() {

	var dotahero int = 1
	for ; dotahero < 7; dotahero++ {
		switch dotahero {
		case 1:
			fmt.Println("齐天大圣")
		case 2:
			fmt.Println("祈求者")
		case 3:
			fmt.Println("风暴之灵")
		case 4:
			fmt.Println("影魔")
		case 5:
			fmt.Println("电炎绝手")
		default:
			fmt.Println("谜团")
		}
	}
	// case 一次判断多个值
	num := 0
	for ; num <= 10; num++ {
		switch num {
		case 1, 3, 5, 7, 9:
			fmt.Println("奇数")
		case 2, 4, 6, 8:
			fmt.Println("偶数")
		}
	}
	//case 判断表达式
	value := 100
	switch { //此处switch后不带变量
	case value > 10 && value < 50:
		fmt.Println("A")
	case value >= 50 && value < 100:
		fmt.Println("B")
	default:
		fmt.Println("C")
	}

	for {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
			if i == 5 {
				goto breakTag //跳出循环至标签处
			}
		}
	}
breakTag: //定义标签
	fmt.Println("结束循环")
}
