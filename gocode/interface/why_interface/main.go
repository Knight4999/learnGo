package main

import "fmt"

type Dog struct {
}

type Cat struct {
}

func (d Dog) move() {
	fmt.Println("狗会跑")
}
func (c Cat) move() {
	fmt.Println("猫会跑")
}

type Person struct {
	name string
}

func (p Person) move() {
	fmt.Println("人会跑")
}

type mover interface {
	move()
}

func da(ary mover) {
	ary.move()
}

func main() {
	/*c1 := Cat{}
	da(c1)
	d1 := Dog{}
	da(d1)

	p1 := Person{"李白"}
	da(p1)*/
	var m1 mover
	c1 := Cat{}
	m1 = c1
	fmt.Printf("%T\n", m1)

}
