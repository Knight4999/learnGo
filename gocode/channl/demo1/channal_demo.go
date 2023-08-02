package main

import (
	"fmt"
)

func main() {

	ch1 := make(chan int) //无缓存分区channal
	//ch2 := make(chan int, 1) //带缓存分区channal
	ch1 <- 10 //发送值
	x := <-ch1
	fmt.Println(x)
	close(ch1)
}
