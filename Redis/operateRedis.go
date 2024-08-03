package main

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
	// 创建 Redis 客户端
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis 服务器地址
		Password: "",               // Redis 密码，如果没有可以留空
		DB:       0,                // 默认数据库
	})

	// 测试连接
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("could not connect to Redis: %v", err)
	}
	fmt.Println("Connected to Redis successfully!")

	// 设置键值对
	err = rdb.Set(ctx, "mykey", "Hello, Redis!", 0).Err()
	if err != nil {
		log.Fatalf("could not set key: %v", err)
	}
	fmt.Println("Key set successfully!")

	// 获取键值对
	val, err := rdb.Get(ctx, "mykey").Result()
	if err != nil {
		log.Fatalf("could not get key: %v", err)
	}
	fmt.Printf("Key value: %s\n", val)

	// 删除键
	err = rdb.Del(ctx, "mykey").Err()
	if err != nil {
		log.Fatalf("could not delete key: %v", err)
	}
	fmt.Println("Key deleted successfully!")
}
