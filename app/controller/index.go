package controller

import (
	"yuns/gins"
	"yuns/cache"
)

func Index(ctx *gins.Context) {
	
	res := cache.RedisInstance.Get("key")
	ctx.Success(res)
}

func Test(ctx *gins.Context) {
	ctx.Faile("数据不能为空")
}

