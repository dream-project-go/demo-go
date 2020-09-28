package redis

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"io"
	"log"
	"strconv"

	r "github.com/go-redis/redis/v8"
)

func ExampleNewClient(ctx context.Context, opt r.Options) *r.Client {
	rdb := r.NewClient(&r.Options{
		Addr:         opt.Addr,
		Password:     opt.Password, // no password set
		DB:           opt.DB,       // use default DB
		PoolSize:     opt.PoolSize,
		DialTimeout:  opt.DialTimeout,
		ReadTimeout:  opt.ReadTimeout,
		WriteTimeout: opt.WriteTimeout,
	})
	return rdb
}

func GetRedis(ctx context.Context, key string, options []r.Options) *r.Client {
	h := md5.New()
	io.WriteString(h, key)
	var guid string
	guid = hex.EncodeToString(h.Sum(nil))
	hh := guid[len(guid)-1:]
	d, err := strconv.ParseUint(hh, 16, 32)
	if err != nil {
		log.Print(err)
	}
	l := len(options)
	mod := int(d) % l
	opt := options[mod]
	client := ExampleNewClient(ctx, opt)
	return client
}
