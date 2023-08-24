package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"os"
	"time"
)

var ctx, cancel = context.WithTimeout(context.Background(), 300*time.Millisecond)

// Do方法

var rds *redis.Client

func initDB() {
	rds = redis.NewClient(&redis.Options{
		Addr:     "192.168.40.196:6379",
		Password: "",
		DB:       0,
		PoolSize: 20,
	})
}

func main() {
	initDB()

	defer cancel()
	//do函数，执行任意命令或自定义命令的 Do 方法，特别是一些 go-redis 库暂时不支持的命令都可以使用该方法执行。
	err := rds.Do(ctx, "SET", "BJHX", 100, "EX", 3600).Err()
	fmt.Println(err)
	// 执行命令获取结果
	val, err := rds.Do(ctx, "get", "BJHX").Result()
	//redis提供了一个redis.Nil,表示值不存在的错误。可以加以判断

	if err != nil {
		if errors.Is(err, redis.Nil) {
			fmt.Println("key dose not exist")
		}
		fmt.Println(" ", err)
		os.Exit(1)
	}
	fmt.Println(val)
}
