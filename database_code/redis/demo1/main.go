package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

// Go 操作redis数据库

var rds *redis.Client

var rds_cluster *redis.ClusterClient

// 初始化连接
func initDB() error {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	//普通连接方式
	rds = redis.NewClient(&redis.Options{
		Addr:     "192.168.40.196:6379",
		Password: "",
		DB:       0,
		PoolSize: 20,
	})
	_, err := rds.Ping(ctx).Result()
	if err != nil {
		fmt.Println("faild connect!")
		return err
	}
	return nil
	// redis.ParseURL 函数从表示数据源的字符串中解析得到 Redis 服务器的配置信息。
	/*opt,err := redis.ParseURL("redis://<user>:<pass>@localhost:6379/<db>")
	if err != nil {
	    panic(err)
	}
	rds = redis.NewClient(opt)*/
}

// TLS连接模式
func initDbByTLS() {
	rds = redis.NewClient(&redis.Options{
		TLSConfig: &tls.Config{
			MinVersion:   tls.VersionTLS12,      //tls版本
			Certificates: []tls.Certificate{},   //证书验证
			ServerName:   "proxy.wzhtheone.xyz", //域名
		},
	})
}

// 哨兵模式
func initDbBySentinel() {
	rds = redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:    "master",
		SentinelAddrs: []string{":6379", ":6380", "6381"},
	})
}

// 集群模式
func initDbByCluster() {
	rds_cluster = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{":7000", ":7001", ":7002", ":7003", ":7004"},
		// 若要根据延迟或随机路由命令，请启用以下命令之一
		// RouteByLatency: true,
		// RouteRandomly: true,
	})
}

// go-redis 库的基本使用。
func doCommand() {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond) //创建一个上下文，
	defer cancel()

	//执行命令获取结果
	val, err := rds.Get(ctx, "class").Result()
	fmt.Println(val, err)

	lists, err := rds.LRange(ctx, "tools", 0, -1).Result()
	fmt.Println(lists, err)

	//先获取到命令对象
	cmder := rds.SMembers(ctx, "name")
	fmt.Println(cmder.Val())
	fmt.Println(cmder.Err())

	//直接执行命令行获取错误
	err = rds.Set(ctx, "cs", 8, 100*time.Millisecond).Err()
	fmt.Println(err)

	//直接执行命令获取值
	value := rds.Get(ctx, "class").Val()
	fmt.Println(value)

}

func main() {
	//初始化数据库连接
	err := initDB()
	if err != nil {
		fmt.Println(err)
		return
	}
	doCommand()
}
