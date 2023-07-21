package main

import (
	"fmt"
	"math"
)

func main() {
	var a int = 1024
	fmt.Printf("%d\n", a)
	fmt.Printf("%b\n", a)
	fmt.Printf("%o\n", a)

	var age uint8 = 255
	fmt.Println(age)

	fmt.Println(math.MaxFloat64)

	str := "\"Google一下\"\n"
	fmt.Print(str)
	fmt.Println("跛豪")
}
