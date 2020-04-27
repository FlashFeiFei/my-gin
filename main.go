package main

import (
	"fmt"
	"github.com/FlashFeiFei/my-gin/common-lib/config"
	"github.com/FlashFeiFei/my-gin/common-lib/db"
	"github.com/FlashFeiFei/my-gin/common-lib/db/mysql"
	my_router "github.com/FlashFeiFei/my-gin/router"
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

	router := new(my_router.MyRouter)
	//goin的性能分析
	//ginpprof.Wrapper(router)

	//运行服务器
	server_config, err := config.GetConfig("server")
	if err != nil {
		log.Fatalln(err)
	}

	server_info := server_config.GetStringMap("servier")
	http.ListenAndServe(fmt.Sprintf(":%d", server_info["port"].(int)), router.RunRouter())
}
