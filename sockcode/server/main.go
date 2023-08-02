package main

import (
	"bufio"
	"fmt"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()
	//针对当前的连接做数据的交互
	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:])
		if err != nil {
			fmt.Printf("read from conn failed:%v\n", err)
			break
		}
		recv := string(buf[:n])
		fmt.Println("接收到的数据：", recv)
		conn.Write([]byte("ok，已收到"))
	}
}

func main() {
	//1.监听端口
	listen, err := net.Listen("tcp", "127.0.0.1:23450")
	if err != nil {
		fmt.Println("监听异常", err)
		return
	}
	//2.接收客户端请求建立链接
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("conn failed", err)
			continue
		}
		//3.创建goroutine处理链接。
		go process(conn)
	}
}
