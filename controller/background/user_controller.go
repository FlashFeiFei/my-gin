package background

import (
	"fmt"
	"github.com/FlashFeiFei/my-gin/controller"
	"github.com/FlashFeiFei/my-gin/help"
	"log"
	"time"
)

type UserController struct {
	controller.BaseController
}

func (c *UserController) HelloWorld() {
	log.Println("休息前的参数")
	log.Println(c.Ctx.Request.FormValue("name"))
	time.Sleep(time.Second * 10)
	log.Println("休息后的参数")
	log.Println(c.Ctx.Request.FormValue("name"))
	log.Println(fmt.Sprintf("context地址是同一个？,地址=%p", c.Ctx))
	help.Gin200SuccessResponse(c.Ctx, "恭喜恭喜,终于进来了！", nil)
}

//这样带上一个id,为毛我觉得这样怪怪的，直接从请求中获取不好吗
func (c *UserController) HelloWorld2(id uint64) {
	help.Gin200SuccessResponse(c.Ctx, fmt.Sprintf("恭喜恭喜,又终于进来了！id=%d", id), nil)
}
