package main

import (
	"fmt"
	"strings"
)

// 字符串常用操作
func main() {
	str := "hello"
	//len 字符串长度
	fmt.Println(len(str))
	str2 := "hello北京"
	fmt.Println(len(str2))

	//拼接字符串
	fmt.Println(str + str2)
	fmt.Sprintf("%s - %s", str, str2)

	//字符串分割
	s := "How do you do"
	s1 := strings.Split(s, " ")
	fmt.Printf("%T\n", s1)
	fmt.Println(s1)

	//字符串是否包含
	name := "李火旺"
	fmt.Println(strings.Contains(name, "白"))

	//判断前缀
	fmt.Println(strings.HasPrefix(s, "How"))
	//判断后缀
	fmt.Println(strings.HasSuffix(s, "do"))

	//判断字符串位置(记录第一次出现的位置)
	fmt.Println(strings.Index(s, "do"))
	//判断字符串最后出现的位置
	fmt.Println(strings.LastIndex(s, "do"))

	//join,字符串拼接
	s2 := []string{"How", "do", "you", "do"}
	fmt.Println(s2)
	fmt.Println(strings.Join(s2, "："))
}
