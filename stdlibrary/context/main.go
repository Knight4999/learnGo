package main

import (
	"fmt"
	"sync"
	"time"
)

// 为什么需要context

var wg sync.WaitGroup

// 初始的例子
func worker() {
	for {
		fmt.Println("worker")
		time.Sleep(time.Second)
	}
	//如何接收外部的命令实现退出
	wg.Done()
}

// 使用全局变量的方法控制
var exit bool

func worker2() {
	for {
		fmt.Println("worker")
		time.Sleep(time.Second)
		//通过判断全局变量来控制循环退出
		if exit {
			break
		}
	}
	//如何接收外部的命令实现退出
	wg.Done()
}

// 使用管道的方式来处理
// 管道方式存在的问题：
// 1. 使用全局变量在跨包调用时不容易实现规范和统一，需要维护一个共用的channel
func worker3(exitChan chan struct{}) {
LOOP:
	for {
		fmt.Println("worker")
		time.Sleep(time.Second)
		select {
		case <-exitChan:
			break LOOP
		default:
		}
	}
	wg.Done()
}

func main() {
	wg.Add(1)
	/* go worker()
	// 如何优雅的实现结束子goroutine */

	/* go worker2()
	time.Sleep(3 * time.Second) // sleep3秒以免程序过快退出
	exit = true                 // 修改全局变量实现子goroutine的退出 */

	var exitChan = make(chan struct{})

	go worker3(exitChan)
	time.Sleep(time.Second * 3) // sleep3秒以免程序过快退出
	exitChan <- struct{}{}      // 给子goroutine发送退出信号
	close(exitChan)

	wg.Wait()
	fmt.Println("over")
}
