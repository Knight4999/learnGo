package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

// 扫描或遍历所有的Key
var rds *redis.Client

var ctx, cancel = context.WithTimeout(context.Background(), 500*time.Millisecond)

func initDB() {
	rds = redis.NewClient(&redis.Options{
		Addr:     "192.168.100.196:6379",
		Password: "",
		DB:       0,
		PoolSize: 20,
	})
}

// scanKeysDemo1 按前缀查找所有key示例
func scanKeysDemo1() {
	defer cancel()
	var cursor uint64
	for {
		var keys []string
		var err error
		// 按前缀扫描key
		keys, cursor, err = rds.Scan(ctx, cursor, "prefix:*", 0).Result()
		if err != nil {
			panic(err)
		}

		for _, key := range keys {
			fmt.Println("key", key)
		}

		if cursor == 0 { // no more keys
			break
		}
	}

}

// scanKeysDemo2 按前缀查找所有key示例
func scanKeysDemo2() {
	defer cancel()

	//安装前缀扫描key
	iter := rds.Scan(ctx, 0, "prefix:*", 0).Iterator()
	for iter.Next(ctx) {
		fmt.Println("keys", iter.Val())
	}
	if err := iter.Err(); err != nil {
		panic(err)
	}
}
func main() {
	//scanKeysDemo1()
	scanKeysDemo2()
}
