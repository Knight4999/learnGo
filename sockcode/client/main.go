package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

// client
func main() {
	//1.与服务端进行连接
	conn, err := net.Dial("tcp", "127.0.0.1:23450")
	if err != nil {
		log.Fatal(err)
	}
	//2.利用连接发送和接受数据
	input := bufio.NewReader(os.Stdin)
	for {
		str, _ := input.ReadString('\n')
		str = strings.TrimSpace(str) //去掉字符串中的空格
		if strings.ToUpper(str) == "Q" {
			return
		}
		//给服务端发送数据
		_, err := conn.Write([]byte(str))
		if err != nil {
			log.Fatal(err)
		}
		//从服务端接受回复的消息
		var buf [1024]byte
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Println("read from conn failed:", err)
			return
		}
		fmt.Println("收到服务端消息：", string(buf[:n]))
	}
}
