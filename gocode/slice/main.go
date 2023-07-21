package main

import (
	"fmt"
	"sort"
)

// 切片
func main() {
	/* var arr = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var a []int
	var b []string
	var c = []bool{true, false}
	fmt.Println(arr, a, b, c)

	//基于数组
	i := arr[1:5] //{2,3,4,5,0,0,0,0,0}
	fmt.Println(i)
	fmt.Printf("%T\n", i)
	fmt.Println(len(i), cap(i))
	//基于切片
	j := i[1:9] //超出部分从底层数组获取新切片 {3,4,5,6,7,8,9,10}
	// j := i[0:3]
	fmt.Println(j)
	fmt.Printf("%T\n", j)
	fmt.Println(len(j), cap(j))
	//make函数
	k := make([]int, 5, 10)
	fmt.Println(k)

	var l []int //声明int切片
	fmt.Println(l, len(l), cap(l), l == nil)
	var m = []int{} //声明并初始化
	fmt.Println(m, len(m), cap(m), m == nil)
	n := make([]int, 0) //make函数
	fmt.Println(n, len(n), cap(n), n == nil)

	//切片赋值拷贝
	o := make([]int, 3)
	p := o //相当于两个切片指向同一个数组
	p[0] = 100
	fmt.Println(o, p) */

	// Append()函数-->切片的扩容
	var a []int
	/* a = make([]int, 4, 20)
	a[0] = 100 */
	a = append(a, 1, 2, 3, 4, 5)
	fmt.Println(a)
	var a2 = []int{11, 12, 13, 14, 15}
	a = append(a, a2...) //如果以切片为参数，需要在变量名后面+...
	fmt.Println(a)

	//动态扩容，每次超出容量会扩容双倍的容量
	var b []int
	for i := 0; i < 10; i++ {
		b = append(b, i)
		fmt.Printf("%v len:%d cap:%d ptr%p\n", b, len(b), cap(b), b)
	}

	var c = make([]int, 10, 20)
	//Copy函数。复制切片
	copy(c, a)
	c[0] = 1000
	fmt.Println(a, c)
	// 切片删除元素
	var city = []string{"北京", "上海", "深圳", "广州", "武汉"}
	city = append(city[0:3], city[4:]...)
	fmt.Println(city)
	//对数组排序
	var n = [...]int{12, 532, 4, 63, 42, 13, 6, 6, 57, 32, 9}
	sort.Ints(n[:])
	fmt.Println(n)
}
