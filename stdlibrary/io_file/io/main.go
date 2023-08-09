package main

import (
	"fmt"
	"io"
	"os"
)

// 文件相关操作

// 使用read函数读取文件数据
func readDemo(file *os.File) {
	//使用read读取数据
	var buf = make([]byte, 1024)
	n, err := file.Read(buf)
	if err == io.EOF {
		fmt.Println("文件读取完毕")
		return
	}
	if err != nil {
		fmt.Println("文件读取失败：", err)
		return
	}
	fmt.Printf("读取了%d字节数据\n", n)
	fmt.Println(string(buf[:n]))
}

// 循环读取数据
func readCycle(file *os.File) {
	var content []byte
	var tmp = make([]byte, 128)

	for {
		n, err := file.Read(tmp)
		if err == io.EOF {
			fmt.Println("文件读取完毕")
			break
		}
		if err != nil {
			fmt.Println("文件读取失败：", err)
			return
		}
		content = append(content, tmp[:n]...)
	}
	fmt.Println(string(content))
}

// 使用write 写入数据
func writeDemo(file *os.File) {
	str := "snake man"
	_, err := file.Write([]byte(str))
	if err != nil {
		fmt.Println("写入失败", err)
		return
	}
	fmt.Println("success！")
	str2 := "elstic man"
	_, err = file.WriteString(str2)
	if err != nil {
		fmt.Println("写入失败", err)
		return
	}
	fmt.Println("success！")
}
func main() {
	//打开或关闭文件
	//file, err := os.Open("E:/log.txt") //os.Open 默认只读
	file, err := os.OpenFile("E:/log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("open file failed", err)
		return
	}
	fmt.Println("打开成功")
	defer file.Close() //关闭文件
	//readDemo(file)
	readCycle(file)

	writeDemo(file)

}
