package main

import (
	"fmt"
	"time"
)

// ticker  定时器

func main() {
	go func() {
		timer := time.NewTimer(time.Second * 1) //仅仅执行一次
		for i := range timer.C {
			fmt.Println("1", i)
		}
	}()

	go func() {
		ticker := time.NewTicker(time.Second * 3) //每到时间就会执行一次，多次执行

		for i := range ticker.C {
			fmt.Println("2", i)
		}
	}()

	tc := time.Tick(time.Second * 5)
	for i := range tc {
		fmt.Println("3", i)
	}
}
