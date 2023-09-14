package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

// redis 事务

/*
redis 是单线程执行命令的，因此单个命令始终是原子的，但是来自不同客户端的两个给定命令可以依次执行，例如在它们之间交替执行。
但是，Multi/exec能够确保在multi/exec两个语句之间的命令之间没有其他客户端正在执行命令。
在这种场景我们需要使用 TxPipeline 或 TxPipelined 方法将 pipeline 命令使用 MULTI 和EXEC包裹起来。
*/
var rds *redis.Client

var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Millisecond)

func initRedisDb() {
	rds = redis.NewClient(&redis.Options{
		Addr:     "192.168.100.196:6379",
		DB:       0,
		PoolSize: 10,
		Password: "",
	})
}

// watchDemo 在key值不变的情况下将其值+1
/*
我们通常搭配 WATCH命令来执行事务操作。从使用WATCH命令监视某个 key 开始，
直到执行EXEC命令的这段时间里，如果有其他用户抢先对被监视的 key 进行了替换、
更新、删除等操作，那么当用户尝试执行EXEC的时候，事务将失败并返回一个错误，用户可以根据这个错误选择重试事务或者放弃事务。
*/
func watchDemo(ctx context.Context, key string) error {
	return rds.Watch(ctx, func(tx *redis.Tx) error {
		n, err := tx.Get(ctx, key).Int()
		if err != nil && err != redis.Nil {
			return err
		}
		// 假设操作耗时5秒
		// 5秒内我们通过其他的客户端修改key，当前事务就会失败
		time.Sleep(5 * time.Second)
		_, err = tx.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
			pipe.Set(ctx, key, n+1, time.Hour)
			return nil
		})
		return err
	}, key)
}
func main() {
	//TxPipeline demo
	pipe := rds.TxPipeline()
	incr := pipe.Incr(ctx, "tx_pipeline_counter")
	pipe.Expire(ctx, "tx_pipeline_counter", time.Hour)
	_, err := pipe.Exec(ctx)
	fmt.Println(incr.Val(), err)

	// TxPipelined demo
	var incr2 *redis.IntCmd
	_, err = rds.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
		incr2 = pipe.Incr(ctx, "tx_pipeline_counter")
		pipe.Expire(ctx, "tx_pipeline_counter", time.Hour)
		return nil
	})
	fmt.Println(incr2.Val(), err)
}
