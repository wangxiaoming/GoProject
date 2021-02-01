package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

//Background返回一个非空的Context。 它永远不会被取消，没有值，也没有期限。
//它通常在main函数，初始化和测试时使用，并用作传入请求的顶级上下文。
var ctx = context.Background()

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "172.31.1.135:7000",
		Password: "",
		DB:       0,
	})

	pong, err := rdb.Ping(ctx).Result()
	fmt.Println(pong)
	if err != nil {
		fmt.Printf("连接redis出错，错误信息：%v", err)
	}
	fmt.Println("成功连接redis")
}
