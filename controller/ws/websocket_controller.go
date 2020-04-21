package ws

import (
	"github.com/FlashFeiFei/my-gin/controller"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)



var hub *Hub //聊天室的房间

func init() {
	if hub == nil {
		hub = newHub() //创建全局唯一的一个房间
		go hub.run()    //做转发信息
	}
}


type WebSocketController struct {
	controller.BaseController
}

//路由页面
func (ws *WebSocketController) Home() {
	ws.Ctx.HTML(http.StatusOK, "ws/home.html", gin.H{
		"title": "聊天室demo，",
	})
}

//加入聊天室
func (ws *WebSocketController) Join() {
	//新连接的用户加入房间
	log.Println(hub)
	log.Printf("地址=%p", hub)
	serveWs(hub, ws.Ctx.Writer, ws.Ctx.Request)
}
