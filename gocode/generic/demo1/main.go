package main

import "fmt"

// 泛型类型

// 一个泛型切片
type Slice[T int | string] []T

// 一个泛型map
type Map[K int | string, V float32 | float64] map[K]V

// 一个泛型结构体
type MyStruct[T int | string] struct {
	Name  string
	value T
}

// 一个泛型接口
type IPrintData[T int | float32 | float64] interface {
	Print(data T)
}

//一个泛型管道

type Mychan[T int | string | bool] chan T

// 泛型类型嵌套
type Wstruct[T int | float64, S []T] struct {
	data     S
	MaxValue T
	MinValue T
}

// 使用interface{}来处理指针类型 类型参数的问题
type Newtype[T interface{ *int | *float64 }] []T

// 泛型方法
func (ret Slice[T]) sum() T {
	var sum T
	for _, value := range ret {
		sum += value
	}
	return sum
}

// 泛型函数
func min[T int | float64](a, b T) T {
	if a <= b {
		return a
	}
	return b
}

// 匿名函数支持使用提前定义好的泛型
func Myfunc[T int | float64](a, b T) {
	f := func(i T, j T) T {
		return i*2 - j*2
	}
	fmt.Println(f(a, b))
}

func main() {
	/*m := min[int](10, 20)
	fmt.Println(m)

	m2 := min(12.23, 34.09)
	fmt.Println(m2)

	//泛型类型实例化
	fmin := min[float64] //实例化泛型函数，得到一个非泛型函数
	m3 := fmin(98.23, 35.97)
	fmt.Println(m3)

	var s = MyStruct[string]{
		"SAD",
		"10086",
	}
	fmt.Printf("%T\n", s)
	fmt.Println(s.Name, s.value)

	var c = make(Mychan[int], 10)
	c <- 100
	x := <-c
	fmt.Println(x)

	var ws = Wstruct[int, []int]{
		data:     make([]int, 5),
		MaxValue: 10,
		MinValue: 1,
	}
	ws.data = append(ws.data, 1, 2, 3, 4, 5)
	fmt.Println(ws.data, ws.MaxValue, ws.MinValue)

	var nt = make(Newtype[*int], 10)
	var i = 10
	var j = 162

	nt[0] = &i
	nt[1] = &j
	fmt.Println(nt)

	var ns Slice[int] = []int{1, 2, 3, 4, 5, 6, 7, 8, 9} //调用泛型方法
	fmt.Println(ns.sum())*/

	//匿名函数不支持自定义泛型
	/*f := func[T int | string](a T) T {
		return a
	}*/

	Myfunc(1, 10)
}
