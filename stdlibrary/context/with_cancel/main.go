package main

import (
	"context"
	"fmt"
	"time"
)

// withCancel函数

func speak(ctx context.Context) {
	for range time.Tick(time.Second) {
		select {
		case <-ctx.Done():
			fmt.Println("我要闭嘴了")
			return
		default:
			fmt.Println("balabalabalabala")
		}
	}
}
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go speak(ctx)
	time.Sleep(10 * time.Second)
	cancel() //十秒后触发终止通知
	time.Sleep(1 * time.Second)
}
