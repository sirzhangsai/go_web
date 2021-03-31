package gins

import (
   "github.com/gin-gonic/gin"
)
type Context struct {

   doneChan chan struct{}
   *gin.Context
   isAPI bool
   stack string
}

func (ctx *Context) reset(ginCtx *gin.Context) {
   ctx.Context = ginCtx
}

type api struct {
   ctx       *Context
   result    apiResult
   rawResult []byte   
}

type apiResult struct {
   Code  int      `json:"code"`
   Msg   string   `json:"msg"`
   Data  interface{} `json:"data"`
}

//成功函数只接受data数据
func (ctx *Context) Success(data interface{}) {

   code := 1
   res := apiResult{
       Code: code,
       Msg: "sucees",
       Data: data,
   }
   ctx.JSON(200, res)
}
//失败只接受失败消息参数
func (ctx *Context) Faile(msg string) {

   code := -1
   res := apiResult{
       Code: code,
       Msg: msg,
       Data: "",
   }
   ctx.JSON(200, res)
}