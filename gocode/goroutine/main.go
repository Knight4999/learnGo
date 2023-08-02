package main

//goroutine  demo
import (
	"fmt"
	"sync"
)

func A() {
	fmt.Println("Hello World")
	wg.Done() //通知计数器 -1
}

func Hello(i int) {
	fmt.Println("hello", i)
	wg.Done()

}

var wg sync.WaitGroup

func main() {

	wg.Add(1000) //开启一个协程，技术拍 +1
	/*for i := 0; i < 1000; i++ {
		go Hello(i)
	}*/
	for i := 0; i < 1000; i++ {
		go func(i int) {
			fmt.Println("hello!", i)
			wg.Done()
		}(i)
	}
	//A()
	fmt.Println("Earth Song")
	wg.Wait() //阻塞，等待所有的协程结束后，再关闭。
}
