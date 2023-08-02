package main

import (
	"fmt"
	"time"
)

// work pool 模式

func worker(id int, jobs <-chan int, results chan<- int) {
	for x := range jobs {
		fmt.Printf("job:%d start\n", id)
		results <- x + 2023
		time.Sleep(time.Millisecond * 500)
		fmt.Printf("job:%d stop\n", id)
	}
}

func main() {
	jobs := make(chan int, 100)    //任务
	results := make(chan int, 100) //返回结果

	//开启3个goroutine
	for i := 0; i < 3; i++ {
		go worker(i, jobs, results)
	}
	//发送5个任务
	for i := 0; i < 5; i++ {
		jobs <- i
	}
	close(jobs)
	//输出结果
	for i := 0; i < 5; i++ {
		ret := <-results
		fmt.Println(ret)
	}
}
