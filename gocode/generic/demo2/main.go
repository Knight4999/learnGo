package main

import "fmt"

//泛型内容 2

// 定义泛型接口
type Int interface {
	~int | int8 | int16 | int32 | int64 //通过使用 ‘ ~ ’ 符号，使所有以 int 为底层类型的类型也都可用于实例化
}

type UInt interface {
	uint | uint8 | uint16 | uint32 | uint64
}

type Float interface {
	float64 | float32
}

// Slice 使用'|'将多个接口组合使用
type Slice[T Int | UInt | Float] []T

// 接口内部也可以组合其他接口

type StringSlice interface {
	Int | UInt | Float | string // 组合了三个接口类型并额外增加了一个 string 类型
}

type NewSlice[T StringSlice] []T

// map 中键的类型必须是可进行 != 和 == 比较的类型，所以使用any不合适，要使用conparable
type mS[KEY comparable, Value any] map[KEY]Value

func main() {
	/*var s = make(NewSlice[uint8], 0, 10)
	fmt.Printf("%T , %#v\n", s, s)*/

	var s1 NewSlice[int]
	fmt.Println(s1)
	type MyInt int
	//var s2 NewSlice[MyInt] //MyInt类型底层是int，但是不是int类型，不满足泛型约束条件，
	var s2 NewSlice[MyInt]
	fmt.Println(s2)

}
