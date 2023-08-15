package main

import (
	"context"
	"fmt"
	"time"
)

// withTimeout / withDeadline 一般用作于超时控制

func NewContextWithTimeout() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 3*time.Second)
}

/*
	func HttpHandler() {
		ctx, cancel := NewContextWithTimeout()
		defer cancel()
		deal(ctx)
	}
*/
func HttpHandler() {
	ctx, cancel := NewContextWithTimeout()
	defer cancel()
	deal(ctx, cancel)
}

// deal 超时结束进程质量
func deal(ctx context.Context, cancel context.CancelFunc) {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		select {
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			return
		default:
			fmt.Printf("deal time is %d \n", i)
			cancel()
		}
	}
}
func main() {
	HttpHandler()
}
