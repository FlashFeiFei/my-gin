package ws

import (
	"github.com/FlashFeiFei/my-gin/controller"
	"github.com/FlashFeiFei/my-gin/controller/ws"
	"github.com/gin-gonic/gin"
)

//websocket路由注册
func RegisterWSRouter(router *gin.Engine) {

	//websocket模块路由
	websocket_router := router.Group("/ws")
	{
		//访问 ws/home 路由进入websocket聊天
		websocket_controller := controller.NewController(new(ws.WebSocketController))
		websocket_router.GET("/:action", websocket_controller)
	}

}
