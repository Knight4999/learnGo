package main

import (
	"fmt"
	"strings"
)

func main() {
	/* //声明不等于初始化
	var a map[string]int
	fmt.Println(a == nil)
	//map的初始化
	a = make(map[string]int, 8)
	fmt.Println(a == nil)

	//map 添加键值对
	a["北京"] = 100
	a["上海"] = 200
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Printf("a:%#v\n", a)
	//声明同时初始化
	b := map[int]string{
		1: "祈求者",
		2: "影魔",
		3: "蓝猫",
	}
	fmt.Printf("b:%#v\n", b)
	fmt.Printf("%T\n", b)

	//判断map中的键值对是否存在
	var scoreMap = make(map[string]int, 8)
	scoreMap["李白"] = 100
	scoreMap["高进"] = 200

	v, ok := scoreMap["李白"]
	if ok {
		fmt.Println("存在", v)

	}

	for k, v := range scoreMap {
		fmt.Println(k, v)
	}

	//删除某个键值对
	delete(scoreMap, "高进")
	for k, v := range scoreMap {
		fmt.Println(k, v)
	} */
	/* var a = make(map[string]int, 100)

	//一次性添加50个键值对
	for i := 0; i < 50; i++ {
		key := fmt.Sprintf("std%02d", i)
		value := rand.Intn(100) //0~99的随机数
		a[key] = value
	}

	for k, v := range a {
		fmt.Println(k, v)
	}

	//按照Key的顺序去排序
	var keys = make([]string, 0, 100)
	for k := range a {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	for _, v := range keys {
		fmt.Println(v, a[v])
	} */

	//元素为map的切片
	var mapSlice = make([]map[int]string, 10, 10)
	/* mapSlice[0] = map[int]string{1: "1998", 2: "2003"}
	mapSlice[1] = map[int]string{3: "2008", 4: "2013"} */
	mapSlice[0] = make(map[int]string, 8)
	mapSlice[0][10] = "AD"
	mapSlice[0][123] = "AS"
	mapSlice[0][145] = "DS"
	for _, v := range mapSlice {
		fmt.Println(v)
	}
	fmt.Println(mapSlice)

	//值为切片的map
	var sliceMap = make(map[string][]int, 8)
	v, ok := sliceMap["中国"]
	if ok {
		fmt.Println(v)
	} else {
		sliceMap["中国"] = make([]int, 8)
		sliceMap["中国"][0] = 100
		sliceMap["中国"][1] = 200
		sliceMap["中国"][3] = 300
		sliceMap["中国"][4] = 400
	}

	for k, v := range sliceMap {
		fmt.Println(k, v)
	}
	A()
}

func A() {
	//统计一个字符串每个单词出现的次数
	//How do you do
	var s = "how do you do"
	var wordCount = make(map[string]int, 10)
	//将字符串拆分为不同的单词
	words := strings.Split(s, " ")

	for _, word := range words {
		_, ok := wordCount[word]
		if ok {
			wordCount[word]++
		} else {
			wordCount[word] = 1
		}
	}

	for k, v := range wordCount {
		fmt.Println(k, v)
	}
}
