package main

import (
	"fmt"
	"github.com/FlashFeiFei/my-gin/common-lib/config"
	"github.com/FlashFeiFei/my-gin/common-lib/db"
	"github.com/FlashFeiFei/my-gin/common-lib/db/mysql"
	"github.com/FlashFeiFei/my-gin/controller"
	"github.com/FlashFeiFei/my-gin/controller/background"
	"github.com/FlashFeiFei/my-gin/controller/oauth"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

//加载各种配置文件
func loadConfig() {
	err := config.LoadConfig("db", "./config", "yaml")
	if err != nil {
		log.Fatal(err)
	}
	err = config.LoadConfig("server", "./config", "yaml")
	if err != nil {
		log.Fatal(err)
	}
}

func loadDB() {
	db_config, _ := config.GetConfig("db")
	db_map := db_config.GetStringMap("mysql")
	db_map_size := len(db_map)
	db_info_list := make([]db.DBInfo, db_map_size)
	for key, item := range db_map {
		db_info_list = append(db_info_list, db.DBInfo{
			Key: key,
			Dsn: item.(map[string]interface{})["dsn"].(string),
		})
	}
	mysql.InitDBConnect(db_info_list...)
}

func init() {
	//加载配置文件
	loadConfig()
	//连接数据库
	//loadDB()
}
func main() {

	// 默认启动方式，包含 Logger、Recovery 中间件
	router := gin.Default()
	//goin的性能分析
	//ginpprof.Wrapper(router)

	router.POST("/oauth/client/addclient", oauth.AddClient)

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

	//运行服务器
	server_config, err := config.GetConfig("server")
	if err != nil {
		log.Fatalln(err)
	}
	server_info := server_config.GetStringMap("servier")
	http.ListenAndServe(fmt.Sprintf(":%d", server_info["port"].(int)), router)
}
