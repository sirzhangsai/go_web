package gins

import "github.com/gin-gonic/gin"

type Middleware struct {
	engine *gin.Engine
	handlers []HandlerFunc
}
// 注册全局中间件到 gin　中
func (m *Middleware) init() {
	if len(m.handlers) <= 0 {
		return
	}
	ginHandlers := newGinHandler(m.handlers...)
	m.engine.Use(ginHandlers...)
}
// Use 添加全局中间件
func (m *Middleware) Use(middlewares ...HandlerFunc) {
	m.handlers = append(m.handlers, middlewares...)
}