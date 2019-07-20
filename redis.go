package main

import (
	"encoding/json"
	"github.com/go-redis/redis"
)

func main() {
	var redis_options *redis.Options = &redis.Options{
		Addr:"172.16.8.216:6379",
		Password:"dobi#!1123Wqs",
		DB:15}
	redis_clien := redis.NewClient(redis_options)
	defer redis_clien.Close()

	var err error
	/*
		//字符串类型  string
		err = redis_clien.Set("one", "no", 0).Err()
		if err != nil {
			panic(err)
		}

		v, err := redis_clien.Get("one").Result()
		if err != nil {
			panic(err)
		}
		fmt.Println(v)
	*/


	//哈希类型  hash
	/*
		err = redis_clien.HSet("hset", "a","aaa").Err()
		err = redis_clien.HSet("hset", "b","bbb").Err()
		if err != nil {
			panic(err)
		}

		h1, err := redis_clien.HGet("hset", "a").Result()	//字符串
		h2, err := redis_clien.HGetAll("hset").Result()	//字典

		fmt.Println(h1, h2)
	*/


	//列表  list
	type stru struct {
		Name string
		Old int
	}
	s1 := stru{"xd", 30}
	s2 := stru{"xd0", 300}
	s11, _ := json.Marshal(s1)
	s22, _ := json.Marshal(s2)
	err = redis_clien.LPush("list",s11).Err()
	err = redis_clien.LPush("list",s22).Err()
	if err != nil {
		panic(err)
	}

	//redis_clien.RPop("list")
}
