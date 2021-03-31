package gins

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"log"
	"context"
)
type Server struct {

	engine  *gin.Engine //gin Engin
	Router  *Router   //gin Router封装
	server     *http.Server // http服务器
	rootCtx    context.Context
	rootCancel    context.CancelFunc
	Config     *Config
	Middleware  *Middleware
}
//创建engine对象
func New() (gs *Server, err error) {

	gs = &Server{
		engine: gin.New(),
	}
	gs.Router = &Router{RouterGroup: &gs.engine.RouterGroup}
	return 
}


//初始化服务配置
func(gs *Server) Init(conf *Config) {

	gs.server = &http.Server{
	  Addr: conf.Addr(),
	  Handler: gs.engine,
	}

	//　加载核心中间件
	gs.engine.Use(core())
	//　初始化路由
	gs.Router.init()
	
}
//开启服务
func(gs *Server) Start() {
	
	err := gs.server.ListenAndServe()

	if err != nil {
		log.Fatal("服务启动失败!")
	}
	
}