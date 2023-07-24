package main

import "fmt"

// 接口嵌套
type animal interface {
	mover
	sayer
}

// 使用值接受者和指针接受者的区别
type mover interface {
	move()
}
type sayer interface {
	say()
}
type person struct {
	name string
	age  int
}

// 使用值接受者实现接口:类型的值和类型的指针都能够保持到接口当中
/*func (p person) move() {
	fmt.Printf("%s再跑\n", p.name)
}*/

// 使用指针类型的接口

func (p *person) move() {
	fmt.Printf("%s再跑\n", p.name)
}

func (p *person) say() {
	fmt.Printf("%s再叫\n", p.name)
}
func main() {
	var m mover
	/*p1 := person{
		"李白",
		20,
	}*/
	p2 := &person{
		"杜甫",
		45,
	}
	//m = p1 //指针类型接受者不可以存储在
	m = p2
	m.move()
	fmt.Println(m)

	var s sayer
	s = p2
	s.say()
	fmt.Println(s)

	var i animal
	i = p2
	i.move()
	i.say()
	fmt.Println(i)
}
