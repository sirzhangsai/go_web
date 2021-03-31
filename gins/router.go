package gins

import (
	"github.com/gin-gonic/gin"
)
type Router struct {

	*gin.RouterGroup
	tree []func(routerGroup *gin.RouterGroup)
	next []*Router
	action map[string]HandlerFunc
}
// init 初始化路由树
func (r *Router) init() {

	for i, l := 0, len(r.tree); i < l; i++ {
		//添加路由
		r.tree[i](r.RouterGroup)
	}
	for x,y := 0, len(r.next); x < y; x++ {
		//设定父级路由,执行自路由初始化
		r.next[x].RouterGroup = r.RouterGroup
		r.next[x].init()
	}
}
//Get　请求
func (r *Router) GET(relativePath string, handlers ...HandlerFunc) *Router {

	r.tree = append(r.tree, func(routerGroup *gin.RouterGroup) {
		routerGroup.GET(relativePath, newGinHandler(handlers...)...)
	})
	return r
}