package gins

import (
	"sync"
	"github.com/gin-gonic/gin"
)

var ctxPool sync.Pool

func init() {
	ctxPool.New = func() interface{} {
		return &Context{}
	}
}

func core() gin.HandlerFunc {
	return func(ginCtx *gin.Context) {
		ctx := ctxPool.New().(*Context)
		ctx.reset(ginCtx)
		ctx.Set("*Context", ctx)
		
		//go func() {
			defer func() {
			e := recover()
			if e != nil {
				ctx.Faile(e.(string))
			}
		}()
		ctx.Next()
	//	close(ctx.doneChan)
	//}()
	 //  <-ctx.doneChan
		ctxPool.Put(ctx)
}
}