package oauth

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/FlashFeiFei/my-gin/common-lib/db/mysql"
	"github.com/FlashFeiFei/my-gin/help"
	"github.com/FlashFeiFei/my-gin/model"
	"github.com/FlashFeiFei/my-gin/model/oauth"
	"time"
)

//添加client方法
func AddClient(c *gin.Context) {
	var client Client
	if err := c.ShouldBind(&client); err != nil {
		help.Gin400NotFoundResponse(c, err, nil)
		return
	}
	now_time := time.Now()
	client_id, _ := GenerateClientId(now_time)
	client_secret, _ := GenerateClientSecret(now_time)
	oauth_cliet := oauth.OauthClient{
		ClientName:   client.ClientName,
		ClientId:     client_id,
		ClientSecret: client_secret,
		RedirectUrl:  client.RedirectUrl,
		BaseModel: model.BaseModel{
			CreatedAt: now_time.Unix(),
			UpdatedAt: now_time.Unix(),
		},
	}
	db := mysql.GetDBConnect("default")
	//create语句，在创建成功后，会查询数据库，把记录重新赋值
	db.Create(&oauth_cliet)
	if oauth_cliet.ID == 0 {
		help.Gin400NotFoundResponse(c, errors.New("新填数据失败"), nil)
		return
	}
	help.Gin200SuccessResponse(c, "添加成功", oauth_cliet)
	return
}
