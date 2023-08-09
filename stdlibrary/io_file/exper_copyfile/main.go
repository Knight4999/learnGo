package main

import (
	"fmt"
	"io"
	"os"
)

//写一个copy文件的函数

func CopyFile(dst, src string) (written int64, err error) {
	//1.打开一个文件
	srcfile, err := os.Open(src)
	if err != nil {
		fmt.Println("打开文件失败:", err)
		return
	}
	defer srcfile.Close()

	//2.以写|创建的方式打开目标文件
	dstfile, err := os.OpenFile(dst, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("打开文件失败:", err)
		return
	}
	defer dstfile.Close()
	return io.Copy(dstfile, srcfile) //调用io.Copy函数
}

func main() {
	_, err := CopyFile("e:/log.txt", "e:/mylog.txt")
	if err != nil {
		fmt.Println("拷贝失败:", err)
		return
	}
	fmt.Println("拷贝成功！")
}
