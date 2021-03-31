package main

import (
	"yuns/gins"
	_ "yuns/routes"
	"yuns/cache"
	"yuns/config"
)

func main() {

	conf := &gins.Config{
		Host: "0.0.0.0",
		Port: 9501,
	}
	//初始化配置
    config.Init()
	//初始化缓存和db
	cache.Init()
	//开启服务
	gins.Run(conf)
}