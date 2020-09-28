package main

import (
	"context"
	// "crypto/md5"
	// "encoding/hex"
	"fmt"
	// "strconv"
	"redisdemo/redis"

	r "github.com/go-redis/redis/v8"
)

var ctx = context.Background()

// "redisdemo/redis"

func main() {

	// cache := cache.ExampleNewClient()
	options := []r.Options{
		{
			Addr:         "192.168.56.106:6379",
			Password:     "", // no password set
			DB:           0,  // use default DB
			PoolSize:     10,
			DialTimeout:  2,
			ReadTimeout:  2,
			WriteTimeout: 2,
		},
	}
	cache := redis.GetRedis(ctx, "test", options)

	fmt.Println(cache.Ping(ctx).Result())

}
