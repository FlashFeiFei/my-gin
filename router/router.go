package router

import (
	"github.com/FlashFeiFei/my-gin/router/background"
	"github.com/FlashFeiFei/my-gin/router/ws"
	"github.com/gin-gonic/gin"
)

type MyRouter struct {
	router_list []RegisterRouter
}

func (mr *MyRouter) registerRouter() {
	mr.router_list = make([]RegisterRouter, 2)
	//路由注册
	mr.router_list[0] = background.RegisterBackgroundUserRouter
	mr.router_list[1] = ws.RegisterWSRouter
}

func (mr *MyRouter) RunRouter() *gin.Engine {

	//默认启动方式，包含 Logger、Recovery 中间件
	router := gin.Default()
	//加载html文件路径
	router.LoadHTMLGlob("templates/**/*")

	//注册路由
	mr.registerRouter()

	for _, register_router := range mr.router_list {
		register_router(router)
	}

	return router
}

//注册路由的函数
type RegisterRouter func(router *gin.Engine)
