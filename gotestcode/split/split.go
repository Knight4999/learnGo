package split

import "strings"

// 字符串切割函数

// Split 将s按照sep进行切割,并返回一个字符串
func Split(s, sep string) (ret []string) {
	index := strings.Index(s, sep)                   //检索子字符串出现的位置，如果返回-1，则表示
	ret = make([]string, 0, strings.Count(s, sep)+1) //根据性能测试，发现此处有多次申请内存空间操作。所以做统一初始化操作
	for index > -1 {
		//将字符串s，按照sep截取，存储在ret中
		ret = append(ret, s[:index])
		s = s[index+len(sep):] //扣除已截取的字符串
		index = strings.Index(s, sep)
	}
	ret = append(ret, s)
	return
}
