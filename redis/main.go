package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"sync"
	"time"
)

var cache *redis.Client
var redisOnce sync.Once

func GetCache() *redis.Client {
	redisOnce.Do(initCache)
	return cache
}

func initCache() {
	cache = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", "127.0.0.1", 6379),
		DB:   0,
	})

	if err := cache.Ping().Err(); err != nil {
		fmt.Println("redis ping err: ", err)
	}
}

type Msg struct {
	Token  string `json:"token"`
	UserId int    `json:"user_id"`
	S      string `json:"s"`
}

type Msg1 struct {
	Token  string `json:"token"`
	UserId int    `json:"user_id"`
}

func main() {
	msg := Msg{
		Token:  "token",
		UserId: 1,
	}
	data, err := json.Marshal(msg)
	fmt.Println(data, err)

	b := GetCache().SetNX("test_1", data, time.Minute)
	fmt.Println(b.Result())

	var (
		msg1  Msg1
		data1 []byte
	)
	s := GetCache().Get("test_1")
	if data1, err = s.Bytes(); err != nil {
		fmt.Println("s.Bytes() err:", err.Error())
		return
	}
	if err = json.Unmarshal(data1, &msg1);err != nil {
		fmt.Println("Unmarshal err:", err.Error())
		return
	}
	fmt.Println(msg1)
}
