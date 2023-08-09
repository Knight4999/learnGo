package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

// bufio 带缓存空间的io流

func main() {
	file, err := os.OpenFile("e:/log.txt", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	//使用bufio封装
	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n') //以换行为条件读取数据
		if err == io.EOF {
			if len(line) != 0 {
				fmt.Println(line)
			}
			fmt.Println("文件读完了")
			break
		}
		if err != nil {
			fmt.Println("read file failed, err:", err)
			return
		}
		fmt.Println(line)
	}

	//使用os.readfile函数来处理
	content, err := os.ReadFile("e:/log.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(content))
	//创建缓存写入对象
	writer := bufio.NewWriter(file)
	str := "北京晚报"
	writer.Write([]byte(str)) //将数据写入缓存区
	writer.WriteString(str)

	writer.Flush() //将缓存区数据写入到文件中

	//使用os.wirterFile
	str2 := "走向共和"
	err = os.WriteFile("E:/log.txt", []byte(str2), 0666)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("写入成功")
}
