package routes

import (
	"yuns/gins"
	"yuns/app/controller"
)

var apiRouters = gins.Instance.Router

func init() {
		apiRouters.GET("/", controller.Index)
		apiRouters.GET("/test", controller.Test)

}