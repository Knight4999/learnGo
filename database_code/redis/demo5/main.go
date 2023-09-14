package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

/*
Redis Pipeline 允许通过使用单个 client-server-client 往返执行多个命令来提高性能。
区别于一个接一个地执行100个命令，你可以将这些命令放入 pipeline 中，然后使用1次读写操
作像执行单个命令一样执行它们。这样做的好处是节省了执行命令的网络往返时间（RTT）。
*/
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

func pipelineDemo() {
	pipe := rds.Pipeline()

	incr := pipe.Incr(ctx, "pipeline_counter")
	pipe.Expire(ctx, "pipeline_counter", time.Hour)

	_, err := pipe.Exec(ctx)
	if err != nil {
		panic(err)
	}

	// 在执行pipe.Exec之后才能获取到结果
	fmt.Println(incr.Val())
}

func pipelinedDemo() {
	var incr *redis.IntCmd
	_, err := rds.Pipelined(ctx, func(pipeliner redis.Pipeliner) error {
		incr = pipeliner.Incr(ctx, "pipeline_counter")
		pipeliner.Expire(ctx, "pipeline_counter", time.Second*60)
		return nil
	})
	if err != nil {
		panic(err)
	}
	// 在pipeline执行后获取到结果
	fmt.Println(incr.Val())
}

func pipelinedDemo2() {
	cmds, err := rds.Pipelined(ctx, func(pipe redis.Pipeliner) error {
		for i := 0; i < 100; i++ {
			pipe.Get(ctx, fmt.Sprintf("key%d", i))
		}
		return nil
	})
	if err != nil {
		panic(err)
	}

	for _, cmd := range cmds {
		fmt.Println(cmd.(*redis.StringCmd).Val())
	}
}
func main() {
	pipelineDemo()
	pipelinedDemo()
	pipelinedDemo2()
}
