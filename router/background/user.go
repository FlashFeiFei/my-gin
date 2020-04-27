package background

import (
	"github.com/FlashFeiFei/my-gin/controller"
	"github.com/FlashFeiFei/my-gin/controller/background"
	"github.com/gin-gonic/gin"
)

//后台路由注册
func RegisterBackgroundUserRouter(router *gin.Engine) {
	//后台路由
	background_router := router.Group("/background")
	{
		user_router := background_router.Group("/user")
		{
			user_controller := controller.NewController(new(background.UserController))
			user_router.GET("/:action", user_controller)     //访问background/user/hello_world?name=1
			user_router.GET("/:action/:id", user_controller) //resetful风格访问background/user/hello_world2/1?name=1
		}
	}
}
