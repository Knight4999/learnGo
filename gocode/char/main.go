package main

import "fmt"

//字符
func main() {
	//byte 相当于uint8 ACCII码
	//rune int32的别名 Unicode
	var c1 byte = 'C'
	var c2 rune = 'C'
	fmt.Println(c1, c2)
	fmt.Printf("c1:%T,c2:%T\n", c1, c2)

	s := "hello北京"
	for i := 0; i < len(s); i++ { //中文部分会乱码
		fmt.Printf("%c\n", s[i])
	}
	fmt.Println("---------------------")
	for _, r := range s { //可以读出字符串中的中文部分
		fmt.Printf("%c", r)
	}
}
