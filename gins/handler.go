package gins

import (
	"github.com/gin-gonic/gin"
)

type HandlerFunc func(*Context)

func newGinHandler(handlers ...HandlerFunc) []gin.HandlerFunc {

	l := len(handlers)
	if l <= 0 {
		return nil
	}

	ginHandlers := make([]gin.HandlerFunc, 0, l)

	for i := 0; i < l; i++ {
		handler := handlers[i]
		ginHandlers = append(ginHandlers, func(ginCtx *gin.Context) {
			val, _ := ginCtx.Get("*Context")
			ctx := val.(*Context)
			//执行原 HandlerFunc
			handler(ctx)
		})
	}
	return ginHandlers
}