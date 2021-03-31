package cache

import (
	cof "yuns/config"
	"github.com/go-redis/redis/v8"
	"context"
	"time"
	"fmt"
)
var ctxs = context.Background()

var RedisInstance Redis

type Redis struct {

	Instan *(redis.Client)
}

func Init() {
	//将来可以在这里实例化多个redis实例
	RedisInstance.Instan = redis.NewClient(&redis.Options{
		Addr: cof.Config.Cache.Redis.Host + ":" + cof.Config.Cache.Redis.Port,
	})
}

// redis get
func (rInstance *Redis) Get(key string)(res string) {

	res, err := rInstance.Instan.Get(ctxs, key).Result()

	if err != nil {
		panic("获取缓存数据时失败!"+ err.Error())
	}
	return 
}

//redis set
func (rInstance *Redis) Set(key, value string)(res string) {

	res, err := rInstance.Instan.Set(ctxs, key, value, 0).Result()
	if err != nil {
		panic("设置缓存时失败!" + err.Error())
	}
	return 
}
//redis setnx
func (rInstance *Redis) SetNX(key, value string, time time.Duration) {

	res, err := rInstance.Instan.SetNX(ctxs, key, value, time).Result()
	if err != nil {
		panic("设置缓存时失败!" + err.Error())
	}
	fmt.Println(res)
}
