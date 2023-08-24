package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

// ZSET 演示

var rds *redis.Client

var ctx, cancel = context.WithTimeout(context.Background(), 500*time.Millisecond)

func initDB() {
	rds = redis.NewClient(&redis.Options{
		Addr:     "192.168.40.196:6379",
		Password: "",
		DB:       0,
		PoolSize: 20,
	})
}

func zestDemo() {
	defer cancel()
	//key
	zestKey := "language_rank"
	//value
	languages := []redis.Z{
		{Score: 90.0, Member: "Golang"},
		{Score: 98.0, Member: "Java"},
		{Score: 95.0, Member: "Python"},
		{Score: 97.0, Member: "JavaScript"},
		{Score: 99.0, Member: "C++"},
	}
	//ZADD
	err := rds.ZAdd(ctx, zestKey, languages...).Err()
	if err != nil {
		fmt.Printf("zadd failed, err:%v\n", err)
		return
	}
	fmt.Println("zadd success")

	// 把Golang的分数加10
	newScore, err := rds.ZIncrBy(ctx, zestKey, 10.0, "Golang").Result()
	if err != nil {
		fmt.Printf("zincrby failed, err:%v\n", err)
		return
	}
	fmt.Printf("Golang's score is %f now.\n", newScore)

	//取分数最高的三个元素
	ret := rds.ZRevRangeWithScores(ctx, zestKey, 0, 2).Val()
	for _, z := range ret {
		fmt.Println(z.Score, z.Member)
	}

	//取95 - 100分的元素
	op := &redis.ZRangeBy{
		Min: "95",
		Max: "100",
	}
}

func main() {
	initDB()
	zestDemo()
}
