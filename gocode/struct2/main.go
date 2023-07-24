package main

import "fmt"

//结构体的继承

type Animal struct {
	name string
}

func (a *Animal) move() {
	fmt.Printf("%s会跑\n", a.name)
}

type Dog struct {
	Feet    int8
	*Animal //匿名嵌套，嵌套结构体指针
}

func (d *Dog) wang() {
	fmt.Printf("%s会汪汪叫\n", d.name)
}
func main() {
	d1 := &Dog{
		4,
		&Animal{
			"乐乐",
		},
	}
	d1.move()
	d1.wang()
}
