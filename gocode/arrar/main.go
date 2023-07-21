package main

import "fmt"

func main() {
	/* var arr [3]int
	var arr2 [4]int
	fmt.Println(arr, arr2)

	//1.初始化
	var city = [10]string{"北京", "上海", "广州", "深圳"}
	fmt.Println(city)

	//2.编译器自动推导
	var bools = [...]bool{true, false, false, true, false}
	fmt.Println(len(bools))

	//3.索引值初始化
	var indexs = [7]int{1, 3, 5}
	indexs[5] = 10
	fmt.Println(indexs)

	var nums = [...]int{1: 10, 5: 9, 10: 1024}
	fmt.Println(len(nums))
	fmt.Println(nums)
	fmt.Printf("%T\n", nums)

	for i, v := range nums {
		fmt.Printf("nums[%d]:%d\n", i, v)
	} */

	//4.多维数组
	arr := [10][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(arr)
	arr2 := [...][3]int{}
	fmt.Println(arr2)
}
