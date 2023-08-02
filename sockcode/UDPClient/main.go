package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

// UDP client

func main() {
	c, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 23450,
	})
	if err != nil {
		fmt.Println("Dial Failed", err)
		return
	}
	defer c.Close()
	input := bufio.NewReader(os.Stdin)
	for {
		s, err := input.ReadString('\n')
		_, err = c.Write([]byte(s))
		if err != nil {
			fmt.Println("发送失败：", err)
			return
		}
		//接受数据
		var buf [1024]byte
		n, addr, err := c.ReadFromUDP(buf[:])
		if err != nil {
			fmt.Println("send 数据错误：", err)
		}
		fmt.Printf("read form %v, mag:%v \n", addr, string(buf[:n]))
	}
}
