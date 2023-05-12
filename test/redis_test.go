package test

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"testing"
)

var ctx = context.Background()

var rdb = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "", // no password set
	DB:       0,  // use default DB
})

func TestRedisSet(t *testing.T) {
	err := rdb.Set(ctx, "mail", "code", 0).Err()
	if err != nil {
		panic(err)
	}
}

func TestRedisGet(t *testing.T) {
	val, err := rdb.Get(ctx, "mail").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("mail", val)
}
