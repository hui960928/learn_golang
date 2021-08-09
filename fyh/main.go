package main

import (
	"context"
	"fmt"
	"learn_golang/fyh/pub_redis"
	"time"
)

var kongRedis *pub_redis.RedisPool

func init() {
	kongRedis = pub_redis.NewRedisPool(&pub_redis.RedisConf{
		Host:     "127.0.0.1",
		Port:     "6379",
		Password: "",
		DB:       3,
		PoolConfig: &pub_redis.PoolConfig{
			MaxIdle:     5000,
			MaxActive:   0,
			IdleTimeout: 5 * time.Second,
			Wait:        false,
		},
	})
}

func GetKongRedis() *pub_redis.RedisPool {
	return kongRedis
}

func main() {
	a, err := GetKongRedis().Get(context.Background(), "456")
	fmt.Println(a)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(a)
}
