package main

import (
	"fmt"
	"sync"
)

// 多个协程并发操作全局变量x

var (
	wg  sync.WaitGroup
	mux sync.Mutex //互斥锁
	x   int64
)

func add() {
	for i := 0; i < 50000; i++ {
		mux.Lock()
		x = x + 1
		mux.Unlock()
	}
	wg.Done()
}

func readNum() {

}
func main() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
}
