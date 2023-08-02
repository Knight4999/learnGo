package main

import (
	"fmt"
	"log"
	"net"
)

// UDP Server
func main() {
	//1.监听端口
	listen, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 23450,
	})
	if err != nil {
		fmt.Println("listen failed：", err)
		return
	}
	defer listen.Close()

	for {
		var buf [1024]byte
		//从UDP中读数据
		n, addr, err := listen.ReadFromUDP(buf[:])
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("接收到的数据：", string(buf[:n]))
		//往UDP中写数据
		_, err = listen.WriteToUDP(buf[:n], addr)
		if err != nil {
			fmt.Println("写入失败", err)
			return
		}
	}
}
