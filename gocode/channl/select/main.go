package main

import "fmt"

// select
func main() {
	ch1 := make(chan int, 1)

	for i := 0; i < 10; i++ {
		select {
		case x := <-ch1:
			fmt.Println(x)
		case ch1 <- i:
		default:
			fmt.Println("do nothing")
		}
	}
}
